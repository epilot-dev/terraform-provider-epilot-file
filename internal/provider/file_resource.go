// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	speakeasy_stringplanmodifier "github.com/epilot-dev/terraform-provider-epilot-file/internal/planmodifiers/stringplanmodifier"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/pkg/models/shared"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FileResource{}
var _ resource.ResourceWithImportState = &FileResource{}

func NewFileResource() resource.Resource {
	return &FileResource{}
}

// FileResource defines the resource implementation.
type FileResource struct {
	client *sdk.SDK
}

// FileResourceModel describes the resource data model.
type FileResourceModel struct {
	ID                   types.String   `tfsdk:"id"`
	AccessControl        types.String   `tfsdk:"access_control"`
	AdditionalProperties types.String   `tfsdk:"additional_properties"`
	Bucket               types.String   `tfsdk:"bucket"`
	CustomDownloadURL    types.String   `tfsdk:"custom_download_url"`
	DocumentType         types.String   `tfsdk:"document_type"`
	FileEntityID         types.String   `tfsdk:"file_entity_id"`
	Filename             types.String   `tfsdk:"filename"`
	Key                  types.String   `tfsdk:"key"`
	MimeType             types.String   `tfsdk:"mime_type"`
	PublicURL            types.String   `tfsdk:"public_url"`
	SizeBytes            types.Int64    `tfsdk:"size_bytes"`
	Tags                 []types.String `tfsdk:"tags"`
	Type                 types.String   `tfsdk:"type"`
	Versions             []Versions     `tfsdk:"versions"`
}

func (r *FileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "File Resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"access_control": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Requires replacement if changed. ; must be one of ["private", "public-read"]; Default: "private"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"private",
						"public-read",
					),
				},
			},
			"additional_properties": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"bucket": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"custom_download_url": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Custom external download url used for the file. Requires replacement if changed. `,
			},
			"document_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Requires replacement if changed. ; must be one of ["document", "document_template", "text", "image", "video", "audio", "spreadsheet", "presentation", "font", "archive", "application", "unknown"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"document",
						"document_template",
						"text",
						"image",
						"video",
						"audio",
						"spreadsheet",
						"presentation",
						"font",
						"archive",
						"application",
						"unknown",
					),
				},
			},
			"file_entity_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `if passed, adds a new version to existing file entity. Requires replacement if changed. `,
			},
			"filename": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"key": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"mime_type": schema.StringAttribute{
				Computed:    true,
				Description: `MIME type of the file`,
			},
			"public_url": schema.StringAttribute{
				Computed:    true,
				Description: `Direct URL for file (public only if file access control is public-read)`,
			},
			"size_bytes": schema.Int64Attribute{
				Computed:    true,
				Description: `File size in bytes`,
			},
			"tags": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `Requires replacement if changed. `,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `Human readable type for file. must be one of ["document", "document_template", "text", "image", "video", "audio", "spreadsheet", "presentation", "font", "archive", "application", "unknown"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"document",
						"document_template",
						"text",
						"image",
						"video",
						"audio",
						"spreadsheet",
						"presentation",
						"font",
						"archive",
						"application",
						"unknown",
					),
				},
			},
			"versions": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
					},
				},
			},
		},
	}
}

func (r *FileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *FileResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var request *shared.SaveFilePayloadV2
	var additionalProperties interface{}
	if !data.AdditionalProperties.IsUnknown() && !data.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(data.AdditionalProperties.ValueString()), &additionalProperties)
	}
	var tags []string = nil
	for _, tagsItem := range data.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	accessControl := new(shared.SaveFilePayloadV2AccessControl)
	if !data.AccessControl.IsUnknown() && !data.AccessControl.IsNull() {
		*accessControl = shared.SaveFilePayloadV2AccessControl(data.AccessControl.ValueString())
	} else {
		accessControl = nil
	}
	customDownloadURL := new(string)
	if !data.CustomDownloadURL.IsUnknown() && !data.CustomDownloadURL.IsNull() {
		*customDownloadURL = data.CustomDownloadURL.ValueString()
	} else {
		customDownloadURL = nil
	}
	documentType := new(shared.SaveFilePayloadV2DocumentType)
	if !data.DocumentType.IsUnknown() && !data.DocumentType.IsNull() {
		*documentType = shared.SaveFilePayloadV2DocumentType(data.DocumentType.ValueString())
	} else {
		documentType = nil
	}
	fileEntityID := new(string)
	if !data.FileEntityID.IsUnknown() && !data.FileEntityID.IsNull() {
		*fileEntityID = data.FileEntityID.ValueString()
	} else {
		fileEntityID = nil
	}
	filename := data.Filename.ValueString()
	bucket := data.Bucket.ValueString()
	key := data.Key.ValueString()
	s3ref := shared.SaveFilePayloadV2S3ref{
		Bucket: bucket,
		Key:    key,
	}
	request = &shared.SaveFilePayloadV2{
		AdditionalProperties: additionalProperties,
		Tags:                 tags,
		AccessControl:        accessControl,
		CustomDownloadURL:    customDownloadURL,
		DocumentType:         documentType,
		FileEntityID:         fileEntityID,
		Filename:             filename,
		S3ref:                s3ref,
	}
	res, err := r.client.Files.SaveFileV2(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.FileEntity == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res.FileEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *FileResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *FileResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *FileResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *FileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource file.")
}
