package seed

import (
	"time"

	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	defaultCreateTimeout = time.Minute * 30
	defaultUpdateTimeout = time.Minute * 30
	defaultDeleteTimeout = time.Minute * 20
)

// Shoot mutex prevents concurrent writes to the CRD
var shootMutex = mutexkv.NewMutexKV()

func ResourceSeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		// Exists: resourceServerExists,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceServerImport,
		// },
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultCreateTimeout),
			Update: schema.DefaultTimeout(defaultUpdateTimeout),
			Delete: schema.DefaultTimeout(defaultDeleteTimeout),
		},
		Schema: map[string]*schema.Schema{
			"metadata": seedMetadataSchema("seed", false),
			"spec":     seedSpecSchema(),
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
