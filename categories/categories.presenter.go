package categories

import (
	"fmt"

	"github.com/9thDuck/ecommerce-api.git/common"
)

type categoryResponse common.Response
type categoriesResponse common.Response
type errorResponse common.Response

func successCreateResponse(category *Category) *categoryResponse {
	return &categoryResponse{
		Message: "Category created successfully",
		Data:    category,
	}
}

func successUpdateResponse(category *Category) *categoryResponse {
	return &categoryResponse{
		Message: "Category updated successfully",
		Data:    category,
	}
}

func successDeleteResponse() *categoryResponse {
	return &categoryResponse{
		Message: "Category deleted successfully",
		Data:    nil,
	}
}

func successGetCategoryResponse(category *Category) *categoryResponse {
	return &categoryResponse{
		Message: "Category retrieved successfully",
		Data:    category,
	}
}

func successGetCategoriesResponse(categories *[]Category) *categoriesResponse {
	return &categoriesResponse{
		Message: fmt.Sprintf("Found %d categories", len(*categories)),
		Data:    categories,
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
		Message: "Category not found",
		Data:    nil,
	}
}

func internalServerErrorResponse() *errorResponse {
	return &errorResponse{
		Message: "An internal server error occurred",
		Data:    nil,
	}
}

func unauthorizedResponse() *errorResponse {
	return &errorResponse{
		Message: "Unauthorized access",
		Data:    nil,
	}
}