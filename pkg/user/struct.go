package user

type UserID struct {
	ID int `json:"id" db:"id"`
}

type LoginMailInfo struct {
	ID       int    `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
}

type LoginMailRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginGoogleRequest struct {
	Email string `json:"email" db:"email"`
}

type RegisterUserRequest struct {
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	Name         string `json:"name" db:"name"`
	RegisterType int    `json:"register_type" db:"register_type"`
}

type AuthCodeRequest struct {
	Email    string `json:"email" db:"email"`
	AuthCode int    `json:"auth_code" db:"auth_code"`
}
