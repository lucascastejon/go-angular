
# Go - Angular

Prepare melancholia user

Create melancholia user

		CREATE ROLE melancholia LOGIN SUPERUSER;
Set melancholia password

		ALTER USER melancholia WITH PASSWORD 'm1e2l3a4';

This is a test application to communicate Angular and Go

### The Angular application is in the /assets folder

### Run

Just run `go run main.go` and it will start a server at :1357 so enter `localhost:1357` and you'll see the angular application

### What will work

For now it is possible to send reviews for each gem, however it is not listing the reviews yet. To make sure it was saved, refresh the page and check in the reviews tab. You should see a list of all the reviews created.

### Database

It is using sqlite to make it simple to use and test. Each time you run the server, the database file `gowels.db` is deleted and created again

### PS

This Angular application was created following the course from CodeSchool Shaping up with AngularJS, a really good one for beginners like myself.

by Gerep
