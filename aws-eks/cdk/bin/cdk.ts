#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "aws-cdk-lib";
import { Config } from "../lib/config";
import { JK8SAwsCdkStack } from "@jdotw/jk8s-aws-cdk";

const dev: Config = {
  fqdn: "api.dev.syrupstock.com",
};

const prod: Config = {
  fqdn: "api.syrupstock.com",
};

const app = new cdk.App({
  context: {
    dev,
    prod,
  },
});

const configName = app.node.tryGetContext("config") || "prod";
const config = app.node.tryGetContext(configName) as Config;

const jk8s = new JK8SAwsCdkStack(app, "Vanilla", {
  name: "isvanilla",
  fqdn: config.fqdn,
});
