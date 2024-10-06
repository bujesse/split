package helpers

import (
	"context"
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
func ConvertToServerTime(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	serverLocation := time.Local
	localTime := parsedTime.In(serverLocation)

	return localTime, nil
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
