package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"Ontas.Gateway/pkg/auth"
	"Ontas.Gateway/pkg/config"
	"Ontas.Gateway/pkg/order"
	"Ontas.Gateway/pkg/product"

)

func main() {
	c, err := config.LoadConfig()

	if err != nil {	
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
