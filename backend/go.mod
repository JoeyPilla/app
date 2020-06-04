module github.com/JoeyPilla/rest-api-example

replace github.com/JoeyPilla/rest-api-example/api => ./api

replace github.com/JoeyPilla/rest-api-example/raspberrypi => ./raspberrypi

go 1.13

require (
	github.com/JoeyPilla/rest-api-example/api v0.0.0-20200509220308-0d67340149da
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.5.2
	github.com/stianeikeland/go-rpio v4.2.0+incompatible
)
