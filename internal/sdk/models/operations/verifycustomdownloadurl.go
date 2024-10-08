// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

// VerifyCustomDownloadURLResponseBody - Download Url matches signature and has not expired
type VerifyCustomDownloadURLResponseBody struct {
	Valid *bool `json:"valid,omitempty"`
}

func (o *VerifyCustomDownloadURLResponseBody) GetValid() *bool {
	if o == nil {
		return nil
	}
	return o.Valid
}

type VerifyCustomDownloadURLResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Download Url matches signature and has not expired
	Object *VerifyCustomDownloadURLResponseBody
}

func (o *VerifyCustomDownloadURLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VerifyCustomDownloadURLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VerifyCustomDownloadURLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VerifyCustomDownloadURLResponse) GetObject() *VerifyCustomDownloadURLResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
