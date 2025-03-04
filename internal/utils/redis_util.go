package utils

import "fmt"

func GenerateUserRedisKey(username, ip string) string {
	return fmt.Sprintf("%v:%v", username, ip)
}
