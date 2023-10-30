// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumBGPNodeConfigOverridesGetter has a method to return a CiliumBGPNodeConfigOverrideInterface.
// A group's client should implement this interface.
type CiliumBGPNodeConfigOverridesGetter interface {
	CiliumBGPNodeConfigOverrides() CiliumBGPNodeConfigOverrideInterface
}

// CiliumBGPNodeConfigOverrideInterface has methods to work with CiliumBGPNodeConfigOverride resources.
type CiliumBGPNodeConfigOverrideInterface interface {
	Create(ctx context.Context, ciliumBGPNodeConfigOverride *v2alpha1.CiliumBGPNodeConfigOverride, opts v1.CreateOptions) (*v2alpha1.CiliumBGPNodeConfigOverride, error)
	Update(ctx context.Context, ciliumBGPNodeConfigOverride *v2alpha1.CiliumBGPNodeConfigOverride, opts v1.UpdateOptions) (*v2alpha1.CiliumBGPNodeConfigOverride, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumBGPNodeConfigOverride, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumBGPNodeConfigOverrideList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPNodeConfigOverride, err error)
	CiliumBGPNodeConfigOverrideExpansion
}

// ciliumBGPNodeConfigOverrides implements CiliumBGPNodeConfigOverrideInterface
type ciliumBGPNodeConfigOverrides struct {
	client rest.Interface
}

// newCiliumBGPNodeConfigOverrides returns a CiliumBGPNodeConfigOverrides
func newCiliumBGPNodeConfigOverrides(c *CiliumV2alpha1Client) *ciliumBGPNodeConfigOverrides {
	return &ciliumBGPNodeConfigOverrides{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumBGPNodeConfigOverride, and returns the corresponding ciliumBGPNodeConfigOverride object, and an error if there is any.
func (c *ciliumBGPNodeConfigOverrides) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumBGPNodeConfigOverride, err error) {
	result = &v2alpha1.CiliumBGPNodeConfigOverride{}
	err = c.client.Get().
		Resource("ciliumbgpnodeconfigoverrides").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumBGPNodeConfigOverrides that match those selectors.
func (c *ciliumBGPNodeConfigOverrides) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumBGPNodeConfigOverrideList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.CiliumBGPNodeConfigOverrideList{}
	err = c.client.Get().
		Resource("ciliumbgpnodeconfigoverrides").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumBGPNodeConfigOverrides.
func (c *ciliumBGPNodeConfigOverrides) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumbgpnodeconfigoverrides").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumBGPNodeConfigOverride and creates it.  Returns the server's representation of the ciliumBGPNodeConfigOverride, and an error, if there is any.
func (c *ciliumBGPNodeConfigOverrides) Create(ctx context.Context, ciliumBGPNodeConfigOverride *v2alpha1.CiliumBGPNodeConfigOverride, opts v1.CreateOptions) (result *v2alpha1.CiliumBGPNodeConfigOverride, err error) {
	result = &v2alpha1.CiliumBGPNodeConfigOverride{}
	err = c.client.Post().
		Resource("ciliumbgpnodeconfigoverrides").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumBGPNodeConfigOverride).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumBGPNodeConfigOverride and updates it. Returns the server's representation of the ciliumBGPNodeConfigOverride, and an error, if there is any.
func (c *ciliumBGPNodeConfigOverrides) Update(ctx context.Context, ciliumBGPNodeConfigOverride *v2alpha1.CiliumBGPNodeConfigOverride, opts v1.UpdateOptions) (result *v2alpha1.CiliumBGPNodeConfigOverride, err error) {
	result = &v2alpha1.CiliumBGPNodeConfigOverride{}
	err = c.client.Put().
		Resource("ciliumbgpnodeconfigoverrides").
		Name(ciliumBGPNodeConfigOverride.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumBGPNodeConfigOverride).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumBGPNodeConfigOverride and deletes it. Returns an error if one occurs.
func (c *ciliumBGPNodeConfigOverrides) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumbgpnodeconfigoverrides").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumBGPNodeConfigOverrides) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumbgpnodeconfigoverrides").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumBGPNodeConfigOverride.
func (c *ciliumBGPNodeConfigOverrides) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPNodeConfigOverride, err error) {
	result = &v2alpha1.CiliumBGPNodeConfigOverride{}
	err = c.client.Patch(pt).
		Resource("ciliumbgpnodeconfigoverrides").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}