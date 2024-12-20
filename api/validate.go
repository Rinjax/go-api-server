package api

import (
	"strings"
	"unicode"
	"github.com/go-playground/validator/v10"
)

//func validate(r any) map[string]string {
	//validate := validator.New(validator.WithRequiredStructEnabled())

	//err :=  validate.Struct(r)
//}

func getJsonKeyName(e validator.FieldError) string {

	s := e.StructNamespace()

	// Find the index of the first dot
	idx := strings.Index(s, ".")

	if idx != -1 {
		// No dot found, use the string as is
		s = s[idx+1:]
	}

	var str strings.Builder
	
	// range over the string and insert underscores into the builder
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			str.WriteRune('_')
		}

		// add the lower version of the rune to the builder
		str.WriteRune(unicode.ToLower(v))
	}

	return str.String()
}