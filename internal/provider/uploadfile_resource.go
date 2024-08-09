// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-file/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UploadFileResource{}
var _ resource.ResourceWithImportState = &UploadFileResource{}

func NewUploadFileResource() resource.Resource {
	return &UploadFileResource{}
}

// UploadFileResource defines the resource implementation.
type UploadFileResource struct {
	client *sdk.SDK
}

// UploadFileResourceModel describes the resource data model.
type UploadFileResourceModel struct {
	FileEntityID types.String                    `tfsdk:"file_entity_id"`
	Filename     types.String                    `tfsdk:"filename"`
	IndexTag     types.String                    `tfsdk:"index_tag"`
	Metadata     map[string]types.String         `tfsdk:"metadata"`
	MimeType     types.String                    `tfsdk:"mime_type"`
	PublicURL    types.String                    `tfsdk:"public_url"`
	S3ref        *tfTypes.SaveFilePayloadV2S3ref `tfsdk:"s3ref"`
	UploadURL    types.String                    `tfsdk:"upload_url"`
}

func (r *UploadFileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_upload_file"
}

func (r *UploadFileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "UploadFile Resource",
		Attributes: map[string]schema.Attribute{
			"file_entity_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `file entity id. Requires replacement if changed. `,
			},
			"filename": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"index_tag": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Used to index the file at the storage layer, which helps when browsing for this file. Requires replacement if changed. `,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(64),
				},
			},
			"metadata": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `Allows passing in custom metadata for the file, expects key-value pairs of string type. Requires replacement if changed. `,
			},
			"mime_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("application/octet-stream"),
				Description: `MIME type of file. Requires replacement if changed. ; Default: "application/octet-stream"`,
			},
			"public_url": schema.StringAttribute{
				Computed:    true,
				Description: `Returned only if file is permanent i.e. file_entity_id is passed`,
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
			"upload_url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *UploadFileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *UploadFileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *UploadFileResourceModel
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

	uploadFilePayload := data.ToSharedUploadFilePayload()
	fileEntityID := new(string)
	if !data.FileEntityID.IsUnknown() && !data.FileEntityID.IsNull() {
		*fileEntityID = data.FileEntityID.ValueString()
	} else {
		fileEntityID = nil
	}
	request := operations.UploadFileV2Request{
		UploadFilePayload: uploadFilePayload,
		FileEntityID:      fileEntityID,
	}
	res, err := r.client.Files.UploadFileV2(ctx, request)
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
	if !(res.FileUpload != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFileUpload(res.FileUpload)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UploadFileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *UploadFileResourceModel
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

func (r *UploadFileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *UploadFileResourceModel
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

func (r *UploadFileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *UploadFileResourceModel
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

func (r *UploadFileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource upload_file.")
}
