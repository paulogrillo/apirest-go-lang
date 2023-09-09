package usecase

import (
    "github.com/paulogrillo/apirest-go-lang/repositories"
)

type CandidateUseCase struct {
	Repository *models.CandidateRepository
}

func NewCandidateUseCase(repository *CandidateRepository) *CandidateUseCase {
	return &CandidateUseCase{
		Repository: repository,
	}
}

func (uc *CandidateUseCase) GetCandidate(ID string) (Candidate, error) {
	return uc.Repository.Get(ID)
}
