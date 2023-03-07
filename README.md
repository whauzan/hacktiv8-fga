# Hacktiv8 Fresh Graduate Academy - Scalable Web Service With Golang
#### Author: [Whauzan](https://www.showwcase.com/whauzan)

This is repository for Fresh Graduate Academy by [Kominfo](https://www.kominfo.go.id/) and [Hacktiv8](https://www.hacktiv8.com/).

## Features

- Hot Reload with [Air](https://github.com/cosmtrek/air) (If run on docker)

## Tech

This repository are using:

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Air](https://github.com/cosmtrek/air)

## Installation

It requires [Go](https://go.dev/) v1.19 or using [Docker](https://www.docker.com/) to run.

Install the dependencies.

```sh
go mod tidy
```

If you want to run the app using [Docker](https://www.docker.com/), skip the previous step.

## Running
```sh
cd cmd
go run main.go
```

If running with docker

```sh
docker compose up
```