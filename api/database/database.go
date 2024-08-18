package database

import (
	"fmt"
	"log"

	"github.co/golang-programming/restaurant/api/entity"
	"gorm.io/driver/postgres"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		&entity.Table{},
		&entity.Staff{},
	)

	tenants[tenantID] = &TenantConfig{
		DB:  db,
		DSN: dsn,
	}

	return db, nil
}

func init() {
	tenantList := []struct {
		TenantID string
		DSN      string
	}{
		// Update DSNs to match PostgreSQL format: "host=<hostname> user=<username> password=<password> dbname=<dbname> port=<port> sslmode=<mode>"
		{"tenant1", "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.dvtazgngnspaajqfvtyx password=HmjYVLfpBQzo1VYK dbname=postgres port=6543 sslmode=disable"},
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
