package shoot

import (
	"strconv"

	gardener_types "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	"github.com/hashicorp/terraform/helper/schema"
)

const gcp string = "gcp"

func GCPShoot() *schema.Resource {
	return &schema.Resource{
		Create: resourceGCPServerCreate,
		Read:   resourceServerRead,
		Exists: resourceServerExists,
		Update: resourceGCPServerUpdate,
		Delete: resourceServerDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kubernetesversion": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"workerscidr": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"zones": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"worker": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"machinetype": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"autoscalermin": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"autoscalermax": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"maxsurge": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"maxunavailable": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"volumesize": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"volumetype": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
func resourceGCPServerCreate(d *schema.ResourceData, m interface{}) error {
	return resourceServerCreate(d, m, gcp)
}

func resourceGCPServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerUpdate(d, m, gcp)
}

func createGCPSpec(spec gardener_types.ShootSpec, d *schema.ResourceData, secretBinding string) gardener_types.ShootSpec {
	spec.Cloud.SecretBindingRef.Name = secretBinding
	spec.Cloud.GCP = &gardener_types.GCPCloud{
		Networks: gardener_types.GCPNetworks{
			Workers: getCidrs("workerscidr", d),
		},
		Workers: getGCPWorkers(d),
		Zones:   getZones(d),
	}
	return spec
}

func getGCPWorkers(d *schema.ResourceData) []gardener_types.GCPWorker {
	numWorkers := d.Get("worker.#").(int)
	resultWorkers := make([]gardener_types.GCPWorker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		var worker = "worker." + strconv.Itoa(i)
		resultWorkers[i] = gardener_types.GCPWorker{
			Worker:     createGardenWorker(worker, d),
			VolumeSize: d.Get(worker + ".volumesize").(string),
			VolumeType: d.Get(worker + ".volumetype").(string),
		}
	}
	return resultWorkers
}

func updateGCPSpec(d *schema.ResourceData, gcpSpec *gardener_types.GCPCloud) {

	if d.HasChange("workerscidr") {
		gcpSpec.Networks.Workers = getCidrs("workerscidr", d)
	}

	gcpSpec.Workers = getGCPWorkers(d)

	if d.HasChange("zones") {
		gcpSpec.Zones = getZones(d)
	}
}

func flattenGCPWorkers(d *schema.ResourceData, workersarray []gardener_types.GCPWorker) {

	if len(workersarray) > 0 {
		workers := make([]interface{}, len(workersarray))
		for i, v := range workersarray {
			m := map[string]interface{}{}

			if v.Name != "" {
				m["name"] = v.Name
			}
			if v.MachineType != "" {
				m["machine_type"] = v.MachineType
			}
			if v.AutoScalerMin != 0 {
				m["auto_scaler_min"] = v.AutoScalerMin
			}
			if v.AutoScalerMax != 0 {
				m["auto_scaler_max"] = v.AutoScalerMax
			}
			if v.MaxSurge != nil {
				m["max_surge"] = v.MaxSurge.IntValue()
			}
			if v.MaxUnavailable != nil {
				m["max_unavailable"] = v.MaxUnavailable.IntValue()
			}
			if len(v.VolumeType) > 0 {
				m["volume_type"] = v.VolumeType
			}
			if len(v.VolumeSize) > 0 {
				m["volume_size"] = v.VolumeSize
			}
			workers[i] = m
		}
		d.Set("worker", workers)
	}
}
