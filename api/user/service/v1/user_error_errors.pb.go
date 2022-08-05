// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserInvalidParams(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_INVALID_PARAMS.String() && e.Code == 400
}

func ErrorUserInvalidParams(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserServiceErrorReason_USER_INVALID_PARAMS.String(), fmt.Sprintf(format, args...))
}

func IsUserRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_RECORD_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserRecordNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserServiceErrorReason_USER_RECORD_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserInvalidPass(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_INVALID_PASS.String() && e.Code == 400
}

func ErrorUserInvalidPass(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserServiceErrorReason_USER_INVALID_PASS.String(), fmt.Sprintf(format, args...))
}

func IsUserContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorUserContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserServiceErrorReason_USER_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsUserMakeTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_MAKE_TOKEN_ERROR.String() && e.Code == 500
}

func ErrorUserMakeTokenError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, UserServiceErrorReason_USER_MAKE_TOKEN_ERROR.String(), fmt.Sprintf(format, args...))
}
