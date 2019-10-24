package args

// CreateUserInput 사용자 생성 Args
type CreateUserInput struct {
	Input CreateUserArgs
}

// CreateUserArgs ...
type CreateUserArgs struct {
	Email    string
	Password string
	Name     string
}
