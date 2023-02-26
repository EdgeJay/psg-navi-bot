resource "aws_s3_bucket" "psgnavibot" {
  bucket = "psgnavibot-dev"

  tags = {
    environment = "dev"
  }
}

resource "aws_s3_bucket_website_configuration" "psgnavibot_website" {
  bucket = aws_s3_bucket.psgnavibot.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

resource "aws_s3_bucket_acl" "psgnavibot_acl" {
  bucket = aws_s3_bucket.psgnavibot.id
  acl    = "private"
}

resource "aws_s3_bucket_versioning" "psgnavibot_versioning" {
  bucket = aws_s3_bucket.psgnavibot.id
  versioning_configuration {
    status = "Disabled"
  }
}

# Upload files needed for bot-frontend
resource "aws_s3_object" "css" {
  for_each = fileset("../../bot-frontend/dist/assets/", "*.css")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "/${random_id.app_version_suffix.hex}/assets/${each.value}"
  source       = "../../bot-frontend/dist/assets/${each.value}"
  etag         = filemd5("../../bot-frontend/dist/assets/${each.value}")
  acl          = "public-read"
  content_type = "text/css"
}

resource "aws_s3_object" "js" {
  for_each = fileset("../../bot-frontend/dist/assets/", "*.js")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "/${random_id.app_version_suffix.hex}/assets/${each.value}"
  source       = "../../bot-frontend/dist/assets/${each.value}"
  etag         = filemd5("../../bot-frontend/dist/assets/${each.value}")
  acl          = "public-read"
  content_type = "application/javascript"
}

# Upload default index and error html
resource "aws_s3_object" "html" {
  for_each = fileset("./static/", "*.html")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "${each.value}"
  source       = "./static/${each.value}"
  etag         = filemd5("./static/${each.value}")
  acl          = "public-read"
  content_type = "text/html"
}

# Upload favicon
resource "aws_s3_object" "favicon" {
  for_each = fileset("./static/", "*.ico")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "${each.value}"
  source       = "./static/${each.value}"
  etag         = filemd5("./static/${each.value}")
  acl          = "public-read"
  content_type = "image/x-icon"
}