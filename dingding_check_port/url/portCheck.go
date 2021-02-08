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
	statusCode := 404

	resp, _ := client.Get(urls)

	fmt.Println("hahahaha", resp)
	if resp == nil || resp.StatusCode != 200 {

		address = "Address:" + " " + "-" + " " + config.Conf.Address + "\n"
		service = "Service:" + " " + "-" + " " + config.Conf.Service + "\n"
		status = "Status:" + " " + "-" + " " + strconv.Itoa(statusCode) + "\n"
		message = "Message:" + " " + "-" + " " + "它离开你了 " + "\n"
		port1 = "Port:" + " " + "-" + " " + strconv.Itoa(config.Conf.Port) + "\n"
		//return string(address + "\n" + service + "\n" + port1 + "\n" + message + "\n" + status)
		return string(address + "\n" + service + "\n" + port1 + "\n" + message + "\n" + status)
	} else {
		fmt.Println("服务正常了")
	}
	return
}
