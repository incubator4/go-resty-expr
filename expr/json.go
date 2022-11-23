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
	"fmt"
)

func (expr *Expr) MarshalJSON() ([]byte, error) {
	switch expr.Type {
	case TYPE_NUMBER:
		return json.Marshal(expr.NumberVal)
	case TYPE_STRING, TYPE_OPERATOR:
		return json.Marshal(expr.StringVal)
	case TYPE_ARRAY:
		return json.Marshal(expr.ArrayVal)
	default:
		return nil, fmt.Errorf("cannot match Expr Type with %d", expr.Type)
	}
}

func (expr *Expr) UnmarshalJSON(buf []byte) error {
	var tmp interface{}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if floatVal, floatOk := tmp.(float64); floatOk {
		expr.Type = TYPE_NUMBER
		expr.NumberVal = floatVal
	} else if stringVal, stringOk := tmp.(string); stringOk {
		switch stringVal {
		case "==", "~=", "~~", "~*", "!", "has", "in", "<", ">", "ipmatch", "AND", "OR":
			expr.Type = TYPE_OPERATOR
		default:
			expr.Type = TYPE_STRING
		}
		expr.StringVal = stringVal
	} else if anyArray, arrayOk := tmp.([]interface{}); arrayOk {
		var arrayVal []Expr
		for _, val := range anyArray {
			t, e := json.Marshal(val)
			if e != nil {
				return e
			}
			var tmpExpr = new(Expr)
			e = json.Unmarshal(t, tmpExpr)
			if e != nil {
				return e
			}
			arrayVal = append(arrayVal, *tmpExpr)
		}
		expr.Type = TYPE_ARRAY
		expr.ArrayVal = arrayVal
	} else {
		return fmt.Errorf("unmarshal error")
	}
	return nil
}
