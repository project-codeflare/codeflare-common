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
	"sigs.k8s.io/kueue/apis/kueue/v1beta1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateKueueLocalQueue(t Test, namespace, clusterQueueName string) *v1beta1.LocalQueue {
	t.T().Helper()

	localQueue := &v1beta1.LocalQueue{
		TypeMeta: metav1.TypeMeta{
			APIVersion: corev1.SchemeGroupVersion.String(),
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "lq-",
			Namespace:    namespace,
		},
		Spec: v1beta1.LocalQueueSpec{
			ClusterQueue: v1beta1.ClusterQueueReference(clusterQueueName),
		},
	}

	localQueue, err := t.Client().Kueue().KueueV1beta1().LocalQueues(namespace).Create(t.Ctx(), localQueue, metav1.CreateOptions{})
	t.Expect(err).NotTo(gomega.HaveOccurred())
	t.T().Logf("Created Kueue LocalQueue %s/%s successfully", localQueue.Namespace, localQueue.Name)

	return localQueue
}
