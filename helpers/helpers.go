package helpers

import "strconv"

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
