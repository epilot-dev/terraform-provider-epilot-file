# epilot-file

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>

<no value>
<!-- Start SDK <no value> -->
To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "0.2.0"
    }
  }
}

provider "epilot-file" {
  # Configuration options
}
```
<!-- End SDK <no value> -->

<no value>
<!-- Start SDK <no value> -->
### Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

### Example

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK <no value> -->

<no value>
<!-- Start SDK <no value> -->

<!-- End SDK <no value> -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "0.8.0"
    }
  }
}

provider "epilot-file" {
  server_url = "..." # Optional
}
```
<!-- End Installation [installation] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-epilot-file`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/epilot-dev/epilot-file" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Start Authentication [security] -->
## Authentication

This provider supports authentication configuration via provider configuration.

Available configuration:

| Provider Attribute | Description |
|---|---|
| `cookie_auth` | Cookie-based session authentication for browser applications.

**When to use:** Browser-based applications that need to:
- Embed file previews directly in `<img>` tags
- Download files without JavaScript token handling
- Access files from HTML elements that cannot set custom headers

**How to establish a session:**
1. Obtain a Bearer token via EpilotAuth
2. Call `GET /v1/files/session` with the Bearer token
3. The server sets an HTTP-only cookie named `token`
4. Subsequent requests automatically include the cookie

**Security note:** The cookie is HTTP-only and secure, protecting against XSS attacks.
. |
| `epilot_auth` | Bearer token authentication using epilot OAuth2 JWT tokens.

**When to use:** Server-to-server integrations, API clients, and programmatic access.

**How to obtain a token:**
1. Use the epilot Auth API to authenticate
2. Include the token in the `Authorization` header: `Authorization: Bearer <token>`

**Example:**
```
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Token contents:** The JWT contains user identity, organization ID, and permissions.
. |
<!-- End Authentication [security] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources


<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Summary [summary] -->
## Summary

File API: The File API enables you to upload, store, manage, and share files within the epilot platform.

## Key Features
- **Upload files** to temporary storage and save them permanently as File entities
- **Generate previews** (thumbnails) for images and documents
- **Create public links** to share private files externally
- **Organize files** into collections for better management
- **Version control** with automatic file versioning on updates

## File Upload Workflow
1. Call `uploadFileV2` to get a pre-signed S3 URL
2. Upload your file directly to S3 using the pre-signed URL (PUT request)
3. Call `saveFileV2` with the S3 reference to create a permanent File entity

## Changelog
<a href="changelog">View API Changelog</a>
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [epilot-file](#epilot-file)
  * [Installation](#installation)
  * [Testing the provider locally](#testing-the-provider-locally)
  * [Authentication](#authentication)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Key Features](#key-features)
  * [File Upload Workflow](#file-upload-workflow)
  * [Changelog](#changelog)

<!-- End Table of Contents [toc] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/epilot-dev/epilot-file" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

Note: To use the dev_overrides, please ensure you run `go build` in this folder. You must have a binary available for terraform to find.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
