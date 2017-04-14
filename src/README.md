# Clean Architecture in Go

[![Go Report Card](https://goreportcard.com/badge/cloudtamer/portal)](https://goreportcard.com/report/cloudtamer/portal)
[![GoDoc](https://godoc.org/cloudtamer/portal?status.svg)](https://godoc.org/cloudtamer/portal)
[![Coverage Status](https://coveralls.io/repos/github/pt-arvind/gocleanarchitecture/badge.svg?branch=master&randid=4)](https://coveralls.io/github/pt-arvind/gocleanarchitecture?branch=master)

Change your $GOPATH to the root of where you install this!

Run it by going to each microservice in the cmd folder and running the main.go

Watch them communicate!


A good example of clean architecture for a web application in Go.

The **domain** folder is for entities without any dependencies.

The **usecase** folder is for business logic that should not change regardless
of the repository or other services below.

The **repository** folder is for only storing and retrieving entities without
any business logic.

The **controller** folder is for the web handlers.

The **lib** folder contains libraries that can be passed in as services to the
use cases and the controllers.

The **lib/boot** folder handles the set up of the services and the route
assignments for the controllers.
