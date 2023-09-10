package main

import (
	"fmt"
	"../src/repository"
	"../src/usecase"
	"../src/models"
)

func main() {
	repository := *repository.NewCandidateRepository()
	usecase := usecase.NewCandidateUseCase(repository)

	candidate1 := models.Candidate{ID: "1", Name: "João", Age: 30, City: "São Paulo"}
	candidate2 := models.Candidate{ID: "2", Name: "Maria", Age: 25, City: "Rio de Janeiro"}

	repository.Create(candidate1)
	repository.Create(candidate2)

	fmt.Println("Candidato 1:")
	candidate, err := usecase.GetCandidate("1")
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID: %s, Nome: %s, Idade: %d, Cidade: %s\n", candidate.ID, candidate.Name, candidate.Age, candidate.City)
	}

	fmt.Println("Candidato 3:")
	_, err = usecase.GetCandidate("3")
	if err != nil {
		fmt.Println(err)
	}
}
