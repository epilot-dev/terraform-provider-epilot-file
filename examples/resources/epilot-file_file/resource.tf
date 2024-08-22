resource "epilot-file_file" "my_file" {
  id                  = "ef7d985c-2385-44f4-9c71-ae06a52264f8"
  title               = "document.pdf"
  access_control      = "private"
  custom_download_url = "https://some-api-url.com/download?file_id=123"
  filename            = "document.pdf"
  mime_type           = "application/pdf"
  source_url          = "https://productengineer-content.s3.eu-west-1.amazonaws.com/product-engineer-checklist.pdf"
  type                = "unknown"
}