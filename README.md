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
- `FROM_ADDRESS`: Your wallet address (sender)
- `TO_ADDRESS`: Recipient wallet address
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

### Direct Run

```bash
./tx-bot
```

The bot will start sending transactions with random amounts between the configured minimum and maximum values, with random sleep times between transactions.

To stop the bot, press Ctrl+C for graceful shutdown.

### Running in Screen Session (Recommended)

For long-term running and easy monitoring, it's recommended to run the bot in a screen session:

1. Create a new named screen session:
```bash
screen -S tx-bot
```

2. Run the bot inside the screen session:
```bash
./tx-bot
```

3. Detach from the screen session (keep the bot running):
- Press `Ctrl+A` followed by `D`

4. Reattach to the screen session later:
```bash
screen -r tx-bot
```

5. List all screen sessions:
```bash
screen -ls
```

6. Kill the screen session (if needed):
```bash
screen -X -S tx-bot quit
```

Using screen allows you to:
- Keep the bot running even if your SSH connection drops
- Monitor the bot's logs and status
- Easily restart or stop the bot when needed
- Run multiple instances with different names

## Security Notes

- Never commit your `.env` file or expose your private keys
- Keep your private keys secure and never share them
- Consider using a dedicated wallet for the bot with limited funds 