terraform {
  required_providers {
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "2.2.2"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.33.0"
    }
  }
}

variable epilot_auth {
  type        = string
  description = "epilot_auth"
}


provider "epilot-file" {
  # Configuration options
  epilot_auth = var.epilot_auth
}

provider "aws" {
  # Configuration options
}