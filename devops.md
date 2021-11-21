# DevOps Challenge

## Introduction

At Prophet we manage high-volume trading systems with minimal downtime windows for maintenance. You need significant experience with designing high-available system architectures and deploying them on cloud infrastructure with terraform.

## Requirements

1. We value a **clean**, **simple**, working solution.
2. The project must run with `terraform apply`. If you need additional `kubectl apply` steps for YAML files this is fine.
3. You should use the latest version of terraform and should have documentation for your project.
4. Candidates must submit the project as a git repository (github.com, bitbucket.com, gitlab.com). The repository must avoid containing the words `prophet`, `betprophet` and `challenge`.
5. Having tests is a strong bonus.
6. The solution must be production ready.

## Problem Statement

1. Build a terraform project to launch a Kubernetes cluster (either bootstrapped or leveraging EKS or DigitalOcean K8S is fine)
2. Build deployments, services, ingresses and other things you think are necessary to run a production-ready instance of Ghost (ghost.org/docs) (You can use the existing Ghost Dockerfile, no need to re-invent the wheel.)

## Bonus Points

1. Highly available database layer.
2. Provisioning object storage for Ghost through terraform.
3. Centralised logging.
4. Anything security related you think should be added.

Questions? We love to answer: techchallenge@betprophet.co
