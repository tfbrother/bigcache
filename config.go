package bigcache

import "time"

// Config for BigCache
type Config struct {
	// Number of cache shards(分片，分区), value must be a power of two
	Shards int
	// Time after which entry can be evicted(驱逐)
	LifeWindow time.Duration
	// Max number of entries in life window. Used only to calculate(计算) initial size for cache shards.
	// When proper value is set then additional memory allocation does not occur.
	MaxEntriesInWindow int
	// Max size of entry in bytes. Used only to calculate initial size for cache shards.
	// 分区里面，缓存的大小的最大字节数。
	MaxEntrySize int
	// Verbose mode prints information about new memory allocation
	Verbose bool
	// Hasher used to map between string keys and unsigned 64bit integers, by default fnv64 hashing is used.
	Hasher Hasher
	// HardMaxCacheSize is a limit for cache size in MB. Cache will not allocate more memory than this limit.
	// It can protect application from consuming all available memory on machine, therefore from running OOM Killer.
	// Default value is 0 which means unlimited size. When the limit is higher than 0 and reached then
	// the oldest entries are overridden for the new ones.
	HardMaxCacheSize int
}

// DefaultConfig initializes config with default values.
// When load for BigCache can be predicted in advance then it is better to use custom config.
func DefaultConfig(eviction time.Duration) Config {
	return Config{
		Shards:             2,
		LifeWindow:         eviction,
		MaxEntriesInWindow: 2 * 10,
		MaxEntrySize:       2,
		Verbose:            true,
		Hasher:             newDefaultHasher(),
		HardMaxCacheSize:   0,
	}
}
