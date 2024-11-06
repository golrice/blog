package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logDir := os.Getenv("LOG_FILE_DIR")
	if logDir == "" {
		logDir = "./logs"
		return
	}

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal("Error creating log directory: ", err)
		return
	}

	currentTime := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", logDir, currentTime)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

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
			fmt.Printf("服务器启动中，监听端口：%s\n", port)

			log.Printf("服务器正在端口 %s 上运行...\n", port)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
