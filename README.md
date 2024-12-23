# Roshambo

An overengineered Rock, Paper, Scissors game built for fun and experimentation.

## Features

- **GraphQL API**
- **DataLoader and Caching** for efficient data fetching
- **Passwordless Authentication** using magic links
- **JWT** for secure session management
- Custom **React UI Components**
- **gRPC Microservices Architecture**:
  - Gateway
  - Account
  - Profile
  - Stats
  - Mail Service

## Technology Stack

### Frontend

- React
- Apollo GraphQL
- Jest
- Storybook
- Styled Components
- Webpack

### Backend

- Golang
- `gorilla/mux`
- GraphQL
- Postgres
- Redis
- gRPC with Protocol Buffers
- Docker

## Installation and Setup

### Clone the Repository

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

Roshambo use Redis for handle secrets and config variables

```sh
docker-up SRV=roshambo_redis
```

Populate environment variables

```sh
make env
```

`make env` command will look for a file `.env.dev.redis` in the root project. An example for a config file is:

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
```

Run the frontend using `yarn` inside the `app` folder or in the root project using `make`

```sh
yarn dev
```

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
