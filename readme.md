# borgdir.media

A simple web-based App to borrow photo equipment.

This application was build during Summer Semester 2018 in subject web programming with Mr. Luigi Lo Iacono at Th Köln.

## Full acces
  * (User) Benutzername: Heidi Hungrig Passwort: pw
  * (User) Benutzername: Timo Test Passwort: pw
  * (Admin) Benutzername: Adam Admin Passwort: passwort


## ToDOs:

* fix Sorts in Combination with Search and Category
* check if Sorts, Search and Category are working in every template (is working in equipment.tmpl)
* fix header in templates for Gast
* fix issue with profile image when profile data is edited
* fix issue with equipment image when equipment data is edited


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Enjoy experimenting!

### Prerequisites

What things you need to install the software and how to install them

```
Go version 1.9 or above
```

### Installing

Follow these steps:

1. Create the following folder structure in your Go source folder ($GOPATH/src)

```‚
$GOPATH/src/borgdir.media
```

2. Copy the provided sources to the above directory


3. Enter the application's source directory

```
cd $GOPATH/src/borgdir.media
```

4. Install DB driver for SQLite

```
go get "github.com/mattn/go-sqlite3"
```

5. Start the server

```
go run main.go
```

5. Access the web application using your browser: [http://localhost:80/](http://localhost:80/)

## Built With

* [Go](https://www.golang.org/) - The Go Programming Language

## Author

* **Mareike Knost**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
