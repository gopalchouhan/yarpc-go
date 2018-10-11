// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpctransport

import (
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcerror"
)

// InboundBadRequestError builds an error which indicates that an inbound
// cannot process a request because it is a bad request.
//
// IsBadRequestError returns true for these errors.
//
// Deprecated: use yarpcerror.Newf with yarpcerror.CodeInvalidArgument instead.
func InboundBadRequestError(err error) error {
	return yarpcerror.Newf(yarpcerror.CodeInvalidArgument, err.Error())
}

// IsBadRequestError returns true if the request could not be processed
// because it was invalid.
//
// Deprecated: use yarpcerror.FromError(err).Code() == yarpcerror.CodeInvalidArgument instead.
func IsBadRequestError(err error) bool {
	return yarpcerror.FromError(err).Code() == yarpcerror.CodeInvalidArgument
}

// IsUnexpectedError returns true if the server panicked or failed to process
// the request with an unhandled error.
//
// Deprecated: use yarpcerror.FromError(err).Code() == yarpcerror.CodeInternal instead.
func IsUnexpectedError(err error) bool {
	return yarpcerror.FromError(err).Code() == yarpcerror.CodeInternal
}

// IsTimeoutError return true if the given error is a TimeoutError.
//
// Deprecated: use yarpcerror.FromError(err).Code() == yarpcerror.CodeDeadlineExceeded instead.
func IsTimeoutError(err error) bool {
	return yarpcerror.FromError(err).Code() == yarpcerror.CodeDeadlineExceeded
}

// UnrecognizedProcedureError returns an error for the given request,
// such that IsUnrecognizedProcedureError can distinguish it from other errors
// coming out of router.Choose.
//
// Deprecated: use yarpcerror.Newf with yarpcerror.CodeUnimplemented instead.
func UnrecognizedProcedureError(req *yarpc.Request) error {
	return yarpcerror.Newf(yarpcerror.CodeUnimplemented, "unrecognized procedure %q for service %q", req.Procedure, req.Service)
}

// IsUnrecognizedProcedureError returns true for errors returned by
// Router.Choose if the router cannot find a handler for the request.
//
// Deprecated: use yarpcerror.FromError(err).Code() == yarpcerror.CodeUnimplemented instead.
func IsUnrecognizedProcedureError(err error) bool {
	return yarpcerror.FromError(err).Code() == yarpcerror.CodeUnimplemented
}