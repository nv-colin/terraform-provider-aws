// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package int32planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LegacyValue returns a plan modifier that prevents `known after apply` during creation plans for
// attributes that must be `Computed,Optional` for legacy value reasons.
func NullValue() planmodifier.Int32 {
	return nullValueModifier{}
}

type nullValueModifier struct{}

func (m nullValueModifier) Description(_ context.Context) string {
	return ""
}

func (m nullValueModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m nullValueModifier) PlanModifyInt32(ctx context.Context, req planmodifier.Int32Request, resp *planmodifier.Int32Response) {
	// Use value from Config if set
	if !req.ConfigValue.IsNull() {
		return
	}

	// Exit if another planmodifier has set the value
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	if req.StateValue.IsNull() {
		resp.PlanValue = types.Int32Null()
		return
	}
}
