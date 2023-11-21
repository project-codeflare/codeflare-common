package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func NewFakeClient(scheme *runtime.Scheme, objects ...client.Object) client.Client {
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).Build()
}

func TestGetRayJob(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = rayv1alpha1.AddToScheme(scheme)

	fakeRayJobs := []client.Object{
		&rayv1alpha1.RayJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-job-1",
				Namespace: "my-namespace",
			},
		},
	}

	fakeClient := NewFakeClient(scheme, fakeRayJobs...)

	rayJob := &rayv1alpha1.RayJob{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "my-job-1", Namespace: "my-namespace"}, rayJob)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	//fmt.Printf("Retrieved job object: %+v\n", rayJob)

	g.Expect(rayJob.Name).To(gomega.Equal("my-job-1"))
	g.Expect(rayJob.Namespace).To(gomega.Equal("my-namespace"))
}
