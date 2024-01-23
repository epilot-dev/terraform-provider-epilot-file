# Upload file to S3
resource "epilot-file_upload_file" "my_uploadfile" {
  filename       = "NishuGoel.png"
  mime_type      = "image/png"
}

output "uploaded_file" {
  value = epilot-file_upload_file.my_uploadfile
}


resource "aws_s3_object" "s3_file_upload" {
  bucket = epilot-file_upload_file.my_uploadfile.s3ref.bucket
  key    = epilot-file_upload_file.my_uploadfile.s3ref.key
  source = "/Users/nishugoel/epilot/terraform blueprints/terraform-provider-epilot-file/examples/resources/epilot-file_upload_file/NishuGoel.png"
}

#  Save uploaded file
resource "epilot-file_file" "my_file" {
  access_control = "private"
  filename       = epilot-file_upload_file.my_uploadfile.filename
  bucket         = epilot-file_upload_file.my_uploadfile.s3ref.bucket
  key            = epilot-file_upload_file.my_uploadfile.s3ref.key
  entity_id      = ""
}


