package client

type service struct {
	sessionExpiredTime int64
	repos              repository
}

func New(repos repository) *service {
	return &service{
		repos: repos,
	}
}
