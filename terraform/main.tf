provider "aws" {
  region = "ap-southeast-1"
}


module "network" {
  source = "./modules/network"

  vpc_cidr_block    = "10.0.0.0/16"
  private_subnet    = ["10.0.1.0/24", "10.0.2.0/24"]
  public_subnet     = ["10.0.4.0/24", "10.0.5.0/24"]
  availability_zone = ["ap-southeast-1a", "ap-southeast-1b"]

}


module "security" {
  source = "./modules/security"
  vpc_id = module.network.vpc_id
}

module "database" {
  source            = "./modules/database"
  subnet_ids        = module.network.private_subnet_ids
  security_group_id = module.security.rds_sg_id

  db_name  = "testdb"
  username = "admin"
  password = "password"
}


module "compute" {
  source            = "./modules/compute"
  security_group_id = module.security.ec2_sg_id
  security_ssm_group_id = module.security.ec2_sg_ssm_id
  ami               = "ami-0c0292c4186d3d1ec"
  subnet_id         = module.network.private_subnet_ids[0]
  instance_profile_name = module.iam.instance_profile_name
}

module "alb" {
  source = "./modules/loadbalancer"

  subnet_ids        = module.network.public_subnet_ids
  security_group_id = module.security.alb_sg_id
  vpc_id            = module.network.vpc_id
  instance_id       = module.compute.instance_id
}

module "iam" {
  source = "./modules/iam"
}