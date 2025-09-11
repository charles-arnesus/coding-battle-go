package system_operation_service

import system_operation_repository "github.com/charles-arnesus/coding-battle-go/repositories/systemOperation"

type systemOperationService struct {
	systemOperationRepository system_operation_repository.SystemOperationRepository
}

func NewSystemOperationService(systemOperationRepository system_operation_repository.SystemOperationRepository) *systemOperationService {
	return &systemOperationService{
		systemOperationRepository: systemOperationRepository,
	}
}
