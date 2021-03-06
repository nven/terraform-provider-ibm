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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func dataSourceIBMSchematicsState() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSchematicsStateRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the workspace for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API.",
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value.",
			},
			"state_store": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_store_json": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMSchematicsStateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{}

	getWorkspaceTemplateStateOptions.SetWID(d.Get("workspace_id").(string))
	getWorkspaceTemplateStateOptions.SetTID(d.Get("template_id").(string))

	_, response, err := schematicsClient.GetWorkspaceTemplateStateWithContext(context, getWorkspaceTemplateStateOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceTemplateStateWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(dataSourceIBMSchematicsStateID(d))

	var stateStore map[string]interface{}
	json.Unmarshal(response.Result.(json.RawMessage), &stateStore)

	b := bytes.NewReader(response.Result.(json.RawMessage))

	decoder := json.NewDecoder(b)
	decoder.UseNumber()
	decoder.Decode(&stateStore)

	statestr := fmt.Sprintf("%v", stateStore)
	d.Set("state_store", statestr)

	stateByte, err := json.MarshalIndent(stateStore, "", "")
	if err != nil {
		return diag.FromErr(err)
	}

	stateStoreJSON := string(stateByte[:])
	d.Set("state_store_json", stateStoreJSON)

	return nil
}

// dataSourceIBMSchematicsStateID returns a reasonable ID for the list.
func dataSourceIBMSchematicsStateID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
