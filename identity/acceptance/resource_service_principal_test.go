package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServicePrincipalResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	var servicePrincipal ScimServicePrincipal

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: createServicePrincipal(),
				Check: resource.ComposeTestCheckFunc(
					testServicePrincipalResourceExists("databricks_service_principal.sp", &servicePrincipal, t),
				),
				Destroy: false,
			},
		},
	})
}

func createServicePrincipal() string {
	return `
	resource "databricks_service_principal" "sp" {
		application_id = "5f8a224f-5228-4ee0-98c6-82596cb69847"
	}
	`
}

func testServicePrincipalResourceExists(key string, sp *ScimServicePrincipal, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[key]
		if !ok {
			return fmt.Errorf("Not found: %s", key)
		}

		conn := common.CommonEnvironmentClient()
		resp, err := NewServicePrincipalsAPI(conn).Read(sp.ID)
		if err != nil {
			return err
		}
		*sp = resp
		return nil
	}
}
