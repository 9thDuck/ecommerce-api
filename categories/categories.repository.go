package categories

import (
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/jackc/pgx/v5/pgconn"
)

type Category entities.Category

func (c *Category) Create() error {
	if err := db.Instance.Create(&c).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}

		return err
	}
	return nil
}

func (c *Category) Update() error {
	if err := db.Instance.Model(&c).Updates(&c).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}
		return err
	}
	return nil
}

func (c *Category) Delete() error {
	if err := db.Instance.Delete(&c).Error; err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}
		return err
	}
	return nil
}

func (c *Category) FindAll() (*[]Category, error) {
	var categories []Category
	
	if err := db.Instance.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (c *Category) FindByID() error {
	return db.Instance.First(c, c.ID).Error
}

