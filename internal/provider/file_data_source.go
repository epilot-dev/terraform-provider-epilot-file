// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-file/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FileDataSource{}
var _ datasource.DataSourceWithConfigure = &FileDataSource{}

func NewFileDataSource() datasource.DataSource {
	return &FileDataSource{}
}

// FileDataSource is the data source implementation.
type FileDataSource struct {
	client *sdk.SDK
}

// FileDataSourceModel describes the data model.
type FileDataSourceModel struct {
	AccessControl     types.String       `tfsdk:"access_control"`
	ACL               *tfTypes.ACL       `tfsdk:"acl"`
	CreatedAt         types.String       `tfsdk:"created_at"`
	CustomDownloadURL types.String       `tfsdk:"custom_download_url"`
	Filename          types.String       `tfsdk:"filename"`
	ID                types.String       `tfsdk:"id"`
	MimeType          types.String       `tfsdk:"mime_type"`
	Org               types.String       `tfsdk:"org"`
	PublicURL         types.String       `tfsdk:"public_url"`
	Purpose           []types.String     `tfsdk:"purpose"`
	ReadableSize      types.String       `tfsdk:"readable_size"`
	S3ref             *tfTypes.S3Ref     `tfsdk:"s3ref"`
	Schema            types.String       `tfsdk:"schema"`
	SizeBytes         types.Int64        `tfsdk:"size_bytes"`
	SourceURL         types.String       `tfsdk:"source_url"`
	Tags              []types.String     `tfsdk:"tags"`
	Title             types.String       `tfsdk:"title"`
	Type              types.String       `tfsdk:"type"`
	UpdatedAt         types.String       `tfsdk:"updated_at"`
	Versions          []tfTypes.FileItem `tfsdk:"versions"`
}

// Metadata returns the data source type name.
func (r *FileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

// Schema defines the schema for the data source.
func (r *FileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "File DataSource",

		Attributes: map[string]schema.Attribute{
			"access_control": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["private", "public-read"]`,
			},
			"acl": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"delete": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"edit": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"view": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Access control list for file entity (readonly)`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"custom_download_url": schema.StringAttribute{
				Computed:    true,
				Description: `Custom external download url used for the file`,
			},
			"filename": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"mime_type": schema.StringAttribute{
				Computed:    true,
				Description: `MIME type of the file`,
			},
			"org": schema.StringAttribute{
				Computed: true,
			},
			"public_url": schema.StringAttribute{
				Computed:    true,
				Description: `Direct URL for file (public only if file access control is public-read)`,
			},
			"purpose": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"readable_size": schema.StringAttribute{
				Computed:    true,
				Description: `Human readable file size`,
			},
			"s3ref": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Computed: true,
					},
					"key": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"schema": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["file"]`,
			},
			"size_bytes": schema.Int64Attribute{
				Computed:    true,
				Description: `File size in bytes`,
			},
			"source_url": schema.StringAttribute{
				Computed:    true,
				Description: `Source URL for the file. Included if the entity was created from source_url, or when ?source_url=true`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["document", "document_template", "text", "image", "video", "audio", "spreadsheet", "presentation", "font", "archive", "application", "unknown"]`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"versions": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"filename": schema.StringAttribute{
							Computed: true,
						},
						"mime_type": schema.StringAttribute{
							Computed: true,
						},
						"readable_size": schema.StringAttribute{
							Computed: true,
						},
						"s3ref": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"bucket": schema.StringAttribute{
									Computed: true,
								},
								"key": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"size_bytes": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *FileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *FileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FileDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	request := operations.GetFileRequest{
		ID: id,
	}
	res, err := r.client.File.GetFile(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FileEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res.FileEntity)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
