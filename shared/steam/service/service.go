package service

type Service interface {
	Interface() string
	CallName() string
	Version() string
	Parameters(params []string) string
}
