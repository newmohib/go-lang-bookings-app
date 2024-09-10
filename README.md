# Booking and Reservations

This is the repository for my bookings and reservations project.

- Built in go Version 1.18.1
- Uses the [chi router](https://github.com/go-chi/chi/v5)
- Uses [alex cdwards SCS](https://github.com/alexedwards/scs/v2) session management
- Uses [nosurf](https://github.com/choi2k/nosurf)

## Project Planning

    1: Deciding what to build
    2: Project Scope
    3: Key functionality

### Ourt Project

    1: Bookings & Reservations
    2: A Bed & Breakfast with two rooms
    3: what do we need to do

### Key Functionality

    1: Showcase the property
    2: Allow for booking a room for one or more night
    3: check a room's availability
    4: Book the room
    5: Notify guest, and notify property owner

### Key Functionality

    1: Have a back end that the owner logs in to
    2: Review existing bookings
    3: Show a calendar of bookings
    4: Change or Cancel a booking

### What will we need

    1: An authentications system
    2: Somewhere to store information (database)
    3: A means of sending notifications (email/text)

### another template which one is better another template

    https://github.com/CloudyKit/jet

### Test coverage view as html into browser
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
