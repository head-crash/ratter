# RATTER - Random REST Test Responser

This is a simple cli tool written in Go that starts a http server to handle requests with a specified path prefix and http method. It randomly returns either a success or an error response. It is configurable via command-line flags.

## Features

- Handles HTTP requests with a specified method and URL prefix.
- Randomly returns a success or error response.
- Configurable error message and server port.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/head-crash/ratter.git
   cd ratter
   ```

2. Build the application:

   ```bash
   go build -o random-error-server main.go
   ```

## Usage

Run the server with the following command-line flags:

- `-valid-prefix`: The valid URL prefix for handling requests.

Optional:

- `-port`: The port to run the server on (default: "8080").
- `-method`: The HTTP method to handle (default: "POST").
- `-check-path-id`: The error message to return if the path ID is missing after the valid prefix. (default: nil)
- `-error-message`: The error message to return on failure (default: "Random error response").

Example:

```bash
./random-error-server -valid-prefix="/api" -error-message="An error occurred" -port="8080" -method="POST" 
```

## How It Works

1. The server listens on the specified port and handles requests with the specified URL prefix.
2. If the request method does not match the specified method, it returns a `405 Method Not Allowed` error.
3. If the request path does not contain a request ID after the valid prefix, it returns a `400 Bad Request` error.
4. The server randomly decides whether to return a success or error response. If an error is chosen, it returns the specified error message with a `400 Bad Request` status.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
