package system_operation_service

func (s *systemOperationService) SetNextDay() (day int) {
	day = s.systemOperationRepository.SetNextDay()
	return
}
