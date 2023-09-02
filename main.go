package main

import (
	"shiftlab-go-rest-api/controller"
	"shiftlab-go-rest-api/db"
	"shiftlab-go-rest-api/repository"
	"shiftlab-go-rest-api/router"
	"shiftlab-go-rest-api/usecase"
	"shiftlab-go-rest-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	workspaceValidator := validator.NewWorkspaceValidator()

	userRepository := repository.NewUserRepository(db)
	workspaceRepository := repository.NewWorkspaceRepository(db)
	shiftRepostitory := repository.NewShiftRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	workspaceUsecase := usecase.NewWorkspaceUsecase(workspaceRepository, workspaceValidator)
	shiftUsecase := usecase.NewShiftUsecase(shiftRepostitory)

	userController := controller.NewUserController(userUsecase)
	workspaceController := controller.NewWorkspaceController(workspaceUsecase)
	shiftController := controller.NewShiftController(shiftUsecase)

	e := router.NewRouter(userController, workspaceController, shiftController)
	e.Logger.Fatal(e.Start(":8080"))
}
