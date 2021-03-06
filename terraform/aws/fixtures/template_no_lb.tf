resource "aws_eip" "jumpbox_eip" {
  depends_on = ["aws_internet_gateway.ig"]
  vpc        = true
}

resource "tls_private_key" "bosh_vms" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "bosh_vms" {
  key_name   = "${var.env_id}_bosh_vms"
  public_key = "${tls_private_key.bosh_vms.public_key_openssh}"
}

output "bosh_vms_key_name" {
  value = "${aws_key_pair.bosh_vms.key_name}"
}

output "bosh_vms_private_key" {
  value     = "${tls_private_key.bosh_vms.private_key_pem}"
  sensitive = true
}

output "external_ip" {
  value = "${aws_eip.jumpbox_eip.public_ip}"
}

output "jumpbox_url" {
  value = "${aws_eip.jumpbox_eip.public_ip}:22"
}

output "director_address" {
  value = "https://${aws_eip.jumpbox_eip.public_ip}:25555"
}

resource "aws_iam_role" "bosh" {
  name = "${var.env_id}_bosh_role"
  path = "/"

  lifecycle {
    create_before_destroy = true
  }

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "bosh" {
  name = "${var.env_id}_bosh_policy"
  path = "/"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:AssociateAddress",
        "ec2:AttachVolume",
        "ec2:CopyImage",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteVolume",
        "ec2:DescribeAddresses",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeRegions",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSnapshots",
        "ec2:DescribeSubnets",
        "ec2:DescribeVolumes",
        "ec2:DetachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:ModifyInstanceAttribute",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "ec2:RegisterImage",
        "ec2:DeregisterImage"
	  ],
	  "Effect": "Allow",
	  "Resource": "*"
    },
	{
	  "Action": [
	    "iam:PassRole"
	  ],
	  "Effect": "Allow",
	  "Resource": "*"
	},
	{
	  "Action": [
	    "elasticloadbalancing:*"
	  ],
	  "Effect": "Allow",
	  "Resource": "*"
	}
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "bosh" {
  role       = "${var.env_id}_bosh_role"
  policy_arn = "${aws_iam_policy.bosh.arn}"
}

resource "aws_iam_instance_profile" "bosh" {
  name = "${var.env_id}-bosh"
  role = "${aws_iam_role.bosh.name}"

  lifecycle {
    ignore_changes = ["name"]
  }
}

output "bosh_iam_instance_profile" {
  value = "${aws_iam_instance_profile.bosh.name}"
}

variable "nat_ami_map" {
  type = "map"

  default = {
    ap-northeast-1 = "ami-10dfc877"
    ap-northeast-2 = "ami-1a1bc474"
    ap-south-1     = "ami-74c1861b"
    ap-southeast-1 = "ami-36af2055"
    ap-southeast-2 = "ami-1e91817d"
    ca-central-1   = "ami-12d36a76"
    eu-central-1   = "ami-9ebe18f1"
    eu-west-1      = "ami-3a849f5c"
    eu-west-2      = "ami-21120445"
    us-east-1      = "ami-d4c5efc2"
    us-east-2      = "ami-f27b5a97"
    us-gov-west-1  = "ami-c39610a2"
    us-west-1      = "ami-b87f53d8"
    us-west-2      = "ami-8bfce8f2"
  }
}

resource "aws_security_group" "nat_security_group" {
  name        = "${var.env_id}-nat-security-group"
  description = "NAT"
  vpc_id      = "${local.vpc_id}"

  tags {
    Name = "${var.env_id}-nat-security-group"
  }

  lifecycle {
    ignore_changes = ["name"]
  }
}

resource "aws_security_group_rule" "nat_to_internet_rule" {
  security_group_id = "${aws_security_group.nat_security_group.id}"

  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "nat_icmp_rule" {
  security_group_id = "${aws_security_group.nat_security_group.id}"

  type        = "ingress"
  protocol    = "icmp"
  from_port   = -1
  to_port     = -1
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "nat_tcp_rule" {
  security_group_id = "${aws_security_group.nat_security_group.id}"

  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.internal_security_group.id}"
}

resource "aws_security_group_rule" "nat_udp_rule" {
  security_group_id = "${aws_security_group.nat_security_group.id}"

  type                     = "ingress"
  protocol                 = "udp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.internal_security_group.id}"
}

resource "aws_instance" "nat" {
  private_ip             = "${cidrhost(aws_subnet.bosh_subnet.cidr_block, 7)}"
  instance_type          = "t2.medium"
  subnet_id              = "${aws_subnet.bosh_subnet.id}"
  source_dest_check      = false
  ami                    = "${lookup(var.nat_ami_map, var.region)}"
  vpc_security_group_ids = ["${aws_security_group.nat_security_group.id}"]

  tags {
    Name  = "${var.env_id}-nat"
    EnvID = "${var.env_id}"
  }
}

resource "aws_eip" "nat_eip" {
  depends_on = ["aws_internet_gateway.ig"]
  instance   = "${aws_instance.nat.id}"
  vpc        = true
}

output "nat_eip" {
  value = "${aws_eip.nat_eip.public_ip}"
}

variable "access_key" {
  type = "string"
}

variable "secret_key" {
  type = "string"
}

variable "region" {
  type = "string"
}

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region     = "${var.region}"
}

