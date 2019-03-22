export MUSKOOTERS_PORT=:3412
export MUSKOOTERS_MYSQL_URL=root:Danicheeta74@tcp(localhost:3306)/muskooters
export MUSKOOTERS_MONGO_USER=root
export MUSKOOTERS_MONGO_PASS=root
export MUSKOOTERS_MONGO_DB_NAME=muskooters
export MUSKOOTERS_MIGRATION_DIR=./migrations
export GOBIN=${PWD}/bin

run:
	go install
	bin/muskooters