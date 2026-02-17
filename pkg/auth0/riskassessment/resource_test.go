package riskassessment_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/jwefers/terraform-provider-auth0/pkg/acctest"
)

const testAccWithFalse = `resource "auth0_risk_assessments" "my_risk_assessments_settings" { enabled = false}`
const testAccWithTrue = `resource "auth0_risk_assessments" "my_risk_assessments_settings" { enabled = true}`

func TestAccRiskAssessment(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testAccWithFalse,
				Check:  resource.TestCheckResourceAttr("auth0_risk_assessments.my_risk_assessments_settings", "enabled", "false"),
			},
			{
				Config: testAccWithTrue,
				Check:  resource.TestCheckResourceAttr("auth0_risk_assessments.my_risk_assessments_settings", "enabled", "true"),
			},
		},
	})
}
