package utils

import "fmt"

func GetUserRedisKey(username, ip string) string {
	return fmt.Sprintf("%v:%v", username, ip)
}
