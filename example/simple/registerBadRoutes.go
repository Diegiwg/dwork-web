package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
	"github.com/Diegiwg/dwork-web/dw/logger"
)

func registerBadRoutes(app *dworkweb.App) {

	str := "Error reported successfully!"

	// PathAlreadyExist
	if err := app.GET("/", nil); err != nil {
		logger.Info(str)
	}

	if err := app.GET("/about", nil); err != nil {
		logger.Info(str)
	}

	if err := app.GET("/user/<int:id>", nil); err != nil {
		logger.Info(str)
	}

	// ParamsConflict
	if err := app.GET("/user/<uuid:id>/all", nil); err != nil {
		logger.Info(str)
	}

	if err := app.GET("/user/<int:id>/project/<uuid:name>", nil); err != nil {
		logger.Info(str)
	}

	// RepeatedParameter
	if err := app.GET("/user/<int:id>/project/<int:id>/edit", nil); err != nil {
		logger.Info(str)
	}

	if err := app.GET("/user/<int:id>/project/<string:name>/link/<uuid:id>", nil); err != nil {
		logger.Info(str)
	}

	// InvalidParamType
	if err := app.GET("/user/<null:id>/project", nil); err != nil {
		logger.Info(str)
	}

	// InvalidParamStruct
	if err := app.GET("/user/<int:id>/project/string:name", nil); err != nil {
		logger.Info(str)
	}
}
