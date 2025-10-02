package iqbalatma_go_jwt_authentication

type JWTSubject[T any] interface {
	GetSubjectKey() string
}
