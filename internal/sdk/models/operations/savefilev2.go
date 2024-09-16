// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/shared"
	"net/http"
)

type SaveFileV2Request struct {
	SaveFilePayloadV2 *shared.SaveFilePayloadV2 `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Update the diff and entity for the custom activity included in the query.
	// Pending state on activity is automatically ended when activity is filled.
	//
	FillActivity *bool `default:"false" queryParam:"style=form,explode=true,name=fill_activity"`
	// When passed true, the response will contain only fields that match the schema, with non-matching fields included in `__additional`
	Strict *bool `default:"false" queryParam:"style=form,explode=true,name=strict"`
}

func (s SaveFileV2Request) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SaveFileV2Request) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SaveFileV2Request) GetSaveFilePayloadV2() *shared.SaveFilePayloadV2 {
	if o == nil {
		return nil
	}
	return o.SaveFilePayloadV2
}

func (o *SaveFileV2Request) GetActivityID() *string {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *SaveFileV2Request) GetFillActivity() *bool {
	if o == nil {
		return nil
	}
	return o.FillActivity
}

func (o *SaveFileV2Request) GetStrict() *bool {
	if o == nil {
		return nil
	}
	return o.Strict
}

type SaveFileV2Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Created File Entity
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
