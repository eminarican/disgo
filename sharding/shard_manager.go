package sharding

import (
	"context"

	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

// ShardManager manages multiple gateway.Gateway connections.
// For more information on sharding see: https://discord.com/developers/docs/topics/gateway#sharding
type ShardManager interface {
	// Logger returns the logger used by the ShardManager.
	Logger() log.Logger

	// Open opens all configured shards.
	Open(ctx context.Context)
	// Close closes all shards.
	Close(ctx context.Context)

	// OpenShard opens a specific shard.
	OpenShard(ctx context.Context, shardID int) error

	// CloseShard closes a specific shard.
	CloseShard(ctx context.Context, shardID int)

	// ShardByGuildID returns the gateway.Gateway for the shard that contains the given guild.
	ShardByGuildID(guildId snowflake.ID) gateway.Gateway

	// Shard returns the gateway.Gateway for the given shard ID.
	Shard(shardID int) gateway.Gateway

	// Shards returns a copy of all shards as a map.
	Shards() map[int]gateway.Gateway
}

// ShardIDByGuild returns the shard ID for the given guildID and shardCount.
func ShardIDByGuild(guildID snowflake.ID, shardCount int) int {
	return int((uint64(guildID) >> 22) % uint64(shardCount))
}
