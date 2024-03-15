package support

import (
	"testing"

	"github.com/onsi/gomega"
	mcadv1beta2 "github.com/project-codeflare/appwrapper/api/v1beta2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestGetAppWrapper(t *testing.T) {
	test := NewTest(t)

	name := "my-appwrapper-1"
	namespace := "my-namespace-1"
	aw := &mcadv1beta2.AppWrapper{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
		Status:     mcadv1beta2.AppWrapperStatus{Phase: mcadv1beta2.AppWrapperRunning},
	}
	awMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(aw)
	test.Expect(err).NotTo(gomega.HaveOccurred())
	unstruct := unstructured.Unstructured{Object: awMap}

	_, err = test.Client().Dynamic().Resource(appWrapperResource).Namespace(namespace).Create(test.ctx, &unstruct, metav1.CreateOptions{})
	test.Expect(err).NotTo(gomega.HaveOccurred())

	aw2 := GetAppWrapper(test, namespace, name)
	test.Expect(aw2.Name).To(gomega.Equal(name))
	test.Expect(aw2.Namespace).To(gomega.Equal(namespace))
	test.Expect(AppWrapperName(aw2)).To(gomega.Equal(name))
	test.Expect(AppWrapperPhase(aw)).To(gomega.Equal(mcadv1beta2.AppWrapperRunning))
}
