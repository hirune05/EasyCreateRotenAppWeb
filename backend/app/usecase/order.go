package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order interface {
	Create(ctx context.Context, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*OrderDTO, error)
        CreateComplex(ctx context.Context, storeID int, storeStaffID int, items []*OrderComplexItem) (*CreateOrderComplexDTO, error)
	Update(ctx context.Context, id string, storeID *int, storeStaffID *int, status *int, pickedUpAt *time.Time) (*OrderDTO, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*OrderDTO, error)
	GetByStoreID(ctx context.Context, id string) ([]*OrderDTO, error)
	GetByStatus(ctx context.Context, storeID, status string) ([]*OrderDTO, error)
	GetAll(ctx context.Context) ([]*OrderDTO, error)
}
type order struct {
	db        *gorm.DB
	orderRepo repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
}

type OrderComplexItem struct {
	ItemID    int     `json:"item_id"`
	Quantity  int     `json:"quantity"`
	Arranges  *string `json:"arranges"`
}

type OrderDTO struct {
	Order *object.Order
}

type CreateOrderComplexDTO struct {
	Order *object.Order
        OrderItems []*object.OrderItem
}

var _ Order = (*order)(nil)

func NewOrder(db *gorm.DB, orderRepo repository.OrderRepository, orderItemRepo repository.OrderItemRepository) *order {
	return &order{
		db:        db,
		orderRepo: orderRepo,
                orderItemRepo: orderItemRepo,
	}
}

func (a *order) Create(ctx context.Context, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*OrderDTO, error) {
	acc, err := object.NewOrder(storeID, storeStaffID, status, pickedUpAt)
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

	if err := a.orderRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &OrderDTO{
		Order: acc,
	}, nil
}

func (a *order) CreateComplex(ctx context.Context, storeID int, storeStaffID int, items []*OrderComplexItem) (*CreateOrderComplexDTO, error) {
	acc, err := object.NewOrder(storeID, storeStaffID, 1, nil)
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



	if err := a.orderRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}


        acc_items := make([]object.OrderItem, len(items))
        for i, item := range items {
                oi, err := object.NewOrderItem(acc.ID, item.ItemID, item.Quantity, item.Arranges)
                if err != nil {
                        return nil, err
                }

                if err := a.orderItemRepo.Create(ctx, tx, oi); err != nil {
                        return nil, err
                }
                acc_items[i] = *oi
        }

        acc.OrderItems = acc_items

	return &CreateOrderComplexDTO{
		Order: acc,
	}, nil
}

func (a *order) Update(ctx context.Context, id string, storeID *int, storeStaffID *int, status *int, pickedUpAt *time.Time) (*OrderDTO, error) {
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

	order, err := a.orderRepo.GetByID(ctx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

        if storeID != nil {
                order.StoreID = *storeID
        }
        if storeStaffID != nil {
                order.StoreStaffID = *storeStaffID
        }
        if status != nil {
                order.Status = *status
        }
        if pickedUpAt != nil {
                order.PickedUpAt = pickedUpAt
        }
	order.UpdatedAt = time.Now()

	if err := a.orderRepo.Update(ctx, tx, order); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &OrderDTO{
		Order: order,
	}, nil
}

func (a *order) Delete(ctx context.Context, id string) error {
	tx := a.db.Begin()
	if tx.Error != nil {
		return tx.Error
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

	if err := a.orderRepo.Delete(ctx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}

func (a *order) GetByID(ctx context.Context, id string) (*OrderDTO, error) {
	order, err := a.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}

	return &OrderDTO{
		Order: order,
	}, nil
}

func (u *order) GetByStoreID(ctx context.Context, id string) ([]*OrderDTO, error) {
	orders, err := u.orderRepo.GetByStoreID(ctx, id)
	if err != nil {
		return nil, err
	}

	var orderDTOs []*OrderDTO
	for _, item := range orders {
		orderDTOs = append(orderDTOs, &OrderDTO{
			Order: &object.Order{
				ID:           item.ID,
				StoreID:      item.StoreID,
				StoreStaffID: item.StoreStaffID,
				Status:       item.Status,
				PickedUpAt:   item.PickedUpAt,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			},
		})
	}

	return orderDTOs, nil
}

func (u *order) GetByStatus(ctx context.Context, storeID, status string) ([]*OrderDTO, error) {
	orders, err := u.orderRepo.GetByStatus(ctx, storeID, status)
	if err != nil {
		return nil, err
	}

	var orderDTOs []*OrderDTO
	for _, item := range orders {
		orderDTOs = append(orderDTOs, &OrderDTO{
			Order: &object.Order{
				ID:           item.ID,
				StoreID:      item.StoreID,
				StoreStaffID: item.StoreStaffID,
				Status:       item.Status,
				PickedUpAt:   item.PickedUpAt,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			},
		})
	}

	return orderDTOs, nil
}

func (r *order) GetAll(ctx context.Context) ([]*OrderDTO, error) {
	var orders []*object.Order

	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all orders: %w", err)
	}

	var orderDTOs []*OrderDTO
	for _, order := range orders {
		orderDTO := &OrderDTO{
			Order: order,
		}
		orderDTOs = append(orderDTOs, orderDTO)
	}

	return orderDTOs, nil
}
