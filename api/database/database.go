package database

import (
	"fmt"
	"log"

	foodEntity "github.co/golang-programming/restaurant/api/food/entity"
	invoiceEntity "github.co/golang-programming/restaurant/api/invoice/entity"
	menuEntity "github.co/golang-programming/restaurant/api/menu/entity"
	noteEntity "github.co/golang-programming/restaurant/api/note/entity"
	orderEntity "github.co/golang-programming/restaurant/api/order/entity"
	tableEntity "github.co/golang-programming/restaurant/api/table/entity"
	userEntity "github.co/golang-programming/restaurant/api/user/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ActiveDB gorm.DB

type TenantConfig struct {
	DB  *gorm.DB
	DSN string
}

var tenants = make(map[string]*TenantConfig)

func initTenantDB(tenantID, dsn string) (*gorm.DB, error) {
	fmt.Println("tenantID, dsn", tenantID, dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&foodEntity.Food{},
		&userEntity.User{},
		&invoiceEntity.Invoice{},
		&orderEntity.Order{},
		&orderEntity.OrderItem{},
		&tableEntity.Table{},
		&noteEntity.Note{},
		&menuEntity.Menu{},
	)

	tenants[tenantID] = &TenantConfig{
		DB:  db,
		DSN: dsn,
	}

	return db, nil
}

func ConnectToDatabase() {
	tenantList := []struct {
		TenantID string
		DSN      string
	}{
		{"tenant1", "root:@tcp(127.0.0.1:3306)/tenant1_db?charset=utf8mb4&parseTime=True&loc=Local"},
		{"tenant2", "root:@tcp(127.0.0.1:3306)/tenant2_db?charset=utf8mb4&parseTime=True&loc=Local"},
	}

	for _, tenant := range tenantList {
		if _, err := initTenantDB(tenant.TenantID, tenant.DSN); err != nil {
			log.Fatalf("Could not connect to tenant database (%s): %v", tenant.TenantID, err)
		}
	}
}

func SetDB(tenantID string) error {
	tenantConfig, exists := tenants[tenantID]
	if !exists {
		return fmt.Errorf("tenant not found")
	}

	ActiveDB = *tenantConfig.DB
	return nil
}
