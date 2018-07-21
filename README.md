# first-go-server
A basic server written in Go. The purpose of this project is just to practice writing Go code and creating servers with Go.

## Usage
`cd` to the project's folder and type `./run [your-username] [your-username's password]` where the username, and password are your MySQL database credentials.
It is assumed you have a table called `godb` in your MySQL database. If you don't, run the statement `create table People( name varchar(50), age INT(3) );` in the MySQL database.