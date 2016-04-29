package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/lord2800/proxmox-api"
)

func main() {
    plugin.Serve(new(Proxmox))
}

func Proxmox() schema.ResourceProvider {
  return &schema.Provider{
    Schema:        map[string]*schema.Schema{
        "host": &schema.Schema{
            Type: schema.TypeString,
            Required: true,
            Description: "The URL to use as the base of the api"
        },
        "user": &schema.Schema{
            Type: schema.TypeString,
            Required: true,
            Description: "The user to authenticate as"
        },
        "pass": &schema.Schema{
            Type: schema.TypeString,
            Required: true,
            Description: "The password for the specified user"
        }
    },
    ResourcesMap:  map[string]*schema.Resource{
        "proxmox_vm": &schema.Resource{
            Schema: map[string]*schema.Schema{
                "name": &schema.Schema{
                    Type: schema.TypeString,
                    Required: true,
                    Description: "The name of the machine"
                },
                "id": &schema.Schema{
                    Type: schema.TypeString,
                    Required: false,
                    Description: "The id of the machine"
                }
            }
            SchemaVersion: 1,
            Create: func(d *schema.ResourceData, meta interface{}) {},
            Read: func(d *schema.ResourceData, meta interface{}) {},
            Update: func(d *schema.ResourceData, meta interface{}) {},
            Delete: func(d *schema.ResourceData, meta interface{}) {},
        }
    },
    ConfigureFunc: func(*schema.ResourceData) (interface{}, error) {},
  }
}
