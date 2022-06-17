/*
Copyright 2021 The OpenEBS Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openebs/api/v3/pkg/apis/cstor/v1"
	scheme "github.com/openebs/api/v3/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CStorVolumesGetter has a method to return a CStorVolumeInterface.
// A group's client should implement this interface.
type CStorVolumesGetter interface {
	CStorVolumes(namespace string) CStorVolumeInterface
}

// CStorVolumeInterface has methods to work with CStorVolume resources.
type CStorVolumeInterface interface {
	Create(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.CreateOptions) (*v1.CStorVolume, error)
	Update(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.UpdateOptions) (*v1.CStorVolume, error)
	UpdateStatus(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.UpdateOptions) (*v1.CStorVolume, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CStorVolume, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CStorVolumeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CStorVolume, err error)
	CStorVolumeExpansion
}

// cStorVolumes implements CStorVolumeInterface
type cStorVolumes struct {
	client rest.Interface
	ns     string
}

// newCStorVolumes returns a CStorVolumes
func newCStorVolumes(c *CstorV1Client, namespace string) *cStorVolumes {
	return &cStorVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cStorVolume, and returns the corresponding cStorVolume object, and an error if there is any.
func (c *cStorVolumes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CStorVolume, err error) {
	result = &v1.CStorVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CStorVolumes that match those selectors.
func (c *cStorVolumes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CStorVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CStorVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cStorVolumes.
func (c *cStorVolumes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cStorVolume and creates it.  Returns the server's representation of the cStorVolume, and an error, if there is any.
func (c *cStorVolumes) Create(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.CreateOptions) (result *v1.CStorVolume, err error) {
	result = &v1.CStorVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cstorvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cStorVolume).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cStorVolume and updates it. Returns the server's representation of the cStorVolume, and an error, if there is any.
func (c *cStorVolumes) Update(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.UpdateOptions) (result *v1.CStorVolume, err error) {
	result = &v1.CStorVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorvolumes").
		Name(cStorVolume.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cStorVolume).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cStorVolumes) UpdateStatus(ctx context.Context, cStorVolume *v1.CStorVolume, opts metav1.UpdateOptions) (result *v1.CStorVolume, err error) {
	result = &v1.CStorVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorvolumes").
		Name(cStorVolume.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cStorVolume).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cStorVolume and deletes it. Returns an error if one occurs.
func (c *cStorVolumes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorvolumes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cStorVolumes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorvolumes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cStorVolume.
func (c *cStorVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CStorVolume, err error) {
	result = &v1.CStorVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cstorvolumes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}