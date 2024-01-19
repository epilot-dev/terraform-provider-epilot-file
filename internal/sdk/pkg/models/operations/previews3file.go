// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/pkg/models/shared"
	"net/http"
)

type PreviewS3FileRequest struct {
	S3Reference *shared.S3Reference `request:"mediaType=application/json"`
	// height
	H *int64 `queryParam:"style=form,explode=true,name=h"`
	// width
	W *int64 `queryParam:"style=form,explode=true,name=w"`
}

func (o *PreviewS3FileRequest) GetS3Reference() *shared.S3Reference {
	if o == nil {
		return nil
	}
	return o.S3Reference
}

func (o *PreviewS3FileRequest) GetH() *int64 {
	if o == nil {
		return nil
	}
	return o.H
}

func (o *PreviewS3FileRequest) GetW() *int64 {
	if o == nil {
		return nil
	}
	return o.W
}

type PreviewS3FileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PreviewS3FileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PreviewS3FileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PreviewS3FileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
