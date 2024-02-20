# start the environment of FreeCar

.PHONY: start
start:
	docker-compose up -d --build 

# stop the environment of FreeCar

.PHONY: stop
stop:
	docker-compose down


# run the user
.PHONY: admin
admin:
	go run ./app/admin

# run the api
.PHONY: api
api:
	go run ./app/api