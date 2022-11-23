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

package main

import (
	"encoding/json"
	"fmt"
	"github.com/incubator4/go-resty-expr/expr"
)

func main() {
	var exprs = expr.And(
		expr.StringExpr("arg_name").Equals(expr.StringExpr("json")),
		expr.Or(
			expr.StringExpr("arg_weight").GreaterThan(expr.NumberExpr(10)),
			expr.StringExpr("arg_height").Not().GreaterThan(expr.NumberExpr(15)),
		),
	).ToArray()

	bytedata, err := json.Marshal(exprs)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytedata))
	// ["AND",["arg_name","==","json"],["OR",["arg_weight","\u003e",10],["arg_height","!","\u003e",15]]]
}
