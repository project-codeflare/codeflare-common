package support

import (
	"testing"

	"github.com/onsi/gomega"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetIngress(t *testing.T) {

	test := NewTest(t)
	// Create a fake client that returns Ingress objects.
	Ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-ingress-1",
			Namespace: "my-namespace",
		},
	}

	test.client.Core().NetworkingV1().Ingresses("my-namespace").Create(test.ctx, Ingress, metav1.CreateOptions{})

	// Call the Ingress function using the fake client
	ingress := GetIngress(test, "my-namespace", "my-ingress-1")

	test.Expect(ingress.Name).To(gomega.Equal("my-ingress-1"))
	test.Expect(ingress.Namespace).To(gomega.Equal("my-namespace"))
}
