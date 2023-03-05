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
  key          = "/assets/${each.value}"
  source       = "../../bot-frontend/dist/assets/${each.value}"
  etag         = filemd5("../../bot-frontend/dist/assets/${each.value}")
  acl          = "public-read"
  content_type = "text/css"
}

resource "aws_s3_object" "js" {
  for_each = fileset("../../bot-frontend/dist/assets/", "*.js")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "/assets/${each.value}"
  source       = "../../bot-frontend/dist/assets/${each.value}"
  etag         = filemd5("../../bot-frontend/dist/assets/${each.value}")
  acl          = "public-read"
  content_type = "application/javascript"
}

resource "aws_s3_object" "svg_icons" {
  for_each = fileset("../../bot-frontend/dist/icons/", "*.svg")

  bucket       = aws_s3_bucket.psgnavibot.bucket
  key          = "/icons/${each.value}"
  source       = "../../bot-frontend/dist/icons/${each.value}"
  etag         = filemd5("../../bot-frontend/dist/icons/${each.value}")
  acl          = "public-read"
  content_type = "image/svg+xml"
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

resource "aws_s3_object" "index_html" {
  for_each = fileset("../../bot-frontend/dist/", "index.html")

  bucket        = aws_s3_bucket.psgnavibot.bucket
  key           = "${each.value}"
  source        = "../../bot-frontend/dist/${each.value}"
  etag          = format("%s.%s", filemd5("../../bot-frontend/dist/${each.value}"), sha256(format("%s-%s", var.app_version, random_id.app_version_suffix.hex)))
  acl           = "public-read"
  content_type  = "text/html"
  cache_control = "no-cache, max-age=0, s-maxage=0"
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

# ============ This section onwards is for article ingestion/processing pipeline ============

resource "aws_s3_bucket" "psgnavibot_articles" {
  bucket = "${var.uploaded_articles_bucket}"

  tags = {
    environment = "dev"
  }
}

resource "aws_s3_bucket_acl" "psgnavibot_articles_acl" {
  bucket = aws_s3_bucket.psgnavibot_articles.id
  acl    = "private"
}

resource "aws_s3_bucket_versioning" "psgnavibot_articles_versioning" {
  bucket = aws_s3_bucket.psgnavibot_articles.id
  versioning_configuration {
    status = "Disabled"
  }
}

resource "aws_s3_bucket_public_access_block" "psgnavibot_articles_block_public" {
  bucket = aws_s3_bucket.psgnavibot_articles.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# S3 event filter
resource "aws_s3_bucket_notification" "psgnavibot_articles_notification" {
  bucket = aws_s3_bucket.psgnavibot_articles.id
  queue {
    queue_arn     = aws_sqs_queue.psgnavibot_articles_uploaded_queue.arn
    events        = ["s3:ObjectCreated:*"]
  }
}
