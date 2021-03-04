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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsJobDataSourceBasic(t *testing.T) {
	jobRefreshToken := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))
	jobCommandObject := "action"
	//jobCommandObjectID := fmt.Sprintf("command_object_id_%d", acctest.RandIntRange(10, 100))
	jobCommandName := "ansible_playbook_run"
	jobCommandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	jobLocation := "us-east"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobDataSourceConfig(jobRefreshToken, jobCommandObject, actionID, jobCommandName, jobCommandParameter, jobLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsJobDataSourceConfig(jobRefreshToken string, jobCommandObject string, jobCommandObjectID, jobCommandName string, jobCommandParameter string, jobLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
			command_object = "%s"
			command_object_id = "%s"
			command_name = "%s"
			command_parameter = "%s"
			location = "%s"
		}

		data "ibm_schematics_job" "schematics_job" {
			job_id = ibm_schematics_job.schematics_job.id
		}
	`, jobRefreshToken, jobCommandObject, jobCommandObjectID, jobCommandName, jobCommandParameter, jobLocation)
}
