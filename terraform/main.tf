resource "aws_ecr_repository" "subscribed_backend_prd" {
  name                 = "subscribed-backend-prd"
  image_tag_mutability = "IMMUTABLE"
  tags                 = local.common_tags

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_db_instance" "subscribed_backend_postgres_prd" {
  instance_class               = "db.t3.micro"
  storage_encrypted            = true
  tags                         = local.common_tags
  performance_insights_enabled = true
  skip_final_snapshot          = true
  copy_tags_to_snapshot        = true
  apply_immediately            = false
  max_allocated_storage        = 1000
}

# TODO: Import the vpc config and the security group
