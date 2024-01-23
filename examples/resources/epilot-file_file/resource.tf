resource "epilot-file_file" "my_file" {
  access_control        = "private"
  filename              = "blaaa.png"
  bucket = "epilot-prod-user-content"
  key    = "66/temp/4d689aeb-1497-4410-a9fe-b36ca9ac4389/nishu.png"
}