/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"testing"

	"github.com/onsi/gomega"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateUserRoleBinding(t *testing.T) {

	test := NewTest(t)

	role := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			APIVersion: rbacv1.SchemeGroupVersion.String(),
			Kind:       "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "role1",
			Namespace: "ns-1",
		},
	}

	rb := CreateUserRoleBinding(test, "ns-1", "user-1", role)

	test.Expect(rb).To(gomega.Not(gomega.BeNil()))
	test.Expect(rb.GenerateName).To(gomega.Equal("rb-"))
	test.Expect(rb.Namespace).To(gomega.Equal("ns-1"))

	test.Expect(rb.RoleRef.APIGroup).To(gomega.Equal(rbacv1.SchemeGroupVersion.Group))
	test.Expect(rb.RoleRef.Kind).To(gomega.Equal("Role"))
	test.Expect(rb.RoleRef.Name).To(gomega.Equal("role1"))

	test.Expect(rb.Subjects[0].APIGroup).To(gomega.Equal(rbacv1.SchemeGroupVersion.Group))
	test.Expect(rb.Subjects[0].Kind).To(gomega.Equal("User"))
	test.Expect(rb.Subjects[0].Name).To(gomega.Equal("user-1"))
}

func TestCreateUserClusterRoleBinding(t *testing.T) {

	test := NewTest(t)

	crole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: rbacv1.SchemeGroupVersion.String(),
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "role1",
		},
	}

	rb := CreateUserClusterRoleBinding(test, "user-1", crole)

	test.Expect(rb).To(gomega.Not(gomega.BeNil()))
	test.Expect(rb.GenerateName).To(gomega.Equal("crb-"))

	test.Expect(rb.RoleRef.APIGroup).To(gomega.Equal(rbacv1.SchemeGroupVersion.Group))
	test.Expect(rb.RoleRef.Kind).To(gomega.Equal("ClusterRole"))
	test.Expect(rb.RoleRef.Name).To(gomega.Equal("role1"))

	test.Expect(rb.Subjects[0].APIGroup).To(gomega.Equal(rbacv1.SchemeGroupVersion.Group))
	test.Expect(rb.Subjects[0].Kind).To(gomega.Equal("User"))
	test.Expect(rb.Subjects[0].Name).To(gomega.Equal("user-1"))
}
