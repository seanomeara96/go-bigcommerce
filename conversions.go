package bigcommerce

/*
I have this idea where you select a parent product and you select some other existing products and they will be recreated as variants
under the selected parent product. All products that became variants will be deleted and redirected

first port of call is the ability to convert  from a product to a variant
*/
func (product *Product) ToVariantCreateParams(parentProductID int) ProductVariantCreateParams {
	return ProductVariantCreateParams{
		CostPrice:              product.CostPrice,
		Price:                  product.SalePrice,
		RetailPrice:            product.RetailPrice,
		Weight:                 product.Weight,
		Width:                  product.Width,
		Height:                 product.Height,
		Depth:                  product.Depth,
		IsFreeShipping:         product.IsFreeShipping,
		FixedCostShippingPrice: product.FixedCostShippingPrice,
		UPC:                    product.UPC,
		InventoryLevel:         product.InventoryLevel,
		InventoryWarningLevel:  product.InventoryWarningLevel,
		BinPickingNumber:       product.BinPickingNumber,
		GTIN:                   product.GTIN,
		MPN:                    product.MPN,
		ProductID:              parentProductID,
		SKU:                    product.SKU,
	}
}

func (product *Product) ToVariantCreateParamsWithOptions(parentProductID int, options *[]VariantOption) ProductVariantCreateParams {
	createVariantParams := product.ToVariantCreateParams(parentProductID)
	createVariantParams.OptionValues = options
	return createVariantParams
}

func (client *V3Client) ProductToProductVariant(parentProductID int, product Product, options *[]VariantOption) (ProductVariant, error) {
	var params ProductVariantCreateParams
	if options != nil {
		params = product.ToVariantCreateParamsWithOptions(parentProductID, options)
	} else {
		params = product.ToVariantCreateParams(parentProductID)
	}
	err := client.DeleteProduct(product.ID)
	if err != nil {
		return ProductVariant{}, err
	}
	return client.CreateProductVariant(parentProductID, params)
}
