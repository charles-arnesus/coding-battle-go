package system_operation_service

type SystemOperationService interface {
	GetCurrentDay() (day int)
	SetNextDay() (day int)
	SetDayToDefault() (day int)
}
