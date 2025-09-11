package system_operation_service

func (r *systemOperationService) SetDayToDefault() (day int) {
	day = r.systemOperationRepository.SetDayToDefault()
	return
}
