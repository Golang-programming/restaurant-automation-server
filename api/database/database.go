package database

import (
	"fmt"
	"log"

	"github.co/golang-programming/restaurant/api/entity"
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
		&entity.Bill{},
		&entity.Customer{},
		&entity.Deal{},
		&entity.DealItem{},
		&entity.Food{},
		&entity.Invoice{},
		&entity.InvoiceItem{},
		&entity.Menu{},
		&entity.Note{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Reservation{},
		&entity.Table{},
		&entity.User{},
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
		{"tenant1", "root:@tcp(127.0.0.1:3306)/restaurant_1?charset=utf8mb4&parseTime=True&loc=Local"},
		{"tenant2", "root:@tcp(127.0.0.1:3306)/restaurant_2?charset=utf8mb4&parseTime=True&loc=Local"},
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
