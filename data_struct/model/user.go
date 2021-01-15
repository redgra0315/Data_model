/*
2 @Auth: Linux
3 @Date: 2021/1/14 14:37
4 */
package model

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}
