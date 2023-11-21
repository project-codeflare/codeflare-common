package support

/*
import (
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	imagev1 "github.com/openshift/api/image/v1"
)

func TestGetImageStream(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Create a fake client that returns different ImageStream objects.
	fakeImageStream := []runtime.Object{
		&imagev1.ImageStream{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-imagestream-1",
				Namespace: "my-namespace",
			},
		},
	}
	fakeClient := NewFakeKubeClientWithObjects(fakeImageStream...)

	test := With(t).(*T)
	test.client = &testClient{
		core: fakeClient,
	}

	// Call the ImageStream function using the fake client
	imageStream := GetImageStream(test, "my-namespace", "my-image-stream")

	// Assertions
	g.Expect(imageStream.Name).To(gomega.Equal("my-image-stream"))
	g.Expect(imageStream.Namespace).To(gomega.Equal("my-namespace"))

}
*/
