# coding-battle-go
Coding Battle - GO!

# Prerequisites
There are tools that needs to be prepared first before running the application:
- If you are in Windows, please make sure that wsl is installed, you can install it via PowerShell, and then install it with:
```
wsl --install
```
- After that, please make sure that docker and docker-compose already installed, you can install it with:
```
sudo snap install docker
```

# How to run the app
This app will run on CLI, make sure you are in this root project folder. You can start the application with:
```
sudo docker-compose run --rm coding-battle-go-app
```
This command will first start a postgres instance, and then it will run the app on the CLI

# Application Features
This app consist of 2 roles, Admin and Passenger, where each roles has different features:
## Admin
These are admin features:
- Register Aircraft
- Add flight destination
- Create flight route
- Enable/Disable Booking Service
- Go to Next Day
- Run Flight
### Register Aircraft
This feature allows admin to register aircraft to be used by the flight.

Admin will have to fill in the aircraft name and seat capacity.

### Add flight destination
This feature allows admin to register flight destination, and later to be used by the flight route

### Create flight route
This feature allows admin to create a new flight route. This flight route will be based on the available aircraft and flight destination.

Admin will have to choose:
- Departure city
- Destination city
- Aircraft
- Departure day (2-365)
- Departure time (Morning / Evening)

### Enable/Disable Booking Service
This feature allows admin to enable or disable booking service.

If booking service is disabled, passenger will not be able to make a booking flight. The default value of booking service is disabled.

### Go to Next Day
This feature allows admin to advance system days to the next day (e.g. if current day is 1, if admin use this feature, the current day will be changed to 2)

After advancing to the next day, the system will shows scheduled flight for the current day.

This feature will also remind admin if there are flights that has not been executed yet, and admin can choose whether to cancel the flight that has not been executed yet or not

### Run Flight

This feature allows admin to execute all the flight that has been scheduled.

It will change the flight Status from SCHEDULED -> DEPARTED -> ARRIVED.

The flight mechanism in this application is when a flight is scheduled in the MORNING time, it will always arrived at EVENING time in the same day. If a flight is scheduled in the EVENING time, it will always arrived at MORNING time in the next day.

Example:
- Departure 1: Jakarta -> Bandung, Day 1, MORNING 
- Arrival 1: Jakarta -> Bandung, Day 1, EVENING
- Departure 2: Bandung -> Jakarta Day 2, EVENING
- Arrival 2: Bandung -> Jakarta Day 3, MORNING

If a flight is scheduled on EVENING time, it will only show the DEPARTURE status, the ARRIVAL status will be shown if the admin Run the flight in the next day.

## Passenger
Passenger needs to be logged in to the application if they want to book a flight.

If the passenger is not registered yet, they will be automatically registered, and the application will ask their name.

If the booking service is disabled, passenger will not be able to book or cancel a flight.

These are passenger features:
- Book a Flight
- Cancel a Flight

### Book a Flight
This feature allows passengers to book a flight.
Passenger will need to choose:
- Departure City
- Destination City
- Departure Day
- Departure Time

And after that, the system will search for the first found flight based on the criterias, the aircraft will be automatically chosen for the passenger.

The system will prioritize direct flight first, if there is no direct flight available, the system will try to find a transit route available based on the criteria.

After that, passenger will confirm the booking.

Booking can only be done for maximum D-1 flights.

### Cancel a Flight
This feature allows passenger to cancel a flight.

The system will show booked flight for the passenger, and then the passenger will have to insert the booking Id to cancel the flight.

Booking cancellation can be done for maximum D-1 flights
