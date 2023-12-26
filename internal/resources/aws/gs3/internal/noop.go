package internal

import "fmt"

type NoOpClient struct {
}

func (c *NoOpClient) WriteToObject(key string, data []byte) error {
	return fmt.Errorf("NoOpClient")
}
