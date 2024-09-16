// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-file/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/validators"
	speakeasy_stringvalidators "github.com/epilot-dev/terraform-provider-epilot-file/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	AccessControl     types.String              `tfsdk:"access_control"`
	ACL               *tfTypes.BaseEntityACL    `tfsdk:"acl"`
	ActivityID        types.String              `tfsdk:"activity_id"`
	Additional        map[string]types.String   `tfsdk:"additional"`
	CreatedAt         types.String              `tfsdk:"created_at"`
	CustomDownloadURL types.String              `tfsdk:"custom_download_url"`
	Filename          types.String              `tfsdk:"filename"`
	FillActivity      types.Bool                `tfsdk:"fill_activity"`
	ID                types.String              `tfsdk:"id"`
	Manifest          []types.String            `tfsdk:"manifest"`
	MimeType          types.String              `tfsdk:"mime_type"`
	Org               types.String              `tfsdk:"org"`
	Owners            []tfTypes.BaseEntityOwner `tfsdk:"owners"`
	PublicURL         types.String              `tfsdk:"public_url"`
	Purpose           []types.String            `tfsdk:"purpose"`
	ReadableSize      types.String              `tfsdk:"readable_size"`
	S3ref             *tfTypes.S3Ref            `tfsdk:"s3ref"`
	Schema            types.String              `tfsdk:"schema"`
	SizeBytes         types.Int64               `tfsdk:"size_bytes"`
	SourceURL         types.String              `tfsdk:"source_url"`
	Strict            types.Bool                `tfsdk:"strict"`
	Tags              []types.String            `tfsdk:"tags"`
	Title             types.String              `tfsdk:"title"`
	Type              types.String              `tfsdk:"type"`
	UpdatedAt         types.String              `tfsdk:"updated_at"`
	Versions          []tfTypes.FileItem        `tfsdk:"versions"`
}

func (r *FileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "File Resource",
		Attributes: map[string]schema.Attribute{
			"access_control": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("private"),
				Description: `Default: "private"; must be one of ["private", "public-read"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"private",
						"public-read",
					),
				},
			},
			"acl": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"delete": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"edit": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"view": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Access control list (ACL) for an entity. Defines sharing access to external orgs or users.`,
			},
			"activity_id": schema.StringAttribute{
				Optional:    true,
				Description: `Activity to include in event feed`,
			},
			"additional": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `Additional fields that are not part of the schema`,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"custom_download_url": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Custom external download url used for the file`,
			},
			"filename": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fill_activity": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
				MarkdownDescription: `Update the diff and entity for the custom activity included in the query.` + "\n" +
					`Pending state on activity is automatically ended when activity is filled.` + "\n" +
					`Default: false`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"manifest": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `Manifest ID used to create/update the entity`,
			},
			"mime_type": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `MIME type of the file`,
			},
			"org": schema.StringAttribute{
				Computed: true,
			},
			"owners": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"org_id": schema.StringAttribute{
							Computed: true,
						},
						"user_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"public_url": schema.StringAttribute{
				Computed:    true,
				Description: `Direct URL for file (public only if file access control is public-read)`,
			},
			"purpose": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"readable_size": schema.StringAttribute{
				Computed:    true,
				Description: `Human readable file size`,
			},
			"s3ref": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"key": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
				},
			},
			"schema": schema.StringAttribute{
				Computed:    true,
				Description: `must be "file"`,
				Validators: []validator.String{
					stringvalidator.OneOf("file"),
				},
			},
			"size_bytes": schema.Int64Attribute{
				Computed:    true,
				Description: `File size in bytes`,
			},
			"source_url": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Source URL for the file. Included if the entity was created from source_url, or when ?source_url=true`,
			},
			"strict": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `When passed true, the response will contain only fields that match the schema, with non-matching fields included in ` + "`" + `__additional` + "`" + `. Default: false`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `must be one of ["document", "document_template", "text", "image", "video", "audio", "spreadsheet", "presentation", "font", "archive", "application", "unknown"]`,
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
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
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

	saveFilePayloadV2 := data.ToSharedSaveFilePayloadV2()
	activityID := new(string)
	if !data.ActivityID.IsUnknown() && !data.ActivityID.IsNull() {
		*activityID = data.ActivityID.ValueString()
	} else {
		activityID = nil
	}
	fillActivity := new(bool)
	if !data.FillActivity.IsUnknown() && !data.FillActivity.IsNull() {
		*fillActivity = data.FillActivity.ValueBool()
	} else {
		fillActivity = nil
	}
	strict := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict = data.Strict.ValueBool()
	} else {
		strict = nil
	}
	request := operations.SaveFileV2Request{
		SaveFilePayloadV2: saveFilePayloadV2,
		ActivityID:        activityID,
		FillActivity:      fillActivity,
		Strict:            strict,
	}
	res, err := r.client.File.SaveFileV2(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FileEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res.FileEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id string
	id = data.ID.ValueString()

	strict1 := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict1 = data.Strict.ValueBool()
	} else {
		strict1 = nil
	}
	request1 := operations.GetFileRequest{
		ID:     id,
		Strict: strict1,
	}
	res1, err := r.client.File.GetFile(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.FileEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res1.FileEntity)
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

	var id string
	id = data.ID.ValueString()

	strict := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict = data.Strict.ValueBool()
	} else {
		strict = nil
	}
	request := operations.GetFileRequest{
		ID:     id,
		Strict: strict,
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

	saveFilePayloadV2 := data.ToSharedSaveFilePayloadV2()
	activityID := new(string)
	if !data.ActivityID.IsUnknown() && !data.ActivityID.IsNull() {
		*activityID = data.ActivityID.ValueString()
	} else {
		activityID = nil
	}
	fillActivity := new(bool)
	if !data.FillActivity.IsUnknown() && !data.FillActivity.IsNull() {
		*fillActivity = data.FillActivity.ValueBool()
	} else {
		fillActivity = nil
	}
	strict := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict = data.Strict.ValueBool()
	} else {
		strict = nil
	}
	request := operations.SaveFileV2Request{
		SaveFilePayloadV2: saveFilePayloadV2,
		ActivityID:        activityID,
		FillActivity:      fillActivity,
		Strict:            strict,
	}
	res, err := r.client.File.SaveFileV2(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FileEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res.FileEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id string
	id = data.ID.ValueString()

	strict1 := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict1 = data.Strict.ValueBool()
	} else {
		strict1 = nil
	}
	request1 := operations.GetFileRequest{
		ID:     id,
		Strict: strict1,
	}
	res1, err := r.client.File.GetFile(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.FileEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedFileEntity(res1.FileEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

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

	activityID := new(string)
	if !data.ActivityID.IsUnknown() && !data.ActivityID.IsNull() {
		*activityID = data.ActivityID.ValueString()
	} else {
		activityID = nil
	}
	var id string
	id = data.ID.ValueString()

	strict := new(bool)
	if !data.Strict.IsUnknown() && !data.Strict.IsNull() {
		*strict = data.Strict.ValueBool()
	} else {
		strict = nil
	}
	request := operations.DeleteFileRequest{
		ActivityID: activityID,
		ID:         id,
		Strict:     strict,
	}
	res, err := r.client.File.DeleteFile(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *FileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
