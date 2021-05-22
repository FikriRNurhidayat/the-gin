package model

type Credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Refresh struct {
	Token string `json:"refresh_token" binding:"required"`
}
