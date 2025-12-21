package utils

// GetIndexFromUserId generates an index from user ID
// This is a simple implementation - in production you might want to use a more sophisticated hashing
func GetIndexFromUserId(userId string) int64 {
	var hash int64
	for _, char := range userId {
		hash = (hash*31 + int64(char)) & 0x7FFFFFFF // Keep it positive
	}
	return hash
}
