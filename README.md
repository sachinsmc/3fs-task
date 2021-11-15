# 3fs-task
This repository contains task related to 3fs assignment for Go Developer

## How to run 

Prerequisite : Docker

Clone the repo

`$ git clone https://github.com/sachinsmc/3fs-task.git`

`$ docker compose up`

This will launch mysql database server, build Dockerfile and run app which will migrate and create DB.

- [Homepage](http://localhost:3003) :  http://localhost:3003  - This show realtime monitoring of the API Server using
Prometheus

- Postman collection link : [https://www.getpostman.com/collections/ef68d581ad490e901479](https://www.getpostman.com/collections/ef68d581ad490e901479)
- Documentation using postman : [https://documenter.getpostman.com/view/1550409/UVC8DRd1](https://documenter.getpostman.com/view/1550409/UVC8DRd1)
- Swagger : http://127.0.0.1:3003/swagger/index.html

- To run test cases `$ go run ./routes`

### What can be improved :

1. Validations
2. Add al test cases which covers every endpoint and scenarios
3. Add certificates for SSL
4. Write comments in code

I wanted to add above but because of lack of time, and I had other interviews along with this.



## Task :

Create a user and group management REST API service. Service should provide a way to list, add, modify and remove users and groups. Each user can belong to at most one group.

Create a REST API service with the following requirements:


- use latest Go release
- use latest go-swagger release
- use docker and Docker Compose
- write API specification in OpenAPI Specification version 2.0
- database of your choice


Data model:

    Groups:
        Name
    Users:
        Email
        Password
        Name




- Write all needed tests.


- A private git repository with full commit history is expected to be part of the delivered solution. 


- Other:
    - if needed, provide additional installation instructions, but there shouldn't be much more than running a simple command to set everything up
    - use best practices all around
