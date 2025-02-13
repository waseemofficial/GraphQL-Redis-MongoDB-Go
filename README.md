<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fhits.dwyl.com%2Fwaseemofficial%2F{env.}.json%3Fcolor%3Dpink"/>



<p align="center" >
<div align="center" >
<img src="https://github.com/waseemofficial/DSA_Python/blob/main/Images/github_logo_blue.png"/>
</div>

<div align="center">
<a href="https://github.com/waseemofficial">
<img src="https://img.shields.io/badge/syed-waseem-93b023?&style=plastic&logo=&logoColor=white"/></a>
<img src="https://img.shields.io/badge/gitlab-%23181717.svg?style=plastic&logo=gitlab&logoColor=white"/>
<img src="https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=plastic&logo=visual-studio-code&logoColor=white"/>
<img src="https://img.shields.io/badge/markdown-%23000000.svg?style=plastic&logo=markdown&logoColor=white"/>
</div></p>


<div align="center">
<a href="https://github.com/waseemofficial">
<img src="https://img.shields.io/github/followers/waseemofficial?label=Follow%20Me&style=social"/>
</a>
<br>

<img src="https://img.shields.io/github/license/waseemofficial/GraphQL-Redis-MongoDB-Go.svg?style=flat"/> <img src="https://img.shields.io/github/languages/top/waseemofficial/GraphQL-Redis-MongoDB-Go?style=flat"/> <img src="https://img.shields.io/github/stars/waseemofficial/GraphQL-Redis-MongoDB-Go.svg?colorB=orange&style=flat"/> <img sec="https://img.shields.io/github/languages/top/waseemofficial/GraphQL-Redis-MongoDB-Go.svg?style=flat"/> <img src="https://img.shields.io/github/languages/code-size/waseemofficial/GraphQL-Redis-MongoDB-Go.svg?style=flat"/> <img src="https://img.shields.io/github/issues-raw/waseemofficial/GraphQL-Redis-MongoDB-Go.svg?style=flat" />
</div>

<div align="center"> 

---

### Languages

