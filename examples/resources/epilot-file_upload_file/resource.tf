resource "epilot-file_upload_file" "my_uploadfile" {
  file_entity_id = "ef7d985c-2385-44f4-9c71-ae06a52264f8"
  filename       = "document.pdf"
  index_tag      = "2f6a377c8e78"
  metadata = {
    "id"      = "..."
    "Jamaica" = "..."
  }
  mime_type = "application/pdf"
}