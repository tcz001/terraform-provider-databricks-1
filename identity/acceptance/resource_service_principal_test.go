package acceptance

import (
	"testing"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccServicePrincipalResource(t * testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	var servicePrincipal ScimServicePrincipal

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: createServicePrincipal()
			}
		}
		Check: resource.ComposeTestCheckFunc(
			testServicePrincipalResourceExists("databricks_service_principal.sp", &sp, t))
		)
	})
}

func createServicePrincipal() string {
	return `
	resource "databricks_service_principal" "sp" {
		application_id = "00000000-0000-0000-0000-000000000000"
	}
	`
}

func testServicePrincipalResourceExists(key string, sp *ScimServicePrincipal, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resource[key]
		if !ok {
			return fmt.Errorf("Not found: %s", key)
		}
	}

	conn := common.CommonEnvironmentClient()
	resp, err :=  NewServicePrisdfncipalsAPI(conn).Read(rs.Primary.ID)
	if err != nil {
		return err
	}
	*sp =  resp
	return nill
}