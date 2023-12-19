module main

go 1.18

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/lib/pq v1.10.6
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
