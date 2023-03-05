provider "aws" {
  region = "ap-southeast-1"
}

variable "app_name" {
  description = "Application name"
  default     = "psg-navi-bot-backend"
}

variable "articles_upload_app_name" {
  description = "Articles upload application name"
  default     = "psg-navi-bot-articles-upload"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "dev"
}

variable "app_version" {
  description = "Application version"
}

variable "app_version_secret" {
  description = "Secret to hash application version for verification purposes"
}

variable "bot_token" {
  description = "API token of Telegram bot"
}

variable "dropbox_app_key" {
  description = "Dropbox app key for PSGNaviBot"
}

variable "dropbox_app_secret" {
  description = "Dropbox app secret for PSGNaviBot"
}

variable "dropbox_refresh_token" {
  description = "Dropbox refresh token for PSGNaviBot"
}

variable "openai_api_key" {
  description = "API key for OpenAI"
}

variable "lambda_invoke_url" {
  description = "Url to invoke Lambda function"
}

variable "cookie_duration" {
  description = "Duration of cookies in seconds"
}

variable "init_token_secret" {
  description = "Secret for comparing init token"
}

variable "telegram_webapp_secret_key" {
  description = "Secret key for Telegram web app for validating initData"
}

variable "menu_session_checksum_secret" {
  description = "Secret key for creating checksum of menu session data"
}

variable "uploaded_articles_bucket" {
  description = "Name of S3 bucket for storing uploaded articles"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
  articles_upload_id = "${lower(var.articles_upload_app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir  = "../../bot-backend/build/dev/bin"
  output_path = "../../bot-backend/build/dev/app.zip"
}

data "archive_file" "lambda_articles_upload_zip" {
  type        = "zip"
  source_dir  = "../../articles-upload/build/dev/bin"
  output_path = "../../articles-upload/build/dev/app.zip"
}

data "local_file" "rsa_private" {
  filename = "../../bot-backend/certs/rsa_private.pem"
}

data "local_file" "rsa_public" {
  filename = "../../bot-backend/certs/rsa_public.pem"
}

data "local_file" "config_admins" {
  filename = "../../config/admins.json"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

resource "random_id" "app_version_suffix" {
  byte_length = 4

  keepers = {
    archive_hash = "${data.archive_file.lambda_zip.output_md5}"
  }
}

output "api_url" {
  value = "https://${aws_cloudfront_distribution.psgnavibot_s3_distribution.domain_name}/api"
}

output "app_version" {
  value = aws_lambda_function.lambda_func.environment[0].variables.app_version
}

output "init_token_secret" {
  value = var.init_token_secret
}

output "cloudfront_domain" {
  value = aws_cloudfront_distribution.psgnavibot_s3_distribution.domain_name
}