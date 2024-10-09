# Event Management System

## Prerequisite 🏆
- Go Version `>= go 1.20`
- PostgreSQL Version `>= version 15.0`

## How To Use
There are 2 ways to do running
### With Docker
1. Copy the example environment file and configure it:
  ```bash
  cp.env.example .env
  ```
2. Build Docker
  ```bash
  docker-compose build --no-cache
  ```
3. Run Docker Compose
  ```bash
  docker compose up -d
  ```

### Without Docker
1. Clone the repository or **Use This Template**
  ```bash
  git clone https://github.com/pankop/event-porto.git
  ```
2. Navigate to the project directory:
  ```bash
  cd go-gin-clean-starter
  ```
3. Copy the example environment file and configure it:
  ```bash
  cp .env.example .env
  ```
4. Configure `.env` with your PostgreSQL credentials:
  ```bash
  DB_HOST=localhost
  DB_USER=postgres
  DB_PASS=
  DB_NAME=
  DB_PORT=5432
  ```
5. Open the terminal and follow these steps:
  - If you haven't downloaded PostgreSQL, download it first.
  - Run:
    ```bash
    psql -U postgres
    ```
  - Create the database according to what you put in `.env` => if using uuid-ossp or auto generate (check file **/entity/user.go**):
    ```bash
    CREATE DATABASE your_database;
    \c your_database
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; // remove default:uuid_generate_v4() if you not use you can uncomment code in user_entity.go
    \q
    ```
6. Run the application:
  ```bash
  go run main.go
  ```

## Run Migrations and Seeder
To run migrations and seed the database, use the following commands:

```bash
go run main.go --migrate --seed
```

#### Migrate Database
To migrate the database schema
```bash
go run main.go --migrate
```
This command will apply all pending migrations to your PostgreSQL database specified in `.env`

#### Seeder Database
To seed the database with initial data:
```bash
go run main.go --seed
```
This command will populate the database with initial data using the seeders defined in your application.


### API Documentation
You can explore the available endpoints and their usage in the [Postman Documentation](https://documenter.getpostman.com/view/29665461/2s9YJaZQCG). This documentation provides a comprehensive overview of the API endpoints, including request and response examples, making it easier to understand how to interact with the API.

### Issue / Pull Request Template

The repository includes templates for issues and pull requests to standardize contributions and improve the quality of discussions and code reviews.

- **Issue Template**: Helps in reporting bugs or suggesting features by providing a structured format to capture all necessary information.
- **Pull Request Template**: Guides contributors to provide a clear description of changes, related issues, and testing steps, ensuring smooth and efficient code reviews.
