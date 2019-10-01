provider "gardener" {
  profile              = "<my-gardener-project>"
  azure_secret_binding = "<my-gardener-azure-secret>"
  kube_file          = "${file("<my-gardener-service-account-kubeconfig>")}"
  /*kube_file          =<<-EOT
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

    EOT*/
}

resource "gardener_azure_shoot" "my-server" {
  name              = "tf-gardener-azure"
  region            = "westeurope"
  kubernetesversion = "1.15.2"
  vnetcidr          = "10.250.0.0/16"
  workercidr        = "10.250.0.0/22"
  worker {
    name           = "cpu-worker"
    machinetype    = "Standard_D2_v3"
    autoscalermin  = 2
    autoscalermax  = 4
    maxsurge       = 1
    maxunavailable = 0
    volumesize     = "35Gi"
    volumetype     = "standard"
  }
}
