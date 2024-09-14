package products

import (
	"fmt"

	"github.com/9thDuck/ecommerce-api.git/common"
)

type productResponse common.Response
type productsResponse common.Response
type errorResponse common.Response

func successCreateResponse(product *Product) *productResponse {
	return &productResponse{
		Message: "Product created successfully",
		Data:    product,
	}
}

func successUpdateResponse(product *Product) *productResponse {
	return &productResponse{
		Message: "Product updated successfully",
		Data:    product,
	}
}

func successDeleteResponse() *productResponse {
	return &productResponse{
		Message: "Product deleted successfully",
		Data:    nil,
	}
}

func successGetProductResponse(product *Product) *productResponse {
	return &productResponse{
		Message: "Product retrieved successfully",
		Data:    product,
	}
}

func successGetProductsResponse(products *[]Product) *productsResponse {
	return &productsResponse{
		Message: fmt.Sprintf("Found %d products", len(*products)),
		Data:    products,
	}
}

func successSearchProductsResponse(products *[]Product) *productsResponse {
	return &productsResponse{
		Message: fmt.Sprintf("Found %d products matching the search criteria", len(*products)),
		Data:    products,
	}
}


func badRequestResponse(errorStr string) *errorResponse {
	return &errorResponse{
		Message: "Bad request: " + errorStr,
		Data:    nil,
	}
}

func notFoundResponse() *errorResponse {
	return &errorResponse{
		Message: "Product not found",
		Data:    nil,
	}
}

func internalServerErrorResponse() *errorResponse {
	return &errorResponse{
		Message: "An internal server error occurred",
		Data:    nil,
	}
}

