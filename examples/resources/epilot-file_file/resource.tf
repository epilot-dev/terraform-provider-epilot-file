resource "epilot-file_file" "my_file" {
  access_control = "public-read"
  acl = {
    delete = [
      "org:456"
    ]
    edit = [
      "org:456"
    ]
    view = [
      "org:456"
    ]
  }
  activity_id = "01F130Q52Q6MWSNS8N2AVXV4JN"
  additional = {
    key = jsonencode("value")
  }
  async               = true
  custom_download_url = "https://some-api-url.com/download?file_id=123"
  delete_temp_file    = false
  filename            = "document.pdf"
  fill_activity       = true
  id                  = "...my_id..."
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  mime_type = "application/pdf"
  purge     = true
  purpose = [
    "8d396871-95a0-4c9d-bb4d-9eda9c35776c",
    "da7cdf9a-01be-40c9-a29c-9a8f9f0de6f8",
  ]
  s3ref = {
    bucket = "epilot-prod-user-content"
    key    = "123/4d689aeb-1497-4410-a9fe-b36ca9ac4389/document.pdf"
  }
  source_url = "https://productengineer-content.s3.eu-west-1.amazonaws.com/product-engineer-checklist.pdf"
  strict     = false
  tags = [
    "tag1",
    "tag2",
  ]
  title = "document.pdf"
  type  = "font"
}