package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
)

func NewFakeKubeClientWithMachines(scheme *runtime.Scheme, objects ...client.Object) client.Client {
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).Build()
}

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
	fakeClient := NewFakeKubeClientWithMachines(scheme, testmachines...)

	machine := &machinev1beta1.MachineSet{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "test-machineset-1", Namespace: "openshift-machine-api"}, machine)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	// Assertions
	g.Expect(machine.Name).To(gomega.Equal("test-machineset-1"))
	g.Expect(machine.Namespace).To(gomega.Equal("openshift-machine-api"))

}



/*
import (

	"testing"

	"github.com/onsi/gomega"
	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/controller-runtime/pkg/client"

    "k8s.io/client-go/kubernetes"

)

func TestGetMachines(t *testing.T) {
	g := gomega.NewWithT(t)

	// Create a fake client with test data
	scheme := runtime.NewScheme()
	_ = machinev1beta1.AddToScheme(scheme)
	testmachines := []client.Object{
		&machinev1beta1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "machine-1",
				Namespace: "default",
			},
			// ...
		},
		&machinev1beta1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "machine-2",
				Namespace: "default",
			},
			// ...
		},
	}
	fakeClient := fake.NewSimpleClientset(testmachines...)

	// Define the machine set name to use in the test
	machineSetName := "my-machine-set"

	// Create a fake test object with the fake client
	test := fakeTest{t, fakeClient}

	// Call the GetMachines function and assert the result
	result := GetMachines(test, machineSetName)
	g.Expect(result).To(gomega.HaveLen(2))
	g.Expect(result[0].Name).To(gomega.Equal("machine-1"))
	g.Expect(result[1].Name).To(gomega.Equal("machine-2"))
}

// Define a fake test object that implements the Test interface
type fakeTest struct {
	*testing.T
	client kubernetes.Interface
}

// Implement the Test interface for the fakeTest object
func (f fakeTest) Client() kubernetes.Interface {
	return f.client
}
*/
/*

func TestGetMachines(t *testing.T) {
	g := gomega.NewWithT(t)

	scheme := runtime.NewScheme()
	_ = machinev1beta1.AddToScheme(scheme)
	// Create a fake client and add some test data
	testmachines := []client.Object{
		&machinev1beta1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "machine-1",
				Namespace: "default",
			},
			// ...
		},
		&machinev1beta1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "machine-2",
				Namespace: "default",
			},
			// ...
		},
	}
	_ = NewFakeKubeClientWithMachines(scheme, testmachines...)
	// Define the machine set name to use in the test
	machineSetName := "my-machine-set"

	test := With(t).(*T)



	// Call the GetMachines function and assert the result
	result := GetMachines(test, machineSetName)
	g.Expect(result).To(gomega.HaveLen(3))
	g.Expect(result[0].Name).To(gomega.Equal("machine-1"))
	g.Expect(result[1].Name).To(gomega.Equal("machine-2"))
	}

*/
/*
import (
    "testing"
	"github.com/stretchr/testify/assert"
)
func TestGetMachines(t *testing.T) {
	t.Run("GetMachines returns machines for a given machine set", func(t *testing.T) {
		machineSetName := "test-machine-set"
		test := With(t).(*T)
		machines := GetMachines(test, machineSetName)
		assert.Len(t, machines, 1)
		assert.Equal(t, machines[0].Name, "test-machine")
	})
}
*/


