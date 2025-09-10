package utils

import "fmt"

func ContainsString(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true // Found the string
		}
	}
	return false // String not found
}

func ConvertInputToIDService(input, userRole string) string {
	return fmt.Sprintf("%s_%s", userRole, input)
}
