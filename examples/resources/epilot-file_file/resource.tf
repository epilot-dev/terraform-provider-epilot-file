resource "epilot-file_file" "my_file" {
  access_control        = "public-read"
  additional_properties = "{ \"see\": \"documentation\" }"
  bucket                = "epilot-files-prod"
  custom_download_url   = "https://both-legging.net"
  document_type         = "archive"
  file_entity_id        = "...my_file_entity_id..."
  filename              = "document.pdf"
  key                   = "123/4d689aeb-1497-4410-a9fe-b36ca9ac4389/document.pdf"
  tags = [
    "...",
  ]
}