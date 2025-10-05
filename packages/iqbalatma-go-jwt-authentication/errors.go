package iqbalatma_go_jwt_authentication

import "errors"

var ErrInvalidTokenType = errors.New("you are using invalid token type to access this resource")
var ErrMissingRequiredAccessTokenVerifierCookie = errors.New("missing required access token verifier cookie")
var ErrInvalidAccessTokenVerifier = errors.New("invalid access token verifier cookie")
