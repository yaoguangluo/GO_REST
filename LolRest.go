package main

import (
	"github.com/backend/code/rest/lib"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"

	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	//"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/martini-contrib/render"
)

var (
	ac    ApiConfig
	x     lib.Orm
	m     *martini.Martini
	fconf *string = flag.String("c", "C:/Users/comforx/Desktop/go/gitHub/src/github.com/backend/code/rest/config.ini", "config file")
)

func main() {
	err := lib.DBC.Read(*fconf)

	orm, err := lib.DBC.InitOrm()
	x = lib.Orm{DB: orm}
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Hello World!")
	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "templates",
		},
	))

	m.Get("/", func(r render.Render, req *http.Request) {
		if req.URL.Query().Get("wait") != "" {
			sleep, _ := strconv.Atoi(req.URL.Query().Get("wait"))
			time.Sleep(time.Duration(sleep) * time.Second)
		}
		r.HTML(200, "index", nil)
	})

	m.Get("/rest", func(r render.Render) {
		//readmysql
		fmt.Println("Hello World!1")
		str := "Vi"
		r.JSON(200, map[string]interface{}{"error": "10001", "msg": x.SelectValuebyName(str)})

	})

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go http.Serve(listener, m)
	log.Println("Listening on 0.0.0.0:" + port)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	listener.Close()
}
