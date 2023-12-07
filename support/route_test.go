package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	routev1 "github.com/openshift/api/route/v1"
)

func NewFakeKubeClientWithRoute(scheme *runtime.Scheme, objects ...client.Object) client.Client {
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).Build()
}

func TestGetRoute(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = routev1.AddToScheme(scheme)

	fakeroute := []client.Object{
		&routev1.Route{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-1",
				Namespace: "my-namespace",
			},
		},
	}
	fakeClient := NewFakeKubeClientWithRoute(scheme, fakeroute...)

	route := &routev1.Route{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "test-1", Namespace: "my-namespace"}, route)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	// Assertions
	g.Expect(route.Name).To(gomega.Equal("test-1"))
	g.Expect(route.Namespace).To(gomega.Equal("my-namespace"))

}
