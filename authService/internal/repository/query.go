package repository

var (
	CreateUserQuery     = "INSERT INTO accounts (username, email, password) VALUES ($1,$2,$3);"
	GetUserbyIdQuery    = "SELECT * from accounts WHERE id = $1 ;"
	GetUserbyEmailQuery = "SELECT * from accounts WHERE email = $1 ;"
)
