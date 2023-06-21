locals {
  bucket_name = var.stage == "prod" ? "${var.project}-apt-repository" : "${var.stage}-${var.project}-apt-repository"
}
