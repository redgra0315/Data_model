/*
2 @Auth: Linux
3 @Date: 2021/1/14 14:28
4 */
package Init

import (
	"Data_model/data_struct/model"
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func InitDb() (err error) {
	dsn := "root:123456@tcp(192.168.20.58)/sql_test?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB filed  err:%v\n!", err)
		return
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return err
}

func QueryRowDemo(age int) (err error) {
	sqlstr := "select id, name, age from user where  id=?"
	var u model.User
	err = DB.Get(&u, sqlstr, age)
	if err != nil {
		fmt.Printf("get filed,err %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
	return
}

func InsertData() {
	sqlstr := "insert into user(name, age) values (?,?)"
	ret, err := DB.Exec(sqlstr, "沙河小王子-1", 24)
	if err != nil {
		errors.New("err")
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func Delete(id int) (err1 error) {
	err := QueryRowDemo(id)
	if err != nil {
		fmt.Println("数据不存在")
		os.Exit(0)

	}
	sqlStr := "delete  from user  where id = ?"
	ret, err := DB.Exec(sqlStr, id)
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		errors.New("删除失败")
		return
	}
	fmt.Printf("删除成功，影响行数为:%d", n)
	return
}
func InsertUserDome() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = DB.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "qimi",
			"age":  23,
		})
	fmt.Println("插入成功！")
	return
}
