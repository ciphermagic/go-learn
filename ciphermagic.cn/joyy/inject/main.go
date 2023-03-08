package main

import (
	"fmt"
	"github.com/facebookarchive/inject"
)

type UserService struct {
	User map[string]string
}

func newUserService() *UserService {
	return &UserService{
		User: map[string]string{"name": "cipher"},
	}
}

type UserController struct {
	UserService *UserService `inject:""`
}

func newUserController() *UserController {
	return &UserController{}
}

func main() {
	service := newUserService()
	controller := newUserController()

	graph := inject.Graph{}
	if err := graph.Provide(
		&inject.Object{
			Value: service,
		},
		&inject.Object{
			Value: controller,
		},
	); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
	fmt.Println(service.User["name"])
	fmt.Println(controller.UserService.User["name"])
}
