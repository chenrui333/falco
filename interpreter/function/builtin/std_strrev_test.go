// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of std.strrev
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrev/
func Test_Std_strrev(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "abc", expect: "cba"},
		{input: "abc日本語", expect: ""},
	}

	for i, tt := range tests {
		ret, err := Std_strrev(
			&context.Context{},
			&value.String{Value: tt.input},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.StringType {
			t.Errorf("[%d] Unexpected return type, expect=STRING, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.String](ret)
		if diff := cmp.Diff(tt.expect, v.Value); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
