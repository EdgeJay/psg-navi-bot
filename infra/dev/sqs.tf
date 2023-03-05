resource "aws_sqs_queue" "psgnavibot_articles_uploaded_queue" {
  name = "psgnavibot-${var.app_env}-articles-uploaded-queue"
  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "arn:aws:sqs:*:*:psgnavibot-${var.app_env}-articles-uploaded-queue",
      "Condition": {
        "ArnEquals": {
            "aws:SourceArn": "${aws_s3_bucket.psgnavibot_articles.arn}"
        }
      }
    }
  ]
}
POLICY
}