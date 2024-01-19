// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RevokePublicLinkRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RevokePublicLinkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RevokePublicLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Revokes a public link successfully.
	Res *string
}

func (o *RevokePublicLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RevokePublicLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RevokePublicLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RevokePublicLinkResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
