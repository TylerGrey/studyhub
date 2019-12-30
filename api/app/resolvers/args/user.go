package args

// CreateUserInput 사용자 생성 Args
type CreateUserInput struct {
	Input CreateUserArgs
}

// CreateUserArgs ...
type CreateUserArgs struct {
	Email     string
	Password  string
	FirstName *string
	LastName  *string
	Nickname  string
	Mobile    string
	Birth     *string
	Gender    *string
}
