# Project-Based Intern: BTPN Syariah Fullstack Developer

## Description

This is a simple REST-API created using GO. Users can login, register, logout, view user list, update user data, and delete user data. In addition, users can view a list of profile photos, add photos, edit profile photos, and delete profile photos.

## Features

- API Documentation available at [API Documentation](/api-docs.md)

## Technology Used

- GO
- MySQL

## Tools Used

- gorilla/mux
- GORM
- GORM MySQL Driver
- jwt-go
- Go Cryptography
- Package validator

## Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/nikenhpsr/task-5-pbi-btpns-niken.git
```

2. Change into the project directory:

```bash
cd task-5-pbi-btpns-niken
```

3. Install the necessary dependencies:

```bash
go get -u github.com/gorilla/mux gorm.io/gorm gorm.io/driver/mysql golang.org/x/crypto github.com/golang-jwt/jwt/v5 github.com/go-playground/validator/v10
```

4. Create a .env file and add your database configuration:

- **PORT**: Port to be used by your application.
- **JWT_SECRET**: The secret key used to generate and verify JWT tokens.
- **DB_USER**: Your base data username.
- **DB_PASSWORD**: Your basic data password.
- **DB_DATABASE**: Database name to be used by your application.
- **DB_HOST**: Your host base data (eg: `localhost` or IP address).
- **DB_PORT**: The port used by your database (eg: `3306` for MySQL).

Make sure to populate these values with information appropriate to your development environment. All these configurations are very important to run the application properly.

Example `.env`:

```env
PORT=3000
JWT_SECRET=my_secret_key
DB_USER=my_username
DB_PASSWORD=my_password
DB_DATABASE=my_db_name
DB_HOST=localhost
DB_PORT=3306
```

5. Run the project:

```bash
go run main.go
```
