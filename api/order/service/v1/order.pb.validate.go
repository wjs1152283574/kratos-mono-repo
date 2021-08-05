// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/order/service/v1/order.proto

package v1

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
)

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
					field:  fmt.Sprintf("ProList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Uid

	return nil
}

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateOrderReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderNo

	return nil
}

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on DeleteOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteOrderRequestValidationError is the validation error returned by
// DeleteOrderRequest.Validate if the designated constraints aren't met.
type DeleteOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrderRequestValidationError) ErrorName() string {
	return "DeleteOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrderRequestValidationError{}

// Validate checks the field values on DeleteOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteOrderReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteOrderReplyValidationError is the validation error returned by
// DeleteOrderReply.Validate if the designated constraints aren't met.
type DeleteOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrderReplyValidationError) ErrorName() string { return "DeleteOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrderReplyValidationError{}

// Validate checks the field values on GetOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetOrderRequestValidationError is the validation error returned by
// GetOrderRequest.Validate if the designated constraints aren't met.
type GetOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderRequestValidationError) ErrorName() string { return "GetOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderRequestValidationError{}

// Validate checks the field values on GetOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetOrderReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetOrderReplyValidationError{
					field:  fmt.Sprintf("ProList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Oid

	return nil
}

// GetOrderReplyValidationError is the validation error returned by
// GetOrderReply.Validate if the designated constraints aren't met.
type GetOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderReplyValidationError) ErrorName() string { return "GetOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderReplyValidationError{}

// Validate checks the field values on ListOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Page

	// no validation rules for Limit

	// no validation rules for Uid

	return nil
}

// ListOrderRequestValidationError is the validation error returned by
// ListOrderRequest.Validate if the designated constraints aren't met.
type ListOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderRequestValidationError) ErrorName() string { return "ListOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderRequestValidationError{}

// Validate checks the field values on ListOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListOrderReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOrderReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListOrderReplyValidationError is the validation error returned by
// ListOrderReply.Validate if the designated constraints aren't met.
type ListOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderReplyValidationError) ErrorName() string { return "ListOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderReplyValidationError{}

// Validate checks the field values on OrderList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrderList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderListValidationError{
					field:  fmt.Sprintf("ProList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for OrderNo

	// no validation rules for Status

	return nil
}

// OrderListValidationError is the validation error returned by
// OrderList.Validate if the designated constraints aren't met.
type OrderListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderListValidationError) ErrorName() string { return "OrderListValidationError" }

// Error satisfies the builtin error interface
func (e OrderListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderListValidationError{}

// Validate checks the field values on CreateOrderRequest_ProList with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateOrderRequest_ProList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Price

	return nil
}

// CreateOrderRequest_ProListValidationError is the validation error returned
// by CreateOrderRequest_ProList.Validate if the designated constraints aren't met.
type CreateOrderRequest_ProListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequest_ProListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequest_ProListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequest_ProListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequest_ProListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequest_ProListValidationError) ErrorName() string {
	return "CreateOrderRequest_ProListValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequest_ProListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest_ProList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequest_ProListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequest_ProListValidationError{}

// Validate checks the field values on GetOrderReply_ProList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetOrderReply_ProList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Price

	return nil
}

// GetOrderReply_ProListValidationError is the validation error returned by
// GetOrderReply_ProList.Validate if the designated constraints aren't met.
type GetOrderReply_ProListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderReply_ProListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderReply_ProListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderReply_ProListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderReply_ProListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderReply_ProListValidationError) ErrorName() string {
	return "GetOrderReply_ProListValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderReply_ProListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderReply_ProList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderReply_ProListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderReply_ProListValidationError{}

// Validate checks the field values on OrderList_ProList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OrderList_ProList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Banner

	// no validation rules for Price

	return nil
}

// OrderList_ProListValidationError is the validation error returned by
// OrderList_ProList.Validate if the designated constraints aren't met.
type OrderList_ProListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderList_ProListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderList_ProListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderList_ProListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderList_ProListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderList_ProListValidationError) ErrorName() string {
	return "OrderList_ProListValidationError"
}

// Error satisfies the builtin error interface
func (e OrderList_ProListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderList_ProList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderList_ProListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderList_ProListValidationError{}
