package bigcommerce

import (
	"fmt"
	"net/url"
	"strconv"
)

// ProductReviewStatus is the moderation status of a product review.
type ProductReviewStatus string

const (
	ProductReviewStatusApproved    ProductReviewStatus = "approved"
	ProductReviewStatusDisapproved ProductReviewStatus = "disapproved"
	ProductReviewStatusPending     ProductReviewStatus = "pending"
)

// ProductReview is a review attached to a catalog product.
type ProductReview struct {
	ID           int                 `json:"id"`
	ProductID    int                 `json:"product_id"`
	Title        string              `json:"title"`
	DateReviewed string              `json:"date_reviewed"`
	Text         string              `json:"text"`
	Status       ProductReviewStatus `json:"status"`
	Rating       int                 `json:"rating"`
	Email        string              `json:"email"`
	Name         string              `json:"name"`
	DateCreated  string              `json:"date_created"`
	DateModified string              `json:"date_modified"`
}

// ProductReviewQueryParams controls filtering, field selection, and pagination.
// Status must point to 0 for pending reviews or 1 for approved reviews.
type ProductReviewQueryParams struct {
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Status        *int   `url:"status,omitempty"`
}

// CreateProductReviewParams contains the writable fields for a new review.
// Title and DateReviewed are required by BigCommerce.
type CreateProductReviewParams struct {
	Title        string              `json:"title"`
	DateReviewed string              `json:"date_reviewed"`
	Text         string              `json:"text,omitempty"`
	Status       ProductReviewStatus `json:"status,omitempty"`
	Rating       int                 `json:"rating,omitempty"`
	Email        string              `json:"email,omitempty"`
	Name         string              `json:"name,omitempty"`
}

// UpdateProductReviewParams contains fields to update. Pointer fields distinguish
// an omitted field from an intentional empty string or zero value.
type UpdateProductReviewParams struct {
	Title        *string              `json:"title,omitempty"`
	DateReviewed *string              `json:"date_reviewed,omitempty"`
	Text         *string              `json:"text,omitempty"`
	Status       *ProductReviewStatus `json:"status,omitempty"`
	Rating       *int                 `json:"rating,omitempty"`
	Email        *string              `json:"email,omitempty"`
	Name         *string              `json:"name,omitempty"`
}

// GetProductReviews retrieves one page of reviews for a product.
func (client *V3Client) GetProductReviews(productID int, params ProductReviewQueryParams) ([]ProductReview, MetaData, error) {
	var response struct {
		Data []ProductReview `json:"data"`
		Meta MetaData        `json:"meta"`
	}

	reviewsURL, err := urlWithQueryParams(client.productReviewsURL(productID), params)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL for product reviews (product ID: %d): %w", productID, err)
	}
	if err := client.Get(reviewsURL, &response); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get reviews for product ID %d: %w", productID, err)
	}

	return response.Data, response.Meta, nil
}

// GetAllProductReviews retrieves every page of reviews for a product.
func (client *V3Client) GetAllProductReviews(productID int, params ProductReviewQueryParams) ([]ProductReview, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 250
	}

	var reviews []ProductReview
	for {
		page, meta, err := client.GetProductReviews(productID, params)
		if err != nil {
			return nil, fmt.Errorf("failed to get all reviews for product ID %d at page %d: %w", productID, params.Page, err)
		}
		reviews = append(reviews, page...)

		if meta.Pagination.TotalPages > 0 {
			if meta.Pagination.CurrentPage >= meta.Pagination.TotalPages {
				break
			}
		} else if len(page) < params.Limit {
			break
		}
		params.Page++
	}

	return reviews, nil
}

// GetProductReview retrieves a single review.
func (client *V3Client) GetProductReview(productID, reviewID int) (ProductReview, error) {
	var response struct {
		Data ProductReview `json:"data"`
	}
	if err := client.Get(client.productReviewURL(productID, reviewID), &response); err != nil {
		return ProductReview{}, fmt.Errorf("failed to get review ID %d for product ID %d: %w", reviewID, productID, err)
	}
	return response.Data, nil
}

// CreateProductReview creates a review for a product.
func (client *V3Client) CreateProductReview(productID int, params CreateProductReviewParams) (ProductReview, error) {
	var response struct {
		Data ProductReview `json:"data"`
	}
	if params.Title == "" || params.DateReviewed == "" {
		return ProductReview{}, fmt.Errorf("invalid params for CreateProductReview: title and date_reviewed are required (product ID: %d)", productID)
	}
	if err := client.Post(client.productReviewsURL(productID), params, &response); err != nil {
		return ProductReview{}, fmt.Errorf("failed to create review for product ID %d: %w", productID, err)
	}
	return response.Data, nil
}

// UpdateProductReview updates a review for a product.
func (client *V3Client) UpdateProductReview(productID, reviewID int, params UpdateProductReviewParams) (ProductReview, error) {
	var response struct {
		Data ProductReview `json:"data"`
	}
	if err := client.Put(client.productReviewURL(productID, reviewID), params, &response); err != nil {
		return ProductReview{}, fmt.Errorf("failed to update review ID %d for product ID %d: %w", reviewID, productID, err)
	}
	return response.Data, nil
}

// DeleteProductReview deletes a review from a product.
func (client *V3Client) DeleteProductReview(productID, reviewID int) error {
	if err := client.Delete(client.productReviewURL(productID, reviewID), nil); err != nil {
		return fmt.Errorf("failed to delete review ID %d for product ID %d: %w", reviewID, productID, err)
	}
	return nil
}

func (client *V3Client) productReviewsURL(productID int) *url.URL {
	return client.constructURL("catalog", "products", strconv.Itoa(productID), "reviews")
}

func (client *V3Client) productReviewURL(productID, reviewID int) *url.URL {
	return client.constructURL("catalog", "products", strconv.Itoa(productID), "reviews", strconv.Itoa(reviewID))
}
