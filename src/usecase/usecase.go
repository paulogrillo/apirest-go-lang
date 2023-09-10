package usecase

import (
	"../repository"
	"../models"
)

type CandidateUseCase struct {
	Repository repository.CandidateRepository
}

func NewCandidateUseCase(repository repository.CandidateRepository) *CandidateUseCase {
	return &CandidateUseCase{
		Repository: repository,
	}
}

func (uc *CandidateUseCase) GetCandidate(ID string) (models.Candidate, error) {
	return uc.Repository.Get(ID)
}
