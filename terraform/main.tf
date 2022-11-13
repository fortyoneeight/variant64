provider "aws" {
  region = "us-east-2"
}

# Container Repositories
resource "aws_ecr_repository" "variant64-server" {
  name = "variant64/server"
}

# ECS Cluster and Services
resource "aws_ecs_cluster" "variant64" {
  name = "variant64"
}

resource "aws_ecs_service" "server" {
  name                               = "server"
  cluster                            = aws_ecs_cluster.variant64.id
  task_definition                    = aws_ecs_task_definition.server.arn
  launch_type                        = "FARGATE"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 0

  network_configuration {
    subnets          = ["${aws_subnet.subnet-public.id}"]
    assign_public_ip = true
    security_groups  = ["${aws_security_group.security-group-server.id}"]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.loadbalancer-target-group.id
    container_name   = "server"
    container_port   = "8000"
  }

  depends_on = [aws_lb_listener.loadbalancer-listener-server]
}

# Server Task Definition
resource "aws_ecs_task_definition" "server" {
  family                   = "server"
  container_definitions    = <<DEFINITION
  [
    {
      "name": "server",
      "image": "${aws_ecr_repository.variant64-server.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 8000,
          "hostPort": 8000
        }
      ],
      "memory": 512,
      "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = aws_iam_role.role-ecs-task-execution.arn
}

# Load Balancer Defintions
resource "aws_lb" "loadbalancer-server" {
  name               = "loadbalancer-server"
  load_balancer_type = "network"
  subnets            = [aws_subnet.subnet-public.id]
}

resource "aws_lb_listener" "loadbalancer-listener-server" {
  load_balancer_arn = aws_lb.loadbalancer-server.arn
  protocol          = "TCP"
  port              = "80"

  default_action {
    target_group_arn = aws_lb_target_group.loadbalancer-target-group.arn
    type             = "forward"
  }
}

resource "aws_lb_target_group" "loadbalancer-target-group" {
  port        = "80"
  protocol    = "TCP"
  vpc_id      = aws_vpc.vpc-main.id
  target_type = "ip"

  health_check {
    protocol            = "TCP"
    interval            = "30"
    healthy_threshold   = 5
    unhealthy_threshold = 5
  }
}

resource "aws_vpc_endpoint_service" "endpoint-server" {
  acceptance_required        = false
  network_load_balancer_arns = [aws_lb.loadbalancer-server.arn]
}

# VPC Definitions
resource "aws_vpc" "vpc-main" {
  cidr_block       = "10.0.0.0/24"
  instance_tenancy = "default"
}

resource "aws_internet_gateway" "gateway-main" {
  vpc_id = aws_vpc.vpc-main.id
}

resource "aws_subnet" "subnet-public" {
  vpc_id     = aws_vpc.vpc-main.id
  cidr_block = "10.0.0.128/26"
}

resource "aws_subnet" "subnet-private" {
  vpc_id     = aws_vpc.vpc-main.id
  cidr_block = "10.0.0.192/26"
}

resource "aws_route_table" "routetable-public" {
  vpc_id = aws_vpc.vpc-main.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gateway-main.id
  }
}

resource "aws_route_table" "routetable-private" {
  vpc_id = aws_vpc.vpc-main.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat-gateway.id
  }
}

resource "aws_route_table_association" "routetable-association-public" {
  subnet_id      = aws_subnet.subnet-public.id
  route_table_id = aws_route_table.routetable-public.id
}

resource "aws_route_table_association" "routetable-association-private" {
  subnet_id      = aws_subnet.subnet-private.id
  route_table_id = aws_route_table.routetable-private.id
}

resource "aws_eip" "main-eip" {
  vpc = true
}

resource "aws_nat_gateway" "nat-gateway" {
  allocation_id = aws_eip.main-eip.id
  subnet_id     = aws_subnet.subnet-public.id
}

resource "aws_route53_zone" "zone-private" {
  name = "variant64.xyz"

  vpc {
    vpc_id = aws_vpc.vpc-main.id
  }
}

resource "aws_route53_record" "record-server" {
  zone_id = aws_route53_zone.zone-private.zone_id
  name    = "server.variant64.xyz"
  type    = "CNAME"
  ttl     = 172800
  records = [aws_lb.loadbalancer-server.dns_name]
}

# Security Group Definitions
resource "aws_security_group" "security-group-vpc" {
  vpc_id = aws_vpc.vpc-main.id
}

resource "aws_security_group_rule" "rule-ingress-80" {
  security_group_id = aws_security_group.security-group-vpc.id
  from_port         = "80"
  to_port           = "80"
  protocol          = "tcp"
  type              = "ingress"
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "rule-ingress-8000" {
  security_group_id = aws_security_group.security-group-vpc.id
  from_port         = "8000"
  to_port           = "8000"
  protocol          = "tcp"
  type              = "ingress"
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group" "security-group-server" {
  vpc_id = aws_vpc.vpc-main.id

  ingress {
    from_port   = 8000
    to_port     = 8000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# IAM Policies
resource "aws_iam_role" "role-ecs-task-execution" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.policy-assume-role.json
}

data "aws_iam_policy_document" "policy-assume-role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "role-ecs-task-execution-policy" {
  role       = aws_iam_role.role-ecs-task-execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}
