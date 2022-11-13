package gql

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func EncodeIntID(filedName string, rawID int) string {
	if strings.Contains(filedName, ":") {
		panic(fmt.Sprintf("Invalid field name: %s", filedName))
	}
	format := fmt.Sprintf("%s:%d", filedName, rawID)
	return base64.StdEncoding.EncodeToString([]byte(format))
}

func EncodeStrID(filedName string, rawID string) string {
	if strings.Contains(filedName, ":") {
		panic(fmt.Sprintf("Invalid field name: %s", filedName))
	}
	format := fmt.Sprintf("%s:%s", filedName, rawID)
	return base64.StdEncoding.EncodeToString([]byte(format))
}

func DecodeIntID(graphID string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(graphID)
	if err != nil {
		return 0, err
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return 0, err
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DecodeStrID(graphID string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(graphID)
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return "", err
	}

	return parts[1], nil
}

func DecodeStrIDPtr(graphID *string) (*string, error) {
	if graphID == nil {
		return nil, nil
	}
	dst, err := DecodeStrID(*graphID)
	if err != nil {
		return nil, err
	}
	return &dst, nil
}

func DecodeIntIDPtr(graphID *string) (*int, error) {
	if graphID == nil {
		return nil, nil
	}
	dst, err := DecodeIntID(*graphID)
	if err != nil {
		return nil, err
	}
	return &dst, nil
}
