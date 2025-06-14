# 💸 avd-simple-transfer
A simple RESTful API written in Go (Gin) for performing internal financial transactions between accounts.

It allows you to:
 - Create new accounts
 - Check account balances
 - Transfer money between accounts

📁 Project Structure
avd-simple-transfer/
├── server/
│   └── main.go                 # Entry point
├── src/
│   ├── handler/                # HTTP handlers (Gin)
│   ├── model/                  # Struct definitions
│   ├── repository/             # DB logic
│   └── service/                # Business logic
├── config/                     # DB connection setup
├── db/                         # Database schema
│   └── migrations.sql
├── go.mod
└── README.md

## 🚀 Getting Started
1. Clone the repo
```bash
git clone https://github.com/your-username/internal-transfer.git
cd internal-transfer
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
The server starts on http://localhost:8080.

## 📦 API Endpoints
### ➕ Create account
POST /accounts
Request body
{
  "account_id": 123,
  "initial_balance": "100.23344"
}

### 📘 Get account balance
GET /accounts/123
Request body: 
{
  "account_id": 123,
  "balance": "100.23344"
}

### 💸 Submit transaction
POST /transactions
Request body: 
{
  "source_account_id": 123,
  "destination_account_id": 456,
  "amount": "100.12345"
}

## 🔐 Assumptions
All accounts use the same currency.
No authentication or authorization is implemented.
Transfers are atomic and transactional at the database level.
Monetary values are stored as NUMERIC(20,8) in PostgreSQL for precision.

## 🛠️ Tech Stack
- Golang (Gin)
- PostgreSQL
- pgx PostgreSQL driver