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
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func dataSourceIBMSchematicsStateRead(d *schema.ResourceData, meta interface{}) error {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}

	getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{}

	getWorkspaceTemplateStateOptions.SetWID(d.Get("w_id").(string))

	getWorkspaceTemplateStateOptions.SetTID(d.Get("t_id").(string))

	_, response, err := schematicsClient.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceTemplateState failed %s\n%s", err, response)
		return err
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
		return err
	}

	stateStoreJSON := string(stateByte[:])
	d.Set("state_store_json", stateStoreJSON)

	return nil
}

// dataSourceIBMSchematicsStateID returns a reasonable ID for the list.
func dataSourceIBMSchematicsStateID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
