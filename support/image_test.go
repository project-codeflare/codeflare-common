package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	imagev1 "github.com/openshift/api/image/v1"
)

func NewFakeKubeClientWithImages(scheme *runtime.Scheme, objects ...client.Object) client.Client {
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).Build()
}

func TestGetImageStream(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = imagev1.AddToScheme(scheme)

	fakeImageStream := []client.Object{
		&imagev1.ImageStream{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-imagestream-1",
				Namespace: "my-namespace",
			},
		},
	}
	fakeClient := NewFakeKubeClientWithImages(scheme, fakeImageStream...)

	image := &imagev1.ImageStream{}

	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "my-imagestream-1", Namespace: "my-namespace"}, image)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	// Assertions
	g.Expect(image.Name).To(gomega.Equal("my-imagestream-1"))
	g.Expect(image.Namespace).To(gomega.Equal("my-namespace"))
}
