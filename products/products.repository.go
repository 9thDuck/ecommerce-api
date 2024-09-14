package products

import (
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/jackc/pgx/v5/pgconn"
)

type Product entities.Product

func (product *Product) create() error {
	if err := db.Instance.Create(&product).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}
		return err
	}
	return nil
}

func (product *Product) update() error {
	if err := db.Instance.Model(&product).Updates(&product).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}
		return err
	}
	return nil
}

func (product *Product) delete() error {
	if err := db.Instance.Delete(&product).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}
		return err
	}
	return nil
}

func (product *Product) get() error {
	if err := db.Instance.First(&product, &product).Error; err != nil {
		return err
	}
	return nil
}

func getAllByOptions(product *Product) (*[]Product, error) {
	var products []Product

	if err := db.Instance.Find(&products, product).Error; err != nil {
		return nil, err
	}

	return &products, nil
}
