package support

import (
	"testing"

	"github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetPods(t *testing.T) {
	test := NewTest(t)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "test-namespace",
		},
	}

	test.client.Core().CoreV1().Pods("test-namespace").Create(test.ctx, pod, metav1.CreateOptions{})

	// Call the GetPods function with the fake client and namespace
	pods := GetPods(test, "test-namespace", metav1.ListOptions{})

	test.Expect(pods).Should(gomega.HaveLen(1), "Expected 1 pod, but got %d", len(pods))
	test.Expect(pods[0].Name).To(gomega.Equal("test-pod"), "Expected pod name 'test-pod', but got '%s'", pods[0].Name)
}

func TestGetNodes(t *testing.T) {
	test := NewTest(t)
	node := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-node",
		},
	}

	test.client.Core().CoreV1().Nodes().Create(test.ctx, node, metav1.CreateOptions{})
	nodes := GetNodes(test)

	test.Expect(nodes).Should(gomega.HaveLen(1), "Expected 1 node, but got %d", len(nodes))
	test.Expect(nodes[0].Name).To(gomega.Equal("test-node"), "Expected node name 'test-node', but got '%s'", nodes[0].Name)

}

func TestGetContainerName(t *testing.T) {
	test := NewTest(t)
	container := corev1.Container{
		Name: "test-container",
	}
	containerName := GetContainerName(test, container)
	test.Expect(containerName).To(gomega.Equal("test-container"), "Expected container name 'test-container', but got '%s'", containerName)
}

func TestGetVolumeName(t *testing.T) {
	test := NewTest(t)
	volume := corev1.Volume{
		Name: "test-volume",
	}
	volumeName := GetVolumeName(test, volume)
	test.Expect(volumeName).To(gomega.Equal("test-volume"), "Expected volume name 'test-volume', but got '%s'", volumeName)
}

func TestGetServiceAccountName(t *testing.T) {
	test := NewTest(t)
	serviceAccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-service-account",
		},
	}
	serviceAccountName := GetServiceAccountName(test, serviceAccount)
	test.Expect(serviceAccountName).To(gomega.Equal("test-service-account"), "Expected service account name 'test-service-account', but got '%s'", serviceAccountName)
}

func TestGetVolumeMountName(t *testing.T) {
	test := NewTest(t)
	volumeMount := corev1.VolumeMount{
		Name: "test-volume-mount",
	}
	volumeMountName := GetVolumeMountName(test, volumeMount)
	test.Expect(volumeMountName).To(gomega.Equal("test-volume-mount"), "Expected volume mount name 'test-volume-mount', but got '%s'", volumeMountName)
}

func TestGetEnvVarName(t *testing.T) {
	test := NewTest(t)
	envVar := corev1.EnvVar{
		Name: "test-env-var",
	}
	envVarName := GetEnvVarName(test, envVar)
	test.Expect(envVarName).To(gomega.Equal("test-env-var"), "Expected env var name 'test-env-var', but got '%s'", envVarName)
}
