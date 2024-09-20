package support

import (
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	imagev1 "github.com/openshift/api/image/v1"
)

func TestGetImageStream(t *testing.T) {

	test := NewTest(t)

	ImageStream := &imagev1.ImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-imagestream-1",
			Namespace: "my-namespace",
		},
	}

	test.client.Image().ImageV1().ImageStreams("my-namespace").Create(test.ctx, ImageStream, metav1.CreateOptions{})

	image := GetImageStream(test, "my-namespace", "my-imagestream-1")

	test.Expect(image.Name).To(gomega.Equal("my-imagestream-1"))
	test.Expect(image.Namespace).To(gomega.Equal("my-namespace"))
}
