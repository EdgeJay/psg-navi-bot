locals {
  s3_origin_id = "psgNavitBotS3Origin"
}

resource "aws_cloudfront_distribution" "psgnavibot_s3_distribution" {
  origin {
    domain_name = aws_s3_bucket_website_configuration.psgnavibot_website.website_endpoint
    # domain_name = aws_s3_bucket.psgnavibot.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    custom_origin_config {
      http_port              = "80"
      https_port             = "443"
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
  }

  enabled             = true
  is_ipv6_enabled     = true

  # logging_config {
  #   include_cookies = true
  #   bucket          = "cf-psgnavibot-dev-logs.s3.amazonaws.com"
  #   prefix          = "cf-psgnavibot-dev"
  # }

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = true

      cookies {
        forward = "all"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  # Cache behavior with precedence 0
  ordered_cache_behavior {
    path_pattern     = "/assets/*"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl                = 0
    default_ttl            = 86400
    max_ttl                = 31536000
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  price_class = "PriceClass_200"
  
  restrictions {
    geo_restriction {
      restriction_type = "none"
      locations        = []
    }
  }

  tags = {
    environment = "dev"
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
