package errors

import "errors"

var LoginCredentialsWrong error = errors.New("User Credentials Wrong.")
var WrongPassword error = errors.New("Your Current Password Is Wrong.")
