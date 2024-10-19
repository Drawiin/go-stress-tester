# Go Stress Tester

## Objective

Create a CLI system in Go to perform load tests on a web service. The user must provide the service URL, the total number of requests, and the number of simultaneous calls.

The system should generate a report with specific information after the tests are executed.

## CLI Parameter Input

- **`-url`**: URL of the service to be tested.
- **`-requests`**: Total number of requests.
- **`-concurrency`**: Number of simultaneous calls.

## Test Execution

- Perform HTTP requests to the specified URL.
- Distribute the requests according to the defined concurrency level.
- Ensure that the total number of requests is fulfilled.

## Report Generation

- Present a report at the end of the tests containing:
    - Total time spent on execution.
    - Total number of requests made.
    - Number of requests with HTTP status 200.
    - Distribution of other HTTP status codes (such as 404, 500, etc.).

## How to Run

### Prerequisites

- Docker installed on your machine.

### Building the Docker Image

To build the Docker image, run the following command:

```sh
docker build -f Dockerfile -t <your-docker-image-name>:latest .
```

### Running the Application

You can use this application by making a call via Docker. For example:

```sh
docker run -it --rm <your-docker-image-name>:latest --url=http://google.com --requests=1000 --concurrency=10
```

### Running from Docker Hub

To ensure you are running the image from Docker Hub and not from a local cache, use the `--pull always` flag:

```sh
docker run -it --rm --pull always <your-docker-image-name>:latest --url=http://google.com --requests=1000 --concurrency=10
```

### Makefile Commands

You can also use the provided `Makefile` for common tasks:

- **Build Docker Image:**

  ```sh
  make build-img
  ```

- **Run Docker Image:**

  ```sh
  make run-img
  ```

- **Run Docker Image from Docker Hub:**

  ```sh
  make run-from-hub
  ```
