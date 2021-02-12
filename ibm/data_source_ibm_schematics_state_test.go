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

func TestAccIBMSchematicsStateDataSourceBasic(t *testing.T) {
	wID := fmt.Sprintf("w_id_%d", acctest.RandIntRange(10, 100))
	tID := fmt.Sprintf("t_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsStateDataSourceConfigBasic(wID, tID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "terraform_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "serial"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "lineage"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "modules.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsStateDataSourceConfigBasic(wID string, tID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_state" "schematics_state" {
			w_id = "%s"
			t_id = "%s"
		}
	`, wID, tID)
}

