resource "aws_lambda_function" "lambda_func" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = local.app_id
  handler          = "app"
  source_code_hash = base64sha256(data.archive_file.lambda_zip.output_path)
  runtime          = "go1.x"
  role             = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      app_env                     = var.app_env
      app_version                 = "${var.app_version}-${random_id.app_version_suffix.hex}"
      app_version_secret          = var.app_version_secret
      lambda_invoke_url           = var.lambda_invoke_url
      cookie_duration             = var.cookie_duration
      telegram_webapp_secret_key  = var.telegram_webapp_secret_key
    }
  }
}

# Assume role setup
resource "aws_iam_role" "lambda_exec" {
  name_prefix = local.app_id

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

data "aws_iam_policy_document" "lambda_ssm_policy_document" {
  statement {
    effect  = "Allow"
    actions = [
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:GetParametersByPath"
    ]

    resources = [
      "${aws_ssm_parameter.dev_params.arn}",
      "${aws_ssm_parameter.dev_dbx_app_key.arn}",
      "${aws_ssm_parameter.dev_dbx_app_secret.arn}",
      "${aws_ssm_parameter.dev_dbx_refresh_token.arn}",
      "${aws_ssm_parameter.dev_openai_api_key.arn}",
      "${aws_ssm_parameter.dev_rsa_private.arn}",
      "${aws_ssm_parameter.dev_rsa_public.arn}",
      "${aws_ssm_parameter.dev_init_token_secret.arn}",
      "${aws_ssm_parameter.dev_config_admins.arn}"
    ]
  }

  statement {
    effect  = "Allow"
    actions = [
      "kms:Decrypt"
    ]

    resources = [
      "arn:aws:kms:ap-southeast-1:818374272882:key/1a7ea267-43ae-40cc-8ead-8e26a0d63149"
    ]
  }
}

resource "aws_iam_policy" "lambda_ssm_policy" {
  name   = "PSGNaviBotLambdaSSMPolicy"
  policy = data.aws_iam_policy_document.lambda_ssm_policy_document.json
}

resource "aws_iam_policy_attachment" "role_attach" {
  name       = "policy-${local.app_id}"
  roles      = [aws_iam_role.lambda_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}

resource "aws_iam_policy_attachment" "custom_policy_attach" {
  name       = "custom-policy-${local.app_id}"
  roles      = [aws_iam_role.lambda_exec.id]
  policy_arn = aws_iam_policy.lambda_ssm_policy.arn
}