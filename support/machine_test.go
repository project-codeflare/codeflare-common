package support

/*import (
    "testing"
    "k8s.io/apimachinery/pkg/runtime"
    
    "k8s.io/client-go/kubernetes/fake"
    "k8s.io/apimachinery/pkg/apis/meta/v1"
    "log"
	
    machinev1beta1 "github.com/openshift/api/machine/v1beta1"
)

func NewFakeKubeClientMachine(objects ...runtime.Object) *fake.Clientset {
    fakeClient := fake.NewSimpleClientset(objects...)
    return fakeClient
}




func TestGetMachineSets(t *testing.T) {
	t.Parallel()

	// Create a mock machineSet object
	machineSet1 := &machinev1beta1.MachineSet{
		ObjectMeta: v1.ObjectMeta{
			Name: "machine-set1",
		},
	}

	// Create a mock clientset
	clientset := NewFakeKubeClientMachine(machineSet1)

	// Call the GetMachineSets function with the clientset
	machineSets, err := GetMachineSets(clientset)

	if err != nil {
		// Handle the error here
		log.Println("Error:", err)
		return
	}

	// Assert that the returned machineSets slice is equal to the expected slice
	RegisterTestingT(t)
	Expect(machineSets).To(HaveLen(1))
	Expect(machineSets[0].Name).To(Equal("machine-set1"))
}
*/


