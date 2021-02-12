/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func dataSourceIBMSchematicsState() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSchematicsStateRead,

		Schema: map[string]*schema.Schema{
			"w_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The workspace ID for the workspace that you want to query.  You can run the GET /workspaces call if you need to look up the  workspace IDs in your IBM Cloud account.",
			},
			"t_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Template ID for which you want to get the values.  Use the GET /workspaces to look up the workspace IDs  or template IDs in your IBM Cloud account.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"terraform_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
			},
			"serial": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"lineage": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
			},
			"modules": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
		},
	}
}


func dataSourceIBMSchematicsStateRead(d *schema.ResourceData, meta interface{}) error {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}

	getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{}


	getWorkspaceTemplateStateOptions.SetWID(d.Get("w_id").(string))

	getWorkspaceTemplateStateOptions.SetTID(d.Get("t_id").(string))

	templateStateStore, response, err := schematicsClient.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceTemplateState failed %s\n%s", err, response)
		return err
	}

	d.SetId(dataSourceIBMSchematicsStateID(d))
	if err = d.Set("version", templateStateStore.Version); err != nil {
		return fmt.Errorf("Error setting version: %s", err)
	}
	if err = d.Set("terraform_version", templateStateStore.TerraformVersion); err != nil {
		return fmt.Errorf("Error setting terraform_version: %s", err)
	}
	if err = d.Set("serial", templateStateStore.Serial); err != nil {
		return fmt.Errorf("Error setting serial: %s", err)
	}
	if err = d.Set("lineage", templateStateStore.Lineage); err != nil {
		return fmt.Errorf("Error setting lineage: %s", err)
	}
	if err = d.Set("modules", templateStateStore.Modules); err != nil {
		return fmt.Errorf("Error setting modules: %s", err)
	}

	return nil
}

// dataSourceIBMSchematicsStateID returns a reasonable ID for the list.
func dataSourceIBMSchematicsStateID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

