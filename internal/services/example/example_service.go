package example

import "log"

type ExampleService struct{}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) DoSomething() {
	log.Println("Service is doing something")
}
