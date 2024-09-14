package categories

func create(category *Category) error {
	return category.Create()
}

func update(category *Category) error {
	return category.Update()
}

func delete(category *Category) error {
	return category.Delete()
}

func get(category *Category) error {
	return category.FindByID()
}

func getAllCategories() (*[]Category, error) {
	var c *Category
	return c.FindAll()
}
