package config

type config struct {
	URLs  []string
	Limit int
}

type Configuration interface {
	GetURLs() []string
	GetParallelRequestLimit() int
}

func (c *config) GetURLs() []string {
	return c.URLs
}

func (c *config) GetParallelRequestLimit() int {
	return c.Limit
}

func Setup(urls []string, limit int) Configuration {

	return &config{
		URLs:  urls,
		Limit: limit,
	}

}
