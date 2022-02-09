package utils

import "fmt"

// GetAvatar300x300URL ...
func GetAvatar300x300URL(mediaID string) string {
	return fmt.Sprintf("https://shelfish-storage-resized.s3.us-east-1.amazonaws.com/avatars/300x300/%s.jpg", mediaID)
}

// GetAvatar600x600URL ...
func GetAvatar600x600URL(mediaID string) string {
	return fmt.Sprintf("https://shelfish-storage-resized.s3.us-east-1.amazonaws.com/avatars/600x600/%s.jpg", mediaID)
}

// GetAvatar500x500URL ...
func GetAvatar500x500URL(mediaID string) string {
	return fmt.Sprintf("https://createhq-storage-resized.s3.us-east-1.amazonaws.com/avatars/500x500/%s.jpg", mediaID)
}