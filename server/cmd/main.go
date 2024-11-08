// main.go
package main

import (
	"log"
	"os"

	"github.com/golrice/blog/internal/config"
	"github.com/golrice/blog/internal/logutils"
	"github.com/golrice/blog/internal/server"

	"github.com/urfave/cli/v2"
)

func main() {
	config.LoadEnv()

	file, err := logutils.SetupLogging()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	app := &cli.App{
		Name:  "blog-server",
		Usage: "启动blog的 HTTP 服务器",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "8080",
				Usage:   "指定服务器监听的端口",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			log.Printf("服务器启动中，监听端口：%s\n", port)
			log.Printf("服务器正在端口 %s 上运行...\n", port)
			server.StartServer(port)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
