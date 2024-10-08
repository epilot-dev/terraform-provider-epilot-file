// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/shared"
	"net/http"
)

type ListPublicLinksForFileRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ListPublicLinksForFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// ListPublicLinksForFileResponseBody - Public links of a file retrieved successfully
type ListPublicLinksForFileResponseBody struct {
	Results []shared.PublicLink `json:"results,omitempty"`
}

func (o *ListPublicLinksForFileResponseBody) GetResults() []shared.PublicLink {
	if o == nil {
		return nil
	}
	return o.Results
}

type ListPublicLinksForFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Public links of a file retrieved successfully
	Object *ListPublicLinksForFileResponseBody
}

func (o *ListPublicLinksForFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPublicLinksForFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPublicLinksForFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPublicLinksForFileResponse) GetObject() *ListPublicLinksForFileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
