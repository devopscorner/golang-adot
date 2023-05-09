# ==========================================================================
#  Resources: EKS / backend.tf (Storing tfstate)
# --------------------------------------------------------------------------
#  Description
# --------------------------------------------------------------------------
#    - S3 Bucket Path
#    - DynamoDB Table
# ==========================================================================

# --------------------------------------------------------------------------
#  Store Path for Terraform State
# --------------------------------------------------------------------------
terraform {
  backend "s3" {
    region         = "us-west-2"
    bucket         = "devopscorner-adot-remote-state"
    dynamodb_table = "devopscorner-adot-state-lock"
    key            = "resources/adot/infra/terraform.tfstate"
    encrypt        = true
  }
}
