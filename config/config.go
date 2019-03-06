package config

type config struct {
	Parallel bool
	URLs     []string
	Limit    int
}

type Configuration interface {
	ParallelRequestsEnabled() bool
	GetURLs() []string
	GetParallelRequestLimit() int
}

func (c *config) ParallelRequestsEnabled() bool {
	return c.Parallel
}

func (c *config) GetURLs() []string {
	return c.URLs
}

func (c *config) GetParallelRequestLimit() int {
	return c.Limit
}
