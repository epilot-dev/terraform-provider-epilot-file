// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UploadFilePublicS3ref struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (o *UploadFilePublicS3ref) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *UploadFilePublicS3ref) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

// UploadFilePublicResponseBody - Pre-signed URL for POST / PUT upload
type UploadFilePublicResponseBody struct {
	S3ref     *UploadFilePublicS3ref `json:"s3ref,omitempty"`
	UploadURL *string                `json:"upload_url,omitempty"`
}

func (o *UploadFilePublicResponseBody) GetS3ref() *UploadFilePublicS3ref {
	if o == nil {
		return nil
	}
	return o.S3ref
}

func (o *UploadFilePublicResponseBody) GetUploadURL() *string {
	if o == nil {
		return nil
	}
	return o.UploadURL
}

type UploadFilePublicResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Pre-signed URL for POST / PUT upload
	Object *UploadFilePublicResponseBody
}

func (o *UploadFilePublicResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFilePublicResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFilePublicResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadFilePublicResponse) GetObject() *UploadFilePublicResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
