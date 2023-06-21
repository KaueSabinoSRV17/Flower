module "apt-repository-bucket" {
  source = "./modules/apt_repository_bucket"

  stage   = var.stage
  project = var.project
}
