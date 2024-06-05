package routes

import (
	"Cargo_Dash/controllers"
	"Cargo_Dash/middleeware"
	fiber2 "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber2.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleeware.IsAuthenticated)

	app.Put("/api/user/info", controllers.UpdateInfo)
	app.Put("/api/user/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)

	app.Get("/api/cargos", controllers.AllCargo)
	app.Post("/api/cargos", controllers.CreateCargo)
	app.Get("/api/cargos/:id", controllers.GetByCargo)
	app.Get("/api/cargos/tracking_number/:tracking_number", controllers.GetCargo)
	app.Put("/api/cargos/:id", controllers.UpdateCargo)
	app.Delete("/api/cargos/:id", controllers.DeleteCargo)
}
