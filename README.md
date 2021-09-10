# roshambo

An overengineering Rock, Paper, Scissors for fun.

## Roshambo Features

- Passwordless with magic link.
- Own gRPC microservices structure (gateway, account, profile, stats, mail).

## Roshambo Tech Tools

### Frontend

- React
- Styled Components
- Apollo GraphQL
- Storybook
- Jest

### Backend

- Golang
- gorilla/mux
- GraphQL
- Postgress
- Redis
- gRPC and protocol buffers
- Docker

## Installation, Dependencies and Running Roshambo

Clone this repo to your local machine using:

```sh
git clone git@github.com:aitorfernandez/roshambo.git
cd roshambo
```

### Docker containers

#### Databases

One database per microservice

```sh
make docker-up SRV=roshambo_account_postgres
make docker-up SRV=roshambo_profile_postgres
make docker-up SRV=roshambo_stat_postgres
```

Database migrations

```sh
make migrate-account ACTION=up
make migrate-profile ACTION=up
make migrate-stat ACTION=up
```

#### Config

Roshambo use Redis for secret and config variables,

```sh
docker-up SRV=roshambo_redis
```

Populate environment variables
```sh
make env
```

`make env` command will look for a file `.env.dev.redis` in the root of the project. An example of config file is:

```
HSET app env "dev"
HSET app secret "123456"

HSET gateway addr ":4040"

HSET account addr ":5010"
HSET account psql "postgres://postgres:postgres@0.0.0.0:5410/roshambo"

HSET mail addr ":5020"

HSET profile addr ":5030"
HSET profile psql "postgres://postgres:postgres@0.0.0.0:5430/roshambo"

HSET stat addr ":5040"
HSET stat psql "postgres://postgres:postgres@0.0.0.0:5440/roshambo"

HSET gmail addr "smtp.gmail.com:587"
HSET gmail host "smtp.gmail.com"
HSET gmail identity ""
HSET gmail password "123456"
HSET gmail username "rosambo@gmail.com"

HSET jwt secret "123456"
HSET jwt expires "3650"
```

#### Running services

Open a new terminal per service to see output logs

```sh
make run-account
```
```sh
make run-profile
```
```sh
make run-stat
```
```sh
make run-gateway
```

### Frontend

Install dependencies using `yarn`

```sh
cd app
yarn install
yarn dev
```

or in the root of the app
```sh
make run-app
```

## Screenshots

### Get Started

<p align="center">
  <img src="screenshots/get-started.png" alt="GetStarted Screen"/>
</p>

### Game

<p align="center">
  <img src="screenshots/game.png" alt="Game Screen"/>
</p>
