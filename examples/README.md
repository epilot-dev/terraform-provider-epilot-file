## Upload file to S3

```
resource "epilot-file_upload_file" "my_uploadfile" {
  filename       = "pumpkin.png"
  mime_type      = "image/png"
}

```

```
output "val" {
  value = epilot-file_upload_file.my_uploadfile
}
```

##  Save uploaded file

```
resource "epilot-file_file" "my_file" {
  access_control = "private"
  filename       = epilot-file_upload_file.my_uploadfile.filename
  bucket         = epilot-file_upload_file.my_uploadfile.s3ref.bucket
  key            = epilot-file_upload_file.my_uploadfile.s3ref.key
}
```

