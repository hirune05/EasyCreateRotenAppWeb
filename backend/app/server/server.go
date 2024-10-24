package server

import (
	"backend/app/config"
	"backend/app/dao"
	"backend/app/domain/object"
	"backend/app/handler"
	"backend/app/usecase"
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() error {

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return errors.New("DB user is not configured")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return errors.New("DB password is not configured")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return errors.New("DB host is not configured")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return errors.New("DB port is not configured")
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/roten-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected successfully")

	err = db.AutoMigrate(
		&object.Store{},
		&object.StoreStaff{},
		&object.Item{},
		&object.Order{},
		&object.OrderItem{},
		&object.Report{},
		&object.Event{},
		&object.AdminUser{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")

	// // シードデータを追加したいときだけコメントアウトを外す
	// if err := seedData(db); err != nil {
	// 	log.Fatalf("failed to seed data: %v", err)
	// }
	// log.Println("Seed data added successfully")

	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)

	orderUseCase := usecase.NewOrder(db, dao.NewOrderRepository(db))
	orderItemUseCase := usecase.NewOrderItem(db, dao.NewOrderItemRepository(db))
	adminUserUseCase := usecase.NewAdminUser(db, dao.NewAdminUserRepository(db))
	storeStaffUseCase := usecase.NewStoreStaff(db, dao.NewStoreStaffRepository(db))
	studentUseCase := usecase.NewStudent(db, dao.NewStudentRepository(db))
	eventUseCase := usecase.NewEvent(db, dao.NewEventRepository(db))

	r := handler.NewRouter(
		orderUseCase, orderItemUseCase, adminUserUseCase, storeStaffUseCase, studentUseCase, eventUseCase)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}

func seedData(db *gorm.DB) error {
	// Eventのシードデータ
	events := []object.Event{
		{Name: "Event 1", Year: 2023, StartTime: time.Date(2023, time.October, 23, 0, 0, 0, 0, time.UTC), EndTime: time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Event 2", Year: 2024, StartTime: time.Date(2024, time.October, 23, 0, 0, 0, 0, time.UTC), EndTime: time.Date(2024, time.October, 31, 0, 0, 0, 0, time.UTC)},
	}
	if err := db.Create(&events).Error; err != nil {
		return err
	}

	// Storeのシードデータ
	stores := []object.Store{
		{Name: "Store A", ImageURL: nil, EventID: events[0].ID},
		{Name: "Store B", ImageURL: nil, EventID: events[1].ID},
	}
	if err := db.Create(&stores).Error; err != nil {
		return err
	}

	// Studentのシードデータ
	student := []object.Student{
		{ID: 1234, Name: "John Doe", Password: "password1"},
		{ID: 5678, Name: "Jane Smith", Password: "password2"},
	}
	if err := db.Create(&student).Error; err != nil {
		return err
	}

	// StoreStaffのシードデータ
	storeStaffs := []object.StoreStaff{
		{StudentID: student[0].ID, Role: 1, StoreID: stores[0].ID},
		{StudentID: student[1].ID, Role: 0, StoreID: stores[1].ID},
	}
	if err := db.Create(&storeStaffs).Error; err != nil {
		return err
	}

	// Itemのシードデータ
	items := []object.Item{
		{StoreID: stores[0].ID, Name: "Item A", Description: nil, Price: 100, ImageURL: nil},
		{StoreID: stores[1].ID, Name: "Item B", Description: nil, Price: 200, ImageURL: nil},
	}
	if err := db.Create(&items).Error; err != nil {
		return err
	}

	// Orderのシードデータ
	now := time.Now()
	orders := []object.Order{
		{StoreID: stores[0].ID, PickedUpAt: &now, Status: 0, StoreStaffID: storeStaffs[0].ID},
		{StoreID: stores[1].ID, PickedUpAt: &now, Status: 1, StoreStaffID: storeStaffs[1].ID},
	}
	if err := db.Create(&orders).Error; err != nil {
		return err
	}

	// OrderItemのシードデータ
	orderItems := []object.OrderItem{
		{OrderID: orders[0].ID, ItemID: items[0].ID, Arranges: nil, Quantity: 2},
		{OrderID: orders[1].ID, ItemID: items[1].ID, Arranges: nil, Quantity: 1},
	}
	if err := db.Create(&orderItems).Error; err != nil {
		return err
	}

	// Reportのシードデータ
	reports := []object.Report{
		{StoreID: &stores[0].ID, StoreStaffID: &storeStaffs[0].ID, Description: "Report 1"},
		{StoreID: &stores[1].ID, StoreStaffID: &storeStaffs[1].ID, Description: "Report 2"},
	}
	if err := db.Create(&reports).Error; err != nil {
		return err
	}

	// AdminUserのシードデータ
	adminUsers := []object.AdminUser{
		{Username: "admin1", Password: "password", Email: "admin1@gmail.com"},
		{Username: "admin2", Password: "password", Email: "admin2@gmail.com"},
	}
	if err := db.Create(&adminUsers).Error; err != nil {
		return err
	}

	return nil
}
