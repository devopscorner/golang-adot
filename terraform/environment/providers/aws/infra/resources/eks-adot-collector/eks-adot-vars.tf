# ==========================================================================
#  Resources: EKS / eks-vars.tf (Spesific Environment)
# --------------------------------------------------------------------------
#  Description
# --------------------------------------------------------------------------
#    - Input Variable for Environment Variables
# ==========================================================================

# ------------------------------------
#  AWS Zone
# ------------------------------------
variable "aws_az" {
  type        = map(string)
  description = "AWS Zone Target Deployment"
  default = {
    lab     = "us-west-2a"
    staging = "us-west-2b"
    prod    = "us-west-2b"
  }
}

# ------------------------------------
#  EKS Cluster
# ------------------------------------
# PEM File from existing
variable "eks_cluster_name" {
  type        = string
  description = "default cluster name"
  default     = "devopscorner"
}
