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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringExpr(t *testing.T) {
	expr := StringExpr("arg_name")
	want := Expr{Type: TYPE_STRING, StringVal: "arg_name"}
	assert.Equal(t, expr, want)
}

func TestOperatorExpr(t *testing.T) {
	expr := OperatorExpr("has")
	want := Expr{Type: TYPE_OPERATOR, StringVal: "has"}
	assert.Equal(t, expr, want)
}

func TestNumberExpr(t *testing.T) {
	expr := NumberExpr(15)
	want := Expr{Type: TYPE_NUMBER, NumberVal: 15}
	assert.Equal(t, expr, want)
}
