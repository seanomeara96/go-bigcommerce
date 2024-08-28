package bigcommerce

import (
	"testing"
)

func TestCreateCoupon(t *testing.T) {
	client, _ := getClient()

	params := CreateUpdateCouponParams{
		Name:    "Test Coupon",
		Type:    "per_item_discount",
		Amount:  "10.00",
		Enabled: true,
		Code:    "TESTCOUPON",
		AppliesTo: AppliesTo{
			IDs:    []int{1},
			Entity: "products",
		},
	}

	coupon, err := client.V2.CreateCoupon(params)
	if err != nil {
		t.Fatalf("Failed to create coupon: %v", err)
	}

	if coupon.Name != params.Name {
		t.Errorf("Expected coupon name %s, got %s", params.Name, coupon.Name)
	}

	// Clean up
	err = client.V2.DeleteCoupon(coupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete test coupon: %v", err)
	}
}

func TestUpdateCoupon(t *testing.T) {
	client, _ := getClient()

	// First, create a coupon
	createParams := CreateUpdateCouponParams{
		Name:    "Test Coupon",
		Type:    "per_item_discount",
		Amount:  "10.00",
		Enabled: true,
		Code:    "TESTCOUPON",
		AppliesTo: AppliesTo{
			IDs:    []int{1},
			Entity: "products",
		},
	}

	coupon, err := client.V2.CreateCoupon(createParams)
	if err != nil {
		t.Fatalf("Failed to create test coupon: %v", err)
	}

	// Update the coupon
	updateParams := CreateUpdateCouponParams{
		Name:   "Updated Test Coupon",
		Amount: "15.00",
	}

	updatedCoupon, err := client.V2.UpdateCoupon(coupon.ID, updateParams)
	if err != nil {
		t.Fatalf("Failed to update coupon: %v", err)
	}

	if updatedCoupon.Name != updateParams.Name {
		t.Errorf("Expected updated coupon name %s, got %s", updateParams.Name, updatedCoupon.Name)
	}

	if updatedCoupon.Amount != updateParams.Amount {
		t.Errorf("Expected updated coupon amount %s, got %s", updateParams.Amount, updatedCoupon.Amount)
	}

	// Clean up
	err = client.V2.DeleteCoupon(coupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete test coupon: %v", err)
	}
}

func TestGetCoupons(t *testing.T) {
	client, _ := getClient()

	params := CouponQueryParams{
		Limit: 10,
	}

	coupons, meta, err := client.V2.GetCoupons(params)
	if err != nil {
		t.Fatalf("Failed to get coupons: %v", err)
	}

	if len(coupons) > params.Limit {
		t.Errorf("Expected at most %d coupons, got %d", params.Limit, len(coupons))
	}

	if meta.Pagination.Count != params.Limit {
		t.Errorf("Expected count %d in metadata, got %d", params.Limit, meta.Pagination.Count)
	}
}

func TestGetCoupon(t *testing.T) {
	client, _ := getClient()

	// First, create a coupon
	createParams := CreateUpdateCouponParams{
		Name:    "Test Coupon",
		Type:    "per_item_discount",
		Amount:  "10.00",
		Enabled: true,
		Code:    "TESTCOUPON",
		AppliesTo: AppliesTo{
			IDs:    []int{1},
			Entity: "products",
		},
	}

	createdCoupon, err := client.V2.CreateCoupon(createParams)
	if err != nil {
		t.Fatalf("Failed to create test coupon: %v", err)
	}

	// Get the coupon
	coupon, err := client.V2.GetCoupon(createdCoupon.ID)
	if err != nil {
		t.Fatalf("Failed to get coupon: %v", err)
	}

	if coupon.ID != createdCoupon.ID {
		t.Errorf("Expected coupon ID %d, got %d", createdCoupon.ID, coupon.ID)
	}

	if coupon.Name != createParams.Name {
		t.Errorf("Expected coupon name %s, got %s", createParams.Name, coupon.Name)
	}

	// Clean up
	err = client.V2.DeleteCoupon(createdCoupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete test coupon: %v", err)
	}
}

func TestDeleteCoupon(t *testing.T) {
	client, _ := getClient()

	// First, create a coupon
	createParams := CreateUpdateCouponParams{
		Name:    "Test Coupon",
		Type:    "per_item_discount",
		Amount:  "10.00",
		Enabled: true,
		Code:    "TESTCOUPON",
		AppliesTo: AppliesTo{
			IDs:    []int{1},
			Entity: "products",
		},
	}

	coupon, err := client.V2.CreateCoupon(createParams)
	if err != nil {
		t.Fatalf("Failed to create test coupon: %v", err)
	}

	// Delete the coupon
	err = client.V2.DeleteCoupon(coupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete coupon: %v", err)
	}

	// Try to get the deleted coupon (should fail)
	_, err = client.V2.GetCoupon(coupon.ID)
	if err == nil {
		t.Errorf("Expected error when getting deleted coupon, got nil")
	}
}

func TestValidateCreateUpdateCoupon(t *testing.T) {
	tests := []struct {
		name    string
		params  CreateUpdateCouponParams
		wantErr bool
	}{
		{
			name: "Valid params",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			params: CreateUpdateCouponParams{
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing type",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing amount",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing code when enabled",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
			},
			wantErr: true,
		},
		{
			name: "Empty AppliesTo.IDs",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					Entity: "products",
				},
			},
			wantErr: true,
		},
		{
			name: "Negative MaxUses",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
				MaxUses: -1,
			},
			wantErr: true,
		},
		{
			name: "Negative MaxUsesPerCustomer",
			params: CreateUpdateCouponParams{
				Name:    "Test Coupon",
				Type:    "per_item_discount",
				Amount:  "10.00",
				Enabled: true,
				Code:    "TESTCOUPON",
				AppliesTo: AppliesTo{
					IDs:    []int{1},
					Entity: "products",
				},
				MaxUsesPerCustomer: -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateUpdateCoupon(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCreateUpdateCoupon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
