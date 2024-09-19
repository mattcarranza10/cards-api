package infrastructure

import (
	"time"

	"cards-api/internal/cards"

	"gorm.io/gorm"
)

type DAO struct {
	ID            string         `gorm:"column:id"`
	CustomerID    string         `gorm:"column:customer_id"`
	Last4Digits   string         `gorm:"column:last_4_digits"`
	SensitiveData string         `gorm:"column:sensitive_data"`
	HolderName    string         `gorm:"column:holder_name"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (d *DAO) TableName() string {
	return "cards"
}

func (d *DAO) toDomain() *cards.Card {
	return &cards.Card{
		ID:            d.ID,
		CustomerID:    d.CustomerID,
		Last4Digits:   d.Last4Digits,
		SensitiveData: d.SensitiveData,
		Details: &cards.CardDetails{
			HolderName: &d.HolderName,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func toDAO(c *cards.Card) *DAO {
	return &DAO{
		ID:            c.ID,
		CustomerID:    c.CustomerID,
		Last4Digits:   c.Last4Digits,
		SensitiveData: c.SensitiveData,
		HolderName:    *c.Details.HolderName,
	}
}
