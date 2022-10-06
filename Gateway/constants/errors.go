package constants

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	badRequestErrMsg           = "bad.request"
	badRequestErrItemMsg       = "error.bad.request"
	badRequestErrItemDetailMsg = "The request can't be processed by the server"
	preconditionErrMsg         = "precondition.failed"

	serviceUnavailableErrMsg           = "service.unavailable"
	successfulMsg                      = "Successful"
	serviceUnavailableErrItemMsg       = "error.internal.server"
	badRequestValidatorMsg             = "error.validator"
	serviceUnavailableErrItemDetailMsg = "Internal Server Error. Please, contact customer support"
)

const (
	CodeBadRequest                     = 400
	CodeBadRequestMethodNotAllow       = 405
	CodeBadRequestContentTypeNotAllow  = 415
	CodeBadRequestHeaderAcceptNotAllow = 406
	CodeClientNotAvailable             = 011
)

type Custom_Error interface {
	Error() string
	GetHTTPCode() int
}

type CustomError struct {
	Code         int        `json:"code"`
	Message      string     `json:"message"`
	ResponseTime time.Time  `json:"responseTime"`
	Errors       *ErrorItem `json:"errors"`
	HTTPCode     int        `json:"_"`
}
type ErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (o CustomError) Error() string {
	return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message, o.Errors)
}
func (o CustomError) GetHTTPCode() int {
	return o.HTTPCode
}

func OK(ctx context.Context) *CustomError {
	return &CustomError{
		Code:         200,
		Message:      successfulMsg,
		ResponseTime: time.Now(),
		Errors:       nil,
		HTTPCode:     http.StatusOK,
	}
}
func InternalServerError(ctx context.Context, err error) *CustomError {
	return &CustomError{
		Code:         500,
		Message:      err.Error(),
		ResponseTime: time.Now(),
		Errors: &ErrorItem{
			Code:    500,
			Message: err.Error(),
		},
		HTTPCode: http.StatusServiceUnavailable,
	}
}

func BadRequest(ctx context.Context) *CustomError {
	return &CustomError{
		Code:         CodeBadRequest,
		Message:      badRequestErrMsg,
		ResponseTime: time.Now(),
		Errors: &ErrorItem{
			Code:    400,
			Message: badRequestErrItemMsg,
			Details: badRequestErrItemDetailMsg,
		},
		HTTPCode: http.StatusBadRequest,
	}
}
func BadRequestValue(ctx context.Context) *CustomError {
	return &CustomError{
		Code:         CodeBadRequest,
		Message:      badRequestErrMsg,
		ResponseTime: time.Now(),
		Errors: &ErrorItem{
			Code:    CodeBadRequest,
			Message: badRequestValidatorMsg,
			Details: badRequestErrItemDetailMsg,
		},
	}
}
