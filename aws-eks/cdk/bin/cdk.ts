#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "aws-cdk-lib";
import * as ec2 from "aws-cdk-lib/aws-ec2";
import { RDSStack } from "../lib/rds-stack";
import { EKSStack } from "../lib/eks-stack";
import { VPCStack } from "../lib/vpc-stack";
import { DNSStack } from "../lib/dns-stack";
import { SecretsStack } from "../lib/secrets-stack";
import { OpenSearchStack } from "../lib/opensearch-stack";

const name = "isvanilla";

const app = new cdk.App();

const vpc = new VPCStack(app, "VPCStack", {});

const secrets = new SecretsStack(app, "SecretsStack", {});

const dns = new DNSStack(app, "DNSStack", {
  name,
});

const rds = new RDSStack(app, "RDSStack", {
  vpc,
});

const cluster = new EKSStack(app, "EKSStack", {
  name,
  vpc,
  rds,
  dns,
});

const opensearch = new OpenSearchStack(app, "OpenSearchStack", { vpc });
