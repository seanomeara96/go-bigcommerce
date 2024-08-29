package bigcommerce

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

const productIdDoesExist = 79

func TestCreateCoupon(t *testing.T) {
	client, _ := getTestClient()

	// Generate a unique coupon name and code
	uniqueSuffix := time.Now().UnixNano()
	uniqueName := fmt.Sprintf("Test Coupon %d", uniqueSuffix)
	uniqueCode := fmt.Sprintf("TEST%d", uniqueSuffix)

	params := CreateCouponParams{
		Name:    uniqueName,
		Type:    "per_item_discount",
		Amount:  "1.00",
		Enabled: true,
		Code:    uniqueCode,
		AppliesTo: &AppliesTo{
			IDs:    []int{productIdDoesExist},
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
	client, _ := getTestClient()

	// First, create a coupon
	coupon, err := client.V2.CreateCoupon(CreateCouponParams{
		Name:    fmt.Sprintf("Test Coupon %d", time.Now().UnixNano()),
		Type:    "per_item_discount",
		Amount:  "1.00",
		Enabled: true,
		Code:    fmt.Sprintf("TEST%d", time.Now().UnixNano()),
		AppliesTo: &AppliesTo{
			IDs:    []int{productIdDoesExist},
			Entity: "products",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create test coupon: %v", err)
	}

	// Update the coupon
	uniqueSuffix := time.Now().UnixNano()
	updateParams := UpdateCouponParams{
		Name:   fmt.Sprintf("Test Updated Coupon %d", uniqueSuffix),
		Amount: "1.00",
		Type:   "per_item_discount",
	}

	updatedCoupon, err := client.V2.UpdateCoupon(coupon.ID, updateParams)
	if err != nil {
		t.Fatalf("Failed to update coupon: %v", err)
	}

	if updatedCoupon.Name != updateParams.Name {
		t.Errorf("Expected updated coupon name %s, got %s", updateParams.Name, updatedCoupon.Name)
	}

	actual, _ := strconv.Atoi(updatedCoupon.Amount)
	expected, _ := strconv.Atoi(updateParams.Amount)
	if actual != expected {
		t.Errorf("Expected updated coupon amount %s, got %s", updateParams.Amount, updatedCoupon.Amount)
	}

	// Clean up
	err = client.V2.DeleteCoupon(coupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete test coupon: %v", err)
	}
}

func TestGetCoupons(t *testing.T) {
	client, _ := getTestClient()

	params := CouponQueryParams{
		Limit: 10,
	}

	coupons, err := client.V2.GetCoupons(params)
	if err != nil {
		t.Fatalf("Failed to get coupons: %v", err)
	}

	if len(coupons) > params.Limit {
		t.Errorf("Expected at most %d coupons, got %d", params.Limit, len(coupons))
	}
}

func TestGetCoupon(t *testing.T) {
	client, _ := getTestClient()

	// First, create a coupon
	createdCoupon, err := client.V2.CreateCoupon(CreateCouponParams{
		Name:    fmt.Sprintf("Test Coupon %d", time.Now().UnixNano()),
		Type:    "per_item_discount",
		Amount:  "1.00",
		Enabled: true,
		Code:    fmt.Sprintf("TEST%d", time.Now().UnixNano()),
		AppliesTo: &AppliesTo{
			IDs:    []int{productIdDoesExist},
			Entity: "products",
		},
	})
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

	if coupon.Name != createdCoupon.Name {
		t.Errorf("Expected coupon name %s, got %s", createdCoupon.Name, coupon.Name)
	}

	// Clean up
	err = client.V2.DeleteCoupon(coupon.ID)
	if err != nil {
		t.Fatalf("Failed to delete test coupon: %v", err)
	}
}

func TestDeleteCoupon(t *testing.T) {
	client, _ := getTestClient()

	// First, create a coupon
	coupon, err := client.V2.CreateCoupon(CreateCouponParams{
		Name:    fmt.Sprintf("Test Coupon %d", time.Now().UnixNano()),
		Type:    "per_item_discount",
		Amount:  "1.00",
		Enabled: true,
		Code:    fmt.Sprintf("TEST%d", time.Now().UnixNano()),
		AppliesTo: &AppliesTo{
			IDs:    []int{productIdDoesExist},
			Entity: "products",
		},
	})
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
