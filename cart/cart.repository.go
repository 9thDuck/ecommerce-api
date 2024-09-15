package cart

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart entities.Cart

func upsertCartItem(userID uuid.UUID, productID uint, quantity int) error {
	cart := Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	result := db.Instance.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "product_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"quantity"}),
	}).Create(&cart)

	return result.Error
}

func GetCartItemsByUserID(tx *gorm.DB, userID uuid.UUID) (*[]common.CartItem, error) {
	if tx == nil {
		tx = db.Instance
	}

	var cartItems []common.CartItem
	err := tx.Where("user_id = ?", userID).Find(&cartItems).Error
	return &cartItems, err
}

func decrement(userID uuid.UUID, productID uint) error {
	result := db.Instance.Exec(`
		WITH updated AS (
			UPDATE cart
			SET quantity = quantity - 1
			WHERE user_id = ? AND product_id = ? AND quantity > 0
			RETURNING *
		)
		DELETE FROM cart
		WHERE id IN (SELECT id FROM updated WHERE quantity = 0)
	`, userID, productID)

	return result.Error
}

func delete(userID uuid.UUID, productID uint) error {
	result := db.Instance.Exec(`
		DELETE FROM cart
		WHERE user_id = ? AND product_id = ?
	`, userID, productID)

	return result.Error
}

func deleteAll(tx *gorm.DB, userID uuid.UUID) error {
	if tx == nil {
		tx = db.Instance
	}
	result := tx.Exec(`
		DELETE FROM cart
		WHERE user_id = ?
	`, userID)

	return result.Error
}
