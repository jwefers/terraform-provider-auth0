package role

import (
	"github.com/auth0/go-auth0/management"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/jwefers/terraform-provider-auth0/pkg/value"
)

func expandRole(data *schema.ResourceData) *management.Role {
	cfg := data.GetRawConfig()

	return &management.Role{
		Name:        value.String(cfg.GetAttr("name")),
		Description: value.String(cfg.GetAttr("description")),
	}
}
