package node

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)



func CheckNodeSync(client *ethclient.Client) bool {
    // Check if node is fully synced
    progress, err := client.SyncProgress(context.Background())
    if err != nil {
        log.Error().Err(err).Msg("Failed to get sync progress")
        return false
    }

    if progress != nil {
        log.Warn().
            Uint64("current", progress.CurrentBlock).
            Uint64("highest", progress.HighestBlock).
            Msg("Node not fully synced")
        return false
    }

    return true
}