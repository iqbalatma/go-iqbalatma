package iqbalatma_go_jwt_authentication

type JWTSubject interface {
	GetSubjectKey() string
}
