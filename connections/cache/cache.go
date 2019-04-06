package cache

type (
	// Cache interface
	Cache interface {
		Ping() error
		Get(key string) (string, error)
		Set(key string, value string, ttl ...int) error
		// ZIncrBy(key string, increment int, member string) error                    // Increment a member of sorter set by `increment`
		// Zrange(key string, start int, stop int, withscores bool) ([]string, error) // get soretd from range in inclusive manner from start to finish
	}
)
