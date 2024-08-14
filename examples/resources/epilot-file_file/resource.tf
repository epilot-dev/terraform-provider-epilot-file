resource "epilot-file_file" "my_file" {
  filename            = "terraform-product-engineer-checklist.pdf"
  access_control      = "private"
  source_url          = "https://productengineer-content.s3.eu-west-1.amazonaws.com/product-engineer-checklist.pdf"
}
