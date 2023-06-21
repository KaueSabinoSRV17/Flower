resource "aws_s3_bucket" "name" {
  bucket        = local.bucket_name
  force_destroy = true
}
