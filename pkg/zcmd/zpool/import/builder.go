/*
Copyright 2019 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pimport

import (
	"fmt"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/openebs/cstor-operators/pkg/zcmd/bin"
	"github.com/pkg/errors"
)

const (
	// Operation defines type of zfs operation
	Operation = "import"
)

//PoolImport defines structure for pool 'Import' operation
type PoolImport struct {
	//cachefile to do import
	Cachefile string

	//directory to search
	Directorylist []string

	//import all pools
	ImportAll bool

	//force import
	ForceImport bool

	//property list
	Property []string

	//pool name or id
	Pool string

	//new pool name
	NewPool string

	// command string
	Command string

	// checks is list of predicate function used for validating object
	checks []PredicateFunc

	// Executor is to execute the commands
	Executor bin.Executor

	// error
	err error
}

// NewPoolImport returns new instance of object PoolImport
func NewPoolImport() *PoolImport {
	return &PoolImport{}
}

// WithCheck add given check to checks list
func (p *PoolImport) WithCheck(check ...PredicateFunc) *PoolImport {
	p.checks = append(p.checks, check...)
	return p
}

// WithCachefile method fills the Cachefile field of PoolImport object.
func (p *PoolImport) WithCachefile(Cachefile string) *PoolImport {
	p.Cachefile = Cachefile
	return p
}

// WithDirectory method fills the Directorylist field of PoolImport object.
func (p *PoolImport) WithDirectory(dir string) *PoolImport {
	p.Directorylist = append(p.Directorylist, dir)
	return p
}

// WithImportAll method fills the ImportAll field of PoolImport object.
func (p *PoolImport) WithImportAll(ImportAll bool) *PoolImport {
	p.ImportAll = ImportAll
	return p
}

// WithForceImport method fills the ForceImport field of PoolImport object.
func (p *PoolImport) WithForceImport(ForceImport bool) *PoolImport {
	p.ForceImport = ForceImport
	return p
}

// WithProperty method fills the Property field of PoolImport object.
func (p *PoolImport) WithProperty(key, value string) *PoolImport {
	if len(value) != 0 {
		p.Property = append(p.Property, fmt.Sprintf("%s=%s", key, value))
	}
	return p
}

// WithPool method fills the Pool field of PoolImport object.
func (p *PoolImport) WithPool(Pool string) *PoolImport {
	p.Pool = Pool
	return p
}

// WithNewPool method fills the NewPool field of PoolImport object.
func (p *PoolImport) WithNewPool(NewPool string) *PoolImport {
	p.NewPool = NewPool
	return p
}

// WithCommand method fills the Command field of PoolImport object.
func (p *PoolImport) WithCommand(Command string) *PoolImport {
	p.Command = Command
	return p
}

// WithExecutor method fills the Executor field of PoolImport object.
func (p *PoolImport) WithExecutor(executor bin.Executor) *PoolImport {
	p.Executor = executor
	return p
}

// Validate is to validate generated PoolImport object by builder
func (p *PoolImport) Validate() *PoolImport {
	for _, check := range p.checks {
		if !check(p) {
			p.err = errors.Wrapf(p.err, "validation failed {%v}", runtime.FuncForPC(reflect.ValueOf(check).Pointer()).Name())
		}
	}
	return p
}

// Execute is to execute generated PoolImport object
func (p *PoolImport) Execute() ([]byte, error) {
	p, err := p.Build()
	if err != nil {
		return nil, err
	}

	if IsExecutorSet()(p) {
		return p.Executor.Execute(p.Command)
	}
	// execute command here
	// #nosec
	return exec.Command(bin.BASH, "-c", p.Command).CombinedOutput()
}

// Build returns the PoolImport object generated by builder
func (p *PoolImport) Build() (*PoolImport, error) {
	var c strings.Builder
	p = p.Validate()
	p.appendCommand(&c, bin.ZPOOL)
	p.appendCommand(&c, fmt.Sprintf(" %s ", Operation))

	if IsPropertySet()(p) {
		for _, v := range p.Property {
			p.appendCommand(&c, fmt.Sprintf(" -o %s ", v))
		}
	}

	if !IsCachefileSet()(p) && IsDirectorylistSet()(p) {
		for _, i := range p.Directorylist {
			if len(i) != 0 {
				p.appendCommand(&c, fmt.Sprintf(" -d %s ", i))
			}
		}
	}

	if IsCachefileSet()(p) {
		p.appendCommand(&c, fmt.Sprintf(" -c %s ", p.Cachefile))
	}

	if IsForceImportSet()(p) {
		p.appendCommand(&c, fmt.Sprintf(" -f "))
	}

	if IsImportAllSet()(p) {
		p.appendCommand(&c, fmt.Sprintf(" -a "))
	} else {
		p.appendCommand(&c, fmt.Sprintf(p.Pool))

		if IsNewPoolSet()(p) {
			p.appendCommand(&c, fmt.Sprintf(" %s ", p.NewPool))
		}
	}

	p.Command = c.String()
	return p, p.err
}

// appendCommand append string to given string builder
func (p *PoolImport) appendCommand(c *strings.Builder, cmd string) {
	_, err := c.WriteString(cmd)
	if err != nil {
		p.err = errors.Wrapf(p.err, "Failed to append cmd{%s} : %s", cmd, err.Error())
	}
}