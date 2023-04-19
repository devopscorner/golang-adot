# ==========================================================================
#  Resources: CICD DynamoDB / variables.tf (Global Environment)
# --------------------------------------------------------------------------
#  Description
# --------------------------------------------------------------------------
#    - Input Variable for Environment Variables
# ==========================================================================

# ------------------------------------
#  AWS Region
# ------------------------------------
variable "aws_region" {
  type        = string
  description = "AWS Region Target Deployment"
  default     = "us-west-2"
}

# ------------------------------------
#  Workspace
# ------------------------------------
variable "env" {
  type        = map(string)
  description = "Workspace Environment Selection"
  default = {
    lab     = "lab"
    staging = "staging"
    prod    = "prod"
  }
}

# ------------------------------------
#  Environment Resources Tags
# ------------------------------------
variable "environment" {
  type        = map(string)
  description = "Target Environment (tags)"
  default = {
    lab     = "RND"
    staging = "STG"
    prod    = "PROD"
  }
}

# ------------------------------------
#  Department Tags
# ------------------------------------
variable "department" {
  type        = string
  description = "Department Owner"
  default     = "DEVOPS"
}

# ------------------------------------
#  KMS Key
# ------------------------------------
variable "kms_key" {
  type        = map(string)
  description = "KMS Key References"
  default = {
    lab     = "arn:aws:kms:us-west-2:YOUR_AWS_ACCOUNT:key/HASH_KEY"
    staging = "arn:aws:kms:us-west-2:YOUR_AWS_ACCOUNT:key/HASH_KEY"
    prod    = "arn:aws:kms:us-west-2:YOUR_AWS_ACCOUNT:key/HASH_KEY"
  }
}

# ------------------------------------
#  KMS Environment
# ------------------------------------
variable "kms_env" {
  type        = map(string)
  description = "KMS Key Environment"
  default = {
    lab     = "RnD"
    staging = "Staging"
    prod    = "Production"
  }
}

# ------------------------------------
#  DynamoDB Go App
# ------------------------------------
variable "goapp_dynamodb_table" {
  type        = string
  description = "Name of dynamodb table to store golang app"
  default     = "dynamodb-golang-adot"
}

# ------------------------------------
#  S3 Bucket Name
# ------------------------------------
variable "bucket_name" {
  type        = string
  description = "CICD Bucket Name"
  default     = "devopscorner-goapp"
}

# ------------------------------------
#  Bucket Terraform State
# ------------------------------------
variable "tfstate_encrypt" {
  type        = bool
  description = "Name of bucket to store tfstate"
  default     = true
}

variable "tfstate_bucket" {
  type        = string
  description = "Name of bucket to store tfstate"
  default     = "devopscorner-adot-remote-state"
}

variable "tfstate_dynamodb_table" {
  type        = string
  description = "Name of dynamodb table to store tfstate"
  default     = "devopscorner-adot-state-lock"
}

variable "tfstate_path" {
  type        = string
  description = "Path .tfstate in Bucket"
  default     = "resources/cicd/terraform.tfstate"
}
