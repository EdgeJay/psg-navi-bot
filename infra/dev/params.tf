resource "aws_ssm_parameter" "dev_params" {
  name        = "/psg_navi_bot/dev/telegram_api_token"
  description = "API token of PSGNaviBot Telegram bot"
  type        = "SecureString"
  value       = var.bot_token

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_dbx_app_key" {
  name        = "/psg_navi_bot/dev/dropbox_app_key"
  description = "Dropbox app key for PSGNaviBot"
  type        = "SecureString"
  value       = var.dropbox_app_key

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_dbx_app_secret" {
  name        = "/psg_navi_bot/dev/dropbox_app_secret"
  description = "Dropbox app secret for PSGNaviBot"
  type        = "SecureString"
  value       = var.dropbox_app_secret

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_dbx_refresh_token" {
  name        = "/psg_navi_bot/dev/dropbox_refresh_token"
  description = "Dropbox refresh token for PSGNaviBot"
  type        = "SecureString"
  value       = var.dropbox_refresh_token

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_openai_api_key" {
  name        = "/psg_navi_bot/dev/openai_api_key"
  description = "API key for OpenAI"
  type        = "SecureString"
  value       = var.openai_api_key

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_rsa_private" {
  name        = "/psg_navi_bot/dev/rsa_private"
  description = "RSA private key for JWT"
  type        = "SecureString"
  value       = "${data.local_file.rsa_private.content}"

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_rsa_public" {
  name        = "/psg_navi_bot/dev/rsa_public"
  description = "RSA public key for JWT"
  type        = "SecureString"
  value       = "${data.local_file.rsa_public.content}"

  tags = {
    environment = "dev"
  }
}