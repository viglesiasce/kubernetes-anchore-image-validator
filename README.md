# Anchore Image Validator for Kubernetes

## THIS IS NOT AN OFFICIAL GOOGLE PRODUCT

## Intro

[Anchore Engine](https://github.com/anchore/anchore-engine) provides a mechanism to scan Docker images and then evaluate them against
a set of policies. This evaluation result can be used to gate a CI pipeline or, as used in this repo,
to gate the deployment of an image into a Kubernetes cluster.

![Anchore Image Validator Architecture](docs/img/anchore-validator-arch.png)

This repository contains a server that can be used as a [Validating Webhook](https://kubernetes.io/docs/admin/admission-controllers/#validatingadmissionwebhook-alpha-in-18-beta-in-19)
in your Kubernetes cluster. After its been configured, Kubernetes will send a request to this server any time a Pod is requested.
The server will get container images out of the PodSpec and check them against the Anchore Engine API to see if they
adhere to the policy that has been defined. If the image does not yet exist in Anchore Engine it will automatically be added
and scanned. The default policy validates that there are no critical security vulnerabilities in the image.

## Quick Start

1. [Install Helm](https://github.com/kubernetes/helm/blob/master/docs/install.md)

1. Run hack/install.sh

1. Follow the instructions output by the chart installation for installing the validating web hook.