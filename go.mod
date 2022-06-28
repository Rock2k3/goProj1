module goProject1

go 1.17

require app2 v0.0.0-00010101000000-000000000000

require (
	github.com/golang-migrate/migrate/v4 v4.15.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace app2 => ./internal/app2
