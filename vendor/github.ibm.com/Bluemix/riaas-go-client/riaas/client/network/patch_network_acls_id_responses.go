// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchNetworkAclsIDReader is a Reader for the PatchNetworkAclsID structure.
type PatchNetworkAclsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchNetworkAclsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchNetworkAclsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchNetworkAclsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchNetworkAclsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchNetworkAclsIDOK creates a PatchNetworkAclsIDOK with default headers values
func NewPatchNetworkAclsIDOK() *PatchNetworkAclsIDOK {
	return &PatchNetworkAclsIDOK{}
}

/*PatchNetworkAclsIDOK handles this case with default header values.

dummy
*/
type PatchNetworkAclsIDOK struct {
	Payload *models.NetworkACL
}

func (o *PatchNetworkAclsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /network_acls/{id}][%d] patchNetworkAclsIdOK  %+v", 200, o.Payload)
}

func (o *PatchNetworkAclsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNetworkAclsIDBadRequest creates a PatchNetworkAclsIDBadRequest with default headers values
func NewPatchNetworkAclsIDBadRequest() *PatchNetworkAclsIDBadRequest {
	return &PatchNetworkAclsIDBadRequest{}
}

/*PatchNetworkAclsIDBadRequest handles this case with default header values.

error
*/
type PatchNetworkAclsIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchNetworkAclsIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /network_acls/{id}][%d] patchNetworkAclsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchNetworkAclsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNetworkAclsIDInternalServerError creates a PatchNetworkAclsIDInternalServerError with default headers values
func NewPatchNetworkAclsIDInternalServerError() *PatchNetworkAclsIDInternalServerError {
	return &PatchNetworkAclsIDInternalServerError{}
}

/*PatchNetworkAclsIDInternalServerError handles this case with default header values.

error
*/
type PatchNetworkAclsIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PatchNetworkAclsIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /network_acls/{id}][%d] patchNetworkAclsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchNetworkAclsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchNetworkAclsIDBody patch network acls ID body
swagger:model PatchNetworkAclsIDBody
*/
type PatchNetworkAclsIDBody struct {

	// The user-defined name for this acl
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this patch network acls ID body
func (o *PatchNetworkAclsIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchNetworkAclsIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchNetworkAclsIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchNetworkAclsIDBody) UnmarshalBinary(b []byte) error {
	var res PatchNetworkAclsIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
