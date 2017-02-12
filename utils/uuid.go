package utils

import uuid "github.com/satori/go.uuid"

var (
	namespaceUsers, _ = uuid.FromString("1f24dbac-7f88-4937-bdb7-5641bd0831c8")
)

func GenerateUUID(input string) string {
	v5 := uuid.NewV5(namespaceUsers, input).String()
	return v5
}
