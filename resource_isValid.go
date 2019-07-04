package main

import (
	"fmt"
	"strconv"

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
				Type:     schema.TypeMap,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(map[string]interface{})
					var errorMessage = "Not Valid"

					// Get assertValue (string)
					assertValue, assertKeyExists := v["assert"]
					if !assertKeyExists {
						errs = append(errs, fmt.Errorf("'test' map must contain an 'assert' key"))
						return
					}

					// Parse assert value (string -> bool)
					assert, err := strconv.ParseBool(assertValue.(string))
					if err != nil {
						errs = append(errs, fmt.Errorf("Your assert must be a bool: %v", assertValue))
						return
					}

					// Get optional error message
					if x, ok := v["error_message"]; ok {
						errorMessage = x.(string)
					}

					// Check assertion
					if !assert {
						errs = append(errs, fmt.Errorf(errorMessage))
					}
					return
				},
			},
		},
	}
}

func resourceIsValidCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	test := d.Get("test").(map[string]interface{})
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
