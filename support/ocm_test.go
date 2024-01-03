package support

/*
import (
	"testing"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	ocmsdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func TestGetMachinePools(t *testing.T) {
	// Create a fake OpenShift client for testing
	fakeClient := fake.NewSimpleClientset(machinePool{
		MachinePool: &cmv1.MachinePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-machinepool",
				Namespace: "test-namespace",
			},
		},
	})

	test := With(t).(*T)
	test.client = &testClient{
		core: fakeClient,
	}

	// Call the GetMachinePools function with the fake client and connection
	machinePools := GetMachinePools(test, &ocmsdk.Connection{})
	test.Expect(machinePools).Should(gomega.HaveLen(1), "Expected 1 machine pool, but got %d", len(machinePools))
}
*/
