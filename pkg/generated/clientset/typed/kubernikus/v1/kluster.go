package v1

import (
	v1 "github.com/sapcc/kubernikus/pkg/apis/kubernikus/v1"
	scheme "github.com/sapcc/kubernikus/pkg/generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KlustersGetter has a method to return a KlusterInterface.
// A group's client should implement this interface.
type KlustersGetter interface {
	Klusters(namespace string) KlusterInterface
}

// KlusterInterface has methods to work with Kluster resources.
type KlusterInterface interface {
	Create(*v1.Kluster) (*v1.Kluster, error)
	Update(*v1.Kluster) (*v1.Kluster, error)
	UpdateStatus(*v1.Kluster) (*v1.Kluster, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Kluster, error)
	List(opts meta_v1.ListOptions) (*v1.KlusterList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Kluster, err error)
	KlusterExpansion
}

// klusters implements KlusterInterface
type klusters struct {
	client rest.Interface
	ns     string
}

// newKlusters returns a Klusters
func newKlusters(c *KubernikusV1Client, namespace string) *klusters {
	return &klusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kluster, and returns the corresponding kluster object, and an error if there is any.
func (c *klusters) Get(name string, options meta_v1.GetOptions) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Klusters that match those selectors.
func (c *klusters) List(opts meta_v1.ListOptions) (result *v1.KlusterList, err error) {
	result = &v1.KlusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested klusters.
func (c *klusters) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kluster and creates it.  Returns the server's representation of the kluster, and an error, if there is any.
func (c *klusters) Create(kluster *v1.Kluster) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("klusters").
		Body(kluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kluster and updates it. Returns the server's representation of the kluster, and an error, if there is any.
func (c *klusters) Update(kluster *v1.Kluster) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("klusters").
		Name(kluster.Name).
		Body(kluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *klusters) UpdateStatus(kluster *v1.Kluster) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("klusters").
		Name(kluster.Name).
		SubResource("status").
		Body(kluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the kluster and deletes it. Returns an error if one occurs.
func (c *klusters) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("klusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *klusters) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kluster.
func (c *klusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("klusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
