package infrastructure

import (
	"cards-api/internal/cards"
	"context"
	"errors"
	"gorm.io/gorm"
)

type GormRepository struct {
	gormDB *gorm.DB
}

func NewGormRepository(gormDB *gorm.DB) *GormRepository {
	return &GormRepository{
		gormDB: gormDB,
	}
}

func (gr *GormRepository) Create(ctx context.Context, card *cards.Card) error {
	dao := toDAO(card)
	return gr.gormDB.WithContext(ctx).Create(dao).Error
}

func (gr *GormRepository) Get(ctx context.Context, id, customerID string) (*cards.Card, error) {
	dao := &DAO{}
	if err := gr.gormDB.WithContext(ctx).First(dao, "id = ? AND customer_id = ?", id, customerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cards.ErrCardNotFound
		}
		return nil, err
	}
	return dao.toDomain(), nil
}

func (gr *GormRepository) Update(ctx context.Context, card *cards.Card) error {
	dao := toDAO(card)
	return gr.gormDB.WithContext(ctx).Updates(dao).Error
}

func (gr *GormRepository) Delete(ctx context.Context, id, customerID string) error {
	return gr.gormDB.WithContext(ctx).Delete(&DAO{}, "id = ? AND customer_id = ?", id, customerID).Error
}
