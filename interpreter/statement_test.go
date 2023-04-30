package interpreter

import (
	"testing"

	_ "github.com/k0kubun/pp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/interpreter/value"
)

func TestDeclareStatement(t *testing.T) {
	tests := []struct {
		name    string
		decl    *ast.DeclareStatement
		expect  value.Value
		isError bool
	}{
		{
			name: "INTEGER value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "INTEGER"},
			},
			expect: &value.Integer{},
		},
		{
			name: "FLOAT value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "FLOAT"},
			},
			expect: &value.Float{},
		},
		{
			name: "BOOL value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "BOOL"},
			},
			expect: &value.Boolean{},
		},
		{
			name: "BACKEND value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "BACKEND"},
			},
			expect: &value.Backend{},
		},
		{
			name: "IP value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "IP"},
			},
			expect: &value.IP{},
		},
		{
			name: "STRING value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "STRING"},
			},
			expect: &value.String{},
		},
		{
			name: "RTIME value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "RTIME"},
			},
			expect: &value.RTime{},
		},
		{
			name: "TIME value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "TIME"},
			},
			expect: &value.Time{},
		},
		{
			name: "ACL value declaration",
			decl: &ast.DeclareStatement{
				Name:      &ast.Ident{Value: "var.foo"},
				ValueType: &ast.Ident{Value: "ACL"},
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		ip := New(nil)
		err := ip.localVars.Declare(tt.decl.Name.Value, tt.decl.ValueType.Value)
		if err != nil {
			if !tt.isError {
				t.Errorf("%s: unexpected error returned: %s", tt.name, err)
			}
			continue
		}

		v, err := ip.localVars.Get(tt.decl.Name.Value)
		if err != nil {
			t.Errorf("%s: %s varible must be declared: %s", tt.name, tt.decl.Name.Value, err)
			continue
		}
		assertValue(t, tt.name, tt.expect, v)
	}
}

func TestReturnStatement(t *testing.T) {
	var exp ast.Expression = &ast.Ident{
		Value: "pass",
		Meta:  &ast.Meta{},
	}
	tests := []struct {
		name   string
		stmt   *ast.ReturnStatement
		expect State
	}{
		{
			name: "should return pass state",
			stmt: &ast.ReturnStatement{
				ReturnExpression: &exp,
			},
			expect: PASS,
		},
	}

	for _, tt := range tests {
		ip := New(nil)
		s := ip.ProcessReturnStatement(tt.stmt)
		if s != tt.expect {
			t.Errorf("%s expects state %s, got %s", tt.name, tt.expect, s)
		}
	}
}
