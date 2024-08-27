# epilot-file_file.my_file:
resource "epilot-file_file" "my_file" {
    access_control      = "private"
    custom_download_url = "https://ucarecdn.com/05f649bf-b70b-4cf8-90f7-2588ce404a08/-/format/auto/"
    filename            = "3rd-party-api-flow.jpg"
    mime_type           = "image/jpeg"
    purpose             = [
        "22fba2b3-9e40-40c1-bb06-e9cd20cea6cf",
        "c12409a0-770a-4009-b725-c84d6f3ccb5b",
    ]
    s3ref               = {
        bucket = "epilot-dev-user-content"
        key    = "739224/7daa1de4-fc77-48b3-8906-8e5a9129be07/3rd-party-api-flow.jpg"
    }
    source_url          = "https://file.dev.sls.epilot.io/v1/files/public/links/739224-yImePD8pFE1k1II0UumDl/3rd-party-api-flow.jpg"
    tags                = []
    title               = "3rd-party-api-flow.jpg"
    type                = "image"
}
