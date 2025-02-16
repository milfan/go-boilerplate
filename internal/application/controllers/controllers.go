package controllers

import (
	mobile_v1_controller "github.com/milfan/golang-gin/internal/application/controllers/mobile/v1"
	web_v1_controller "github.com/milfan/golang-gin/internal/application/controllers/web/v1"
)

type (
	Controllers struct {
		// v1
		V1Controller V1Controllers
	}

	V1Controllers struct {
		MobileControllers MobileV1Controllers
		WebControllers    WebV1Controllers
	}

	MobileV1Controllers struct {
		EmployeeController mobile_v1_controller.IEmployeeController
	}

	WebV1Controllers struct {
		EmployeeController web_v1_controller.IEmployeeController
	}
)

func LoadControllers() Controllers {

	mobileV1Controller := MobileV1Controllers{
		EmployeeController: mobile_v1_controller.NewEmployeeController(),
	}

	webV1Controllers := WebV1Controllers{
		EmployeeController: web_v1_controller.NewEmployeeController(),
	}

	return Controllers{
		V1Controller: V1Controllers{
			MobileControllers: mobileV1Controller,
			WebControllers:    webV1Controllers,
		},
	}
}
