provider "gardener" {
  profile            = "<my-gardener-project>"
  gcp_secret_binding = "<my-gardener-gcp-secret>"
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

resource "gardener_gcp_shoot" "my-server" {
  name              = "tf-gardener-gcp"
  region            = "europe-west3"
  zones             = ["europe-west3-b"]
  workerscidr       = ["10.250.0.0/19"]
  kubernetesversion = "1.15.2"
  worker {
    name           = "cpu-worker"
    machinetype    = "n1-standard-4"
    autoscalermin  = 2
    autoscalermax  = 2
    maxsurge       = 1
    maxunavailable = 0
    volumesize     = "20Gi"
    volumetype     = "pd-standard"
  }
}
