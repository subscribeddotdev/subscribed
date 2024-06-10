resource "aws_ecr_repository" "subscribed_backend_prd" {
  name                 = "subscribed-backend-prd"
  image_tag_mutability = "IMMUTABLE"
  tags = local.common_tags

  image_scanning_configuration {
    scan_on_push = true
  }
}