![Python](https://img.shields.io/badge/-Python-000?&logo=Python)
![JavaScript](https://img.shields.io/badge/-JavaScript-000?&logo=JavaScript)
![Golang](https://img.shields.io/badge/-Golang-000?&logo=Go)
![Java](https://img.shields.io/badge/-Java-000?&logo=jdk)
![Solidity](https://img.shields.io/badge/-Solidity-000?&logo=Solidity)
![SQL](https://img.shields.io/badge/-SQL-000?&logo=MySQL)
![Bash](https://img.shields.io/badge/-Bash-000?&logo=gnu-bash&logoColor=white)
![Bash](https://img.shields.io/badge/-markdown-000?&logo=markdown)



### Technologies

![Docker](https://img.shields.io/badge/-Docker-000?&logo=Docker)
![Linux](https://img.shields.io/badge/-Linux-000?&logo=Linux)
![Node.js](https://img.shields.io/badge/-Node.js-000?&logo=node.js)
![React](https://img.shields.io/badge/-React-000?&logo=React)
![Redis](https://img.shields.io/badge/-Redis-000?&logo=Redis)
![Cypress](https://img.shields.io/badge/-Postman-000?&logo=Postman)
![Cypress](https://img.shields.io/badge/-Cypress-000?&logo=Cypress)
![GitHub](https://img.shields.io/badge/-GitHub-000?&logo=GitHub)
![GitHub](https://img.shields.io/badge/-Selenium-000?&logo=Selenium)
![GitHub](https://img.shields.io/badge/-Regex-000?&logo=Regex)
![GithubActions](https://img.shields.io/badge/-GithubActions-000?&logo=GithubActions)
</div>

---

# GraphQL, Redis, and MongoDB with Go

Welcome to the **GraphQL-Redis-MongoDB-Go** repository! This project demonstrates how to build a **high-performance, scalable backend system** using **GraphQL** for API querying, **Redis** for caching, and **MongoDB** as the primary database, all implemented in **Go (Golang)**. This repository serves as a practical example of integrating these technologies to create efficient and modern web applications.

---

## 🚀 Key Features
- **GraphQL API**: Flexible and efficient querying with GraphQL.
- **Redis Caching**: Improve performance by caching frequently accessed data.
- **MongoDB Database**: Scalable NoSQL database for storing and retrieving data.
- **Go Backend**: High-performance backend built with Go.
- **Modular Design**: Clean and organized codebase for easy understanding and extension.

---

## 🛠️ Technologies and Tools Used
- **Programming Language**: Go (Golang)
- **GraphQL Library**: [gqlgen](https://gqlgen.com/)
- **Database**: MongoDB
- **Caching**: Redis
- **API Testing**: GraphQL Playground
- **Dependency Management**: Go Modules

---

## 📂 Repository Structure
Here’s an overview of the repository structure:

```
GraphQL-Redis-MongoDB-Go/
├── graph/                  # GraphQL schema and resolvers
├── models/                 # Data models and MongoDB interactions
├── redis/                  # Redis caching implementation
├── main.go                 # Entry point of the application
├── go.mod                  # Go module dependencies
├── README.md               # This file
└── .env.example            # Environment variables template
```

---

## 🧑‍💻 Getting Started

### Prerequisites
- Go 1.16 or higher
- MongoDB (local or cloud instance)
- Redis (local or cloud instance)
- GraphQL Playground (for testing)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/waseemofficial/GraphQL-Redis-MongoDB-Go.git
   cd GraphQL-Redis-MongoDB-Go
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up environment variables:
   - Rename `.env.example` to `.env`.
   - Update the `.env` file with your MongoDB and Redis connection details:
     ```
     MONGODB_URI=mongodb://localhost:27017
     REDIS_ADDR=localhost:6379
     ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. Open GraphQL Playground at `http://localhost:8080` to interact with the API.

---

## 🧑‍💻 Example Code and Insights

### GraphQL Schema
The GraphQL schema defines the structure of the API. Here’s an example schema for a `User` type:

```graphql
type User {
    id: ID!
    name: String!
    email: String!
    age: Int!
}

type Query {
    users: [User!]!
    user(id: ID!): User
}

type Mutation {
    createUser(name: String!, email: String!, age: Int!): User!
}
```

### Resolver Implementation
Resolvers handle the logic for fetching and manipulating data. Here’s an example resolver in Go:

```go
package graph

import (
    "context"
    "GraphQL-Redis-MongoDB-Go/models"
    "GraphQL-Redis-MongoDB-Go/redis"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
    DB    *mongo.Database
    Cache *redis.Client
}

func (r *Resolver) Users(ctx context.Context) ([]*models.User, error) {
    // Check Redis cache first
    cachedUsers, err := r.Cache.Get("users")
    if err == nil {
        return cachedUsers, nil
    }

    // Fetch from MongoDB if not in cache
    var users []*models.User
    cursor, err := r.DB.Collection("users").Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, &user)
    }

    // Cache the result in Redis
    r.Cache.Set("users", users)

    return users, nil
}
```

### Redis Caching
Redis is used to cache frequently accessed data, reducing database load and improving performance. Here’s how caching is implemented:

```go
package redis

import (
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "context"
)

type Client struct {
    client *redis.Client
}

func NewClient(addr string) *Client {
    return &Client{
        client: redis.NewClient(&redis.Options{
            Addr: addr,
        }),
    }
}

func (c *Client) Get(key string) (interface{}, error) {
    val, err := c.client.Get(context.Background(), key).Result()
    if err != nil {
        return nil, err
    }

    var result interface{}
    json.Unmarshal([]byte(val), &result)
    return result, nil
}

func (c *Client) Set(key string, value interface{}) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return c.client.Set(context.Background(), key, data, 0).Err()
}
```

---

## 📜 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments
- [gqlgen](https://gqlgen.com/) for the GraphQL library.
- The Go and open-source communities for their invaluable tools and libraries.

---

## 📞 Contact
For questions or feedback, feel free to reach out:
- **GitHub**: [waseemofficial](https://github.com/waseemofficial)
- **Email**: [Your Email]
