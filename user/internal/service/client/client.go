package client

type service struct {
	repos repository
}

func New(repos repository) *service {
	return &service{
		repos: repos,
	}
}
