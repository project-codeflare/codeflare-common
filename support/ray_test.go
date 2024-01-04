package support

import (
	"testing"

	"github.com/onsi/gomega"
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetRayJob(t *testing.T) {

	test := NewTest(t)

	RayJob := &rayv1alpha1.RayJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-job-1",
			Namespace: "my-namespace",
		},
	}

	test.client.Ray().RayV1().RayJobs("my-namespace").Create(test.ctx, RayJob, metav1.CreateOptions{})

	rayJob := GetRayJob(test, "my-namespace", "my-job-1")
	test.Expect(rayJob.Name).To(gomega.Equal("my-job-1"))
	test.Expect(rayJob.Namespace).To(gomega.Equal("my-namespace"))
}

func TestGetRayCluster(t *testing.T) {

	test := NewTest(t)

	RayCluster := &rayv1alpha1.RayCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-cluster-1",
			Namespace: "my-namespace",
		},
	}

	test.client.Ray().RayV1().RayClusters("my-namespace").Create(test.ctx, RayCluster, metav1.CreateOptions{})
	raycluster := GetRayCluster(test, "my-namespace", "my-cluster-1")

	test.Expect(raycluster.Name).To(gomega.Equal("my-cluster-1"))
	test.Expect(raycluster.Namespace).To(gomega.Equal("my-namespace"))
}
