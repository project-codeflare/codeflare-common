package support

import (
	"os"
	"testing"
)

func TestGetCodeFlareSDKVersion(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestSdkVersion, "1.2.3")

	// Get the version.
	version := GetCodeFlareSDKVersion()

	// Assert that the version is correct.
	if version != "1.2.3" {
		t.Errorf("Expected version 1.2.3, but got %s", version)
	}
}

func TestGetRayVersion(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestRayVersion, "1.4.5")

	// Get the version.
	version := GetRayVersion()

	// Assert that the version is correct.
	if version != "1.4.5" {
		t.Errorf("Expected version 1.4.5, but got %s", version)
	}
}

func TestGetRayImage(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestRayImage, "ray/ray:latest")

	// Get the image.
	image := GetRayImage()

	// Assert that the image is correct.
	if image != "ray/ray:latest" {
		t.Errorf("Expected image ray/ray:latest, but got %s", image)
	}
}

func TestGetPyTorchImage(t *testing.T) {
	// Set the environment variable.
	os.Setenv(CodeFlareTestPyTorchImage, "pytorch/pytorch:latest")

	// Get the image.
	image := GetPyTorchImage()

	// Assert that the image is correct.
	if image != "pytorch/pytorch:latest" {
		t.Errorf("Expected image pytorch/pytorch:latest, but got %s", image)
	}
}

func TestGetOsdClusterID(t *testing.T) {
	os.Setenv(OsdClusterID, "my-cluster-id")
	clusterId, ok := GetOsdClusterId()
	if !ok {
		t.Errorf("Expected GetOsdClusterId() to return true, but got false.")
	}
	if clusterId != "my-cluster-id" {
		t.Errorf("Expected GetOsdClusterId() to return 'my-cluster-id', but got '%s'.", clusterId)
	}

}
func TestGetInstascaleOcmSecret(t *testing.T) {
	// Set the Instascale OCM secret environment variable.
	os.Setenv(InstaScaleOcmSecret, "default/instascale-ocm-secret")
	// Get the Instascale OCM secret namespace and secret name.
	namespace, secretName := GetInstascaleOcmSecret()

	// Verify that the namespace and secret name are correct.
	if namespace != "default" || secretName != "instascale-ocm-secret" {
		t.Errorf("Expected GetInstascaleOcmSecret() to return 'default/instascale-ocm-secret', but got '%s/%s'.", namespace, secretName)
	}

}
