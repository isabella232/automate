// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/legacy/legacy.proto

package legacy

import policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"

func init() {
	policyv2.MapMethodTo("/chef.automate.api.legacy.LegacyDataCollector/Status", "infra:status", "infra:ingest:get", "GET", "/events/data-collector", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
