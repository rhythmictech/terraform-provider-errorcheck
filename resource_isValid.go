package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIsValid() *schema.Resource {
	return &schema.Resource{
		Create: resourceIsValidCreate,
		Read:   resourceIsValidRead,
		Update: resourceIsValidUpdate,
		Delete: resourceIsValidDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"test": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(bool)
					if v != true {
						errs = append(errs, fmt.Errorf("Not Valid"))
					}
					return
				},
			},
		},
	}
}

func resourceIsValidCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	test := d.Get("test").(string)
	d.SetId(name)
	d.Set("name", name)
	d.Set("test", test)
	return resourceIsValidRead(d, m)
}

func resourceIsValidRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceIsValidUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIsValidRead(d, m)
}

func resourceIsValidDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
