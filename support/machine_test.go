package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
)

func TestGetMachineSets(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = machinev1beta1.AddToScheme(scheme)

	testmachines := []client.Object{
		&machinev1beta1.MachineSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-machineset-1",
				Namespace: "openshift-machine-api",
			},
		},
	}
	fakeClient := NewFakeKubeClientWithScheme(scheme, testmachines...)

	machine := &machinev1beta1.MachineSet{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "test-machineset-1", Namespace: "openshift-machine-api"}, machine)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	// Assertions
	g.Expect(machine.Name).To(gomega.Equal("test-machineset-1"))
	g.Expect(machine.Namespace).To(gomega.Equal("openshift-machine-api"))

}
