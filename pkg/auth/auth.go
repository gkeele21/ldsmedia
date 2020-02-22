package auth

const (
//secret = "superSecretSecret"
)

//JWTPayload holds the payload portion of the web token
type JWTPayload struct {
	UserId int   `json:"uId"`
	Exp    int64 `json:"exp"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAuthForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"userId"`
}
