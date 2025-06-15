# Transaction Bot

A Go-based bot that sends periodic transactions on a custom EVM blockchain.

## Features

- Configurable transaction amounts
- Random sleep times between transactions
- Graceful shutdown handling
- Environment-based configuration
- Error handling with exponential backoff

## Prerequisites

- Go 1.21 or later
- Access to an EVM-compatible blockchain node
- Private key with sufficient funds

## Configuration

Copy the `env.example` file to `.env` and fill in your configuration:

```bash
cp env.example .env
```

Edit the `.env` file with your specific values:

- `WEB3_PROVIDER_URL`: Your blockchain node URL
- `FROM_ADDRESS`: Your wallet address
- `PRIVATE_KEY`: Your wallet's private key
- `CHAIN_ID`: Network chain ID
- `GAS_LIMIT`: Maximum gas per transaction
- `GAS_PRICE_GWEI`: Gas price in Gwei
- `MIN_TRANSFER_AMOUNT`: Minimum transfer amount
- `MAX_TRANSFER_AMOUNT`: Maximum transfer amount
- `MIN_SLEEP_TIME`: Minimum time between transactions (seconds)
- `MAX_SLEEP_TIME`: Maximum time between transactions (seconds)
- `MIN_LONG_SLEEP_TIME`: Minimum time to sleep after errors (seconds)
- `MAX_LONG_SLEEP_TIME`: Maximum time to sleep after errors (seconds)

## Building

```bash
go mod download
go mod tidy
go build
```

## Running

```bash
./tx-bot
```

The bot will start sending transactions with random amounts between the configured minimum and maximum values, with random sleep times between transactions.

To stop the bot, press Ctrl+C for graceful shutdown.

## Security Notes

- Never commit your `.env` file or expose your private keys
- Keep your private keys secure and never share them
- Consider using a dedicated wallet for the bot with limited funds 