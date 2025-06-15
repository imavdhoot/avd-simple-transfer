# 💸 avd-simple-transfer
A simple RESTful API written in Go (Gin) for performing internal financial transactions between accounts.

It allows you to:
 - Create new accounts
 - Check account balances
 - Transfer money between accounts

📁 Project Structure
```
avd-simple-transfer/
├── server/
│   └── main.go                 # Entry point
├── src/
│   ├── constant/               # constant values used in project
│   ├── dto/                    # HTTP request and response dto
│   ├── handler/                # HTTP handlers (Gin)
│   ├── model/                  # Struct definitions
│   ├── repository/             # DB logic
│   └── service/                # Business logic
│   └── utils/                  # Utility & helpers
├── config/                     # DB connection setup
├── db/                         # Database schema
│   └── migrations.sql
├── test/                       # Testsuite
├── go.mod
└── README.md
```

## 🚀 Getting Started
1. Clone the repo from the github and change directory to cloned repo
```bash
git clone https://github.com/imavdhoot/avd-simple-transfer.git
cd avd-simple-transfer
```
2. Set up PostgreSQL
Create a database named transfers (or change the name in config/db.go) and run the migration command below.
```bash
psql -U postgres -d transfers -f db/migrations.sql
```
Ensure your PostgreSQL user and password match the connection string inside .env file and also
Your user have required permissions like CREATE TABLE etc.

3. Install dependencies
```bash
go mod tidy
```
4. Run the app
```bash
go run server/main.go
```
The server starts on http://localhost:8080

## 📦 API Endpoints
If you are running on local then use Host http://localhost:8080
### 📗 1. Create account
- HTTP Method: POST
- URI: /api/v1/accounts
- Request body:
```bash
{
  "account_id": 123,
  "initial_balance": 100.23344
}
```
- Response
  - Http status: 201 on account creation
  - Body: empty body for successful creation of account
  - on error refer to [here](#️-errorful-response)

### 📘 2. Get account balance
- HTTP Method: GET
- URI: /api/v1/accounts/:accountId
  - eg. /api/v1/accounts/123

- Response Body: 
  - Http status 200 if successful
```bash
{
  "account_id": 123,
  "initial_balance": 100.23344
}
```
### 💸 3. Submit transaction
- HTTP Method: POST
- URI: /api/v1/transactions
- Request body:
```bash
{
  "source_account_id": 123,
  "destination_account_id": 456,
  "amount": 100.12345
}
```
- Response Body: 
  - successful Http status 200
```bash
{
  "transaction_id": 4,
  "message": "success",
  "status": 200,
  "created_at": "2025-06-16T00:29:42+08:00",
  "request_id": "78726603-701d-425e-9fa0-f0a96870d1c1"
}
```

### ❌ Errorful Response
- In case of 4xx or 5xx errors following will be the response body
```bash
{
  "error": "account not found",                                // human readble error message
  "code": "ACCOUNT_NOT_FOUND",                                 // error code for tracing
  "status": 404,                                               // http status depending on type of error occurred
  "request_id": "8b71a739-8dee-4f96-8bca-572326597707"         // request_id
  "details": {                                                 // details
    "AccountID": "is required",
  }
}
```

## ✅ Running Tests
We use a dedicated test database (`transfers_test`) for safe, repeatable tests.
Make sure your PostgreSQL user also have required permissions to do these migrations.
For help can use this [cheatsheet for PostgreSQL](https://quickref.me/postgres.html)
#### 1. Create the test database and apply migrations (from terminal):

```bash
# Create test DB (run from your terminal)
createdb transfers_test

# Apply schema to test DB (also run from your terminal)
psql -U postgres -d transfers_test -f db/migrations.sql
```
#### 2. To run unit tests for API endpoints:
```bash
go test ./test/...
```

## 🔐 Assumptions
- All accounts use the same currency.
- No authentication or authorization is implemented.
- Transfers are atomic and transactional at the database level.
- Monetary values are stored as NUMERIC(20,8) in PostgreSQL for precision.
- 'Create Account' API and 'Submit transaction' API supposed to have string amount value as per assignment but used float instead and realized very late in assignment.
- String values are much better to avoid rounding off errors, but needs an extra strconv.ParseFloat on every request
- precision of amounts are still protected with gorm NUMERIC(20,8) fields

## 🛠️ Tech Stack
- Golang (Gin)
- PostgreSQL
- Gorm ORM