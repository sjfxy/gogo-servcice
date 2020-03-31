package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	service "github.com/sjfxy/gogo-service/service"
)

//上面的处理方式仅仅做测试的使用方式
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
