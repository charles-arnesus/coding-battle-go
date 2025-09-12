package command

import (
	"fmt"
	"os"
	"time"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type RunFlightCommand struct {
	systemOperation system_operation_service.SystemOperationService
	flightService   flight_service.FlightService
}

func NewRunFlightCommand(systemOperation system_operation_service.SystemOperationService, flightService flight_service.FlightService) *RunFlightCommand {
	return &RunFlightCommand{
		systemOperation: systemOperation,
		flightService:   flightService,
	}
}

func (h *RunFlightCommand) Execute() (err error) {

	currentDay := h.systemOperation.GetCurrentDay()

	fmt.Println("=== RUN FLIGHT ===")
	fmt.Printf("Running flights for day %d...\n", currentDay)

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithConfig(tablewriter.Config{
			Header: tw.CellConfig{
				Formatting: tw.CellFormatting{MergeMode: tw.MergeBoth},
				Alignment:  tw.CellAlignment{Global: tw.AlignCenter},
			},
			Stream: tw.StreamConfig{Enable: true},
		}),
	)

	// Start streaming
	if err := table.Start(); err != nil {
		fmt.Printf("Start failed: %v", err)
	}

	defer table.Close()

	header := []string{"#", "Flight Route", "Aircraft Name", "Passenger", "Total Passenger", "Flight Status"}
	table.Header(header)

	err = h.runFlightRoute(currentDay-1, utils.EVENING, utils.DEPARTED, table)
	if err != nil {
		return
	}

	err = h.runFlightRoute(currentDay, utils.MORNING, utils.SCHEDULED, table)
	if err != nil {
		return
	}

	time.Sleep(2000 * time.Millisecond)

	err = h.runFlightRoute(currentDay, utils.MORNING, utils.DEPARTED, table)
	if err != nil {
		return
	}

	err = h.runFlightRoute(currentDay, utils.EVENING, utils.SCHEDULED, table)
	if err != nil {
		return
	}

	return
}

func (h *RunFlightCommand) ID() string {
	return utils.AdminRunFlightID
}

func (h *RunFlightCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}

func (h *RunFlightCommand) runFlightRoute(currentDay int, time, status string, table *tablewriter.Table) (err error) {
	params := flight_model.GetFlightRouteByRequest{
		DepartureDay: currentDay,
		Status:       status,
	}

	if time != "" {
		params.DepartureTime = time
	}

	if status != "" {
		params.Status = status
	}

	flightRoutes, err := h.flightService.GetFlightRouteByParams(params)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	// show the records to table
	for _, flightRoute := range flightRoutes {
		if len(flightRoute.FlightRouteSeat) == 0 {
			continue
		}

		var passengerStr string
		for i, flightRouteSeat := range flightRoute.FlightRouteSeat {
			passengerStr += fmt.Sprintf("%s (Seat #%d)", flightRouteSeat.User.Name, flightRouteSeat.SeatNumber)
			if i < len(flightRoute.FlightRouteSeat)-1 {
				passengerStr += ",\n"
			}
		}

		route := fmt.Sprintf("%s -> %s (%s)", flightRoute.DepartureCity.Name, flightRoute.DestinationCity.Name, flightRoute.DepartureTime)
		nextStatus := utils.ConvertToNextStatus(flightRoute.Status)
		flightRoute.Status = nextStatus

		// append the record to table
		var record = []string{fmt.Sprintf("Day - %d", currentDay), route, flightRoute.Aircraft.Name, passengerStr, fmt.Sprintf("%d", len(flightRoute.FlightRouteSeat)), nextStatus}
		table.Append(record)

		// update the record
		updatedRecord := flight_model.UpsertFlightRouteRequest{
			FlightRoute: flightRoute,
		}
		_ = h.flightService.UpdateFlightRouteStatus(updatedRecord)
	}

	return
}
