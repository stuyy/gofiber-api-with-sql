package utils

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Validator *validator.Validate
