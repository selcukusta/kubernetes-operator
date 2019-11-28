// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagement":       schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagement(ref),
		"github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementSpec":   schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagementSpec(ref),
		"github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementStatus": schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagementStatus(ref),
	}
}

func schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NetCoreConfigManagement is the Schema for the netcoreconfigmanagements API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementSpec", "github.com/selcukusta/cm-operator/pkg/apis/selcukusta/v1alpha1.NetCoreConfigManagementStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagementSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NetCoreConfigManagementSpec defines the desired state of NetCoreConfigManagement",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_selcukusta_v1alpha1_NetCoreConfigManagementStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NetCoreConfigManagementStatus defines the observed state of NetCoreConfigManagement",
				Type:        []string{"object"},
			},
		},
	}
}
