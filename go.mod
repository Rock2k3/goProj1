module goProject1

go 1.17

require (
	app2 v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0 // indirect
)

replace app2 => ./internal/app2
