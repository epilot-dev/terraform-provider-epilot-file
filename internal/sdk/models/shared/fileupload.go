// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type S3ref struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (o *S3ref) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *S3ref) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

type FileUpload struct {
	// Returned only if file is permanent i.e. file_entity_id is passed
	PublicURL *string `json:"public_url,omitempty"`
	S3ref     *S3ref  `json:"s3ref,omitempty"`
	UploadURL *string `json:"upload_url,omitempty"`
}

func (o *FileUpload) GetPublicURL() *string {
	if o == nil {
		return nil
	}
	return o.PublicURL
}

func (o *FileUpload) GetS3ref() *S3ref {
	if o == nil {
		return nil
	}
	return o.S3ref
}

func (o *FileUpload) GetUploadURL() *string {
	if o == nil {
		return nil
	}
	return o.UploadURL
}