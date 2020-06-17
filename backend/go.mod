module github.com/JoeyPilla/rest-api-example

replace github.com/JoeyPilla/rest-api-example/api => ./api

replace github.com/JoeyPilla/rest-api-example/raspberrypi => ./raspberrypi

replace github.com/JoeyPilla/rest-api-example/ingredient => ./ingredient

replace github.com/JoeyPilla/rest-api-example/recipe => ./recipe

replace github.com/JoeyPilla/rest-api-example/drink => ./drink

replace github.com/JoeyPilla/rest-api-example/global => ./global

go 1.13

require (
	github.com/JoeyPilla/rest-api-example/api v0.0.0-20200509220308-0d67340149da
	github.com/JoeyPilla/rest-api-example/drink v0.0.0-00010101000000-000000000000 // indirect
	github.com/JoeyPilla/rest-api-example/global v0.0.0-00010101000000-000000000000
	github.com/JoeyPilla/rest-api-example/ingredient v0.0.0-00010101000000-000000000000
	github.com/JoeyPilla/rest-api-example/recipe v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/jsha/minica v1.0.2 // indirect
	github.com/lib/pq v1.5.2
	github.com/stianeikeland/go-rpio v4.2.0+incompatible
)
