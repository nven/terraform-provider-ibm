// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
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
	actionName := fmt.Sprintf("acc-test-schematics-actions_%s", acctest.RandString(10))
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "tf-acc-test-schematics-actions"
			location = "us-east"
			resource_group = "default"
		}
	`, actionName)
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
		//getActionOptions.SetProfile("detailed")

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
			return fmt.Errorf("schematics_action still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_action (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
