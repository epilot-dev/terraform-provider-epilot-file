// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/shared"
	"net/http"
)

type SaveFileRequest struct {
	SaveFilePayload *shared.SaveFilePayload `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for updated entity to become available in Search API. Useful for large migrations
	Async *bool `default:"false" queryParam:"style=form,explode=true,name=async"`
}

func (s SaveFileRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SaveFileRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SaveFileRequest) GetSaveFilePayload() *shared.SaveFilePayload {
	if o == nil {
		return nil
	}
	return o.SaveFilePayload
}

func (o *SaveFileRequest) GetActivityID() *string {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *SaveFileRequest) GetAsync() *bool {
	if o == nil {
		return nil
	}
	return o.Async
}

type SaveFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Created File Entity
	FileEntity *shared.FileEntity
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SaveFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SaveFileResponse) GetFileEntity() *shared.FileEntity {
	if o == nil {
		return nil
	}
	return o.FileEntity
}

func (o *SaveFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SaveFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
