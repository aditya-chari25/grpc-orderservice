# Three server Communication (Inventory, Order, User microservice)

This project is a microservices-based e-commerce system built with **Go**, **gRPC**, and a RESTful **API Gateway**.

### Services:
- `InventoryService` – manages products and stock
- `OrderService` – handles order placement
- `UserService` – stores user info
- `API Gateway` – exposes REST endpoints and routes requests to the appropriate services

---

## Project Structure
lfdtassgn/ <br>
├── cmd/<br>
│ ├── api-gateway/ # HTTP API Gateway <br>
│ ├── inventory/ # gRPC InventoryService<br>
│ ├── order/ # gRPC OrderService<br>
│ └── user/ # gRPC UserService<br>
├── internal/<br>
│ ├── inventory/ # Inventory business logic<br>
│ ├── order/ # Order logic and gRPC clients to inventory + user<br>
│ └── user/ # User logic<br>
├── services/<br>
│ ├── inventory/ # gRPC client + HTTP handler<br>
│ ├── order/ # gRPC client + HTTP handler<br>
│ └── user/ # gRPC client + HTTP handler<br>
├── pkg/proto/ # Protobuf files and generated Go code<br>
├── proto/ # this is used for generating the protobuf files<br>

## How Services Inter-communicate

### `OrderService → InventoryService`
- Before placing an order, `OrderService` calls `CheckStock(product_id, quantity)`
- After a successful check, it calls `UpdateStock(product_id, quantity)`
  
### `OrderService → UserService`
- It also calls `GetUser(user_id)` to validate the user exists before creating the order

All these are done using internal gRPC clients injected into the `OrderServer`.


```bash
go run cmd/inventory/main.go 
go run cmd/order/main.go  
go run cmd/user/main.go  
go run cmd/api-gateway/main.go 
```

Available endpoints <br>
/check_stock
```bash
curl -X POST http://localhost:8080/check_stock \
  -H "Content-Type: application/json" \
  -d '{"product_id": 1, "quantity": 1}'
```
/update_stock
```bash
curl -X POST http://localhost:8080/update_stock \
  -H "Content-Type: application/json" \
  -d '{"product_id": 101, "quantity": 2}'
```
/create_order
This triggers gRPC calls from OrderService to:
>* UserService.GetUser(user_id)
>* InventoryService.CheckStock(product_id)
>* InventoryService.UpdateStock(product_id)

```bash
curl -X POST http://localhost:8080/create_order \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {
        "product_id": 101,
        "quantity": 1,
        "unit_price": 500.0,
        "total_price": 500.0
      }
    ]
  }'
```





