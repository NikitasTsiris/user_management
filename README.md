# Dockerized gRPC User Management Application

This is a simple gRPC-based user management application consisting of a client and server. The client communicates with the server to register users to a database. The client provides the name and age of the user to the server, which adds the user to the database and responds with a success message.

## Prerequisites

Before running the application, make sure you have Docker installed on your system. You can download Docker from the official website: [https://www.docker.com/](https://www.docker.com/)

## Building the Docker Images

To build the client and server containers, follow these steps:

1. Clone the [repository](https://github.com/NikitasTsiris/user_management) and navigate to the root directory of the project.
2. Open a terminal or command prompt.

To build the server container, run the following command:

```bash
docker build . -t user_management_server --target server
```

To build the client container, run the following command:

```bash
docker build . -t user_management_client --target client
```


## Running the Application

To run the application, follow these steps:

1. Start the server container by running the following command:

```bash
docker run --rm user_management_server --port 50051:50051
```


2. Open a separate terminal or command prompt.

3. Instantiate the client container by running the following command:

```bash
docker run --rm --network host user_management_client
```

The client will communicate with the server, allowing you to register users to the database. Provide the name and age of the user to the client, and the server will add the user to the database and respond with a success message.

Note: The `--rm` flag is used to automatically remove the containers after they exit.

That's it! You now have a simple Dockerized gRPC user management application up and running.

Feel free to customize or modify the application based on your requirements. For more details on the implementation, refer to the source code and documentation in the project repository.
