package initialization

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	command "github.com/charles-arnesus/coding-battle-go/command/admin"
	handler "github.com/charles-arnesus/coding-battle-go/handlers"
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	flight_repository "github.com/charles-arnesus/coding-battle-go/repositories/flight"
	passenger_repository "github.com/charles-arnesus/coding-battle-go/repositories/passenger"
	authentication_service "github.com/charles-arnesus/coding-battle-go/services/authentication"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func Start() {
	db := dbInitialization()
	if db == nil {
		panic("no database connection")
	}

	flightRepository := flight_repository.NewFlightRepository(db)
	passengerRepository := passenger_repository.NewPassengerRepository()

	authenticationService := authentication_service.NewAuthenticationService(passengerRepository)
	flightService := flight_service.NewFlightService(flightRepository)

	handler := handler.NewHandler()
	handler.RegisterCommand(command.NewRegisterAircraftCommand(flightService))
	handler.RegisterCommand(command.NewAddDestinationCommand(flightService))

	// ini nanti panggil function logged user yang di auth service
	loggedUser := user_model.User{}

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
				fmt.Println("Error reading input:", err)
				continue
			}

			input = strings.TrimSpace(input)
			loginDto := &user_model.LoginDto{}
			if input == "1" {
				loginDto.Role = utils.RoleAdmin
			} else {
				fmt.Print("Enter username: ")
				usernameInput, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}
				loginDto.Username = usernameInput
			}

			// masuk ke command login kirim parameter
			authenticationService.LoginUser(loginDto)
			// hasil dari login di tampung ke loggedUser
			loggedUser = user_model.User{
				Role: utils.RoleAdmin,
			}
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
