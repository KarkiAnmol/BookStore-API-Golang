
# BookStore API

## Overview
The BookStore API is a RESTful service implemented in Go, using Gorilla Mux for routing and GORM for object-relational mapping to a MySQL database. It provides a simple way to manage a collection of books, including operations to create, retrieve, update, and delete books.

## Features
- RESTful endpoints for managing books.
- Integration with MySQL for data persistence.
- Structured using best practices for Go project organization.

## Getting Started

### Prerequisites
- Go (1.14 or newer)
- MySQL (5.7 or newer)

## Installation

### 1. Clone the repository:
```bash
git clone https://github.com/KarkiAnmol/BookStore-API-Golang.git
cd BookStore-API-Golang
```

### 2. Install Go dependencies:
```bash
go mod tidy
```

### 3. Configure MySQL:
Configuring MySQL involves several steps, from installation to setting up users and databases. Below, I'll outline a basic guide to get you started with MySQL configuration, particularly focusing on what's necessary for a development environment like the one you're working with for your Go project.

##### Step 1: Installing MySQL

First, you need to install MySQL. The installation steps vary depending on your operating system. Here's how you can install MySQL on Ubuntu. For other systems, please refer to the official documentation.

```sh
sudo apt update
sudo apt install mysql-server
```

##### Step 2: Securing MySQL

After installing MySQL, it's recommended to run the security script that comes with MySQL. This script will help you secure your MySQL installation.

```sh
sudo mysql_secure_installation
```

This script will prompt you to configure security options, including setting up a root password, removing anonymous users, disallowing root login remotely, and removing the test database.

##### Step 3: Logging into MySQL

To interact with MySQL, you can log in using the MySQL client. If you're using the root user with the `auth_socket` plugin, you don't need a password when logging in from the terminal:

```sh
sudo mysql -u root
```

If you've set a password for the root user or are using another user that requires a password, log in like this:

```sh
mysql -u username -p
```

##### Step 4: Creating a Database

Inside the MySQL client, you can create a new database that your application will use. For example, to create a database named `simplerest`:

```sql
CREATE DATABASE simplerest;
```

##### Step 5: Creating a User and Granting Permissions

For security reasons, it's best not to use the root account for your applications. Instead, create a new user and give it access to the database you just created.

```sql
CREATE USER 'anmol'@'localhost' IDENTIFIED BY 'mypassword123';
GRANT ALL PRIVILEGES ON simplerest.* TO 'anmol'@'localhost';
FLUSH PRIVILEGES;
```

Replace `'anmol'` and `'mypassword123'` with your preferred username and password. This command grants the user `anmol` all privileges on the `simplerest` database, allowing them to execute all operations.

##### Step 6: Exiting MySQL

Once you've set up your database and user, you can exit the MySQL client:

```sql
EXIT;
```

##### Step 7: Updating Your Go Application

Make sure your Go application's database connection string matches the credentials and database name you've just set up. For example, in `app.go`:

```go
gorm.Open("mysql", "anmol:mypassword123@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
```

These steps cover basic MySQL setup and configuration to get you started with developing a Go application that interacts with a MySQL database. For more complex configurations, refer to the MySQL documentation.

### Running the Application

1. Start the server:
```bash
go run main.go
```
The API server will start on `localhost:8083`.

### API Endpoints

| Method | Endpoint            | Description                         |
|--------|---------------------|-------------------------------------|
| POST   | `/book/`            | Create a new book.                  |
| GET    | `/book/`            | Retrieve all books.                 |
| GET    | `/book/{id}`        | Retrieve a book by its ID.          |
| PUT    | `/book/{id}`        | Update a book by its ID.            |
| DELETE | `/book/{id}`        | Delete a book by its ID.            |

## Development

To contribute to this project, please create a fork and submit a pull request.

## License

This project is open-source.
