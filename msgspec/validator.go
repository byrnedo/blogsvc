package msgspec

import validator "gopkg.in/bluesuncorp/validator.v8"

var V *validator.Validate

func init() {
	V = validator.New(&validator.Config{TagName: "validate"})
}
