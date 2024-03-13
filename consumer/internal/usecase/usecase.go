package usecase

type UsecaseInputDto interface{}

type UsecaseOutputDto interface{}

type Usecase interface {
	Execute(input UsecaseInputDto) (UsecaseOutputDto, error)
}
