#Makefile

build:
	/usr/local/go/bin/go build ./datacore/src/datacore/main.go
	mv ./datacore/main ./datacore/build/datacore

docker-datacore:
	docker build -t datacore:1.0 --no-cache -f ./datacore/Dockerfile .

docker-basicpythonprocessor:
	docker build -t basicpythonprocessor:1.0 --no-cache -f ./basicPythonProcessor/Dockerfile .

datarestful-up:
	docker-compose up

datarestful-down:
	docker-compose down

datarestful-show:
	docker-compose logs

datacore-service-up:
	docker-compose up datacore

processor-service-up:
	docker-compose up pythonprocessor

mongo-service-up:
	docker-compose up mongodb

datacore-service-down:
	docker-compose down datacore

processor-service-down:
	docker-compose down pythonprocessor

mongo-service-down:
	docker-compose down mongodb

datacore-service-show:
	docker-compose logs datacore

processor-service-show:
	docker-compose logs pythonprocessor

mongo-service-show:
		docker-compose logs mongodb
