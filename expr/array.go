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

func ExprFromStrings(stringArray []string) *Expr {
	e := &Expr{Type: TYPE_ARRAY}
	for _, str := range stringArray {
		e.AppendString(str)
	}
	return e
}

func ExprArrayFromStrings(stringArray []string) []Expr {
	exprArray := []Expr{}
	for _, str := range stringArray {
		exprArray = append(exprArray, StringExpr(str))
	}
	return exprArray
}

// StringArrayExpr function create a new array type Expr by string slice.
func (e *Expr) StringArrayExpr(strs ...string) Expr {
	return ArrayExpr(ExprArrayFromStrings(strs)...)
}

func (e Expr) ToArray() []Expr {
	if e.Type != TYPE_ARRAY {
		panic("Expr must be TYPE_ARRAY")
	} else {
		return e.ArrayVal
	}
}
