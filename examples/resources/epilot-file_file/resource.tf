resource "epilot-file_file" "my_file" {
  access_control = "private"
  filename       = epilot-file_upload_file.my_uploadfile.filename
  bucket         = epilot-file_upload_file.my_uploadfile.s3ref.bucket
  key            = epilot-file_upload_file.my_uploadfile.s3ref.key
}