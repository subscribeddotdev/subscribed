resource "aws_amplify_app" "subscribed_webapp" {
  name       = "subscribed-webapp"
  platform   = "WEB_COMPUTE"
  repository = "https://github.com/subscribeddotdev/subscribed-webapp"
  build_spec = <<-EOT
            version: 1
            frontend:
                phases:
                    preBuild:
                        commands:
                            - 'npm ci --cache .npm --prefer-offline'
                    build:
                        commands:
                            - 'npm run build'
                artifacts:
                    baseDirectory: .next
                    files:
                        - '**/*'
                cache:
                    paths:
                        - '.next/cache/**/*'
                        - '.npm/**/*'
        EOT

  custom_rule {
    source = "/<*>"
    status = "404-200"
    target = "/index.html"
  }

  tags = local.common_tags
}

