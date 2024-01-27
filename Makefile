# start the environment of FreeCar

#.PHONY: start
#start:
#	docker-compose up -d

# stop the environment of FreeCar

#.PHONY: stop
#stop:
#	docker-compose down


# run the user
.PHONY: user
user:
	go run ./app/user

# run the api
.PHONY: api
api:
	go run ./app/api