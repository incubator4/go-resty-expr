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

func TestOperatorNotAnd(t *testing.T) {
	expr := NotAnd(
		StringExpr("a").Equals(StringExpr("b")),
		StringExpr("c").NotEquals(StringExpr("d")),
	)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["!AND",["a","==","b"],["c","~=","d"]]`, string(byteData))
}

func TestOperatorNotOr(t *testing.T) {
	expr := NotOr(
		StringExpr("arg_weight").GreaterThan(NumberExpr(10)),
		StringExpr("arg_height").Not().GreaterThan(NumberExpr(15)),
	)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["!OR",["arg_weight",">",10],["arg_height","!",">",15]]`, string(byteData))
}

func TestOperatorRegularMatch(t *testing.T) {
	expr := StringExpr("arg_name").RegularMatch(StringExpr("[a-z]+"))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_name", "~~", "[a-z]+"]`, string(byteData))
}

func TestOperatorInsensitiveRegularMatch(t *testing.T) {
	expr := StringExpr("arg_name").InsensitiveRegularMatch(StringExpr("[a-z]+"))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_name", "~*", "[a-z]+"]`, string(byteData))
}

func TestOperatorIn(t *testing.T) {
	expr := StringExpr("arg_name").In(ArrayExpr(
		StringExpr("1"),
		StringExpr("2"),
	))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["arg_name", "in", ["1","2"]]`, string(byteData))
}

func TestOperatorHas(t *testing.T) {
	expr := StringExpr("graphql_root_fields").Has(StringExpr("repo"))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["graphql_root_fields", "has", "repo"]`, string(byteData))
}

func TestOperatorIPMatch(t *testing.T) {
	expr := StringExpr("remote_addr").IPMatch(ArrayExpr(
		StringExpr("127.0.0.1"),
		StringExpr("192.168.0.0/16"),
	))
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `["remote_addr", "ipmatch", ["127.0.0.1","192.168.0.0/16"]]`, string(byteData))
}
