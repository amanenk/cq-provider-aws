ALTER TABLE "aws_ec2_instance_network_interface_ipv6_addresses" RENAME COLUMN "instance_network_interface_cq_id" TO "instance_network_interface_id";

ALTER TABLE "aws_ec2_instance_network_interface_private_ip_addresses" RENAME COLUMN "instance_network_interface_cq_id" TO "instance_network_interface_id";
