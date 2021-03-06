/*
Copyright 2018 The Federation v2 Authors.

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
package fake

import (
	federatedscheduling "github.com/kubernetes-sigs/federation-v2/pkg/apis/federatedscheduling"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicaSchedulingPreferences implements ReplicaSchedulingPreferenceInterface
type FakeReplicaSchedulingPreferences struct {
	Fake *FakeFederatedscheduling
	ns   string
}

var replicaschedulingpreferencesResource = schema.GroupVersionResource{Group: "federatedscheduling.k8s.io", Version: "", Resource: "replicaschedulingpreferences"}

var replicaschedulingpreferencesKind = schema.GroupVersionKind{Group: "federatedscheduling.k8s.io", Version: "", Kind: "ReplicaSchedulingPreference"}

// Get takes name of the replicaSchedulingPreference, and returns the corresponding replicaSchedulingPreference object, and an error if there is any.
func (c *FakeReplicaSchedulingPreferences) Get(name string, options v1.GetOptions) (result *federatedscheduling.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicaschedulingpreferencesResource, c.ns, name), &federatedscheduling.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), err
}

// List takes label and field selectors, and returns the list of ReplicaSchedulingPreferences that match those selectors.
func (c *FakeReplicaSchedulingPreferences) List(opts v1.ListOptions) (result *federatedscheduling.ReplicaSchedulingPreferenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicaschedulingpreferencesResource, replicaschedulingpreferencesKind, c.ns, opts), &federatedscheduling.ReplicaSchedulingPreferenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &federatedscheduling.ReplicaSchedulingPreferenceList{}
	for _, item := range obj.(*federatedscheduling.ReplicaSchedulingPreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSchedulingPreferences.
func (c *FakeReplicaSchedulingPreferences) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicaschedulingpreferencesResource, c.ns, opts))

}

// Create takes the representation of a replicaSchedulingPreference and creates it.  Returns the server's representation of the replicaSchedulingPreference, and an error, if there is any.
func (c *FakeReplicaSchedulingPreferences) Create(replicaSchedulingPreference *federatedscheduling.ReplicaSchedulingPreference) (result *federatedscheduling.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicaschedulingpreferencesResource, c.ns, replicaSchedulingPreference), &federatedscheduling.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), err
}

// Update takes the representation of a replicaSchedulingPreference and updates it. Returns the server's representation of the replicaSchedulingPreference, and an error, if there is any.
func (c *FakeReplicaSchedulingPreferences) Update(replicaSchedulingPreference *federatedscheduling.ReplicaSchedulingPreference) (result *federatedscheduling.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicaschedulingpreferencesResource, c.ns, replicaSchedulingPreference), &federatedscheduling.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSchedulingPreferences) UpdateStatus(replicaSchedulingPreference *federatedscheduling.ReplicaSchedulingPreference) (*federatedscheduling.ReplicaSchedulingPreference, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicaschedulingpreferencesResource, "status", c.ns, replicaSchedulingPreference), &federatedscheduling.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), err
}

// Delete takes name of the replicaSchedulingPreference and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSchedulingPreferences) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicaschedulingpreferencesResource, c.ns, name), &federatedscheduling.ReplicaSchedulingPreference{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSchedulingPreferences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicaschedulingpreferencesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &federatedscheduling.ReplicaSchedulingPreferenceList{})
	return err
}

// Patch applies the patch and returns the patched replicaSchedulingPreference.
func (c *FakeReplicaSchedulingPreferences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federatedscheduling.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicaschedulingpreferencesResource, c.ns, name, data, subresources...), &federatedscheduling.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), err
}
