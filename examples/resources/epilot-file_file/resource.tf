resource "epilot-file_file" "my_file" {
  id             = "ef7d985c-2385-44f4-9c71-ae06a52264f8"
  access_control = "private"
  filename       = "document.pdf"
  mime_type      = "application/pdf"
  public_url     = "https://epilot-files-prod.s3.eu-central-1.amazonaws.com/123/4d689aeb-1497-4410-a9fe-b36ca9ac4389/document.pdf"
  size_bytes     = 1
  type           = "font"
}