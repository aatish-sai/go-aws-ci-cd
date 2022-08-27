## Integrating with GitHub Actions - CICD pipeline to Deploy Go Application to Amazon ECS 
[![Go CICD](https://github.com/aatish-sai/go-aws-ci-cd/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/aatish-sai/go-aws-ci-cd/actions/workflows/main.yml)

This repository uses GitHub Actions to push the docker image to the ECR, update the task definition and deploy the latest revision in the pre-existing ECS service.

## Solution Overview

The solution utilizes following services:

1. GitHub Actions
2. AWS ECR
3. AWS ECS

## Prerequisites

* An AWS account CLI credentials with proper access to ECS and ECR
* ECS cluster and service