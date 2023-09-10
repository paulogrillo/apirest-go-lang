package repository

import (
	"errors"
	"../models"
)

type CandidateRepository struct {
	Candidates map[string]models.Candidate
}

func NewCandidateRepository() *CandidateRepository {
	return &CandidateRepository{
		Candidates: make(map[string]models.Candidate),
	}
}

func (r *CandidateRepository) Create(candidate models.Candidate) error {
	if _, exists := r.Candidates[candidate.ID]; exists {
		return errors.New("Candidato já existe")
	}
	r.Candidates[candidate.ID] = candidate
	return nil
}

func (r *CandidateRepository) Get(ID string) (models.Candidate, error) {
	candidate, exists := r.Candidates[ID]
	if !exists {
		return models.Candidate{}, errors.New("Candidato não encontrado")
	}
	return candidate, nil
}

func (r *CandidateRepository) Update(ID string, candidate models.Candidate) error {
	if _, exists := r.Candidates[ID]; !exists {
		return errors.New("Candidato não encontrado")
	}
	r.Candidates[ID] = candidate
	return nil
}

func (r *CandidateRepository) Delete(ID string) error {
	if _, exists := r.Candidates[ID]; !exists {
		return errors.New("Candidato não encontrado")
	}
	delete(r.Candidates, ID)
	return nil
}
