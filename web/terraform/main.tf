terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.54.1"
    }
  }

  backend "s3" {
    bucket = "subscribed-tf-infra"
    key    = "subscribed-webapp"
    region = "eu-west-1"
  }
}

