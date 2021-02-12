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

func TestAccIBMSchematicsJobDataSourceBasic(t *testing.T) {
	jobID := fmt.Sprintf("job_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobDataSourceConfigBasic(jobID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_options.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "duration"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "results_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsJobDataSourceAllArgs(t *testing.T) {
	jobID := fmt.Sprintf("job_id_%d", acctest.RandIntRange(10, 100))
	profile := "summary"

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobDataSourceConfig(jobID, profile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_options.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "duration"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "results_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsJobDataSourceConfig(jobID string, profile string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_job" "schematics_job" {
			job_id = "%s"
			profile = "%s"
		}
	`, jobID, profile)
}

func testAccCheckIBMSchematicsJobDataSourceConfigBasic(jobID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_job" "schematics_job" {
			job_id = "%s"
		}
	`, jobID)
}

