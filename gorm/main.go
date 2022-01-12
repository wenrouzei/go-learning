package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//create()
	query()
}
func create() {
	var num uint32
	limitCh := make(chan struct{}, 50)
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for i := 0; i <= 5000000; i++ {
		limitCh <- struct{}{}
		go insert(i, db, limitCh, &num)
		if num%1000 == 0 {
			fmt.Println(num)
		}
	}
}

func insert(i int, db *gorm.DB, limitCh <-chan struct{}, num *uint32) {
	defer func() {
		atomic.AddUint32(num, 1)
		<-limitCh
	}()
	menu := Menu{
		ParentId: i,
		Type:     RandomInt(10000, 1000000),
		Order:    RandomInt(50000, 1000000),
		Title:    RandString(10),
		Icon:     RandString(15),
		Uri:      RandString(50),
		Header:   RandString(20),
	}
	if err := db.Table("menu").Create(&menu).Error; err != nil {
		panic(err)
	}
}

type Menu struct {
	Id       uint   `gorm:"column:id" json:"id" form:"id"`
	ParentId int    `gorm:"column:parent_id" json:"parent_id" form:"parent_id"`
	Type     int    `gorm:"column:type" json:"type" form:"type"`
	Order    int    `gorm:"column:order" json:"order" form:"order"`
	Title    string `gorm:"column:title" json:"title" form:"title"`
	Icon     string `gorm:"column:icon" json:"icon" form:"icon"`
	Uri      string `gorm:"column:uri" json:"uri" form:"uri"`
	Header   string `gorm:"column:header" json:"header" form:"header"`
	//CreatedAt int64  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	//UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}

func query() {
	ctx, cancel := context.WithCancel(context.TODO())
	var num uint32
	a := make(chan *Menu, 1000)
	var wg sync.WaitGroup
	for i := 0; i <= 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case row, ok := <-a:
					if ok {
						time.Sleep(10 * time.Millisecond)
						atomic.AddUint32(&num, 1)
						if row.Id < 0 {

						}
					} else {
						return
					}
				}
			}
		}()
	}
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Table("menu").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		select {
		case <-ctx.Done():
			return
		default:
			if num > 0 && num%1000 == 0 {
				fmt.Println(num)
			}
			row := Menu{}
			db.ScanRows(rows, &row)
			a <- &row
		}
	}
	wg.Wait()
	cancel()
}

//CREATE TABLE `menu` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`parent_id` int(11) unsigned NOT NULL DEFAULT '0',
//`type` int(10) unsigned NOT NULL DEFAULT '0',
//`order` int(11) unsigned NOT NULL DEFAULT '0',
//`title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
//`icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
//`uri` varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
//`header` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
