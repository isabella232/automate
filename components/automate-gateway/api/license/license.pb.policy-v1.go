// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/license/license.proto

package license

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.license.License/ApplyLicense", "license", "apply", "POST", "/license/apply", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ApplyLicenseReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "license":
					return m.License
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.license.License/GetStatus", "license:status", "read", "GET", "/license/status", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.license.License/RequestLicense", "license", "request", "POST", "/license/request", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*RequestLicenseReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "first_name":
					return m.FirstName
				case "last_name":
					return m.LastName
				case "email":
					return m.Email
				default:
					return ""
				}
			})
		}
		return ""
	})
}
