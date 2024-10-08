// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/internal/utils"
)

type SaveCustomFilePayloadAccessControl string

const (
	SaveCustomFilePayloadAccessControlPrivate    SaveCustomFilePayloadAccessControl = "private"
	SaveCustomFilePayloadAccessControlPublicRead SaveCustomFilePayloadAccessControl = "public-read"
)

func (e SaveCustomFilePayloadAccessControl) ToPointer() *SaveCustomFilePayloadAccessControl {
	return &e
}
func (e *SaveCustomFilePayloadAccessControl) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private":
		fallthrough
	case "public-read":
		*e = SaveCustomFilePayloadAccessControl(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveCustomFilePayloadAccessControl: %v", v)
	}
}

type SaveCustomFilePayload struct {
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
	ID                   *string `json:"_id,omitempty"`
	// Manifest ID used to create/update the entity
	Manifest      []string                            `json:"_manifest,omitempty"`
	Purpose       []string                            `json:"_purpose,omitempty"`
	Tags          []string                            `json:"_tags,omitempty"`
	AccessControl *SaveCustomFilePayloadAccessControl `default:"private" json:"access_control"`
	// Custom external download url used for the file
	CustomDownloadURL *string `json:"custom_download_url,omitempty"`
	// Deprecated, use _id instead
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	FileEntityID *string `json:"file_entity_id,omitempty"`
	Filename     *string `json:"filename,omitempty"`
	// MIME type of the file
	MimeType *string `json:"mime_type,omitempty"`
	// List of entities to relate the file to
	Relations []FileRelationItem `json:"relations,omitempty"`
	Type      *FileType          `json:"type,omitempty"`
}

func (s SaveCustomFilePayload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SaveCustomFilePayload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SaveCustomFilePayload) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SaveCustomFilePayload) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SaveCustomFilePayload) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *SaveCustomFilePayload) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *SaveCustomFilePayload) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SaveCustomFilePayload) GetAccessControl() *SaveCustomFilePayloadAccessControl {
	if o == nil {
		return nil
	}
	return o.AccessControl
}

func (o *SaveCustomFilePayload) GetCustomDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.CustomDownloadURL
}

func (o *SaveCustomFilePayload) GetFileEntityID() *string {
	if o == nil {
		return nil
	}
	return o.FileEntityID
}

func (o *SaveCustomFilePayload) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *SaveCustomFilePayload) GetMimeType() *string {
	if o == nil {
		return nil
	}
	return o.MimeType
}

func (o *SaveCustomFilePayload) GetRelations() []FileRelationItem {
	if o == nil {
		return nil
	}
	return o.Relations
}

func (o *SaveCustomFilePayload) GetType() *FileType {
	if o == nil {
		return nil
	}
	return o.Type
}
