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

func ConvertToNextStatus(status string) (nextStatus string) {
	switch status {
	case SCHEDULED:
		nextStatus = DEPARTED
	case DEPARTED:
		nextStatus = ARRIVED
	case ARRIVED:
		nextStatus = ARRIVED
	default:
		nextStatus = CANCELLED
	}

	return
}
