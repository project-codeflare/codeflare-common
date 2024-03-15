/*
Copyright 2023.

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
	mcadv1beta2 "github.com/project-codeflare/appwrapper/api/v1beta2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var appWrapperResource = mcadv1beta2.GroupVersion.WithResource("appwrappers")

func AppWrapper(t Test, namespace string, name string) func(g gomega.Gomega) *mcadv1beta2.AppWrapper {
	return func(g gomega.Gomega) *mcadv1beta2.AppWrapper {
		unstruct, err := t.Client().Dynamic().Resource(appWrapperResource).Namespace(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		aw := &mcadv1beta2.AppWrapper{}
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstruct.UnstructuredContent(), aw)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return aw
	}
}

func GetAppWrapper(t Test, namespace string, name string) *mcadv1beta2.AppWrapper {
	t.T().Helper()
	return AppWrapper(t, namespace, name)(t)
}

func AppWrappers(t Test, namespace string) func(g gomega.Gomega) []*mcadv1beta2.AppWrapper {
	return func(g gomega.Gomega) []*mcadv1beta2.AppWrapper {
		aws, err := t.Client().Dynamic().Resource(appWrapperResource).Namespace(namespace).List(t.Ctx(), metav1.ListOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())

		awsp := []*mcadv1beta2.AppWrapper{}
		for _, v := range aws.Items {
			aw := &mcadv1beta2.AppWrapper{}
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(v.UnstructuredContent(), aw)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			awsp = append(awsp, aw)
		}

		return awsp
	}
}

func AppWrapperName(aw *mcadv1beta2.AppWrapper) string {
	return aw.Name
}

func AppWrapperPhase(aw *mcadv1beta2.AppWrapper) mcadv1beta2.AppWrapperPhase {
	return aw.Status.Phase
}
