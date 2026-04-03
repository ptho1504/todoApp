resource "aws_instance" "app" {
  ami                    = var.ami
  instance_type          = var.instance_type
  subnet_id              = var.subnet_id
  vpc_security_group_ids = [var.security_group_id]

  associate_public_ip_address = false


  user_data = <<-EOF
            #!/bin/bash
            yum update -y
            yum install -y docker
            systemctl start docker
            systemctl enable docker

            docker run -d -p 80:80 nginx
            EOF

  tags = {
    Name = "app-ec2"
  }
}
