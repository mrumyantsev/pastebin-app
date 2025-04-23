# Pastebin

Pastebin is a web application where anyone can store any text online for easy sharing.

## Features

- The application can store text data up to 10 megabytes.
- The shorl URL address will be generated for paste after saving.
- When paste expires it will be removed automatically from the system.

![Pastebin](./pastebin-web-ui.png)
*Pastebin Web Interface*

## Implementation

The maximum size a paste can be is 512 kilobytes (0.5 megabytes). This is should be enough for almost any piece of text or script, and it prevents people from jamming the backend servers. Loggined users can edit or delete anything they pasted and also allowed to create pastes up to 10 megabytes. Expired pastes will be removed permanently.

![Architecture](./pastebin-arch-diagram.png)
*Pastebin Architecture Diagram*

## Prerequisites

- [Docker](https://docs.docker.com/desktop/linux/install/)

## Usage

Build the project with Docker Compose.

```
make
```

Run the project with Docker Compose.

```
make run
```

Build backend service locally (build the binary).

```
make build-local
```

Run backend service locally (run the binary).

```
make run-local
```

Run backend service in a temporary directory (fast).

```
make run-fast
```

## Testing

Run all the unit-tests.

```
make test
```

Build a PostgreSQL container with Docker.

```
make build-postgres-container
```

Build a MongoDB container with Docker.

```
make build-mongo-container
```

Run a PostgreSQL container with Docker.

```
make run-postgres-container
```

Run a MongoDB container with Docker.

```
make run-mongo-container
```
