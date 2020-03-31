package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	service "www.cloudnative.com/gogo-servcice/service"
)

//测试部署
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	server := service.NewServer(appEnv)
	server.Run(":" + port)
}
