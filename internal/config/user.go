package config

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	return write(*c)
}
