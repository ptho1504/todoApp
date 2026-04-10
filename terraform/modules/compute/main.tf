resource "aws_instance" "app" {
  ami                    = var.ami
  instance_type          = var.instance_type
  subnet_id              = var.subnet_id
  vpc_security_group_ids = [var.security_group_id, var.security_ssm_group_id]
  iam_instance_profile = var.instance_profile_name

  associate_public_ip_address = false


  user_data = <<-EOF
            #!/bin/bash
            EOF

  tags = {
    Name = "app-ec2"
  }
}
