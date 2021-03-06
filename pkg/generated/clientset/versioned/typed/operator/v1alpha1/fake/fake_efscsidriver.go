// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/bertinatto/aws-efs-csi-driver-operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEFSCSIDrivers implements EFSCSIDriverInterface
type FakeEFSCSIDrivers struct {
	Fake *FakeCsiV1alpha1
	ns   string
}

var efscsidriversResource = schema.GroupVersionResource{Group: "csi.efs.aws.com", Version: "v1alpha1", Resource: "efscsidrivers"}

var efscsidriversKind = schema.GroupVersionKind{Group: "csi.efs.aws.com", Version: "v1alpha1", Kind: "EFSCSIDriver"}

// Get takes name of the eFSCSIDriver, and returns the corresponding eFSCSIDriver object, and an error if there is any.
func (c *FakeEFSCSIDrivers) Get(name string, options v1.GetOptions) (result *v1alpha1.EFSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(efscsidriversResource, c.ns, name), &v1alpha1.EFSCSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EFSCSIDriver), err
}

// List takes label and field selectors, and returns the list of EFSCSIDrivers that match those selectors.
func (c *FakeEFSCSIDrivers) List(opts v1.ListOptions) (result *v1alpha1.EFSCSIDriverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(efscsidriversResource, efscsidriversKind, c.ns, opts), &v1alpha1.EFSCSIDriverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EFSCSIDriverList{ListMeta: obj.(*v1alpha1.EFSCSIDriverList).ListMeta}
	for _, item := range obj.(*v1alpha1.EFSCSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eFSCSIDrivers.
func (c *FakeEFSCSIDrivers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(efscsidriversResource, c.ns, opts))

}

// Create takes the representation of a eFSCSIDriver and creates it.  Returns the server's representation of the eFSCSIDriver, and an error, if there is any.
func (c *FakeEFSCSIDrivers) Create(eFSCSIDriver *v1alpha1.EFSCSIDriver) (result *v1alpha1.EFSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(efscsidriversResource, c.ns, eFSCSIDriver), &v1alpha1.EFSCSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EFSCSIDriver), err
}

// Update takes the representation of a eFSCSIDriver and updates it. Returns the server's representation of the eFSCSIDriver, and an error, if there is any.
func (c *FakeEFSCSIDrivers) Update(eFSCSIDriver *v1alpha1.EFSCSIDriver) (result *v1alpha1.EFSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(efscsidriversResource, c.ns, eFSCSIDriver), &v1alpha1.EFSCSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EFSCSIDriver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEFSCSIDrivers) UpdateStatus(eFSCSIDriver *v1alpha1.EFSCSIDriver) (*v1alpha1.EFSCSIDriver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(efscsidriversResource, "status", c.ns, eFSCSIDriver), &v1alpha1.EFSCSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EFSCSIDriver), err
}

// Delete takes name of the eFSCSIDriver and deletes it. Returns an error if one occurs.
func (c *FakeEFSCSIDrivers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(efscsidriversResource, c.ns, name), &v1alpha1.EFSCSIDriver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEFSCSIDrivers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(efscsidriversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EFSCSIDriverList{})
	return err
}

// Patch applies the patch and returns the patched eFSCSIDriver.
func (c *FakeEFSCSIDrivers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EFSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(efscsidriversResource, c.ns, name, pt, data, subresources...), &v1alpha1.EFSCSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EFSCSIDriver), err
}
