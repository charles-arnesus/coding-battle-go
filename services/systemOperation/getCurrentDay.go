package system_operation_service

func (s *systemOperationService) GetCurrentDay() (day int) {
	day = s.systemOperationRepository.GetCurrentDay()
	return
}
