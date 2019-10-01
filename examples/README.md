# Gardener Provider Examples

## Overview
This folder contains a set of examples which use Gardener services to deploy [aws](https://github.com/kyma-incubator/terraform-provider-gardener/tree/master/examples/aws), [gcp](https://github.com/kyma-incubator/terraform-provider-gardener/tree/master/examples/gcp) and [azure](https://github.com/kyma-incubator/terraform-provider-gardener/tree/master/examples/azure) clusters. Although the configurations for these providers differ, they have the following common section:

```bash
provider "gardener" {
  profile            = "<my-gardener-project>"
  <provider>_secret_binding = "<my-gardener-<provider>-secret>"
  kube_file          = "${file("<my-gardener-service-account-kubeconfig>")}"
}
```
You can pass the kube_file using the raw text alone as follows:
```bash
kube_path          =<<-EOT
    kind: Config
    clusters:
      - cluster:
          certificate-authority-data: >-
            <certificate-authority-data>
          server: "https://gardener.garden.canary.k8s.ondemand.com"
        name: garden
    users:
      - user:
          token: >-
            <token>
        name: robot
    contexts:
      - context:
          cluster: garden
          user: robot
          namespace: garden-<profile>
        name: garden-<profile>-robot
    current-context: garden-<profile>-robot

    EOT
```

This section includes the following parameters:
* **profile** - the profile you want to deploy to in gardener. 
* **<provider>_secret_binding** - the provider secret binding defined for the profile that you want to use. There might be more than one secret per provider for a profile.
* **kube_file** - the raw string of the kube config file. 

## Installation
Follow these steps to run an example:
1. Clone the `terraform-provider-gardener` repository.
2. Go to `terraform-provider-gardener/examples/{example_name}`.
3. Run  `terraform apply`. 
