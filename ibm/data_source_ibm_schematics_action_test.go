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
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIBMSchematicsActionDataSourceBasic(t *testing.T) {
	actionID := fmt.Sprintf("action_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_readme_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "trigger_record_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "namespace"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "playbook_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsActionDataSourceAllArgs(t *testing.T) {
	actionID := fmt.Sprintf("action_id_%d", acctest.RandIntRange(10, 100))
	profile := "summary"

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfig(actionID, profile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_readme_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "trigger_record_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "namespace"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "playbook_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionDataSourceConfig(actionID string, profile string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_action" "schematics_action" {
			action_id = "%s"
			profile = "%s"
		}
	`, actionID, profile)
}

func testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_action" "schematics_action" {
			action_id = "%s"
		}
	`, actionID)
}

