# 📦 Full-Stack Todo App — React + Go + MySQL + Terraform on AWS

A simple, well-structured project designed to help you learn **Terraform** and **AWS** fundamentals.

## 🗂️ Project Structure

```
project/
├── frontend/          # React app (runs on EC2 or S3+CloudFront)
├── backend/           # Go REST API (runs on EC2)
├── terraform/         # Infrastructure as Code
│   ├── modules/       # Reusable Terraform modules
│   │   ├── vpc/       # Networking
│   │   ├── ec2/       # App servers
│   │   └── rds/       # MySQL database
│   └── environments/
│       └── dev/       # Dev environment config
└── docs/              # Architecture diagrams & notes
```

## 🚀 Tech Stack

| Layer       | Technology         |
|-------------|-------------------|
| Frontend    | React + Vite       |
| Backend     | Go (Gin framework) |
| Database    | MySQL 8.0          |
| Infra       | Terraform + AWS    |

## 🧑‍💻 Local Development

### 1. Start MySQL
```bash
docker run -d \
  --name mysql-local \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=todos \
  -p 3306:3306 \
  mysql:8.0
```

### 2. Start Backend
```bash
cd backend
cp .env.example .env   # edit DB credentials
go mod tidy
go run main.go
# API runs at http://localhost:8080
```

### 3. Start Frontend
```bash
cd frontend
npm install
npm run dev
# App runs at http://localhost:5173
```

## ☁️ Deploy to AWS with Terraform

```bash
cd terraform/environments/dev
cp terraform.tfvars.example terraform.tfvars   # fill in your values
terraform init
terraform plan
terraform apply
```

## 📚 What You'll Learn

- **VPC** — private networking, subnets, security groups
- **EC2** — virtual machines, user data scripts, key pairs
- **RDS** — managed MySQL, private subnet placement
- **Terraform modules** — reusable infra components
- **Outputs & Variables** — passing values between modules
# todoApp
