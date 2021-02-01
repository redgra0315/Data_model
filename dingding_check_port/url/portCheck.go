/*
2 @Auth: Linux
3 @Date: 2021/1/28 15:30
4 */
package url

import (
	"Data_model/dingding_check_port/config"
	"fmt"
	_ "github.com/Depado/bfchroma"
	_ "github.com/alecthomas/chroma/formatters/html"
	"net/http"
	"strconv"
	"time"
)

//   发送告警结构体
var (
	address string
	service string
	status  string
	message string
	port1   string
	CodeId  int = 200
)

// 实现发送告警的message
func Demo1() (mess string) {
	_, _ = config.CheckPath()
	url := "http://" + config.Conf.Address
	urls := url + ":" + strconv.Itoa(config.Conf.Port) + config.Conf.Path
	fmt.Println(urls)
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, _ := client.Get(urls)
	if resp == nil {
		fmt.Println("服务报错")
		CodeId = 404
	}
	if CodeId != 200 {
		address = "Address:" + " " + "-" + " " + config.Conf.Address + "\t\n"
		service = "Service:" + " " + "-" + " " + config.Conf.Service + "\t\n"
		status = "Status:" + " " + "-" + " " + strconv.Itoa(CodeId) + "\t\n"
		message = "Message:" + " " + "-" + " " + "它离开你了 " + "\t\n"
		port1 = "Port:" + " " + "-" + " " + strconv.Itoa(config.Conf.Port) + "\t\n"
		fmt.Println(strconv.Itoa(CodeId))
		return string(address + "\n" + service + "\n" + port1 + "\n" + message + "\n" + status)
	} else {
		fmt.Println("服务正常了")
	}
	return
}
