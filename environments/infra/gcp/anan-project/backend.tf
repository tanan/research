terraform {
  backend "gcs" {
    bucket = "anan-tfstate"
    prefix = "terraform/state"
  }
}