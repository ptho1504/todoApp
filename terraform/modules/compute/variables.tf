variable "ami" {
  description = "AMI ID"
}

variable "instance_type" {
  default = "t3.micro"
}

variable "subnet_id" {
  description = "Private subnet ID"
}

variable "security_group_id" {
  description = "EC2 security group"
}

variable "security_ssm_group_id" {
  description = "EC2 security group"
}


variable "instance_profile_name" {
  description = "Profile Name"
}