package support

import (
	"testing"

	"github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPods(t *testing.T) {
	// Create a fake Kubernetes client for testing
	fakeClient := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "test-namespace",
		},
	})

	test := With(t).(*T)
	test.client = &testClient{
		core: fakeClient,
	}

	// Call the GetPods function with the fake client and namespace
	pods := GetPods(test, "test-namespace", metav1.ListOptions{})

	test.Expect(pods).Should(gomega.HaveLen(1), "Expected 1 pod, but got %d", len(pods))
	test.Expect(pods[0].Name).To(gomega.Equal("test-pod"), "Expected pod name 'test-pod', but got '%s'", pods[0].Name)
}

func TestGetNodes(t *testing.T) {
	// Create a fake Kubernetes client for testing
	fakeClient := fake.NewSimpleClientset(&corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-node",
		},
	})

	test := With(t).(*T)
	test.client = &testClient{
		core: fakeClient,
	}
	nodes := GetNodes(test)

	test.Expect(nodes).Should(gomega.HaveLen(1), "Expected 1 node, but got %d", len(nodes))
	test.Expect(nodes[0].Name).To(gomega.Equal("test-node"), "Expected node name 'test-node', but got '%s'", nodes[0].Name)

}
