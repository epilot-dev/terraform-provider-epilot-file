terraform {
  required_providers {
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "0.8.0"
    }
  }
}

provider "epilot-file" {
  server_url = "..." # Optional
}