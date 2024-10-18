package validation

import (
	"encoding/json"
	"errors"
	"modulo/src/exeption"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidationError(validation_err error) *exeption.Erro {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return exeption.NewBadRequest("O tipo que quer passar esta diferente do que voce recebe")
	} else if errors.As(validation_err, &jsonValidationError) {
		erroCauses := []exeption.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := exeption.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			erroCauses = append(erroCauses, cause)
		}
		return exeption.NewBadRequestValidation("Erro ao validar o campo", erroCauses)
	} else {
		return exeption.NewBadRequest("Erro ao fazer a transformacao do translate")
	}
}
