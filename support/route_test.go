package support

import (
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	routev1 "github.com/openshift/api/route/v1"
)

func TestGetRoute(t *testing.T) {

	test := NewTest(t)

	route := &routev1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-1",
			Namespace: "my-namespace",
		},
	}

	test.client.Route().RouteV1().Routes("my-namespace").Create(test.ctx, route, metav1.CreateOptions{})

	routes := GetRoute(test, "my-namespace", "test-1")

	test.Expect(routes.Name).To(gomega.Equal("test-1"))
	test.Expect(routes.Namespace).To(gomega.Equal("my-namespace"))

}
