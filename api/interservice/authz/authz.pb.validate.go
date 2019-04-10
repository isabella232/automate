// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interservice/authz/authz.proto

package authz

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on GetVersionReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetVersionReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetVersionReqValidationError is the validation error returned by
// GetVersionReq.Validate if the designated constraints aren't met.
type GetVersionReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVersionReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVersionReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVersionReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVersionReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVersionReqValidationError) ErrorName() string { return "GetVersionReqValidationError" }

// Error satisfies the builtin error interface
func (e GetVersionReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVersionReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVersionReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVersionReqValidationError{}

// Validate checks the field values on IsAuthorizedReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *IsAuthorizedReq) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSubjects()) < 1 {
		return IsAuthorizedReqValidationError{
			field:  "Subjects",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSubjects() {
		_, _ = idx, item

		if !_IsAuthorizedReq_Subjects_Pattern.MatchString(item) {
			return IsAuthorizedReqValidationError{
				field:  fmt.Sprintf("Subjects[%v]", idx),
				reason: "value does not match regex pattern \"^(?:team|user):(?:local|ldap|saml):[^:*]+$|^token:[^:*]+$|^tls:service:[^:*]+:[^:*]+$\"",
			}
		}

	}

	if !_IsAuthorizedReq_Resource_Pattern.MatchString(m.GetResource()) {
		return IsAuthorizedReqValidationError{
			field:  "Resource",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*$\"",
		}
	}

	if !_IsAuthorizedReq_Action_Pattern.MatchString(m.GetAction()) {
		return IsAuthorizedReqValidationError{
			field:  "Action",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*$\"",
		}
	}

	return nil
}

// IsAuthorizedReqValidationError is the validation error returned by
// IsAuthorizedReq.Validate if the designated constraints aren't met.
type IsAuthorizedReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsAuthorizedReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsAuthorizedReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsAuthorizedReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsAuthorizedReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsAuthorizedReqValidationError) ErrorName() string { return "IsAuthorizedReqValidationError" }

// Error satisfies the builtin error interface
func (e IsAuthorizedReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsAuthorizedReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsAuthorizedReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsAuthorizedReqValidationError{}

var _IsAuthorizedReq_Subjects_Pattern = regexp.MustCompile("^(?:team|user):(?:local|ldap|saml):[^:*]+$|^token:[^:*]+$|^tls:service:[^:*]+:[^:*]+$")

var _IsAuthorizedReq_Resource_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*$")

var _IsAuthorizedReq_Action_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*$")

// Validate checks the field values on IsAuthorizedResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *IsAuthorizedResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Authorized

	return nil
}

// IsAuthorizedRespValidationError is the validation error returned by
// IsAuthorizedResp.Validate if the designated constraints aren't met.
type IsAuthorizedRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsAuthorizedRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsAuthorizedRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsAuthorizedRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsAuthorizedRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsAuthorizedRespValidationError) ErrorName() string { return "IsAuthorizedRespValidationError" }

// Error satisfies the builtin error interface
func (e IsAuthorizedRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsAuthorizedResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsAuthorizedRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsAuthorizedRespValidationError{}

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Action

	// no validation rules for Id

	// no validation rules for Resource

	// no validation rules for Effect

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyValidationError) ErrorName() string { return "PolicyValidationError" }

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyValidationError{}

// Validate checks the field values on CreatePolicyReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreatePolicyReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_CreatePolicyReq_Action_Pattern.MatchString(m.GetAction()) {
		return CreatePolicyReqValidationError{
			field:  "Action",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*(?::[*])?$|^[*]$\"",
		}
	}

	if len(m.GetSubjects()) < 1 {
		return CreatePolicyReqValidationError{
			field:  "Subjects",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSubjects() {
		_, _ = idx, item

		if !_CreatePolicyReq_Subjects_Pattern.MatchString(item) {
			return CreatePolicyReqValidationError{
				field:  fmt.Sprintf("Subjects[%v]", idx),
				reason: "value does not match regex pattern \"^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$\"",
			}
		}

	}

	if !_CreatePolicyReq_Resource_Pattern.MatchString(m.GetResource()) {
		return CreatePolicyReqValidationError{
			field:  "Resource",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*(?::[*])?$|^[*]$\"",
		}
	}

	return nil
}

// CreatePolicyReqValidationError is the validation error returned by
// CreatePolicyReq.Validate if the designated constraints aren't met.
type CreatePolicyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePolicyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePolicyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePolicyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePolicyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePolicyReqValidationError) ErrorName() string { return "CreatePolicyReqValidationError" }

