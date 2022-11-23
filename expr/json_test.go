/*
 * Copyright 2022 incubator4
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package expr

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalStringVal(t *testing.T) {
	expr := StringExpr("arg_name")
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `"arg_name"`, string(byteData))
}

func TestUnmarshalStringVal(t *testing.T) {
	byteData := []byte(`"arg_name"`)
	expr := NewExpr()
	Want := StringExpr("arg_name")
	err := json.Unmarshal(byteData, &expr)
	assert.Nil(t, err, "checking given error")
	assert.Equal(t, expr, Want)
}

func TestMarshalIntNumberVal(t *testing.T) {
	expr := NumberExpr(15)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `15`, string(byteData))
}

func TestUnmarshalIntNumberVal(t *testing.T) {
	byteData := []byte(`15`)
	expr := NewExpr()
	Want := NumberExpr(15)
	err := json.Unmarshal(byteData, &expr)
	assert.Nil(t, err, "checking given error")
	assert.Equal(t, expr, Want)
}

func TestMarshalFloatNumberVal(t *testing.T) {
	expr := NumberExpr(1234.5678)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `1234.5678`, string(byteData))
}

func TestUnmarshalFloatNumberVal(t *testing.T) {
	byteData := []byte(`1234.5678`)
	expr := NewExpr()
	Want := NumberExpr(1234.5678)
	err := json.Unmarshal(byteData, &expr)
	assert.Nil(t, err, "checking given error")
	assert.Equal(t, expr, Want)
}

func TestMarshalOperatorVal(t *testing.T) {
	expr := OperatorExpr(`<`)
	byteData, err := json.Marshal(&expr)
	assert.Nil(t, err, "checking given error")
	assert.JSONEq(t, `"<"`, string(byteData))
}

func TestUnmarshalOperatorVal(t *testing.T) {
	byteData := []byte(`"<"`)
	expr := NewExpr()
	Want := OperatorExpr(`<`)
	err := json.Unmarshal(byteData, &expr)
	assert.Nil(t, err, "checking given error")
	assert.Equal(t, expr, Want)
}
