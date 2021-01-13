/*
2 @Auth: Linux
3 @Date: 2021/1/13 14:02
4 */
package main

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price int    `json:"price"`
}

type pro struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Code      string    `json:"code"`
	Price     int       `json:"price"`
}
