package support

import (
    "testing"
	"github.com/onsi/gomega"
    networkingv1 "k8s.io/api/networking/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    
)



func TestGetIngress(t *testing.T) {
	
    tests := []struct {
        namespace string
        name      string
        expected  *networkingv1.Ingress
    }{
        {"my-namespace", "my-ingress", &networkingv1.Ingress{
            ObjectMeta: metav1.ObjectMeta{
                Name:      "my-ingress",
                Namespace: "my-namespace",
            },
            Spec: networkingv1.IngressSpec{
                Rules: []networkingv1.IngressRule{
                    {
                        Host: "my-ingress.example.com",
                        IngressRuleValue: networkingv1.IngressRuleValue{
                            HTTP: &networkingv1.HTTPIngressRuleValue{
                                Paths: []networkingv1.HTTPIngressPath{
                                    {
                                        Path: "/",
                                        Backend: networkingv1.IngressBackend{
                                            Service: &networkingv1.IngressServiceBackend{
                                                Name: "my-service",
                                                Port: networkingv1.ServiceBackendPort{
                                                    Number: 80,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        }},
    }
    tt := With(t)
    for _, test := range tests {
        actual := GetIngress(tt , test.namespace, test.name)
        gomega.Expect(actual).To(gomega.Equal(test.expected))
    }
}


