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
	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kueuev1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

func CreateKueueResourceFlavor(t Test, resourceFlavorSpec kueuev1beta1.ResourceFlavorSpec) *kueuev1beta1.ResourceFlavor {
	t.T().Helper()

	resourceFlavor := &kueuev1beta1.ResourceFlavor{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kueuev1beta1.SchemeGroupVersion.String(),
			Kind:       "ResourceFlavor",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "rf-",
		},
		Spec: resourceFlavorSpec,
	}

	resourceFlavor, err := t.Client().Kueue().KueueV1beta1().ResourceFlavors().Create(t.Ctx(), resourceFlavor, metav1.CreateOptions{})
	t.Expect(err).NotTo(gomega.HaveOccurred())
	t.T().Logf("Created Kueue ResourceFlavor %s successfully", resourceFlavor.Name)

	return resourceFlavor
}

func CreateKueueClusterQueue(t Test, clusterQueueSpec kueuev1beta1.ClusterQueueSpec) *kueuev1beta1.ClusterQueue {
	t.T().Helper()

	clusterQueue := &kueuev1beta1.ClusterQueue{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kueuev1beta1.SchemeGroupVersion.String(),
			Kind:       "ClusterQueue",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "cq-",
		},
		Spec: clusterQueueSpec,
	}

	clusterQueue, err := t.Client().Kueue().KueueV1beta1().ClusterQueues().Create(t.Ctx(), clusterQueue, metav1.CreateOptions{})
	t.Expect(err).NotTo(gomega.HaveOccurred())
	t.T().Logf("Created Kueue ClusterQueue %s successfully", clusterQueue.Name)

	return clusterQueue
}

func CreateKueueLocalQueue(t Test, namespace, clusterQueueName string) *kueuev1beta1.LocalQueue {
	t.T().Helper()

	localQueue := &kueuev1beta1.LocalQueue{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kueuev1beta1.SchemeGroupVersion.String(),
			Kind:       "LocalQueue",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "lq-",
			Namespace:    namespace,
		},
		Spec: kueuev1beta1.LocalQueueSpec{
			ClusterQueue: kueuev1beta1.ClusterQueueReference(clusterQueueName),
		},
	}

	localQueue, err := t.Client().Kueue().KueueV1beta1().LocalQueues(namespace).Create(t.Ctx(), localQueue, metav1.CreateOptions{})
	t.Expect(err).NotTo(gomega.HaveOccurred())
	t.T().Logf("Created Kueue LocalQueue %s/%s successfully", localQueue.Namespace, localQueue.Name)

	return localQueue
}