// Error satisfies the builtin error interface
func (e CreatePolicyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePolicyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePolicyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePolicyReqValidationError{}

var _CreatePolicyReq_Action_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*(?::[*])?$|^[*]$")

var _CreatePolicyReq_Subjects_Pattern = regexp.MustCompile("^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$")

var _CreatePolicyReq_Resource_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*(?::[*])?$|^[*]$")

// Validate checks the field values on CreatePolicyResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreatePolicyResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePolicyRespValidationError{
				field:  "Policy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreatePolicyRespValidationError is the validation error returned by
// CreatePolicyResp.Validate if the designated constraints aren't met.
type CreatePolicyRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePolicyRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePolicyRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePolicyRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePolicyRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePolicyRespValidationError) ErrorName() string { return "CreatePolicyRespValidationError" }

// Error satisfies the builtin error interface
func (e CreatePolicyRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePolicyResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePolicyRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePolicyRespValidationError{}

// Validate checks the field values on ListPoliciesReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListPoliciesReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListPoliciesReqValidationError is the validation error returned by
// ListPoliciesReq.Validate if the designated constraints aren't met.
type ListPoliciesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPoliciesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPoliciesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPoliciesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPoliciesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPoliciesReqValidationError) ErrorName() string { return "ListPoliciesReqValidationError" }

// Error satisfies the builtin error interface
func (e ListPoliciesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPoliciesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPoliciesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPoliciesReqValidationError{}

// Validate checks the field values on ListPoliciesResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListPoliciesResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPolicies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPoliciesRespValidationError{
					field:  fmt.Sprintf("Policies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListPoliciesRespValidationError is the validation error returned by
// ListPoliciesResp.Validate if the designated constraints aren't met.
type ListPoliciesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPoliciesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPoliciesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPoliciesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPoliciesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPoliciesRespValidationError) ErrorName() string { return "ListPoliciesRespValidationError" }

// Error satisfies the builtin error interface
func (e ListPoliciesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPoliciesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPoliciesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPoliciesRespValidationError{}

// Validate checks the field values on DeletePolicyReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeletePolicyReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_DeletePolicyReq_Id_Pattern.MatchString(m.GetId()) {
		return DeletePolicyReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$\"",
		}
	}

	return nil
}

// DeletePolicyReqValidationError is the validation error returned by
// DeletePolicyReq.Validate if the designated constraints aren't met.
type DeletePolicyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePolicyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePolicyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePolicyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePolicyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePolicyReqValidationError) ErrorName() string { return "DeletePolicyReqValidationError" }

// Error satisfies the builtin error interface
func (e DeletePolicyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePolicyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePolicyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePolicyReqValidationError{}

var _DeletePolicyReq_Id_Pattern = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

// Validate checks the field values on DeletePolicyResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeletePolicyResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeletePolicyRespValidationError{
				field:  "Policy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeletePolicyRespValidationError is the validation error returned by
// DeletePolicyResp.Validate if the designated constraints aren't met.
type DeletePolicyRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePolicyRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePolicyRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePolicyRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePolicyRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePolicyRespValidationError) ErrorName() string { return "DeletePolicyRespValidationError" }

// Error satisfies the builtin error interface
func (e DeletePolicyRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePolicyResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePolicyRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePolicyRespValidationError{}

// Validate checks the field values on PurgeSubjectFromPoliciesReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PurgeSubjectFromPoliciesReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_PurgeSubjectFromPoliciesReq_Subject_Pattern.MatchString(m.GetSubject()) {
		return PurgeSubjectFromPoliciesReqValidationError{
			field:  "Subject",
			reason: "value does not match regex pattern \"^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$\"",
		}
	}

	return nil
}

// PurgeSubjectFromPoliciesReqValidationError is the validation error returned
// by PurgeSubjectFromPoliciesReq.Validate if the designated constraints
// aren't met.
type PurgeSubjectFromPoliciesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurgeSubjectFromPoliciesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurgeSubjectFromPoliciesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurgeSubjectFromPoliciesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurgeSubjectFromPoliciesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurgeSubjectFromPoliciesReqValidationError) ErrorName() string {
	return "PurgeSubjectFromPoliciesReqValidationError"
}

// Error satisfies the builtin error interface
func (e PurgeSubjectFromPoliciesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurgeSubjectFromPoliciesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurgeSubjectFromPoliciesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurgeSubjectFromPoliciesReqValidationError{}

var _PurgeSubjectFromPoliciesReq_Subject_Pattern = regexp.MustCompile("^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$")

// Validate checks the field values on PurgeSubjectFromPoliciesResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PurgeSubjectFromPoliciesResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PurgeSubjectFromPoliciesRespValidationError is the validation error returned
// by PurgeSubjectFromPoliciesResp.Validate if the designated constraints
// aren't met.
type PurgeSubjectFromPoliciesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurgeSubjectFromPoliciesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurgeSubjectFromPoliciesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurgeSubjectFromPoliciesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurgeSubjectFromPoliciesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurgeSubjectFromPoliciesRespValidationError) ErrorName() string {
	return "PurgeSubjectFromPoliciesRespValidationError"
}

