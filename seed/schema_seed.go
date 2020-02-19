package seed

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dNSResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ingressDomain": {
				Type:        schema.TypeString,
				Description: "Domain is the external available domain of the Seed cluster.",
				Optional:    true,
			},
		},
	}
}

func seedSpecSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "ShootSpec is the specification of a Shoot. (see https://github.com/gardener/gardener/blob/master/pkg/apis/garden/v1beta1/types.go#L609)",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dns": {
					Type:        schema.TypeList,
					Description: "DNS contains information about the DNS settings of the Seed.",
					Optional:    true,
					MaxItems:    1,
					Elem:        dNSResource(),
				},
			},
		},
	}
}
