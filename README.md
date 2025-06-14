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
│   ├── dto/                    # HTTP request and response dto
│   ├── handler/                # HTTP handlers (Gin)
│   ├── model/                  # Struct definitions
│   ├── repository/             # DB logic
│   └── service/                # Business logic
│   └── utils/                  # Utility & helpers
├── config/                     # DB connection setup
├── db/                         # Database schema
│   └── migrations.sql
├── go.mod
└── README.md
```

## 🚀 Getting Started
1. Clone the repo
```bash
git clone https:        //github.com/your-username/avd-simple-transfer.git
cd avd-simple-transfer
```
2. Set up PostgreSQL
Create a database named transfers (or change the name in config/db.go) and run the migration:
```bash
psql -U postgres -d transfers -f db/migrations.sql
```
Ensure your PostgreSQL user and password match the connection string inside config/db.go.

3. Install dependencies
```bash
go mod tidy
```
4. Run the app
```bash
go run server/main.go
```
The server starts on http:        //localhost:8080.

## 📦 API Endpoints
### 📗 Create account
- HTTP Method: POST
- URI: /accounts
- Request body:
```bash
{
	"account_id": 123,
	"initial_balance": "100.23344"
}
```
- Response
	- Http status: 201 on account creation
	- Body: empty body for successful creation of account
	- on error refer to [here](#errorful-response)

### 📘 Get account balance
- HTTP Method: GET
- URI: /accounts/:accountId
	- eg. /accounts/123

- Response Body: 
	- Http status 200 if successful
```bash
{

}
```
### 💸 Submit transaction
- HTTP Method: POST
- URI: /transactions
- Request body:
```bash
{
	"source_account_id": 123,
	"destination_account_id": 456,
	"amount": "100.12345"
}
```
- Response Body: 
	- successful Http status 200
```bash
{

}
```

### ❌ Errorful response
- In case of 4xx or 5xx errors following will be the response body
```bash
{
		"error": "account not found",                                // human readble error message
		"code": "ACCOUNT_NOT_FOUND",                                 // error code for tracing
		"status": 404,                                               // http status depending on type of error occurred
		"request_id": "8b71a739-8dee-4f96-8bca-572326597707"         // request_id
		"details": {                                                 // details
				"AccountID": "is required",
				"InitialBalance": "is required"
		}
}
```

## 🔐 Assumptions
- All accounts use the same currency.
- No authentication or authorization is implemented.
- Transfers are atomic and transactional at the database level.
- Monetary values are stored as NUMERIC(20,8) in PostgreSQL for precision.

## 🛠️ Tech Stack
- Golang (Gin)
- PostgreSQL
- Gorm ORM