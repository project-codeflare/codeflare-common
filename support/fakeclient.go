package support

import (
	"testing"

	fakeray "github.com/ray-project/kuberay/ray-operator/pkg/client/clientset/versioned/fake"

	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeCore "k8s.io/client-go/kubernetes/fake"
	fakekueue "sigs.k8s.io/kueue/client-go/clientset/versioned/fake"

	fakeimage "github.com/openshift/client-go/image/clientset/versioned/fake"
	fakeMachine "github.com/openshift/client-go/machine/clientset/versioned/fake"
	fakeroute "github.com/openshift/client-go/route/clientset/versioned/fake"
)

func NewTest(t *testing.T) *T {
	fakeCoreClient := fakeCore.NewSimpleClientset()
	fakemachineClient := fakeMachine.NewSimpleClientset()
	fakeimageClient := fakeimage.NewSimpleClientset()
	fakerouteClient := fakeroute.NewSimpleClientset()
	fakerayClient := fakeray.NewSimpleClientset()
	fakekueueClient := fakekueue.NewSimpleClientset()
	fakedynamicClient := fakeDynamic.NewSimpleDynamicClient(runtime.NewScheme())

	test := With(t).(*T)
	test.client = &testClient{
		core:    fakeCoreClient,
		machine: fakemachineClient,
		image:   fakeimageClient,
		route:   fakerouteClient,
		ray:     fakerayClient,
		kueue:   fakekueueClient,
		dynamic: fakedynamicClient,
	}
	return test
}
