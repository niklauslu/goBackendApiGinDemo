package utils

import uuid "github.com/satori/go.uuid"

func GenerateUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil).String()
	return u1
}
