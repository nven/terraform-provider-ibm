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

	schematicsv1 "github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMSchematicsActionBasic(t *testing.T) {
	var conf schematicsv1.Action

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
			name = "tf-acc-test-schematics-actions"
			description = "tf-acc-test-schematics-actions"
			location = "us-east"
			resource_group = "default"
		}
	`)
}

func testAccCheckIBMSchematicsActionConfig(name string, description string, location string, resourceGroup string, sourceReadmeURL string, sourceType string, commandParameter string, targetsIni string, triggerRecordID string, xGithubToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
			source_readme_url = "%s"
			source_type = "%s"
			command_parameter = "%s"
			targets_ini = "%s"
			trigger_record_id = "%s"
			X-Github-token = "%s"
		}
	`, name, description, location, resourceGroup, sourceReadmeURL, sourceType, commandParameter, targetsIni, triggerRecordID, xGithubToken)
}

func testAccCheckIBMSchematicsActionExists(n string, obj schematicsv1.Action) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)
		getActionOptions.SetProfile("detailed")

		action, _, err := schematicsClient.GetAction(getActionOptions)
		if err != nil {
			return err
		}

		obj = *action
		return nil
	}
}

func testAccCheckIBMSchematicsActionDestroy(s *terraform.State) error {
	schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_action" {
			continue
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetAction(getActionOptions)

		if err == nil {
			return fmt.Errorf("Complete Action details with user inputs and system generated data. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Complete Action details with user inputs and system generated data. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
