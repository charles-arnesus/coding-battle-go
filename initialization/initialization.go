package initialization

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	admin_command "github.com/charles-arnesus/coding-battle-go/command/admin"
	passenger_command "github.com/charles-arnesus/coding-battle-go/command/passenger"
	handler "github.com/charles-arnesus/coding-battle-go/handlers"
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	booking_repository "github.com/charles-arnesus/coding-battle-go/repositories/booking"
	flight_repository "github.com/charles-arnesus/coding-battle-go/repositories/flight"
	system_operation_repository "github.com/charles-arnesus/coding-battle-go/repositories/systemOperation"
	user_repository "github.com/charles-arnesus/coding-battle-go/repositories/user"
	authentication_service "github.com/charles-arnesus/coding-battle-go/services/authentication"
	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func Start() {
	db := dbInitialization()
	if db == nil {
		panic("no database connection")
	}

	flightRepository := flight_repository.NewFlightRepository(db)
	userRepository := user_repository.NewUserRepository(db)
	bookingRepository := booking_repository.NewBookingRepository(db)
	systemOperationRepository := system_operation_repository.NewSystemOperationRepository()

	authenticationService := authentication_service.NewAuthenticationService(userRepository)
	flightService := flight_service.NewFlightService(flightRepository)
	bookingService := booking_service.NewBookingRepository(bookingRepository, flightRepository)
	systemOperationService := system_operation_service.NewSystemOperationService(systemOperationRepository)

	handler := handler.NewHandler()
	// Register admin command
	handler.RegisterCommand(admin_command.NewRegisterAircraftCommand(flightService))
	handler.RegisterCommand(admin_command.NewAddDestinationCommand(flightService))
	handler.RegisterCommand(admin_command.NewCreateFlightRouteCommand(flightService, systemOperationService))
	handler.RegisterCommand(admin_command.NewSetBookingSystemCommand(bookingService))
	handler.RegisterCommand(admin_command.NewGoToNextDayCommand(systemOperationService, flightService))

	// Register passenger command
	handler.RegisterCommand(passenger_command.NewBookFlightCommand(authenticationService, bookingService, flightService, systemOperationService))
	handler.RegisterCommand(passenger_command.NewCancelFlightCommand(bookingService))

	// ini nanti panggil function logged user yang di auth service
	loggedUser := user_model.User{}
	_ = systemOperationService.SetDayToDefault()

	for {
		reader := bufio.NewReaderSize(os.Stdin, 1)

		// cek apakah ada user yang login
		if loggedUser.Username == "" {

			fmt.Println("Login as:")
			fmt.Printf("1. %s\n", utils.RoleAdminLabel)
			fmt.Printf("2. %s\n", utils.RolePassengerLabel)

			//jika tidak, masuk ke halaman login dan lakukan login
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(utils.ErrInputInvalid)
				continue
			}

			input = strings.TrimSpace(input)
			loginDto := &user_model.LoginDto{}
			switch input {
			case "1":
				loginDto.Role = utils.RoleAdmin
			case "2":
				fmt.Print("Enter username: ")
				usernameInput, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(utils.ErrInputInvalid)
					continue
				}
				loginDto.Username = usernameInput
			default:
				fmt.Println(utils.ErrCommandInvalid)
				continue
			}

			// masuk ke command login kirim parameter
			err = authenticationService.LoginUser(loginDto)
			if err != nil {
				fmt.Println(err.Error())
				if strings.Contains(err.Error(), utils.RecordNotFoundMessage) {
					fmt.Println("username not found, creating new user...")

					fmt.Print("Enter name: ")
					nameInput, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println(utils.ErrInputInvalid)
						continue
					}
					registerUser := user_model.User{
						Username: loginDto.Username,
						Name:     nameInput,
						Role:     utils.RolePassenger,
					}
					err = authenticationService.RegisterUser(registerUser)
					if err != nil {
						fmt.Println("Error creating new user:", err)
						continue
					}
				} else {
					fmt.Println("Error when trying to login:", err)
					continue
				}
			}
			// hasil dari login di tampung ke loggedUser
			loggedUser = authenticationService.GetLoggedUser()
		}

		if loggedUser.Role == utils.RoleAdmin {
			AdminPage()
		} else {
			PassengerPage()
		}

		fmt.Printf("99. %s\n", utils.ExitLabel)
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(utils.ErrInputInvalid)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "99" {
			fmt.Println(utils.ExitSuccessMessage)
			fmt.Println()

			//redirect ke logout
			loggedUser = user_model.User{}
			continue
		}

		err = handler.ExecuteCommand(input, loggedUser.Role)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		fmt.Println()
		fmt.Println()
	}
}

func AdminPage() {
	fmt.Println("========= ADMIN PANEL ==========")

	for idx, v := range utils.MenuAdmin {
		fmt.Printf("%d. %s\n", idx+1, v)
	}
}

func PassengerPage() {
	fmt.Println("========= PASSENGER PANEL ==========")

	for idx, v := range utils.MenuPassenger {
		fmt.Printf("%d. %s\n", idx+1, v)
	}
}
