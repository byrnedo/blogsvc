default: builddocker

setup:
	go get -u

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main ./main.go

builddocker: buildgo
	docker build -t byrnedo/blogsvc .

dev-env:
	docker rm -f dev_consul dev_registrator
	(cd ./_environments/ && capitan -d up)

run: dev-env
	    docker run \
			-p 8080:80 byrnedo/blogsvc

