// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package testable_receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostTestReceiversReader is a Reader for the PostTestReceivers structure.
type PostTestReceiversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTestReceiversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTestReceiversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTestReceiversBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTestReceiversRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTestReceiversInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTestReceiversOK creates a PostTestReceiversOK with default headers values
func NewPostTestReceiversOK() *PostTestReceiversOK {
	return &PostTestReceiversOK{}
}

/*
PostTestReceiversOK describes a response with status code 200, with default header values.

Successfully tested all receivers
*/
type PostTestReceiversOK struct {
	Payload *PostTestReceiversOKBody
}

// IsSuccess returns true when this post test receivers o k response has a 2xx status code
func (o *PostTestReceiversOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post test receivers o k response has a 3xx status code
func (o *PostTestReceiversOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post test receivers o k response has a 4xx status code
func (o *PostTestReceiversOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post test receivers o k response has a 5xx status code
func (o *PostTestReceiversOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post test receivers o k response a status code equal to that given
func (o *PostTestReceiversOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostTestReceiversOK) Error() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversOK  %+v", 200, o.Payload)
}

func (o *PostTestReceiversOK) String() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversOK  %+v", 200, o.Payload)
}

func (o *PostTestReceiversOK) GetPayload() *PostTestReceiversOKBody {
	return o.Payload
}

func (o *PostTestReceiversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTestReceiversOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTestReceiversBadRequest creates a PostTestReceiversBadRequest with default headers values
func NewPostTestReceiversBadRequest() *PostTestReceiversBadRequest {
	return &PostTestReceiversBadRequest{}
}

/*
PostTestReceiversBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostTestReceiversBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post test receivers bad request response has a 2xx status code
func (o *PostTestReceiversBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post test receivers bad request response has a 3xx status code
func (o *PostTestReceiversBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post test receivers bad request response has a 4xx status code
func (o *PostTestReceiversBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post test receivers bad request response has a 5xx status code
func (o *PostTestReceiversBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post test receivers bad request response a status code equal to that given
func (o *PostTestReceiversBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTestReceiversBadRequest) Error() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversBadRequest  %+v", 400, o.Payload)
}

func (o *PostTestReceiversBadRequest) String() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversBadRequest  %+v", 400, o.Payload)
}

func (o *PostTestReceiversBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostTestReceiversBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTestReceiversRequestTimeout creates a PostTestReceiversRequestTimeout with default headers values
func NewPostTestReceiversRequestTimeout() *PostTestReceiversRequestTimeout {
	return &PostTestReceiversRequestTimeout{}
}

/*
PostTestReceiversRequestTimeout describes a response with status code 408, with default header values.

Request time out
*/
type PostTestReceiversRequestTimeout struct {
	Payload string
}

// IsSuccess returns true when this post test receivers request timeout response has a 2xx status code
func (o *PostTestReceiversRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post test receivers request timeout response has a 3xx status code
func (o *PostTestReceiversRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post test receivers request timeout response has a 4xx status code
func (o *PostTestReceiversRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post test receivers request timeout response has a 5xx status code
func (o *PostTestReceiversRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post test receivers request timeout response a status code equal to that given
func (o *PostTestReceiversRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTestReceiversRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTestReceiversRequestTimeout) String() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTestReceiversRequestTimeout) GetPayload() string {
	return o.Payload
}

func (o *PostTestReceiversRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTestReceiversInternalServerError creates a PostTestReceiversInternalServerError with default headers values
func NewPostTestReceiversInternalServerError() *PostTestReceiversInternalServerError {
	return &PostTestReceiversInternalServerError{}
}

/*
PostTestReceiversInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostTestReceiversInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post test receivers internal server error response has a 2xx status code
func (o *PostTestReceiversInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post test receivers internal server error response has a 3xx status code
func (o *PostTestReceiversInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post test receivers internal server error response has a 4xx status code
func (o *PostTestReceiversInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post test receivers internal server error response has a 5xx status code
func (o *PostTestReceiversInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post test receivers internal server error response a status code equal to that given
func (o *PostTestReceiversInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTestReceiversInternalServerError) Error() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTestReceiversInternalServerError) String() string {
	return fmt.Sprintf("[POST /receivers/test][%d] postTestReceiversInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTestReceiversInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostTestReceiversInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostTestReceiversOKBody post test receivers o k body
swagger:model PostTestReceiversOKBody
*/
type PostTestReceiversOKBody struct {

	// receiver
	Receiver string `json:"receiver,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this post test receivers o k body
func (o *PostTestReceiversOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post test receivers o k body based on context it is used
func (o *PostTestReceiversOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTestReceiversOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTestReceiversOKBody) UnmarshalBinary(b []byte) error {
	var res PostTestReceiversOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
