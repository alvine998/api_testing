package user

type ServiceUser interface {
}

type serviceUser struct {
	repository Repository
}

func NewServiceUser(repository Repository) *serviceUser {
	return &serviceUser{repository}
}
