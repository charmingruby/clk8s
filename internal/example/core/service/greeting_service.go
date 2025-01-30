package service

import (
	"github.com/charmingruby/bob/internal/example/core/model"
)

type GreetingParams struct {
	Name string
}

type GreetingResult struct {
	ID string
}

func (s *Service) Greeting(params GreetingParams) (GreetingResult, error) {
	example := model.NewExample(model.NewExampleInput{
		Name: params.Name,
	})

	if err := s.exampleRepository.Store(example); err != nil {
		return GreetingResult{}, err
	}

	return GreetingResult{
		ID: example.ID,
	}, nil
}
