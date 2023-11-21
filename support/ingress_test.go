package support

import (
	"testing"

	"github.com/onsi/gomega"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func NewFakeKubeClientForIngress(objects ...runtime.Object) *fake.Clientset {
	fakeClient := fake.NewSimpleClientset(objects...)
	return fakeClient
}

func TestGetIngress(t *testing.T) {

	g := gomega.NewGomegaWithT(t)
	// Create a fake client that returns different Ingress objects.
	fakeIngress := []runtime.Object{
		&networkingv1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-ingress-1",
				Namespace: "my-namespace",
			},
		},
	}
	fakeClient := NewFakeKubeClientForIngress(fakeIngress...)

	test := With(t).(*T)
	test.client = &testClient{
		core: fakeClient,
	}


	// Call the Ingress function using the fake client
	ingressFunc := Ingress(test, "my-namespace", "my-ingress-1")
	ingress := ingressFunc(g)

	//fmt.Printf("Retrieved ingress object: %+v\n", ingress)

	// Assertions
	g.Expect(ingress.Name).To(gomega.Equal("my-ingress-1"))
	g.Expect(ingress.Namespace).To(gomega.Equal("my-namespace"))
}
