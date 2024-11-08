# Credit Card Validation Webserver
A lightweight web server for validating credit card numbers, written in Go.

## Table of Contents

- [Requirements](#requirements)
- [Building the Webserver](#building-the-webserver)
- [Running the Server](#running-the-server)
- [Endpoints](#endpoints)
  - [Root Endpoint (`/`)](#root-endpoint-)
- [Customisation](#customisation)
  - [Changing the Port](#changing-the-port)
  - [Changing the Path](#changing-the-path)

## Requirements

- Go 1.18 or later.
- A terminal or command prompt to run the application.

## Building the Webserver

1. Clone the repository to your local machine:
   ```bash
   $ git clone https://github.com/ibnaleem/cc-validation-webserver.git
   $ cd cc-validation-webserver 
   ```

2. Initialise Go modules if you haven't done so:
   ```bash
   $ go mod tidy
   ```

3. Build the webserver:
   ```bash
   $ go build
   ```

4. This will create the binary `./webserver`.

## Running the Server

```bash
$ ./webserver
```

By default, the server will start on port `3333`:

```
:: Webserver started on port 3333 ::
```

You can then access the server at `http://127.0.0.1:3333/`.

## Endpoints

### Root Endpoint (`/`)

- **Method**: `POST`
- **Content-Type**: `application/json`
- **Request Payload**:
  
  The server expects a JSON payload containing a `credit-card` field. The value of this field should be a string representing the credit card number you want to validate.

  Example request:
  ```json
  {
      "credit-card": "378282246310005"
  }
  ```

- **Response**:
  
  The server will respond with a string value (`true` or `false`) depending on whether the credit card number is valid according to the [Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm).

  - `true`: If the credit card number is valid.
  - `false`: If the credit card number is invalid.

  Example response:
  ```json
  true
  ```

### Error Responses:

- **Invalid JSON**: If the provided JSON is malformed.
  - Status code: `400 Bad Request`
  - Example response:
    ```json
    {"error": "Invalid JSON"}
    ```

- **Missing Credit Card Field**: If the JSON does not contain the `credit-card` field.
  - Status code: `400 Bad Request`
  - Example response:
    ```json
    {"error": "Missing expected field in JSON"}
    ```

- **Invalid Credit Card Number**: If the `credit-card` field contains a number that cannot be converted to an integer or fails the Luhn check.
  - Status code: `400 Bad Request`
  - Example response:
    ```json
    {"error": "Invalid credit card number"}
    ```

## Customisation

### Changing the Port

By default, the server listens on port `3333`. You can change this by setting a different value for the `PORT` variable in the [`webserver.go`](https://github.com/ibnaleem/cc-validation-webserver/blob/main/webserver.go) file:

```go
var PORT string = "3333"
```

Change `"3333"` to any valid port number you prefer (e.g., `"8080"`). Note that `PORT` expect a type of `string`, meaning `var PORT string = 8080` is clearly **invalid.** Simply encapsulate the port in quotes `""`. 

### Changing the Path

To customise the root endpoint (`/`) to a different one, modify the `http.HandleFunc` line in the [`webserver.go`](https://github.com/ibnaleem/cc-validation-webserver/blob/main/webserver.go) file:

```go
http.HandleFunc("/", getRoot)
```

For example, to change the path to `/validate`, update it as follows:

```go
http.HandleFunc("/validate", getRoot)
```

Now, the server will expect requests at `http://127.0.0.1:3333/validate`.

To ensure consistency and avoid confusion, it's best to also update the function name to match the new path. So, if you change the path to `/validate` for example, rename the function to reflect this change:

```go
func getValidate(w http.ResponseWriter, r *http.Request) {...}
```

Additionally, update the `fmt.Printf()` to reflect the new path:

```go
if r.Header.Get("Content-Type") == "" {
    fmt.Printf("[%s] on validate (/validate) with cURL\n", r.Method)
} else {
    fmt.Printf("[%s] on validate (/validate) with header %s\n", r.Method, r.Header.Get("Content-Type"))
}
```
## License

This project is licensed under the GNU General Public License - see the [LICENSE](https://github.com/ibnaleem/cc-validation-webserver/blob/main/LICENSE) file for details.
