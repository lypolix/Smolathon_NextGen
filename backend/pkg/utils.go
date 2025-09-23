package pkg

import (
    "crypto/rand"
    "encoding/hex"
    "log"
)

func GenerateRandomString(n int) string {
    bytes := make([]byte, n)
    if _, err := rand.Read(bytes); err != nil {
        log.Fatal(err)
    }
    return hex.EncodeToString(bytes)
}

func Contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}
