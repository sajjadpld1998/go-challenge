# **User Segmentation and Estimation Service**

## **Overview**
This project consists of two microservices:
1. **User Segmentation Service (USS):** Receives `(user_id, segment)` data, processes it, and forwards it to the Estimation Service.
2. **Estimation Service (ES):** Stores the `(user_id, segment)` data in a Redis database with a TTL of two weeks and provides a REST API to estimate the number of users in a specific segment.

### **Features**
- **User Segmentation Service:**
    - Accepts `user_id` and `segment` pairs via a REST API.
    - Forwards the data to the Estimation Service.

- **Estimation Service:**
    - Stores `(user_id, segment)` pairs in Redis with a two-week expiration (TTL).
    - Provides an API to count the number of users in a specific segment.
    - Ensures high scalability using Redis for efficient storage and querying.

---

## **Architecture**
### **Service 1: User Segmentation Service (USS)**
- **Responsibilities:**
    - Expose a REST API endpoint to accept JSON payloads with `user_id` and `segment`.
    - Forward the data to the Estimation Service via a REST API.

- **API Endpoints:**
    1. `POST /segment`
        - **Request Body:**
          ```json
          {
            "segment": "sport",
            "user_id": "u104010"
          }
          ```
        - **Description:** Sends `(user_id, segment)` data to the Estimation Service.



### **Service 2: Estimation Service (ES)**
- **Responsibilities:**
    - Store `(user_id, segment)` data in Redis with a TTL of 2 weeks.
    - Expose a REST API to estimate the number of users in a specific segment.

- **API Endpoints:**
    1. `POST /segment`
        - **Request Body:**
          ```json
          {
            "segment": "sport",
            "user_id": "u104010"
          }
          ```
        - **Description:** Receives `(user_id, segment)` data from the User Segmentation Service and stores it in Redis.

    2. `GET /segment/:segment`
        - **Path Parameter:** `segment` (e.g., `sport`).
        - **Response:**
          ```json
          {
            "segment": "sport",
            "user_count": 12345
          }
          ```
        - **Description:** Returns the count of users in the given segment.

---

## **Technology Stack**
- **Programming Language:** Go
- **Datastore:** Redis (for fast and efficient storage with TTL support)
- **Frameworks:** Gin (for REST API implementation)
- **Containerization:** Docker (for running services independently)

---

## **Run Services**
You can run the services either separately with `go run` or using `Docker`. Before running the services, make sure to configure them properly.

### **Configuration**
Each service includes a configuration file template named `app/configs.yaml.example`.
To set up your configurations:
1. Copy the `app/configs.yaml.example` file.
2. Rename the copy to `app/configs.yaml`.
3. Edit the `configs.yaml` file with your specific configurations, such as Redis host, port, or other service settings.

### **Swagger**
Each service has swagger page for testing.

**USS**
```
http://localhost:8282/swagger/index.html
```

**ES**
```
http://localhost:8484/swagger/index.html
```

---


## **Scalability Considerations**
1. Redis for Storage:

   - Redis is used for storing (user_id, segment) pairs with a TTL of 2 weeks. It is highly efficient for read/write operations at scale.

2. Decoupled Architecture:

   - USS and ES are decoupled, allowing each service to scale independently.
   
3. Horizontal Scaling:

   - Both services can be replicated to handle higher traffic, with load balancing across instances.
   
4. Asynchronous Processing:

   - If future requirements demand higher reliability, you can introduce a message broker (e.g., RabbitMQ or Kafka) for asynchronous communication between USS and ES.


---


## **Future Improvements**
1. **Message Broker Integration:**

   - Replace direct REST calls between USS and ES with RabbitMQ or Kafka for better reliability and decoupling.

2. **Authentication:**

   - Add API authentication (e.g., JWT or API keys) to restrict access.
   
3. **Metrics and Monitoring:**

   - Integrate Prometheus and Grafana for monitoring system performance and Redis usage.

4. **Job Queue:**
   - Queue send data pairs in USS and send them with a task queue or job queue. we can use this in ES for getting data from USS
   
5. **gRPC Integration:**
   - Replace the REST communication between USS and ES with gRPC for improved performance and lower latency in high-throughput environments.


---


## **Why REST API and Redis?**
### **Why REST API for Receiving Data?**
1. **Simplicity and Universality:**

   - REST is widely supported and easily consumable by various clients (e.g., frontend apps, other microservices, third-party tools).

   - It uses HTTP, which is a well-understood protocol, making integration straightforward.
   
2. **Ease of Development and Debugging:**

   - REST APIs are easy to implement with frameworks like Gin (Go), and debugging tools like Postman or curl make testing seamless.
   
   - Log and monitor incoming requests effortlessly.
   
3. **Flexibility:**

   - REST APIs allow you to extend functionality in the future without overhauling the architecture. For example, adding new endpoints or parameters is straightforward.
   
4. **Stateless Communication:**

   - REST is inherently stateless, which simplifies scalability. Multiple instances of the Estimation Service can handle incoming requests without worrying about session consistency.
   
5. **Interoperability:**

   - REST works with any language or platform, ensuring future systems can integrate with the service easily.


### **Why Not gRPC or Other Protocols?**
- **gRPC**: While gRPC provides better performance, it adds complexity, such as managing Protobuf schemas, requiring HTTP/2, and debugging challenges.

- **SOAP**: REST is less verbose and easier to implement compared to SOAP.

- **GraphQL**: GraphQL is more suited for querying complex, nested data rather than simple task-based APIs like this.


### **Why Redis for the Database?**
1. **High Performance:**

   - Redis is an in-memory database optimized for extremely fast read and write operations. This is crucial for handling millions of (user_id, segment) pairs and ensuring low-latency queries.
   
2. **Built-In TTL Support:**

   - Redis has native support for Time-to-Live (TTL), making it ideal for storing temporary data like (user_id, segment) pairs that expire after two weeks. TTL is applied automatically without requiring complex query logic.

3. **Scalability:**

   - Redis supports horizontal scaling through clustering, allowing the system to handle millions of keys and higher traffic as the system grows.
   
4. **Atomic Operations:**

   - Redis provides atomic commands, ensuring data consistency even with concurrent writes and reads.
   
5. **Ease of Integration:**

   - Redis is easy to integrate with Go using libraries like go-redis, which provide high-level APIs for managing connections and performing operations.
   
6. **Minimal Overhead:**

   - Redis does not require heavy indexing or schema design, making it lightweight and efficient for key-value-based data storage.


### **Why Not a Relational Database?**
- **Performance:** A relational database (e.g., PostgreSQL) would require indexing and more complex queries to support the expiration logic, which is less efficient for this use case.

- **Scalability:** Relational databases typically require more resources to handle millions of rows with frequent writes and TTL enforcement.