package expr

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalOperatorEquals(t *testing.T) {
	expr := StringExpr("arg_name").Equals(StringExpr("json"))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_name","==","json"]`, string(byteData))
}

func TestUnmarshalOperatorEquals(t *testing.T) {
	byteData := []byte(`["arg_name","==","json"]`)
	expr := NewExpr()
	Want := StringExpr("arg_name").Equals(StringExpr("json"))
	err := json.Unmarshal(byteData, &expr)
	assert.Nil(t, err, "checking given error")
	assert.Equal(t, expr, Want)
}

func TestOperatorNotEquals(t *testing.T) {
	expr := StringExpr("arg_name").NotEquals(StringExpr("json"))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_name","~=","json"]`, string(byteData))
}

func TestOperatorGreaterThan(t *testing.T) {
	expr := StringExpr("arg_weight").GreaterThan(NumberExpr(10))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_weight",">",10]`, string(byteData))
}

func TestOperatorAnd(t *testing.T) {
	expr := And(
		StringExpr("a").Equals(StringExpr("b")),
		StringExpr("c").NotEquals(StringExpr("d")),
	)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["AND",["a","==","b"],["c","~=","d"]]`, string(byteData))
}

func TestOperatorOr(t *testing.T) {
	expr := Or(
		StringExpr("arg_weight").GreaterThan(NumberExpr(10)),
		StringExpr("arg_height").Not().GreaterThan(NumberExpr(15)),
	)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["OR",["arg_weight",">",10],["arg_height","!",">",15]]`, string(byteData))
}
