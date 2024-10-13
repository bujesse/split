package helpers

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func StringToUintPointer(s string) (*uint, error) {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil, err
	}
	u := uint(val)
	return &u, nil
}

func StringToUint(s string) (uint, error) {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

// Convert example: 2024-09-14T13:53:00.000Z to server time
// Fallback to date-only format (YYYY-MM-DD)
func ParseDate(dateStr string) (*time.Time, error) {
	if parsedDate, err := time.Parse(time.RFC3339, dateStr); err == nil {
		localTime := parsedDate.In(time.Local)
		return &localTime, nil
	}

	if parsedDate, err := time.Parse("2006-01-02", dateStr); err == nil {
		return &parsedDate, nil
	}

	return nil, fmt.Errorf("Invalid date format: %s", dateStr)
}

func DeepCopyMap[K comparable, V any](original map[K]V) map[K]V {
	copy := make(map[K]V)
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func GetContextUserID(ctx context.Context) string {
	if username, ok := ctx.Value("currentUserID").(string); ok {
		return username
	}
	return ""
}
