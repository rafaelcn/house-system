::
:: Build script file for windows 10 (only tested on this platform) and go1.13
:: where go must be available system wide (on %PATH%).
::
:: Failing to meet the criteria given above will result in errors.
::
:: Rafael Campos Nunes <rafaelnunes@engineer.com>
::

@echo off

SET CC=go
SET SRC=src/main.go src/server.go src/rest.go src/response.go src/report.go src/entities.go src/page.go src/queries.go src/config.go src/database.go
SET OUT=bin

%CC% build -o %OUT%/rest.exe %SRC%