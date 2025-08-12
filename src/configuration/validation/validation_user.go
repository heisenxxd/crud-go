package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
	resterr "github.com/heisenxxd/crud-go/crud-go/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translator, _ = unt.GetTranslator("en")
		en_translator.RegisterDefaultTranslations(val, translator)
	}
}

func ValidateUserError(
	validation_err error ) *resterr.Resterr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid field type")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(translator),
				Field: e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return resterr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}
	return resterr.NewBadRequestError("Error trying to convert fields")
}