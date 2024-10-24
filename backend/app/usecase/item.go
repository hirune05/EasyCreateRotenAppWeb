package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Item interface {
        Create(ctx context.Context, storeID int, name string, description *string, price int, imageURL *string) (*ItemDTO, error)
	GetByID(ctx context.Context, id string) (*ItemDTO, error)
	GetByStoreID(ctx context.Context, id string) (*GetByStoreIDDTO, error)
        GetAll(ctx context.Context) ([]*ItemDTO, error)
}
type item struct {
	db        *gorm.DB
	itemRepo repository.ItemRepository
}

type GetByStoreIDDTO struct {
        StoreAndItems *repository.GetByStoreIDDTO
}

type ItemDTO struct {
	Item *object.Item
}

var _ Item = (*item)(nil)

func NewItem(db *gorm.DB, itemRepo repository.ItemRepository) *item {
	return &item{
		db:        db,
		itemRepo: itemRepo,
	}
}

func (a *item) Create(ctx context.Context, storeID int, name string, description *string, price int, imageURL *string) (*ItemDTO, error) {
	acc, err := object.NewItem(storeID, name, description, price, imageURL)
	if err != nil {
		return nil, err
	}

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.itemRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &ItemDTO{
		Item: acc,
	}, nil
}

func (a *item) GetByID(ctx context.Context, id string) (*ItemDTO, error) {
	item, err := a.itemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get item by id: %w", err)
	}

	return &ItemDTO{
		Item: item,
	}, nil
}

func (u *item) GetByStoreID(ctx context.Context, id string) (*GetByStoreIDDTO, error) {
	items, err := u.itemRepo.GetByStoreID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetByStoreIDDTO{
		StoreAndItems: items,
	}, nil
}
func (r *item) GetAll(ctx context.Context) ([]*ItemDTO, error) {
	var items []*object.Item

	if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all items: %w", err)
	}

	var itemDTOs []*ItemDTO
	for _, item := range items {
		itemDTO := &ItemDTO{
			Item: item,
		}
		itemDTOs = append(itemDTOs, itemDTO)
	}

	return itemDTOs, nil
}
