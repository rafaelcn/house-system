# Rest

This is a smart system to control your house.

# Build

To build the project you first have to check some boxes. As this project was
made using `go1.13` you have to check your go version for that.

With go1.13 installed you either use the `Makefile` on Linux or `build.bat` on
Windows to build the project and that will generate a binary on the bin folder
of the project.

Furthermore, you have to configure the _Rest_ database using a PostgreSQL and
the `.sql` files that are inside the *database* folder.

# Usage

Just execute the binary generated on build and view the project running from
any available device on the network.