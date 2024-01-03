package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	imagev1 "github.com/openshift/api/image/v1"
)

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
	fakeClient := NewFakeKubeClientWithScheme(scheme, fakeImageStream...)

	image := &imagev1.ImageStream{}

	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "my-imagestream-1", Namespace: "my-namespace"}, image)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(image.Name).To(gomega.Equal("my-imagestream-1"))
	g.Expect(image.Namespace).To(gomega.Equal("my-namespace"))
}
