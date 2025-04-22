package container

import (
	"myapp-go/src/controllers"
	"myapp-go/src/database"
	"myapp-go/src/repositories"
	"myapp-go/src/services"
)
type Container struct {
	UserController   			*controllers.UserController
	AuthController   			*controllers.AuthController
	CustomerController   	*controllers.CustomerController
	BankAccountController   	*controllers.BankAccountController
	PocketController   	*controllers.PocketController
}

func InitContainer() *Container {

	db := database.DB

	// Inisialisasi Repository
	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	bankAccountRepo := repositories.NewBankAccountRepository(db)
	pocketRepo := repositories.NewPocketRepository(db)

	// Inisialisasi Service
	userService := services.NewUserService(userRepo, sessionRepo)
	authService := services.NewAuthService(userRepo, sessionRepo)
	customerService := services.NewCustomerService(customerRepo, userRepo, bankAccountRepo)
	bankAccountService := services.NewBankAccountService(bankAccountRepo, customerRepo)
	pocketService := services.NewPocketService(pocketRepo, bankAccountRepo)

	// Inisialisasi Handler
	userController:= controllers.NewUserController(userService)
	authController:= controllers.NewAuthController(authService)
	customerController:= controllers.NewCustomerController(customerService)
	bankAccountController:= controllers.NewBankAccountController(bankAccountService)
	pocketController:= controllers.NewPocketController(pocketService)

	return &Container{
		UserController:    userController,
		AuthController:    authController,
		CustomerController: customerController,
		BankAccountController: bankAccountController,
		PocketController: pocketController,
	}
}
