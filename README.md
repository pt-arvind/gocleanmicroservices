# Clean Architecture in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/pt-arvind/gocleanarchitecture)](https://goreportcard.com/report/github.com/pt-arvind/gocleanarchitecture)
[![GoDoc](https://godoc.org/github.com/pt-arvind/gocleanarchitecture?status.svg)](https://godoc.org/github.com/pt-arvind/gocleanarchitecture)
[![Coverage Status](https://coveralls.io/repos/github/pt-arvind/gocleanarchitecture/badge.svg?branch=master&randid=4)](https://coveralls.io/github/pt-arvind/gocleanarchitecture?branch=master)

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