// Error satisfies the builtin error interface
func (e PurgeSubjectFromPoliciesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurgeSubjectFromPoliciesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurgeSubjectFromPoliciesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurgeSubjectFromPoliciesRespValidationError{}

// Validate checks the field values on FilterAuthorizedPairsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FilterAuthorizedPairsReq) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSubjects()) < 1 {
		return FilterAuthorizedPairsReqValidationError{
			field:  "Subjects",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSubjects() {
		_, _ = idx, item

		if !_FilterAuthorizedPairsReq_Subjects_Pattern.MatchString(item) {
			return FilterAuthorizedPairsReqValidationError{
				field:  fmt.Sprintf("Subjects[%v]", idx),
				reason: "value does not match regex pattern \"^(?:(?:team|user):(?:local|ldap|saml)|token|tls:service:[^:*]+):[^:*]+$\"",
			}
		}

	}

	for idx, item := range m.GetPairs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterAuthorizedPairsReqValidationError{
					field:  fmt.Sprintf("Pairs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FilterAuthorizedPairsReqValidationError is the validation error returned by
// FilterAuthorizedPairsReq.Validate if the designated constraints aren't met.
type FilterAuthorizedPairsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterAuthorizedPairsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterAuthorizedPairsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterAuthorizedPairsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterAuthorizedPairsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterAuthorizedPairsReqValidationError) ErrorName() string {
	return "FilterAuthorizedPairsReqValidationError"
}

// Error satisfies the builtin error interface
func (e FilterAuthorizedPairsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterAuthorizedPairsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterAuthorizedPairsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterAuthorizedPairsReqValidationError{}

var _FilterAuthorizedPairsReq_Subjects_Pattern = regexp.MustCompile("^(?:(?:team|user):(?:local|ldap|saml)|token|tls:service:[^:*]+):[^:*]+$")

// Validate checks the field values on FilterAuthorizedPairsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FilterAuthorizedPairsResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPairs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterAuthorizedPairsRespValidationError{
					field:  fmt.Sprintf("Pairs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FilterAuthorizedPairsRespValidationError is the validation error returned by
// FilterAuthorizedPairsResp.Validate if the designated constraints aren't met.
type FilterAuthorizedPairsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterAuthorizedPairsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterAuthorizedPairsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterAuthorizedPairsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterAuthorizedPairsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterAuthorizedPairsRespValidationError) ErrorName() string {
	return "FilterAuthorizedPairsRespValidationError"
}

// Error satisfies the builtin error interface
func (e FilterAuthorizedPairsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterAuthorizedPairsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterAuthorizedPairsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterAuthorizedPairsRespValidationError{}

// Validate checks the field values on FilterAuthorizedProjectsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FilterAuthorizedProjectsResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FilterAuthorizedProjectsRespValidationError is the validation error returned
// by FilterAuthorizedProjectsResp.Validate if the designated constraints
// aren't met.
type FilterAuthorizedProjectsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterAuthorizedProjectsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterAuthorizedProjectsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterAuthorizedProjectsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterAuthorizedProjectsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterAuthorizedProjectsRespValidationError) ErrorName() string {
	return "FilterAuthorizedProjectsRespValidationError"
}

// Error satisfies the builtin error interface
func (e FilterAuthorizedProjectsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterAuthorizedProjectsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterAuthorizedProjectsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterAuthorizedProjectsRespValidationError{}

// Validate checks the field values on Pair with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Pair) Validate() error {
	if m == nil {
		return nil
	}

	if !_Pair_Resource_Pattern.MatchString(m.GetResource()) {
		return PairValidationError{
			field:  "Resource",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*$\"",
		}
	}

	if !_Pair_Action_Pattern.MatchString(m.GetAction()) {
		return PairValidationError{
			field:  "Action",
			reason: "value does not match regex pattern \"^[a-z][^:*]*(?::[^:*]+)*$\"",
		}
	}

	return nil
}

// PairValidationError is the validation error returned by Pair.Validate if the
// designated constraints aren't met.
type PairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PairValidationError) ErrorName() string { return "PairValidationError" }

// Error satisfies the builtin error interface
func (e PairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PairValidationError{}

var _Pair_Resource_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*$")

var _Pair_Action_Pattern = regexp.MustCompile("^[a-z][^:*]*(?::[^:*]+)*$")
