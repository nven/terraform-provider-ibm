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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMSchematicsWorkspaceDataSourceBasic(t *testing.T) {
	wID := workspaceID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic(wID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "applied_shareddata_ids.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_repo.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "type.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status_msg.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic(wID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_workspace" "schematics_workspace" {
			w_id = "%s"
		}
	`, wID)
}
