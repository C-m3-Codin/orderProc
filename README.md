# Order Processing Service

## Overview
This service handles order requests asynchronously using two internal queues and processes them with different sets of worker threads. It saves order data to a PostgreSQL database and provides metrics on order processing.

## API Endpoints

### Create an Order
Creates a new order and enqueues it for processing.

```bash
curl --location 'http://localhost:8080/order' --header 'Content-Type: application/json' --data '{
           "user_id": "user123",
           "item_ids": "item1,item2,item3",
           "total_amount": 99.99
         }'
```

**Response:**
```
Received order ab079da1-e078-4d40-b36d-739f06a242fe
```

---

### Get Order Details
Fetches details of an existing order by its ID.

```bash
curl --location 'http://localhost:8080/order/ab079da1-e078-4d40-b36d-739f06a242fe'
```

**Response:**
```json
{
  "order_id": "ab079da1-e078-4d40-b36d-739f06a242fe",
  "user_id": "user123",
  "item_ids": ["item1", "item2", "item3"],
  "total_amount": 99.99,
  "status": "processed"
}
```

---

### Metrics
Provides metrics on order processing.

- **Pending Orders:** Total number of orders currently pending processing.
- **Total Orders:** Total number of orders received.
- **Average Processing Time:** Average time taken to process an order.

```bash
curl --location 'http://localhost:8080/metrics/pending'
```

**Response:**
```json
{
  "PendingCount": 0,
  "Proccessed": 0,
  "Completed": 15831,
  "TotalCount": 15831,
  "AverageProcessingTime": 0
}
```

## Architecture
1. **Order Reception:** An order request is received and a unique order ID is generated.
2. **Queueing (Pending Status):** The order is added to the first internal queue and saved in a concurrent-safe hashmap with a `pending` status.
3. **Database Storage (Processing Status):** Workers from the first queue pick orders and save them into the PostgreSQL database. Once saved, the order status changes to `processing` and the order is pushed to a second internal queue.
4. **Processing Simulation (Completed Status):** Workers from the second queue pick orders and simulate order processing by waiting a random time between 1 to 10 seconds. Once the processing is complete, the order status changes to `completed`.

## Design Tradeoffs
- **Asynchronous Processing:**
  - Pros: Higher throughput, better scalability under load.
  - Cons: Increased complexity, potential for higher latency on individual requests.
- **Two-Queue Model:**
  - Pros: Clear separation of concerns between order creation and processing.
  - Cons: Requires more coordination between different worker pools.
- **Concurrent-Safe Hashmap:**
  - Pros: Fast in-memory access and safe in a concurrent environment.
  - Cons: Limited by memory size, potential for stale data if not synchronized properly.
- **Synchronous Worker Model:**
  - Pros: Simple and predictable processing flow.
  - Cons: Limited by worker capacity and potential bottlenecks.

## Benchmark Results
Benchmarking performed using ApacheBench:

```
Server Hostname:        localhost
Server Port:            8080

Document Path:          /order
Document Length:        54 bytes

Concurrency Level:      2000
Time taken for tests:   9.078 seconds
Complete requests:      8000
Failed requests:        0
Total transferred:      1416000 bytes
Total body sent:        1728000
Requests per second:    881.28 [#/sec] (mean)
Time per request:       2269.423 [ms] (mean)
Time per request:       1.135 [ms] (mean, across all concurrent requests)
Transfer rate:          338.23 [Kbytes/sec] total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       5
Processing:   748 2058 541.3   2184    2887
Waiting:        1 1120 652.1   1015    2673
Total:        748 2059 541.3   2184    2887

Percentage of the requests served within a certain time (ms)
  50%   2184
  66%   2279
  75%   2401
  80%   2432
  90%   2769
  95%   2806
  98%   2827
  99%   2840
 100%   2887 (longest request)
```

## Requirements
- Golang
- PostgreSQL (must be running)
- Docker (optional)

## How to Run
```bash
go run main.go
```

Ensure PostgreSQL is running and accessible. Configure database connection settings in your environment variables or configuration file.
