package model

type UserRegisterIn struct {
	Channel  int
	Passport string
	Password string
	Confirm  string
}

type UserRegisterOut struct {
}

type UserLoginIn struct {
	Passport string
	Password string
}

type UserLoginOut struct {
}
