/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "Hybrid_Cloud/pkg/apis/sync/v1alpha1"
	scheme "Hybrid_Cloud/pkg/client/sync/v1alpha1/clientset/versioned/scheme"
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SyncsGetter has a method to return a SyncInterface.
// A group's client should implement this interface.
type SyncsGetter interface {
	Syncs(namespace string) SyncInterface
}

// SyncInterface has methods to work with Sync resources.
type SyncInterface interface {
	Create(ctx context.Context, sync *v1alpha1.Sync, opts v1.CreateOptions) (*v1alpha1.Sync, error)
	Update(ctx context.Context, sync *v1alpha1.Sync, opts v1.UpdateOptions) (*v1alpha1.Sync, error)
	UpdateStatus(ctx context.Context, sync *v1alpha1.Sync, opts v1.UpdateOptions) (*v1alpha1.Sync, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Sync, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SyncList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Sync, err error)
	SyncExpansion
}

// syncs implements SyncInterface
type syncs struct {
	client rest.Interface
	ns     string
}

// newSyncs returns a Syncs
func newSyncs(c *HcpV1alpha1Client, namespace string) *syncs {
	return &syncs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sync, and returns the corresponding sync object, and an error if there is any.
func (c *syncs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Sync, err error) {
	result = &v1alpha1.Sync{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("syncs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Syncs that match those selectors.
func (c *syncs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SyncList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SyncList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("syncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested syncs.
func (c *syncs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("syncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sync and creates it.  Returns the server's representation of the sync, and an error, if there is any.
func (c *syncs) Create(ctx context.Context, sync *v1alpha1.Sync, opts v1.CreateOptions) (result *v1alpha1.Sync, err error) {
	result = &v1alpha1.Sync{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("syncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sync).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sync and updates it. Returns the server's representation of the sync, and an error, if there is any.
func (c *syncs) Update(ctx context.Context, sync *v1alpha1.Sync, opts v1.UpdateOptions) (result *v1alpha1.Sync, err error) {
	result = &v1alpha1.Sync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("syncs").
		Name(sync.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sync).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *syncs) UpdateStatus(ctx context.Context, sync *v1alpha1.Sync, opts v1.UpdateOptions) (result *v1alpha1.Sync, err error) {
	result = &v1alpha1.Sync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("syncs").
		Name(sync.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sync).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sync and deletes it. Returns an error if one occurs.
func (c *syncs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("syncs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *syncs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("syncs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sync.
func (c *syncs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Sync, err error) {
	result = &v1alpha1.Sync{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("syncs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
