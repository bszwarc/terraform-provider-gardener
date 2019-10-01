provider "gardener" {
  profile            = "<my-gardener-project>"
  aws_secret_binding = "<my-gardener-aws-secret>"
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

resource "gardener_aws_shoot" "my-server" {
  name              = "tf-gardener-aws"
  region            = "eu-west-1"
  zones             = ["eu-west-1a"]
  workerscidr       = ["10.250.0.0/17"]
  internalscidr     = ["10.250.112.0/22"]
  publicscidr       = ["10.250.96.0/22"]
  vpccidr = "10.250.0.0/16"
  kubernetesversion = "1.15.2"
  worker {
    name           = "cpu-worker"
    machinetype    = "m5.large"
    autoscalermin  = 2
    autoscalermax  = 2
    maxsurge       = 1
    maxunavailable = 0
    volumesize     = "20Gi"
    volumetype     = "gp2"
  }
}
