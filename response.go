package main

type (
	// Response ...
	Response struct {
		Status uint16
		Content interface{}
	}

	Error struct {
		Code int
		Description string
	}
)

const (
	StatusOk uint16 = 0x0000
	StatusError uint16 = 0x00FF

	ErrorIncompleteRequest int = 0x00F1
	ErrorNotFullfilledRequest int = 0x00F2
	ErrorDatabaseResponse int = 0x00F3
	ErrorNotAuthorized int = 0x00F4
)