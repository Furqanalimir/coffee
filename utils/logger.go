package utils

import "log"

func LogError(msg any) {
	log.Printf("Error: %v", msg)
}
