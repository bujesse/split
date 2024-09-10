package models

import (
	"time"

	_ "gorm.io/driver/sqlite"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	Username string `gorm:"size:100;unique"`
	Email    string `gorm:"size:255;unique"`
	Password string
}

type Category struct {
	BaseModel
	Name        string   `gorm:"size:100"`
	Icon        string   `gorm:"size:255"`
	Description string   `gorm:"size:255"`
	Type        string   `gorm:"size:255"`
	Tags        []string `gorm:"type:text"`
}

type Currency struct {
	Code               string  `gorm:"size:3;primaryKey"`
	Name               string  `gorm:"size:100"`
	LatestFxRateUSD    float64 `gorm:"default:1.0"`
	IsBaseCurrency     bool    `gorm:"default:false"`
	TwoCharCountryCode string  `gorm:"size:2"`
}

type Settlement struct {
	BaseModel
	SettledByID    uint      `gorm:"index"`
	SettledBy      User      `gorm:"foreignKey:SettledByID"`
	Amount         float64   `gorm:"type:decimal(10,2);not null"`
	CurrencyCode   string    `gorm:"size:3;not null;default:'USD'"`
	Currency       Currency  `gorm:"foreignKey:CurrencyCode;references:Code"`
	SettlementDate time.Time `gorm:"autoCreateTime"`
	Notes          string    `gorm:"size:255"`
}

type Expense struct {
	BaseModel
	Title         string         `gorm:"size:200"`
	Description   string         `gorm:"size:255"`
	Amount        float64        `gorm:"type:decimal(10,2);not null"`
	CurrencyCode  string         `gorm:"size:3;not null;default:'USD'"`
	Currency      Currency       `gorm:"foreignKey:CurrencyCode;references:Code"`
	Notes         string         `gorm:"size:255"`
	CategoryID    *uint          `gorm:"index";default:19`
	Category      Category       `gorm:"foreignKey:CategoryID"`
	SettlementID  *uint          `gorm:"index"`
	Settlement    Settlement     `gorm:"foreignKey:SettlementID"`
	PaidByID      uint           `gorm:"index;not null"`
	PaidBy        User           `gorm:"foreignKey:PaidByID"`
	PaidDate      time.Time      `gorm:"autoCreateTime"`
	CreatedByTask bool           `gorm:"default:false"`
	CreatedByID   uint           `gorm:"index;not null"`
	CreatedBy     User           `gorm:"foreignKey:CreatedByID"`
	UpdatedByID   *uint          `gorm:"index"`
	UpdatedBy     User           `gorm:"foreignKey:UpdatedByID"`
	ExpenseSplits []ExpenseSplit `gorm:"foreignKey:ExpenseID"`
}

type SplitType string

const (
	Pct SplitType = "pct"
	Amt SplitType = "amt"
)

type ExpenseSplit struct {
	BaseModel
	ExpenseID    uint      `gorm:"index"`
	Expense      Expense   `gorm:"foreignKey:ExpenseID"`
	UserID       uint      `gorm:"index"`
	User         User      `gorm:"foreignKey:UserID"`
	SplitType    SplitType `gorm:"size:3"`
	SplitValue   float64   `gorm:"type:decimal(10,2);not null"`
	CurrencyCode string    `gorm:"index"`
	Currency     Currency  `gorm:"foreignKey:CurrencyCode;references:Code"`
}

type FxRate struct {
	BaseModel
	FromCurrencyCode string   `gorm:"size:3;not null"`
	ToCurrencyCode   string   `gorm:"size:3;not null"`
	Rate             float64  `gorm:"not null"`
	Date             string   `gorm:"not null"`
	FromCurrency     Currency `gorm:"foreignKey:FromCurrencyCode;references:Code"`
	ToCurrency       Currency `gorm:"foreignKey:ToCurrencyCode;references:Code"`
}
