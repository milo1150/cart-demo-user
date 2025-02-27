package utils

import "golang.org/x/crypto/bcrypt"

// How cost Works
// The bcrypt algorithm uses 2^cost iterations internally.
// A higher cost exponentially increases the hashing time.
// If cost < MinCost (default is 4), it is automatically set to DefaultCost (10 in Go’s bcrypt package).
// 10 (default): A good balance between security and performance.
// 12: Recommended for modern web applications.
// 14+: For highly sensitive systems, but may impact performance.
//
// Example Cost vs. Performance
// Cost Factor	Approximate Time (on modern CPUs)
// 4	~1ms (very fast, not secure)
// 10	~100ms (default, balanced)
// 12	~300ms (more secure)
// 14	~1s (very secure, slower)
func HashPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}

func CheckHashPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
