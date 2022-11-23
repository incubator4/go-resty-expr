package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendExpr(t *testing.T) {
	expr := ArrayExpr(StringExpr("arg_name")).
		AppendExpr(StringExpr("new_string"), StringExpr("new_string_2"))

	want := Expr{Type: TYPE_ARRAY, ArrayVal: []Expr{
		{Type: TYPE_STRING, StringVal: "arg_name"},
		{Type: TYPE_STRING, StringVal: "new_string"},
		{Type: TYPE_STRING, StringVal: "new_string_2"},
	}}
	assert.Equal(t, expr, want)
}

func TestCreateArrayExpr(t *testing.T) {
	expr := StringExpr("arg_name").
		CreateArrayExpr(StringExpr("new_string"), StringExpr("new_string_2"))

	want := Expr{Type: TYPE_ARRAY, ArrayVal: []Expr{
		{Type: TYPE_STRING, StringVal: "arg_name"},
		{Type: TYPE_STRING, StringVal: "new_string"},
		{Type: TYPE_STRING, StringVal: "new_string_2"},
	}}
	assert.Equal(t, expr, want)
}
