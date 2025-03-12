package utils

// Min returns the smaller of two integers
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Contains checks if a string slice contains a specific string
func Contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}