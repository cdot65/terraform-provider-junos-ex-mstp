
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlProtocolsMstpMstiName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_msti  struct {
			XMLName xml.Name `xml:"msti"`
			V_name  *string  `xml:"name,omitempty"`
		} `xml:"protocols>mstp>msti"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsMstpMstiNameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsMstpMstiName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_msti.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsMstpMstiNameRead(d,m)
}

func junosProtocolsMstpMstiNameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsMstpMstiName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_msti.V_name)

	return nil
}

func junosProtocolsMstpMstiNameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsMstpMstiName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_msti.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosProtocolsMstpMstiNameRead(d,m)
}

func junosProtocolsMstpMstiNameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosProtocolsMstpMstiName() *schema.Resource {
	return &schema.Resource{
		Create: junosProtocolsMstpMstiNameCreate,
		Read: junosProtocolsMstpMstiNameRead,
		Update: junosProtocolsMstpMstiNameUpdate,
		Delete: junosProtocolsMstpMstiNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_msti. ",
			},
		},
	}
}