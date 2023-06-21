terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.0.1"
    }
  }

  backend "s3" {
    bucket  = "flower-terraform-state"
    key     = "terraform.tfstate"
    encrypt = true
    region  = "us-east-1"
    profile = "default"
  }
}
