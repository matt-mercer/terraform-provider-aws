package globalaccelerator_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccGlobalAcceleratorCustomRoutingListener_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_globalaccelerator_custom_routing_listener.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, globalaccelerator.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCustomRoutingAcceleratorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRoutingListenerConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRoutingAcceleratorExists(ctx, rName),
					resource.TestCheckResourceAttr(resourceName, "port_range.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "port_range.0.from_port", "443"),
					resource.TestCheckResourceAttr(resourceName, "port_range.0.to_port", "443"),
					resource.TestCheckResourceAttr(resourceName, "port_range.1.from_port", "10000"),
					resource.TestCheckResourceAttr(resourceName, "port_range.1.to_port", "30000"),
				),
			},
		},
	})
}

func testAccCustomRoutingListenerConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_globalaccelerator_custom_routing_accelerator" "test" {
  name = %[1]q
}

resource "aws_globalaccelerator_custom_routing_listener" "test" {
  accelerator_arn = aws_globalaccelerator_custom_routing_accelerator.test.id

  port_range = {
    from_port = 443
    to_port   = 443
  }

  port_range = {
    from_port = 10000
    to_port   = 30000
  }
}
`, rName)
}
