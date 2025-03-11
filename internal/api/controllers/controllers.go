package api_controllers

import web_controller "github.com/milfan/go-boilerplate/internal/api/controllers/web"

type (
	Controllers struct {
		WebControllers web_controller.WebControllers
	}
)

func LoadControllers() Controllers {

	return Controllers{
		WebControllers: web_controller.RegisterWebController(),
	}
}
