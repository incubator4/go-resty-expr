# go-resty-expr

[![GitHub top language](https://img.shields.io/github/languages/top/incubator4/go-resty-expr?style=flat)](https://kotlinlang.org/)
[![GitHub Action](https://github.com/incubator4/mirai-console-saucenao/workflows/Gradle%20CI/badge.svg?branch=main)](https://github.com/<OWNER>/<REPOSITORY>/actions/workflows/<WORKFLOW_FILE>/badge.svg)
[![GitHub](https://img.shields.io/github/license/incubator4/mirai-console-saucenao?style=flat)](https://github.com/incubator4/go-resty-expr/blob/main/LICENSE)
[![visitors](https://visitor-badge.glitch.me/badge?page_id=incubator4.mirai%2Dconsole%2Dsaucenao)]()
go-resty-expr is a toolkit for creating golang struct by expressions from [lua-resty-expr](https://github.com/api7/lua-resty-expr)

---
It includes an expression syntax to avoid use nest interface{} create json like array.

```go
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
```