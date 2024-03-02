package internal

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/use_case"
)

func NewContactRepository(db *postgres.Database) repository.ContactRepository {
	return &repository.ContactRepositoryImpl{DB: db}
}

func NewContactUseCase(repo repository.ContactRepository) use_case.ContactUseCase {
    return use_case.ContactUseCaseImpl{ContactRepository: repo}
}

func NewContactDelivery(useCase use_case.ContactUseCase) delivery.ContactDelivery {
	return delivery.NewContactDelivery(useCase)
}
