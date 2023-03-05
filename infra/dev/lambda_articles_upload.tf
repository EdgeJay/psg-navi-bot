# Event source from SQS
resource "aws_lambda_event_source_mapping" "psgnavibot_articles_event_source_mapping" {
    event_source_arn = aws_sqs_queue.psgnavibot_articles_uploaded_queue.arn
    enabled          = true
    function_name    = aws_lambda_function.articles_upload.arn
    batch_size       = 1
}

# Lambda function declaration
resource "aws_lambda_function" "articles_upload" {
    filename         = data.archive_file.lambda_articles_upload_zip.output_path
    function_name    = local.articles_upload_id
    handler          = "app"
    source_code_hash = data.archive_file.lambda_articles_upload_zip.output_base64sha256
    runtime          = "go1.x"
    role             = aws_iam_role.lambda_exec_articles_upload.arn
    timeout          = 120

    environment {
      variables = {
        app_env                  = var.app_env
        uploaded_articles_bucket = var.uploaded_articles_bucket
      }
    }
}

# Lambda function role
resource "aws_iam_role" "lambda_exec_articles_upload" {
    name_prefix = local.articles_upload_id
    
    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_policy" "lambda_exec_articles_upload_policy" {
    name = "policy-articles-upload-${local.app_id}"
 
    policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject",
        "s3:PutObject",
        "s3:ListBucket"
      ],
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.psgnavibot_articles.arn}/*"
    },
    {
      "Action": [
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage",
        "sqs:GetQueueAttributes"
      ],
      "Effect": "Allow",
      "Resource": "${aws_sqs_queue.psgnavibot_articles_uploaded_queue.arn}"
    },
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Action": [
        "ssm:GetParameter",
        "ssm:GetParameters",
        "ssm:GetParametersByPath"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_ssm_parameter.dev_openai_api_key.arn}"
      ]
    }
  ]
}
EOF
}

# Role to Policy attachment
resource "aws_iam_role_policy_attachment" "lambda_exec_articles_upload_policy_attachment" {
    role        = aws_iam_role.lambda_exec_articles_upload.id
    policy_arn  = aws_iam_policy.lambda_exec_articles_upload_policy.arn
}

# CloudWatch Log Group for the Lambda function
resource "aws_cloudwatch_log_group" "lambda_exec_articles_upload_loggroup" {
    name              = "/aws/lambda/${aws_lambda_function.articles_upload.function_name}"
    retention_in_days = 7
}
