package system_operation_repository

type SystemOperationRepository interface {
	GetCurrentDay() (day int)
	SetNextDay() (day int)
	SetDayToDefault() (day int)
}
