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

func NewExpr() Expr {
	return Expr{}
}

// AppendExpr function appends element to the end of this Expr ArrayVal
func (e Expr) AppendExpr(expr ...Expr) Expr {
	if e.Type != TYPE_ARRAY {
		return ArrayExpr(e).AppendExpr(expr...)
	} else {
		return Expr{
			Type:     TYPE_ARRAY,
			ArrayVal: append(e.ArrayVal, expr...),
		}
	}
}

// AppendString function is a simple way to append string variable.
func (e Expr) AppendString(str string) Expr {
	return e.AppendExpr(StringExpr(str))
}

func (e Expr) AppendStrings(strs ...string) {
	for _, str := range strs {
		e.AppendString(str)
	}
}

// CreateArrayExpr function create a new array type Expr contains this expr and input exprs
func (e Expr) CreateArrayExpr(exprArray ...Expr) Expr {
	newExpr := Expr{Type: TYPE_ARRAY, ArrayVal: []Expr{e}}
	return newExpr.AppendExpr(exprArray...)
}
