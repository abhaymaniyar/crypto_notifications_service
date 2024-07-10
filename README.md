### Requirements

We would like you to create a crypto notification app. The features it should include:

Create a server which will be able to take in the following rest APIs

Function Requirements:

- extensible to support other coins
- extensible to support other channels
- list sent notifications based on status (paginated)
- persist those notifications

Assumptions:

- email functionality is already in place
- auth layer in place

### Discussion Notes

- Create a notification. Line items may include current price of BTC, (nudging a user to invest > 10ms)
  market trade volume, intra day high price, market cap
- Send a notification to an email
- Wishlisted (BTC)
- List sent notifications (sent, outstanding, failed etc.)
- Delete a notification


- Trigger the notification sync (order, wishlisted, invested)
- Trigger the notification async (high_chance)

### Prerequisites

What things you need to install the software and how to install them:

- Go (version 1.18 or later)
- PostgreSQL (or any preferred database, adjust accordingly)

## Deployment

A step by step series of examples that tell you how to get a development environment running:

1. Clone the repository:

```sh
git clone https://github.com/abhaymaniyar/crypto_notifications_service.git
```

2. Create a dev environment file from the sample file

 ```sh
 $cp env.sample development.env
 ```

3. Modify the values in `development.env` to match your needs.

```sh
./run-server.sh [env-file]
```

This command will load environment variables of the specified env file and run the server on the port specified in the
env variables file.

### Detailed API Spec with sample responses

### Postman Collection

### Built With

- httprouter - The web framework used
- GORM - ORM library for Go
- Goose - DB migration library
- Viper - Configuration management

### Key elements

- Logging and error alerts
- Pagination support for get requests
- Standardised request response formats
- Thorough API contract doc with sample responses
- DB migrations with goose
- Config management with Viper
- Possible middleware chaining

### Improvements (things omitted in interest of time)

- add unit test cases (coverage > 85%)
- add concurrency for better scale handling
- add more custom error types instead of returning generic 500 in case of any DB error for instance.