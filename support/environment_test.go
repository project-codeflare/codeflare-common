package support

import (
	"fmt"
	"os"
	"testing"

	"github.com/onsi/gomega"
)

func TestGetCodeFlareSDKVersion(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestSdkVersion, "1.2.3")

	// Get the version.
	version := GetCodeFlareSDKVersion()

	// Assert that the version is correct.
	if version != "1.2.3" {
		gomega.Expect(version).To(gomega.Equal("1.2.3"), "Expected version 1.2.3, but got %s", version)

	}
}

func TestGetRayVersion(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestRayVersion, "1.4.5")

	// Get the version.
	version := GetRayVersion()

	// Assert that the version is correct.
	if version != "1.4.5" {
		gomega.Expect(version).To(gomega.Equal("1.2.3"), "Expected version 1.4.5, but got %s", version)
	}
}

func TestGetRayImage(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestRayImage, "ray/ray:latest")

	// Get the image.
	image := GetRayImage()

	// Assert that the image is correct.
	if image != "ray/ray:latest" {
		gomega.Expect(image).To(gomega.Equal("ray/ray:latest"), "Expected image ray/ray:latest, but got %s", image)

	}
}

func TestGetPyTorchImage(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestPyTorchImage, "pytorch/pytorch:latest")

	// Get the image.
	image := GetPyTorchImage()

	// Assert that the image is correct.
	if image != "pytorch/pytorch:latest" {
		gomega.Expect(image).To(gomega.Equal("pytorch/pytorch:latest"), "Expected image pytorch/pytorch:latest, but got %s", image)

	}
}

func TestGetClusterID(t *testing.T) {
	os.Setenv(ClusterID, "my-cluster-id")
	clusterId, ok := GetClusterId()
	if !ok {
		gomega.Expect(ok).To(gomega.BeTrue(), "Expected GetClusterId() to return true, but got false.")
	}
	if clusterId != "my-cluster-id" {
		gomega.Expect(clusterId).To(gomega.Equal("my-cluster-id"), "Expected GetClusterId() to return 'my-cluster-id', but got '%s'.", clusterId)
	}
}

func TestGetInstascaleOcmSecret(t *testing.T) {
	// Set the Instascale OCM secret environment variable.
	os.Setenv(InstaScaleOcmSecret, "default/instascale-ocm-secret")
	// Get the Instascale OCM secret namespace and secret name.
	namespace, secretName := GetInstascaleOcmSecret()

	// Verify that the namespace and secret name are correct.
	if namespace != "default" || secretName != "instascale-ocm-secret" {
		gomega.Expect(fmt.Sprintf("%s/%s", namespace, secretName)).To(
			gomega.Equal("default/instascale-ocm-secret"),
			"Expected GetInstascaleOcmSecret() to return 'default/instascale-ocm-secret', but got '%s/%s'.",
			namespace, secretName,
		)

	}

}

func TestGetClusterType(t *testing.T) {
	tests := []struct {
		name        string
		envVarValue string
		expected    ClusterType
	}{
		{
			name:        "OSD cluster",
			envVarValue: "OSD",
			expected:    OsdCluster,
		},
		{
			name:        "OCP cluster",
			envVarValue: "OCP",
			expected:    OcpCluster,
		},
		{
			name:        "Hypershift cluster",
			envVarValue: "HYPERSHIFT",
			expected:    HypershiftCluster,
		},
		{
			name:        "KIND cluster",
			envVarValue: "KIND",
			expected:    KindCluster,
		},
	}
	ttt := With(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(ClusterTypeEnvVar, tt.envVarValue)
			actual := GetClusterType(ttt) // Pass tt as an argument to GetClusterType
			if actual != tt.expected {
				gomega.Expect(actual).To(
					gomega.Equal(tt.expected),
					"Expected GetClusterType() to return %v, but got %v", tt.expected, actual,
				)

			}
		})
	}
}
