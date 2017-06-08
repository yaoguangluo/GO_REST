package lib

import (
	"fmt"
	//"log"

	"github.com/go-xorm/xorm"
	//"github.com/wsxiaoys/terminal/color"
)

type Orm struct {
	DB *xorm.Engine
}

type lolroler struct {
	Id   int    `xorm:"INT(9)"`
	Name string `'xorm:"VARCHAR(50)"`
}

func (x *Orm) SelectValuebyName(lolname string) int {
	fmt.Println("Hello World!2")
	var lolrolers []lolroler
	err := x.DB.Sql("select l.id as Id,l.name as Name from lolroler as l where l.name = ?", lolname).Find(&lolrolers)
	if err != nil {
		fmt.Println(111)
		fmt.Println(err)
		fmt.Println(222)
		return 0
	}

	return lolrolers[0].Id
}
