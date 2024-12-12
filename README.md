```bash
https://github.com/threatofwar/geekbakes.git
```
```bash
cd geekbakes
```
```bash
go run main.go
```
```bash
curl -X POST http://localhost:8080/login   -H "Content-Type: application/json"   -d '{"username": "user1", "password": "password1"}'
```
```bash
curl -X POST http://localhost:8080/invoices   -H "Authorization: Bearer <your_jwt_token_here>"   -H "Content-Type: application/json"   -d '{                                                   "invoice_number": "INV-001",
    "date": "2024-12-12T00:00:00Z",
    "invoice_to": "Customer Name",
    "description": "Chocolate Chip Cookies",
    "quantity": 1,
    "unit_price": 4.0,
    "payment_method": "Bank Transfer",
    "account_number": "1234567890",
    "pay_by": "2024-12-30"
  }'
```
