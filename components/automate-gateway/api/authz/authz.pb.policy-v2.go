// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/authz/authz.proto

package authz

import (
	request "github.com/chef/automate/components/automate-gateway/api/authz/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/auth/policies/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/CreatePolicy", "iam:policies", "iam:policies:create", "POST", "/auth/policies", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreatePolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "action":
					return m.Action
				case "resource":
					return m.Resource
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/ListPolicies", "iam:policies", "iam:policies:list", "GET", "/auth/policies", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/DeletePolicy", "iam:policies:{id}", "iam:policies:delete", "DELETE", "/auth/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeletePolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/IntrospectAll", "iam:introspect", "iam:introspect:getAll", "GET", "/auth/introspect", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/IntrospectSome", "iam:introspect", "iam:introspect:getSome", "POST", "/auth/introspect_some", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/Introspect", "iam:introspect", "iam:introspect:get", "POST", "/auth/introspect", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.IntrospectReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "path":
					return m.Path
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.authz.Authorization/IntrospectAllProjects", "iam:introspect", "iam:introspect:getAllProjects", "GET", "/auth/introspect_projects", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