resource "aws_default_security_group" "default_security_group" {
  vpc_id = "${local.vpc_id}"
}

resource "aws_security_group" "internal_security_group" {
  name        = "${var.env_id}-internal-security-group"
  description = "Internal"
  vpc_id      = "${local.vpc_id}"

  tags {
    Name = "${var.env_id}-internal-security-group"
  }

  lifecycle {
    ignore_changes = ["name"]
  }
}

resource "aws_security_group_rule" "internal_security_group_rule_tcp" {
  security_group_id = "${aws_security_group.internal_security_group.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 0
  to_port           = 65535
  self              = true
}

resource "aws_security_group_rule" "internal_security_group_rule_udp" {
  security_group_id = "${aws_security_group.internal_security_group.id}"
  type              = "ingress"
  protocol          = "udp"
  from_port         = 0
  to_port           = 65535
  self              = true
}

resource "aws_security_group_rule" "internal_security_group_rule_icmp" {
  security_group_id = "${aws_security_group.internal_security_group.id}"
  type              = "ingress"
  protocol          = "icmp"
  from_port         = -1
  to_port           = -1
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "internal_security_group_rule_allow_internet" {
  security_group_id = "${aws_security_group.internal_security_group.id}"
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "internal_security_group_rule_ssh" {
  security_group_id        = "${aws_security_group.internal_security_group.id}"
  type                     = "ingress"
  protocol                 = "TCP"
  from_port                = 22
  to_port                  = 22
  source_security_group_id = "${aws_security_group.jumpbox.id}"
}

output "internal_security_group" {
  value = "${aws_security_group.internal_security_group.id}"
}

variable "bosh_inbound_cidr" {
  default = "0.0.0.0/0"
}

resource "aws_security_group" "bosh_security_group" {
  name        = "${var.env_id}-bosh-security-group"
  description = "BOSH Director"
  vpc_id      = "${local.vpc_id}"

  tags {
    Name = "${var.env_id}-bosh-security-group"
  }

  lifecycle {
    ignore_changes = ["name", "description"]
  }
}

output "bosh_security_group" {
  value = "${aws_security_group.bosh_security_group.id}"
}

resource "aws_security_group_rule" "bosh_security_group_rule_tcp_ssh" {
  security_group_id = "${aws_security_group.bosh_security_group.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 22
  to_port           = 22
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "bosh_security_group_rule_tcp_bosh_agent" {
  security_group_id = "${aws_security_group.bosh_security_group.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 6868
  to_port           = 6868
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "bosh_security_group_rule_uaa" {
  security_group_id = "${aws_security_group.bosh_security_group.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 8443
  to_port           = 8443
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "bosh_security_group_rule_tcp_director_api" {
  security_group_id = "${aws_security_group.bosh_security_group.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 25555
  to_port           = 25555
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "bosh_security_group_rule_tcp" {
  security_group_id        = "${aws_security_group.bosh_security_group.id}"
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.internal_security_group.id}"
}

resource "aws_security_group_rule" "bosh_security_group_rule_udp" {
  security_group_id        = "${aws_security_group.bosh_security_group.id}"
  type                     = "ingress"
  protocol                 = "udp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.internal_security_group.id}"
}

resource "aws_security_group_rule" "bosh_security_group_rule_allow_internet" {
  security_group_id = "${aws_security_group.bosh_security_group.id}"
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group" "jumpbox" {
  name        = "${var.env_id}-jumpbox-security-group"
  description = "Jumpbox"
  vpc_id      = "${local.vpc_id}"

  tags {
    Name = "${var.env_id}-jumpbox-security-group"
  }

  lifecycle {
    ignore_changes = ["name", "description"]
  }
}

output "jumpbox_security_group" {
  value = "${aws_security_group.jumpbox.id}"
}

resource "aws_security_group_rule" "jumpbox_ssh" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 22
  to_port           = 22
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "jumpbox_rdp" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 3389
  to_port           = 3389
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "jumpbox_agent" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 6868
  to_port           = 6868
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "jumpbox_credhub" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 8844
  to_port           = 8844
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "jumpbox_director" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 25555
  to_port           = 25555
  cidr_blocks       = ["${var.bosh_inbound_cidr}"]
}

resource "aws_security_group_rule" "jumpbox_egress" {
  security_group_id = "${aws_security_group.jumpbox.id}"
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "bosh_internal_security_rule_tcp" {
  security_group_id        = "${aws_security_group.internal_security_group.id}"
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.bosh_security_group.id}"
}

resource "aws_security_group_rule" "bosh_internal_security_rule_udp" {
  security_group_id        = "${aws_security_group.internal_security_group.id}"
  type                     = "ingress"
  protocol                 = "udp"
  from_port                = 0
  to_port                  = 65535
  source_security_group_id = "${aws_security_group.bosh_security_group.id}"
}

output "internal_cidr" {
  value = "${aws_subnet.bosh_subnet.cidr_block}"
}

resource "aws_subnet" "bosh_subnet" {
  vpc_id     = "${local.vpc_id}"
  cidr_block = "${cidrsubnet(var.vpc_cidr, 8, 0)}"

  tags {
    Name = "${var.env_id}-bosh-subnet"
  }
}

resource "aws_route_table" "bosh_route_table" {
  vpc_id = "${local.vpc_id}"
}

resource "aws_route" "bosh_route_table" {
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.ig.id}"
  route_table_id         = "${aws_route_table.bosh_route_table.id}"
}

resource "aws_route_table_association" "route_bosh_subnets" {
  subnet_id      = "${aws_subnet.bosh_subnet.id}"
  route_table_id = "${aws_route_table.bosh_route_table.id}"
}

output "bosh_subnet_id" {
  value = "${aws_subnet.bosh_subnet.id}"
}

output "bosh_subnet_availability_zone" {
  value = "${aws_subnet.bosh_subnet.availability_zone}"
}

variable "availability_zones" {
  type = "list"
}

resource "aws_subnet" "internal_subnets" {
  count             = "${length(var.availability_zones)}"
  vpc_id            = "${local.vpc_id}"
  cidr_block        = "${cidrsubnet(var.vpc_cidr, 4, count.index+1)}"
  availability_zone = "${element(var.availability_zones, count.index)}"

  tags {
    Name = "${var.env_id}-internal-subnet${count.index}"
  }

  lifecycle {
    ignore_changes = ["cidr_block", "availability_zone"]
  }
}

resource "aws_route_table" "internal_route_table" {
  vpc_id = "${local.vpc_id}"
}

resource "aws_route" "internal_route_table" {
  destination_cidr_block = "0.0.0.0/0"
  instance_id            = "${aws_instance.nat.id}"
  route_table_id         = "${aws_route_table.internal_route_table.id}"
}

resource "aws_route_table_association" "route_internal_subnets" {
  count          = "${length(var.availability_zones)}"
  subnet_id      = "${element(aws_subnet.internal_subnets.*.id, count.index)}"
  route_table_id = "${aws_route_table.internal_route_table.id}"
}

output "internal_az_subnet_id_mapping" {
  value = "${
	  zipmap("${aws_subnet.internal_subnets.*.availability_zone}", "${aws_subnet.internal_subnets.*.id}")
	}"
}

output "internal_az_subnet_cidr_mapping" {
  value = "${
	  zipmap("${aws_subnet.internal_subnets.*.availability_zone}", "${aws_subnet.internal_subnets.*.cidr_block}")
	}"
}

variable "env_id" {
  type = "string"
}

variable "short_env_id" {
  type = "string"
}

variable "vpc_cidr" {
  type    = "string"
  default = "10.0.0.0/16"
}

resource "aws_internet_gateway" "ig" {
  vpc_id = "${local.vpc_id}"
}

output "vpc_id" {
  value = "${local.vpc_id}"
}

resource "aws_flow_log" "bbl" {
  log_group_name = "${aws_cloudwatch_log_group.bbl.name}"
  iam_role_arn   = "${aws_iam_role.flow_logs.arn}"
  vpc_id         = "${local.vpc_id}"
  traffic_type   = "REJECT"
}

resource "aws_cloudwatch_log_group" "bbl" {
  name_prefix = "${var.short_env_id}-log-group"
}

resource "aws_iam_role" "flow_logs" {
  name = "${var.env_id}-flow-logs-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "vpc-flow-logs.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "flow_logs" {
  name = "${var.env_id}-flow-logs-policy"
  role = "${aws_iam_role.flow_logs.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_kms_key" "kms_key" {
  enable_key_rotation = true
}

output "kms_key_arn" {
  value = "${aws_kms_key.kms_key.arn}"
}

variable "existing_vpc_id" {
  type        = "string"
  default     = ""
  description = "Optionally use an existing vpc"
}

locals {
  vpc_count = "${length(var.existing_vpc_id) > 0 ? 0 : 1}"
  vpc_id    = "${length(var.existing_vpc_id) > 0 ? var.existing_vpc_id : join(" ", aws_vpc.vpc.*.id)}"
}

resource "aws_vpc" "vpc" {
  count                = "${local.vpc_count}"
  cidr_block           = "${var.vpc_cidr}"
  instance_tenancy     = "default"
  enable_dns_hostnames = true

  tags {
    Name = "${var.env_id}-vpc"
  }
}
