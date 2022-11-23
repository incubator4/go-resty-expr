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

func StringExpr(str string) Expr {
	return Expr{
		Type:      TYPE_STRING,
		StringVal: str,
	}
}

func OperatorExpr(o string) Expr {
	return Expr{
		Type:      TYPE_OPERATOR,
		StringVal: o,
	}
}

func NumberExpr(number float64) Expr {
	return Expr{Type: TYPE_NUMBER, NumberVal: number}
}

func ArrayExpr(expr ...Expr) Expr {
	return Expr{Type: TYPE_ARRAY, ArrayVal: expr}
}
