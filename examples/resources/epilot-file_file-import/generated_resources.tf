# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "91b9fd89-05ac-464d-8ac5-1379824471ac"
resource "epilot-file_file" "my_file" {
  access_control = "private"
  acl = {
    delete = ["org_739224"]
    edit   = ["org_739224"]
    view   = ["org_739224"]
  }
  additional          = null
  custom_download_url = null
  filename            = "8z87al.jpg"
  mime_type           = "image/jpeg"
  purpose             = []
  s3ref = {
    bucket = "epilot-dev-user-content"
    key    = "739224/69f04890-fa8d-4d2e-9347-4ac384da107b/8z87al.jpg"
  }
  source_url = "https://file.dev.sls.epilot.io/v1/files/public/links/739224-taqDBUyNjSX7c_n2v6Zvv/8z87al.jpg"
  strict     = null
  tags       = []
  title      = "8z87al.jpg"
  type       = "image"
}
