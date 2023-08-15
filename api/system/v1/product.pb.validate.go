// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/system/v1/product.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRequestMultiError, or nil if none found.
func (m *CreateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 50 {
		err := CreateProductRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	if len(errors) > 0 {
		return CreateProductRequestMultiError(errors)
	}

	return nil
}

// CreateProductRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRequestMultiError) AllErrors() []error { return m }

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on CreateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductReplyMultiError, or nil if none found.
func (m *CreateProductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateProductReplyMultiError(errors)
	}

	return nil
}

// CreateProductReplyMultiError is an error wrapping multiple validation errors
// returned by CreateProductReply.ValidateAll() if the designated constraints
// aren't met.
type CreateProductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductReplyMultiError) AllErrors() []error { return m }

// CreateProductReplyValidationError is the validation error returned by
// CreateProductReply.Validate if the designated constraints aren't met.
type CreateProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductReplyValidationError) ErrorName() string {
	return "CreateProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductReplyValidationError{}

// Validate checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProductRequestMultiError, or nil if none found.
func (m *UpdateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateProductRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 50 {
		err := UpdateProductRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	if len(errors) > 0 {
		return UpdateProductRequestMultiError(errors)
	}

	return nil
}

// UpdateProductRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProductRequestMultiError) AllErrors() []error { return m }

// UpdateProductRequestValidationError is the validation error returned by
// UpdateProductRequest.Validate if the designated constraints aren't met.
type UpdateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductRequestValidationError) ErrorName() string {
	return "UpdateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductRequestValidationError{}

// Validate checks the field values on UpdateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProductReplyMultiError, or nil if none found.
func (m *UpdateProductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateProductReplyMultiError(errors)
	}

	return nil
}

// UpdateProductReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateProductReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateProductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProductReplyMultiError) AllErrors() []error { return m }

// UpdateProductReplyValidationError is the validation error returned by
// UpdateProductReply.Validate if the designated constraints aren't met.
type UpdateProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductReplyValidationError) ErrorName() string {
	return "UpdateProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductReplyValidationError{}

// Validate checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductRequestMultiError, or nil if none found.
func (m *DeleteProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProductRequestMultiError(errors)
	}

	return nil
}

// DeleteProductRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProductRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductRequestMultiError) AllErrors() []error { return m }

// DeleteProductRequestValidationError is the validation error returned by
// DeleteProductRequest.Validate if the designated constraints aren't met.
type DeleteProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductRequestValidationError) ErrorName() string {
	return "DeleteProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductRequestValidationError{}

// Validate checks the field values on DeleteProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductReplyMultiError, or nil if none found.
func (m *DeleteProductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteProductReplyMultiError(errors)
	}

	return nil
}

// DeleteProductReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteProductReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteProductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductReplyMultiError) AllErrors() []error { return m }

// DeleteProductReplyValidationError is the validation error returned by
// DeleteProductReply.Validate if the designated constraints aren't met.
type DeleteProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductReplyValidationError) ErrorName() string {
	return "DeleteProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductReplyValidationError{}

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductRequestMultiError, or nil if none found.
func (m *GetProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProductRequestMultiError(errors)
	}

	return nil
}

// GetProductRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductRequestMultiError) AllErrors() []error { return m }

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on GetProductReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductReplyMultiError, or nil if none found.
func (m *GetProductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProductReplyValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProductReplyMultiError(errors)
	}

	return nil
}

// GetProductReplyMultiError is an error wrapping multiple validation errors
// returned by GetProductReply.ValidateAll() if the designated constraints
// aren't met.
type GetProductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductReplyMultiError) AllErrors() []error { return m }

// GetProductReplyValidationError is the validation error returned by
// GetProductReply.Validate if the designated constraints aren't met.
type GetProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductReplyValidationError) ErrorName() string { return "GetProductReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductReplyValidationError{}

// Validate checks the field values on ListProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductRequestMultiError, or nil if none found.
func (m *ListProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Parent

	// no validation rules for PageSize

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListProductRequestMultiError(errors)
	}

	return nil
}

// ListProductRequestMultiError is an error wrapping multiple validation errors
// returned by ListProductRequest.ValidateAll() if the designated constraints
// aren't met.
type ListProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductRequestMultiError) AllErrors() []error { return m }

// ListProductRequestValidationError is the validation error returned by
// ListProductRequest.Validate if the designated constraints aren't met.
type ListProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductRequestValidationError) ErrorName() string {
	return "ListProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductRequestValidationError{}

// Validate checks the field values on ListProductReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListProductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductReplyMultiError, or nil if none found.
func (m *ListProductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProductReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProductReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProductReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListProductReplyMultiError(errors)
	}

	return nil
}

// ListProductReplyMultiError is an error wrapping multiple validation errors
// returned by ListProductReply.ValidateAll() if the designated constraints
// aren't met.
type ListProductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductReplyMultiError) AllErrors() []error { return m }

// ListProductReplyValidationError is the validation error returned by
// ListProductReply.Validate if the designated constraints aren't met.
type ListProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductReplyValidationError) ErrorName() string { return "ListProductReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductReplyValidationError{}
