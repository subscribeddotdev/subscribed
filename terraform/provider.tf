terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.53.0"
    }
  }

  backend "s3" {
    bucket = "subscribed-backend-prd"
    key    = "state"
    region = "eu-west-1"
  }
}
