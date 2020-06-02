#
# Build make file for Linux and go1.13 where go must be available system wide
# (on $path).
#
# Rafael Campos Nunes <rafaelnunes@engineer.com>
#

CC=`which go`
SRC=src/main.go src/server.go src/rest.go src/response.go src/report.go \
src/entities.go src/page.go src/queries.go src/config.go src/database.go

OUT=bin

all:
	$(CC) build -o "$(OUT)/rest.out" $(SRC)