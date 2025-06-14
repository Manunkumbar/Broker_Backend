# Broker_Backend
A backend service for a broker platform to manage client holdings, trades, and user operations.  
Built using **Go (Gin)** and **PostgreSQL**.

---

## Project Structure  

```bash
broker-backend/
├── config/
├── controllers/
├── models/
├── routes/
├── services/
├── tests/
├── utils/
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 🚀 Backend Setup  

### 📦 Prerequisites  

- **Go 1.22+**
- **PostgreSQL 15+**

---

### ⚙️ Environment Configuration  

Create a `.env` file or update `config/config.go` with the following environment variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=broker_db
```

---

### 🐘 Database Setup  

Create your database and run the following SQL schema to set up the required tables:

```sql
-- Table: public.users

CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- Table: public.holdings

CREATE TABLE IF NOT EXISTS public.holdings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.users(id),
    stock_symbol TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    average_price NUMERIC(10, 2) NOT NULL
);

-- Table: public.positions

CREATE TABLE IF NOT EXISTS public.positions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.users(id),
    stock_symbol TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    pnl NUMERIC(10, 2) NOT NULL
);

-- Table: public.orderbook

CREATE TABLE IF NOT EXISTS public.orderbook (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.users(id),
    stock_symbol TEXT NOT NULL,
    order_type TEXT NOT NULL CHECK (order_type IN ('BUY', 'SELL')),
    quantity INTEGER NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('OPEN', 'FILLED', 'CANCELLED')),
    created_at TIMESTAMP DEFAULT now()
);
```

---

### Install Go Dependencies  

In the project root directory:

```bash
go mod tidy
```

---

###  Run the Backend  

```bash
go run main.go
```

The API server will be running at:

```
http://localhost:6001
```

---

## Features  

- REST APIs for Holdings, Positions, Order Book, and Users
- PostgreSQL integration
- Clean MVC structure (models, controllers, services)

---
