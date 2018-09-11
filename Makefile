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

show-datacore-service:
	docker-compose logs datacore

show-processor-service:
	docker-compose logs pythonprocessor

show-processor-service:
		docker-compose logs pythonprocessor

show-mongo-service:
		docker-compose logs mongodb

clean:
	rm -rf ./output/*
	rm ./generadorP3
