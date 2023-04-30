// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"regexp"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/function/shared"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Querystring_regfilter_except_Name = "querystring.regfilter_except"

var Querystring_regfilter_except_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Querystring_regfilter_except_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Querystring_regfilter_except_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_regfilter_except_ArgumentTypes[i] {
			return errors.TypeMismatch(Querystring_regfilter_except_Name, i+1, Querystring_regfilter_except_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.regfilter_except
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-regfilter-except/
func Querystring_regfilter_except(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_regfilter_except_Validate(args); err != nil {
		return value.Null, err
	}

	v := value.Unwrap[*value.String](args[0])
	name := value.Unwrap[*value.String](args[1])

	query, err := shared.ParseQuery(v.Value)
	if err != nil {
		return value.Null, errors.New(
			Querystring_regfilter_except_Name, "Failed to parse query: %s, error: %s", v.Value, err.Error(),
		)
	}

	var matchErr error
	query.Filter(func(key string) bool {
		matched, err := regexp.MatchString(name.Value, key)
		if err != nil {
			matchErr = errors.New(
				Querystring_regfilter_except_Name, "Invalid regexp pattern: %s, error: %s", name.Value, err.Error(),
			)
		}
		return matched
	})

	if matchErr != nil {
		return value.Null, matchErr
	}
	return &value.String{Value: query.String()}, nil
}
