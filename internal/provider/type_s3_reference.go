// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type S3Reference struct {
	Bucket types.String `tfsdk:"bucket"`
	Key    types.String `tfsdk:"key"`
}
