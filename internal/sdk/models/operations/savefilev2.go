// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/shared"
	"net/http"
)

type SaveFileV2Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Updated File Entity
	FileEntity *shared.FileEntity
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SaveFileV2Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SaveFileV2Response) GetFileEntity() *shared.FileEntity {
	if o == nil {
		return nil
	}
	return o.FileEntity
}

func (o *SaveFileV2Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SaveFileV2Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
