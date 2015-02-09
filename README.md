
# Go - Angular

This is a test application to communicate Angular and Go with Sign-up

### The Go application is in the main.go file
### The Angular application is in the /assets folder

### Run

Just run `go run main.go` and it will start a server at :1357 so enter `localhost:1357` and you'll see the angular application


### Database

It is using PostgreSQL on version (9.3.5). 

# requirements
### Install PostgreSQL and Config
    Debian:
		$ sudo apt-get install postgresql pgadmin3

    ArchLinux (manjaro):
		$ sudo pacman -S postgresql pgadmin3

    User postgres
		$ sudo -i -u postgres

    Postgres SQL
		$ psql
		
    Create melancholia user
		$ CREATE ROLE goangular LOGIN SUPERUSER;
		
    Set melancholia password
		$ ALTER USER goangular WITH PASSWORD 'an13gul45ar';
    Create Database
		$ CREATE DATABASE angularjs;

More information: http://bullingox.blogspot.com.br/2014/07/install-postgresql-in-linux.html

### Install GOlang and Config

https://golang.org/doc/install
