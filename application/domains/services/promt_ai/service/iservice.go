package service

//go:generate mockgen -package=mocks -destination=../mocks/iservice_mock.go -source=iservice.go
type IService interface {
	Ask(question string) (answer string, err error)
}
