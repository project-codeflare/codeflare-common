package support

import (
	"testing"

	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
)

func TestGetMachineSets(t *testing.T) {
	test := NewTest(t)

	machine := &machinev1beta1.MachineSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machineset-1",
			Namespace: "openshift-machine-api",
		},
	}

	test.client.Machine().MachineV1beta1().MachineSets("openshift-machine-api").Create(test.ctx, machine, metav1.CreateOptions{})

	machines, _ := GetMachineSets(test)

	test.Expect(machines).To(gomega.HaveLen(1))
	test.Expect(machines[0].Name).To(gomega.Equal("test-machineset-1"))
	test.Expect(machines[0].Namespace).To(gomega.Equal("openshift-machine-api"))

}

func TestMachineSet(t *testing.T) {
	test := NewTest(t)

	machine := &machinev1beta1.MachineSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machineset-1",
			Namespace: "openshift-machine-api",
		},
	}

	test.client.Machine().MachineV1beta1().MachineSets("openshift-machine-api").Create(test.ctx, machine, metav1.CreateOptions{})

	machineSet := MachineSet(test, "openshift-machine-api", "test-machineset-1")
	test.Expect(machineSet(test).Name).To(gomega.Equal("test-machineset-1"))

}
