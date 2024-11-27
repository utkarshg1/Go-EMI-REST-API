# EMI Calculator REST API

A simple REST API built with [Gin](https://github.com/gin-gonic/gin) to calculate Equated Monthly Installment (EMI) based on loan details.

---

## Endpoints

### 1. **Welcome Endpoint**

- **URL**: `/`
- **Method**: `GET`
- **Description**: Returns a welcome message.

**Sample Response**:

```json
{
  "message": "Welcome to the EMI Calculator API!"
}
```

### 2. **Calculate EMI**

- **URL**: `/calculate-emi`
- **Method**: `POST`
- **Description**: Accepts loan details and returns EMI calculations.

**Request Body**:

```json
{
  "principal": 100000, // Loan amount
  "period": 24, // Loan tenure in months
  "rate_of_interest": 8.1 // Annual interest rate in %
}
```

**Sample Response**:

```json
{
  "data": {
    "EMI": 14911,
    "Interest": 431445,
    "Amount": 1431445
  },
  "message": "EMI Calculations successful"
}
```

---

## How to Run

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/emi-rest-api.git
   cd emi-rest-api
   ```

2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```
   The server will run at `http://localhost:8080`.

---
