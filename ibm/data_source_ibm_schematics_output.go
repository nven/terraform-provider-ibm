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
	"github.com/IBM/cloud-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func dataSourceIBMSchematicsOutput() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSchematicsOutputRead,

		Schema: map[string]*schema.Schema{
			"w_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API.",
			},
		},
	}
}

func dataSourceIBMSchematicsOutputRead(d *schema.ResourceData, meta interface{}) error {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}

	getWorkspaceOutputsOptions := &schematicsv1.GetWorkspaceOutputsOptions{}

	getWorkspaceOutputsOptions.SetWID(d.Get("w_id").(string))

	outputValues, response, err := schematicsClient.GetWorkspaceOutputs(getWorkspaceOutputsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceOutputs failed %s\n%s", err, response)
		return err
	}

	d.SetId(dataSourceIBMSchematicsOutputID(d))

	return nil
}

// dataSourceIBMSchematicsOutputID returns a reasonable ID for the list.
func dataSourceIBMSchematicsOutputID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
