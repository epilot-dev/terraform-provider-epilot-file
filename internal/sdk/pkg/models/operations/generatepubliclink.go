// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GeneratePublicLinkRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GeneratePublicLinkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GeneratePublicLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the public link which can be used to access the file later
	Res *string
}

func (o *GeneratePublicLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GeneratePublicLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GeneratePublicLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GeneratePublicLinkResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}