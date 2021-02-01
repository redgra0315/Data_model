/*
2 @Auth: Linux
3 @Date: 2021/1/27 16:32
4 */
package Dingtalk

import (
	"Data_model/dingding_check_port/config"
	"Data_model/dingding_check_port/url"
	"fmt"
	"github.com/CatchZeng/dingtalk"
)

// 发送告警
func DingtalkMessage() {
	_, err := config.CheckPath()
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Conf.Address)
	fmt.Println(config.Conf.Token)
	client := dingtalk.NewClient(config.Conf.Token, config.Conf.Secret)
	title := "服务告警"
	text := url.Demo1()
	var d = []string{"18787072734"}
	mes := dingtalk.NewMarkdownMessage().SetMarkdown(title, text).SetAt(d, true)
	client.Send(mes)
}
