package models

import (
	"time"

	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)

// BaseModel is an abstract model that other models can embed
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User model
type User struct {
	BaseModel
	Username string `gorm:"size:100;unique"`
	Email    string `gorm:"size:255;unique"`
	Password string
}

// Category model
type Category struct {
	BaseModel
	Name        string   `gorm:"size:100"`
	Icon        string   `gorm:"size:255"` // Assuming this is a URL or file path
	Description string   `gorm:"size:255"`
	Tags        []string `gorm:"type:text"` // Tags as a comma-separated string
}

// Currency model
type Currency struct {
	BaseModel
	Code            string  `gorm:"size:3;primaryKey"`
	Name            string  `gorm:"size:100"`
	LatestFxRateUSD float64 `gorm:"default:1.0"`
	IsBaseCurrency  bool    `gorm:"default:false"`
}

// Settlement model
type Settlement struct {
	BaseModel
	SettledByID    *uint     `gorm:"index"` // Foreign key for User
	SettledBy      User      `gorm:"foreignKey:SettledByID"`
	SettlementDate time.Time `gorm:"autoCreateTime"`
	Notes          string    `gorm:"size:255"`
}

// Expense model
type Expense struct {
	BaseModel
	Title         string     `gorm:"size:200"`
	Description   string     `gorm:"size:255"`
	Amount        float64    `gorm:"type:decimal(10,2);not null"`
	CurrencyID    string     `gorm:"index"`
	Currency      Currency   `gorm:"foreignKey:CurrencyID"`
	Notes         string     `gorm:"size:255"`
	CategoryID    *uint      `gorm:"index"` // Foreign key for Category
	Category      Category   `gorm:"foreignKey:CategoryID"`
	SettlementID  *uint      `gorm:"index"` // Foreign key for Settlement
	Settlement    Settlement `gorm:"foreignKey:SettlementID"`
	CreatedByTask bool       `gorm:"default:false"`
	CreatedByID   *uint      `gorm:"index"` // Foreign key for User
	CreatedBy     User       `gorm:"foreignKey:CreatedByID"`
	UpdatedByID   *uint      `gorm:"index"` // Foreign key for User
	UpdatedBy     User       `gorm:"foreignKey:UpdatedByID"`
}

// ExpenseOwed model
type ExpenseOwed struct {
	BaseModel
	ExpenseID  uint     `gorm:"index"` // Foreign key for Expense
	Expense    Expense  `gorm:"foreignKey:ExpenseID"`
	UserID     uint     `gorm:"index"` // Foreign key for User
	User       User     `gorm:"foreignKey:UserID"`
	SplitType  string   `gorm:"size:20"` // Use enum values: "PCT" or "AMT"
	Amount     float64  `gorm:"type:decimal(10,2);not null"`
	Percentage float64  `gorm:"type:decimal(5,2);not null"`
	CurrencyID string   `gorm:"index"` // Foreign key for Currency
	Currency   Currency `gorm:"foreignKey:CurrencyID"`
}
