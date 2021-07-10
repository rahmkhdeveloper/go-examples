package handlers

import (
	"log"

	"github.com/gofiber/fiber"
	"mydomain.com/companyname/gofiber-example/model"
)

func Register(c *fiber.Ctx) {
	c.Send("Hello World!")
}

func GetUser(c *fiber.Ctx) {
	user := new(model.User)   //&{ }
	err := c.BodyParser(user) //&{Rahul Mukherjee}
	if err != nil {
		log.Fatal("Error Occurred!", err)
	}

	var iu model.IUser = user

	iu.WriteMyInitials(iu.GetMyInitials())
	c.Send(user.WriteMyNameAndInitials())
}
