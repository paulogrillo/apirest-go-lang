package main

import (
	"fmt"
    "github.com/paulogrillo/apirest-go-lang/crm/models"
    "github.com/paulogrillo/apirest-go-lang/crm/repositories"
    "github.com/paulogrillo/apirest-go-lang/crm/usecase"
)

func main() {
	repository := repositories.repositoryNewCandidateRepository()
	usecase := usecase.NewCandidateUseCase(repository)

	candidate1 := models.Candidate{ID: "1", Name: "João", Age: 30, City: "São Paulo"}
	candidate2 := models.Candidate{ID: "2", Name: "Maria", Age: 25, City: "Rio de Janeiro"}

	repositories.Create(candidate1)
	repositories.Create(candidate2)

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
