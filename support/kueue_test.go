/*
Copyright 2024.

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

package support

import (
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kueuev1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

func TestCreateKueueResourceFlavor(t *testing.T) {
	test := NewTest(t)

	rfSpec := kueuev1beta1.ResourceFlavorSpec{}

	rf := CreateKueueResourceFlavor(test, rfSpec)

	test.Expect(rf).To(gomega.Not(gomega.BeNil()))
	test.Expect(rf.GenerateName).To(gomega.Equal("rf-"))
}

func TestCreateKueueClusterQueue(t *testing.T) {
	test := NewTest(t)

	cqSpec := kueuev1beta1.ClusterQueueSpec{
		NamespaceSelector: &metav1.LabelSelector{},
	}

	cq := CreateKueueClusterQueue(test, cqSpec)

	test.Expect(cq).To(gomega.Not(gomega.BeNil()))
	test.Expect(cq.GenerateName).To(gomega.Equal("cq-"))
}

func TestCreateKueueLocalQueue(t *testing.T) {
	test := NewTest(t)

	lq := CreateKueueLocalQueue(test, "ns-1", "cq-1")

	test.Expect(lq).To(gomega.Not(gomega.BeNil()))
	test.Expect(lq.GenerateName).To(gomega.Equal("lq-"))
	test.Expect(lq.Namespace).To(gomega.Equal("ns-1"))
	test.Expect(lq.Spec.ClusterQueue).To(gomega.Equal(kueuev1beta1.ClusterQueueReference("cq-1")))
}
