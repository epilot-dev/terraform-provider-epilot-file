// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteFilePayload struct {
	S3ref S3Reference `json:"s3ref"`
}

func (o *DeleteFilePayload) GetS3ref() S3Reference {
	if o == nil {
		return S3Reference{}
	}
	return o.S3ref
}
