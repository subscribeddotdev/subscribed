terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.53.0"
    }
  }

  backend "s3" {
    bucket = "subscribed-tf-infra"
    key    = "subscribed-backend"
    region = "eu-west-1"
  }
}
