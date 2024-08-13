// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *FileDataSourceModel) RefreshFromSharedFileEntity(resp *shared.FileEntity) {
	if resp.ACL == nil {
		r.ACL = nil
	} else {
		r.ACL = &ACL{}
		r.ACL.Delete = nil
		for _, v := range resp.ACL.Delete {
			r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
		}
		r.ACL.Edit = nil
		for _, v := range resp.ACL.Edit {
			r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
		}
		r.ACL.View = nil
		for _, v := range resp.ACL.View {
			r.ACL.View = append(r.ACL.View, types.StringValue(v))
		}
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	} else {
		r.CreatedAt = types.StringNull()
	}
	r.Org = types.StringPointerValue(resp.Org)
	if resp.Schema != nil {
		r.Schema = types.StringValue(string(*resp.Schema))
	} else {
		r.Schema = types.StringNull()
	}
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	r.Title = types.StringPointerValue(resp.Title)
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	if resp.AccessControl != nil {
		r.AccessControl = types.StringValue(string(*resp.AccessControl))
	} else {
		r.AccessControl = types.StringNull()
	}
	r.CustomDownloadURL = types.StringPointerValue(resp.CustomDownloadURL)
	r.Filename = types.StringPointerValue(resp.Filename)
	r.ID = types.StringPointerValue(resp.ID)
	r.MimeType = types.StringPointerValue(resp.MimeType)
	r.PublicURL = types.StringPointerValue(resp.PublicURL)
	r.ReadableSize = types.StringPointerValue(resp.ReadableSize)
	if resp.S3ref == nil {
		r.S3ref = nil
	} else {
		r.S3ref = &S3Ref{}
		r.S3ref.Bucket = types.StringValue(resp.S3ref.Bucket)
		r.S3ref.Key = types.StringValue(resp.S3ref.Key)
	}
	r.SizeBytes = types.Int64PointerValue(resp.SizeBytes)
	r.SourceURL = types.StringPointerValue(resp.SourceURL)
	if resp.Type != nil {
		r.Type = types.StringValue(string(*resp.Type))
	} else {
		r.Type = types.StringNull()
	}
	if len(r.Versions) > len(resp.Versions) {
		r.Versions = r.Versions[:len(resp.Versions)]
	}
	for versionsCount, versionsItem := range resp.Versions {
		var versions1 FileItem
		versions1.Filename = types.StringPointerValue(versionsItem.Filename)
		versions1.MimeType = types.StringPointerValue(versionsItem.MimeType)
		versions1.ReadableSize = types.StringPointerValue(versionsItem.ReadableSize)
		if versionsItem.S3ref == nil {
			versions1.S3ref = nil
		} else {
			versions1.S3ref = &S3Ref{}
			versions1.S3ref.Bucket = types.StringValue(versionsItem.S3ref.Bucket)
			versions1.S3ref.Key = types.StringValue(versionsItem.S3ref.Key)
		}
		versions1.SizeBytes = types.Int64PointerValue(versionsItem.SizeBytes)
		if versionsCount+1 > len(r.Versions) {
			r.Versions = append(r.Versions, versions1)
		} else {
			r.Versions[versionsCount].Filename = versions1.Filename
			r.Versions[versionsCount].MimeType = versions1.MimeType
			r.Versions[versionsCount].ReadableSize = versions1.ReadableSize
			r.Versions[versionsCount].S3ref = versions1.S3ref
			r.Versions[versionsCount].SizeBytes = versions1.SizeBytes
		}
	}
}
