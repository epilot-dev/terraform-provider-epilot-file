// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/pkg/utils"
)

type SaveFilePayloadV2AccessControl string

const (
	SaveFilePayloadV2AccessControlPrivate    SaveFilePayloadV2AccessControl = "private"
	SaveFilePayloadV2AccessControlPublicRead SaveFilePayloadV2AccessControl = "public-read"
)

func (e SaveFilePayloadV2AccessControl) ToPointer() *SaveFilePayloadV2AccessControl {
	return &e
}

func (e *SaveFilePayloadV2AccessControl) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private":
		fallthrough
	case "public-read":
		*e = SaveFilePayloadV2AccessControl(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveFilePayloadV2AccessControl: %v", v)
	}
}

type SaveFilePayloadV2DocumentType string

const (
	SaveFilePayloadV2DocumentTypeDocument         SaveFilePayloadV2DocumentType = "document"
	SaveFilePayloadV2DocumentTypeDocumentTemplate SaveFilePayloadV2DocumentType = "document_template"
	SaveFilePayloadV2DocumentTypeText             SaveFilePayloadV2DocumentType = "text"
	SaveFilePayloadV2DocumentTypeImage            SaveFilePayloadV2DocumentType = "image"
	SaveFilePayloadV2DocumentTypeVideo            SaveFilePayloadV2DocumentType = "video"
	SaveFilePayloadV2DocumentTypeAudio            SaveFilePayloadV2DocumentType = "audio"
	SaveFilePayloadV2DocumentTypeSpreadsheet      SaveFilePayloadV2DocumentType = "spreadsheet"
	SaveFilePayloadV2DocumentTypePresentation     SaveFilePayloadV2DocumentType = "presentation"
	SaveFilePayloadV2DocumentTypeFont             SaveFilePayloadV2DocumentType = "font"
	SaveFilePayloadV2DocumentTypeArchive          SaveFilePayloadV2DocumentType = "archive"
	SaveFilePayloadV2DocumentTypeApplication      SaveFilePayloadV2DocumentType = "application"
	SaveFilePayloadV2DocumentTypeUnknown          SaveFilePayloadV2DocumentType = "unknown"
)

func (e SaveFilePayloadV2DocumentType) ToPointer() *SaveFilePayloadV2DocumentType {
	return &e
}

func (e *SaveFilePayloadV2DocumentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "document":
		fallthrough
	case "document_template":
		fallthrough
	case "text":
		fallthrough
	case "image":
		fallthrough
	case "video":
		fallthrough
	case "audio":
		fallthrough
	case "spreadsheet":
		fallthrough
	case "presentation":
		fallthrough
	case "font":
		fallthrough
	case "archive":
		fallthrough
	case "application":
		fallthrough
	case "unknown":
		*e = SaveFilePayloadV2DocumentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveFilePayloadV2DocumentType: %v", v)
	}
}

type SaveFilePayloadV2S3ref struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (o *SaveFilePayloadV2S3ref) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *SaveFilePayloadV2S3ref) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

type SaveFilePayloadV2 struct {
	AdditionalProperties interface{}                     `additionalProperties:"true" json:"-"`
	Tags                 []string                        `json:"_tags,omitempty"`
	AccessControl        *SaveFilePayloadV2AccessControl `default:"private" json:"access_control"`
	// Custom external download url used for the file
	CustomDownloadURL *string                        `json:"custom_download_url,omitempty"`
	DocumentType      *SaveFilePayloadV2DocumentType `json:"document_type,omitempty"`
	// if passed, adds a new version to existing file entity
	FileEntityID *string                `json:"file_entity_id,omitempty"`
	Filename     string                 `json:"filename"`
	S3ref        SaveFilePayloadV2S3ref `json:"s3ref"`
}

func (s SaveFilePayloadV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SaveFilePayloadV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SaveFilePayloadV2) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SaveFilePayloadV2) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SaveFilePayloadV2) GetAccessControl() *SaveFilePayloadV2AccessControl {
	if o == nil {
		return nil
	}
	return o.AccessControl
}

func (o *SaveFilePayloadV2) GetCustomDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.CustomDownloadURL
}

func (o *SaveFilePayloadV2) GetDocumentType() *SaveFilePayloadV2DocumentType {
	if o == nil {
		return nil
	}
	return o.DocumentType
}

func (o *SaveFilePayloadV2) GetFileEntityID() *string {
	if o == nil {
		return nil
	}
	return o.FileEntityID
}

func (o *SaveFilePayloadV2) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *SaveFilePayloadV2) GetS3ref() SaveFilePayloadV2S3ref {
	if o == nil {
		return SaveFilePayloadV2S3ref{}
	}
	return o.S3ref
}
