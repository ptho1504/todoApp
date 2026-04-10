provider "aws" {
  region = "ap-southeast-1"
}

resource "aws_vpc" "vpc" {
  cidr_block = var.vpc_cidr_block

  tags = {
    Name = "main"
  }
}

// Private subnet X-A
resource "aws_subnet" "private_subnet" {
  count = length(var.private_subnet)

  vpc_id            = aws_vpc.vpc.id
  cidr_block        = var.private_subnet[count.index]
  availability_zone = var.availability_zone[count.index % length(var.availability_zone)]

  tags = {
    "Name" = "private-subnet"
  }
}

// Public subnet X-A
resource "aws_subnet" "public_subnet" {
  count = length(var.public_subnet)

  vpc_id            = aws_vpc.vpc.id
  cidr_block        = var.public_subnet[count.index]
  availability_zone = var.availability_zone[count.index % length(var.availability_zone)]

  tags = {
    "Name" = "public-subnet"
  }
}


// Internet gateway
resource "aws_internet_gateway" "gateway" {
  vpc_id = aws_vpc.vpc.id

  tags = {
    "Name" = "gateway"
  }
}


// Route table - PUBLIC
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gateway.id
  }

  tags = {
    "Name" = "public"
  }
}

// Association with public
resource "aws_route_table_association" "public_association" {
  for_each       = { for k, v in aws_subnet.public_subnet : k => v }
  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

// EIP
resource "aws_eip" "nat" {
  domain = "vpc"
}

// NAT
resource "aws_nat_gateway" "public" {
  depends_on    = [aws_internet_gateway.gateway]
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public_subnet[0].id

  tags = {
    Name = "public NAT"
  }
}



// Route table - PRIVATE
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gateway.id
  }
}

// Association with private
resource "aws_route_table_association" "private_association" {
  for_each       = { for k, v in aws_subnet.private_subnet : k => v }
  subnet_id      = each.value.id
  route_table_id = aws_route_table.private.id
}
