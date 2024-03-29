// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *UploadFileResourceModel) ToSharedUploadFilePayload() *shared.UploadFilePayload {
	filename := r.Filename.ValueString()
	mimeType := new(string)
	if !r.MimeType.IsUnknown() && !r.MimeType.IsNull() {
		*mimeType = r.MimeType.ValueString()
	} else {
		mimeType = nil
	}
	out := shared.UploadFilePayload{
		Filename: filename,
		MimeType: mimeType,
	}
	return &out
}

func (r *UploadFileResourceModel) RefreshFromSharedFileUpload(resp *shared.FileUpload) {
	r.PublicURL = types.StringPointerValue(resp.PublicURL)
	if resp.S3ref == nil {
		r.S3ref = nil
	} else {
		r.S3ref = &S3Reference{}
		r.S3ref.Bucket = types.StringValue(resp.S3ref.Bucket)
		r.S3ref.Key = types.StringValue(resp.S3ref.Key)
	}
	r.UploadURL = types.StringPointerValue(resp.UploadURL)
}
