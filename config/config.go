package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Web3ProviderURL     string
	FromAddress         string
	ToAddress           string
	PrivateKey          string
	ChainID             int64
	GasLimit            uint64
	GasPriceGwei        int64
	MinTransferAmount   float64
	MaxTransferAmount   float64
	MinSleepTime        int
	MaxSleepTime        int
	MinLongSleepTime    int
	MaxLongSleepTime    int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	chainID, _ := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	gasLimit, _ := strconv.ParseUint(os.Getenv("GAS_LIMIT"), 10, 64)
	gasPriceGwei, _ := strconv.ParseInt(os.Getenv("GAS_PRICE_GWEI"), 10, 64)
	minTransferAmount, _ := strconv.ParseFloat(os.Getenv("MIN_TRANSFER_AMOUNT"), 64)
	maxTransferAmount, _ := strconv.ParseFloat(os.Getenv("MAX_TRANSFER_AMOUNT"), 64)
	minSleepTime, _ := strconv.Atoi(os.Getenv("MIN_SLEEP_TIME"))
	maxSleepTime, _ := strconv.Atoi(os.Getenv("MAX_SLEEP_TIME"))
	minLongSleepTime, _ := strconv.Atoi(os.Getenv("MIN_LONG_SLEEP_TIME"))
	maxLongSleepTime, _ := strconv.Atoi(os.Getenv("MAX_LONG_SLEEP_TIME"))

	return &Config{
		Web3ProviderURL:     os.Getenv("WEB3_PROVIDER_URL"),
		FromAddress:         os.Getenv("FROM_ADDRESS"),
		ToAddress:           os.Getenv("TO_ADDRESS"),
		PrivateKey:          os.Getenv("PRIVATE_KEY"),
		ChainID:             chainID,
		GasLimit:            gasLimit,
		GasPriceGwei:        gasPriceGwei,
		MinTransferAmount:   minTransferAmount,
		MaxTransferAmount:   maxTransferAmount,
		MinSleepTime:        minSleepTime,
		MaxSleepTime:        maxSleepTime,
		MinLongSleepTime:    minLongSleepTime,
		MaxLongSleepTime:    maxLongSleepTime,
	}, nil
} 