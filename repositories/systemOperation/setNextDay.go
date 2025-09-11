package system_operation_repository

import "github.com/charles-arnesus/coding-battle-go/utils"

func (r *systemOperationRepository) SetNextDay() (day int) {
	currentDay++

	if currentDay == utils.MaxDaysInYear+1 {
		currentDay = 1
	}

	day = currentDay
	return
}
