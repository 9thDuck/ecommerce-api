package products

import "github.com/9thDuck/ecommerce-api.git/db"

func create(product *Product) error {
	return product.create()
}

func update(product *Product) error {
	return product.update()
}

func delete(product *Product) error {
	return product.delete()
}

func get(product *Product) error {
	return product.get()
}



func getProductsByOptions(options *Product) (*[]Product, error) {
	return getAllByOptions(options)
}

func getProductsBySearchCriteria(criteria *map[string]interface{}) (*[]Product, error) {
	var products []Product
	query := db.Instance.Model(&Product{})

    for key, value := range *criteria {
        switch key {
        case "name", "description":
            query = query.Where("to_tsvector('english', "+key+") @@ plainto_tsquery('english', ?)", value)
        case "price_gte":
            query = query.Where("price >= ?", value)
        case "price_lte":
            query = query.Where("price <= ?", value)
        case "category_id":
            query = query.Where("category_id = ?", value)
        case "in_stock":
            if value.(bool) {
                query = query.Where("stock > 0")
            } else {
                query = query.Where("stock = 0")
            }
        }
    }

    if err := query.Find(&products).Error; err != nil {
        return nil, err
    }

    return &products, nil
}

func getAllProducts() (*[]Product, error) {
    return getProductsByOptions(&Product{})
}

