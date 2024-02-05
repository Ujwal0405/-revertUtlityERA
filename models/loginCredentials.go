package models

//go:generate ffjson $GOFILE

//LoginCredentials - Login keeps information of all Logins (3_2)
type LoginCredentials struct {
	LoginID     string `json:"loginId" validate:"required"`
	LearnerID   string `json:"learnerId" validate:"required"`
	Password    string `json:"password" validate:"required"`
	LoginState  int    `json:"loginState" validate:"required"`
	LearnerCode string `json:"learnerCode" validate:"required"`
}
