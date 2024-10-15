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
	Code               string    `gorm:"size:3;primaryKey"`
	Name               string    `gorm:"size:100"`
	LatestFxRateUSD    float64   `gorm:"default:1.0"`
	IsBaseCurrency     bool      `gorm:"default:false"`
	TwoCharCountryCode string    `gorm:"size:2"`
	FxRateUpdatedAt    time.Time `gorm:"autoUpdateTime"`
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
	SettledToZero  bool      `gorm:"default:false"`
}

type Expense struct {
	BaseModel
	Title              string           `gorm:"size:200"`
	Description        string           `gorm:"size:255"`
	Amount             float64          `gorm:"type:decimal(10,2);not null"`
	CurrencyCode       string           `gorm:"size:3;not null;default:'USD'"`
	Currency           Currency         `gorm:"foreignKey:CurrencyCode;references:Code"`
	Notes              string           `gorm:"size:255"`
	CategoryID         *uint            `gorm:"index"`
	Category           Category         `gorm:"foreignKey:CategoryID"`
	PaidByID           uint             `gorm:"index;not null"`
	PaidBy             User             `gorm:"foreignKey:PaidByID"`
	PaidDate           time.Time        `gorm:"autoCreateTime"`
	CreatedByID        uint             `gorm:"index;not null"`
	CreatedBy          User             `gorm:"foreignKey:CreatedByID"`
	UpdatedByID        *uint            `gorm:"index"`
	UpdatedBy          User             `gorm:"foreignKey:UpdatedByID"`
	ExpenseSplits      []ExpenseSplit   `gorm:"foreignKey:ExpenseID"`
	ScheduledExpenseID *uint            `gorm:"index"`
	ScheduledExpense   ScheduledExpense `gorm:"foreignKey:ScheduledExpenseID"`
}

type RecurrenceTypes string

const (
	Daily   RecurrenceTypes = "daily"
	Weekly  RecurrenceTypes = "weekly"
	Monthly RecurrenceTypes = "monthly"
	Yearly  RecurrenceTypes = "yearly"
)

type ScheduledExpense struct {
	BaseModel
	Title        string   `gorm:"size:200"`
	Description  string   `gorm:"size:255"`
	Amount       float64  `gorm:"type:decimal(10,2);not null"`
	CurrencyCode string   `gorm:"size:3;not null;default:'USD'"`
	Currency     Currency `gorm:"foreignKey:CurrencyCode;references:Code"`
	Notes        string   `gorm:"size:255"`
	CategoryID   *uint    `gorm:"index"`
	Category     Category `gorm:"foreignKey:CategoryID"`
	PaidByID     uint     `gorm:"index;not null"`
	PaidBy       User     `gorm:"foreignKey:PaidByID"`
	CreatedByID  uint     `gorm:"index;not null"`
	CreatedBy    User     `gorm:"foreignKey:CreatedByID"`
	UpdatedByID  *uint    `gorm:"index"`
	UpdatedBy    User     `gorm:"foreignKey:UpdatedByID"`

	SplitByID  uint      `gorm:"index"`
	SplitBy    User      `gorm:"foreignKey:SplitByID"`
	SplitType  SplitType `gorm:"size:3"`
	SplitValue float64   `gorm:"type:decimal(10,2);not null"`

	RecurrenceType     RecurrenceTypes `gorm:"size:20;not null"`
	RecurrenceInterval int             `gorm:"not null;default:1"`
	StartDate          time.Time       `gorm:"not null"`
	NextDueDate        *time.Time      `gorm:"index"`
	EndDate            *time.Time      `gorm:"index"`
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
	FromCurrencyCode string    `gorm:"size:3;not null"`
	ToCurrencyCode   string    `gorm:"size:3;not null"`
	Rate             float64   `gorm:"not null"`
	Date             time.Time `gorm:"not null"`
	FromCurrency     Currency  `gorm:"foreignKey:FromCurrencyCode;references:Code"`
	ToCurrency       Currency  `gorm:"foreignKey:ToCurrencyCode;references:Code"`
}
