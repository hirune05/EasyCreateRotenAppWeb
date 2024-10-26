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
        CreateComplex(ctx context.Context, storeID int, storeStaffID int, items []*OrderComplexItem) (*OrderDTO, error)
	Update(ctx context.Context, id string, storeID *int, storeStaffID *int, status *int, pickedUpAt *time.Time) (*OrderDTO, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*OrderDTO, error)
	GetByStoreID(ctx context.Context, id string) (*OrdersDTO, error)
	GetByStatus(ctx context.Context, storeID, status string) (*OrdersDTO, error)
	GetAll(ctx context.Context) (*OrdersDTO, error)
}
type order struct {
	db        *gorm.DB
	orderRepo repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
}

type OrderComplexItem struct {
	ItemID    int     `json:"itemId"`
	Quantity  int     `json:"quantity"`
	Arranges  *string `json:"arranges"`
}

type OrderDTO struct {
	Order *object.Order
}

type OrdersDTO struct {
	Orders []*object.Order
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

func (a *order) CreateComplex(ctx context.Context, storeID int, storeStaffID int, items []*OrderComplexItem) (*OrderDTO, error) {
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

        acc_items := make([]*object.OrderItem, len(items))
        for i, item := range items {
                oi, err := object.NewOrderItem(acc.ID, item.ItemID, item.Quantity, item.Arranges)
                if err != nil {
                        return nil, err
                }

                if err := a.orderItemRepo.Create(ctx, tx, oi); err != nil {
                        return nil, err
                }
                acc_items[i] = oi
        }

        acc.OrderItems = acc_items

	return &OrderDTO{
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

func (u *order) GetByID(ctx context.Context, id string) (*OrderDTO, error) {
	order, err := u.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}

	return &OrderDTO{
		Order: order,
	}, nil
}

func (u *order) GetByStoreID(ctx context.Context, id string) (*OrdersDTO, error) {
	orders, err := u.orderRepo.GetByStoreID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &OrdersDTO{
		Orders: orders,
	}, nil
}

func (u *order) GetByStatus(ctx context.Context, storeID, status string) (*OrdersDTO, error) {
	orders, err := u.orderRepo.GetByStatus(ctx, storeID, status)
	if err != nil {
		return nil, err
	}

	return &OrdersDTO{
		Orders: orders,
	}, nil
}

func (u *order) GetAll(ctx context.Context) (*OrdersDTO, error) {
	orders, err := u.orderRepo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all orders: %w", err)
	}

	return &OrdersDTO{
                Orders: orders,
        }, nil
}
