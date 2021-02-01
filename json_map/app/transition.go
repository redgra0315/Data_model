/*
2 @Auth: Linux
3 @Date: 2021/1/19 16:25
4 */
package app

import (
	"Data_model/json_map/model"
	"encoding/json"
	"fmt"
)

func JsonMap() {
	u1 := model.UserInfo{Name: "qqq", Age: 20}
	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for i, v := range m {
		fmt.Printf("key:%v values:%v\n", i, v)
	}
}
