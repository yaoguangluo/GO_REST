package lib

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	"github.com/go-xorm/xorm"
)

var (
	DBC DBConf
)

type DBConf struct {
	host     string
	user     string
	pwd      string
	database string
	port     string
	x        *XormConf
}

type XormConf struct {
	ShowSQL      bool
	ShowErr      bool
	ShowInfo     bool
	ShowWarn     bool
	ShowDebug    bool
	MaxConns     int
	MaxOpenConns int
	MaxIdleConns int
}

// 初始化ezgit
// 必须在init函数中执行
func InitGit(file string) {
	var err error
	Git, err = ezgit.NewGitByFile(file)
	if err != nil {
		panic(err)
	}
}

// 初始化数据库设置
// 必须在init函数中执行
func (c *DBConf) Read(file string) error {
	file = "C:/Users/comforx/Desktop/go/gitHub/src/github.com/backend/code/rest/config.ini"
	config, err := goconfig.LoadConfigFile(file)

	if err != nil {
		panic("no database conf file")
	}
	c.host, err = config.GetValue("mysql", "host")
	if err != nil {
		c.host = "35.165.250.72"
	}
	c.user, err = config.GetValue("mysql", "user")
	if err != nil {
		c.user = "root"
	}
	c.pwd, err = config.GetValue("mysql", "pwd")
	if err != nil {
		c.user = "fengyue1985"
	}
	c.database, err = config.GetValue("mysql", "database")
	if err != nil {
		c.database = "good"
	}
	c.port, err = config.GetValue("mysql", "port")
	if err != nil {
		c.port = "3306"
	}
	c.x = new(XormConf)
	c.x.Read(file)
	return nil
}
func (c *XormConf) Read(file string) error {
	file = "C:/Users/comforx/Desktop/go/gitHub/src/github.com/backend/code/rest/config.ini"
	config, err := goconfig.LoadConfigFile(file)
	if err != nil {
		panic("no xorm conf file")
	}
	c.ShowSQL, _ = config.Bool("xorm", "ShowSQL")
	c.ShowErr, _ = config.Bool("xorm", "ShowErr")
	c.ShowInfo, _ = config.Bool("xorm", "ShowInfo")
	c.ShowWarn, _ = config.Bool("xorm", "ShowWarn")
	c.ShowDebug, _ = config.Bool("xorm", "ShowDebug")
	c.MaxConns, err = config.Int("xorm", "MaxConns")
	if err != nil {
		c.MaxConns = 10
	}
	c.MaxOpenConns, err = config.Int("xorm", "MaxOpenConns")
	if err != nil {
		c.MaxOpenConns = 10
	}
	c.MaxIdleConns, err = config.Int("xorm", "MaxIdleConns")
	if err != nil {
		c.MaxIdleConns = 10
	}
	return nil
}

func (c *DBConf) InitOrm() (*xorm.Engine, error) {

	conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.user, c.pwd, c.host, c.port, c.database)
	//conf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", c.user, c.pwd, c.host, c.database)
	orm, err := xorm.NewEngine("mysql", conf)
	orm.ShowSQL = c.x.ShowSQL
	orm.ShowErr = c.x.ShowErr
	orm.ShowInfo = c.x.ShowInfo
	orm.ShowWarn = c.x.ShowWarn
	orm.ShowDebug = c.x.ShowDebug

	orm.SetMaxConns(c.x.MaxConns)
	orm.SetMaxOpenConns(c.x.MaxOpenConns)
	orm.SetMaxIdleConns(c.x.MaxIdleConns)

	return orm, err
}
