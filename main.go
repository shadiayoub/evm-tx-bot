package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tx-bot/blockchain"
	"tx-bot/config"
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create blockchain client
	client, err := blockchain.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create blockchain client: %v", err)
	}

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Create a done channel to signal when the main loop is finished
	done := make(chan struct{})

	// Start main loop in a goroutine
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				log.Println("Main loop stopped")
				return
			default:
				// Send transaction
				if err := client.SendTransaction(ctx); err != nil {
					log.Printf("Failed to send transaction: %v", err)
					// Sleep for a longer time on error
					select {
					case <-ctx.Done():
						return
					case <-time.After(client.GetRandomLongSleepTime()):
					}
					continue
				}

				log.Println("Transaction sent successfully")
				// Sleep for random time between transactions
				select {
				case <-ctx.Done():
					return
				case <-time.After(client.GetRandomSleepTime()):
				}
			}
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	log.Println("Shutting down...")
	cancel()

	// Wait for main loop to finish with timeout
	select {
	case <-done:
		log.Println("Shutdown complete")
	case <-time.After(5 * time.Second):
		log.Println("Shutdown timed out")
	}
} 