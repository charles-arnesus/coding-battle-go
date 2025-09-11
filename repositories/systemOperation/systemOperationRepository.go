package system_operation_repository

var currentDay int

type systemOperationRepository struct {
}

func NewSystemOperationRepository() *systemOperationRepository {
	return &systemOperationRepository{}
}
