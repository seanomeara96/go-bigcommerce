package bigcommerce

import (
	"encoding/json"
	"testing"
)

func TestValidateCreateUpdateProductVariantOptions(t *testing.T) {
	params := CreateUpdateProductVariantOptions{
		ProductID:   1,
		DisplayName: "Size",
		Type:        "dropdown",
		OptionValues: []*Option{
			{Label: "Small", SortOrder: 0},
		},
	}

	if err := ValidateCreateUpdateProductVariantOptions(params); err != nil {
		t.Fatal(err)
	}
}

func TestVariantOptionCanSendExplicitFalse(t *testing.T) {
	params := CreateUpdateProductVariantOptions{
		ProductID:    1,
		DisplayName:  "Size",
		Type:         "dropdown",
		Config:       &Config{CheckedByDefault: BoolPtr(false)},
		OptionValues: []*Option{{Label: "Small", IsDefault: BoolPtr(false)}},
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"product_id":1,"display_name":"Size","type":"dropdown","config":{"checked_by_default":false},"option_values":[{"is_default":false,"label":"Small","sort_order":0,"value_data":{"is_default":false,"label":"","sort_order":0,"value_data":null,"id":0}}]}`
	if string(body) != expected {
		t.Fatalf("expected %s, got %s", expected, string(body))
	}
}
