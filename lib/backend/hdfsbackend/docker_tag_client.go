package hdfsbackend

import (
	"fmt"
	"io"

	"code.uber.internal/infra/kraken/lib/backend/nameparse"
)

// DockerTagClient is an HDFS client for uploading / downloading tags to a docker
// registry.
type DockerTagClient struct {
	config Config
	client *client
}

// NewDockerTagClient creates a new DockerTagClient.
func NewDockerTagClient(config Config) (*DockerTagClient, error) {
	config, err := config.applyDefaults()
	if err != nil {
		return nil, fmt.Errorf("invalid config: %s", err)
	}
	return &DockerTagClient{client: newClient(config), config: config}, nil
}

func (c *DockerTagClient) path(name string) (string, error) {
	return nameparse.RepoTagPath(c.config.RootDirectory, name)
}

// Download downloads the value of name into dst. name should be in the
// format "repo:tag".
func (c *DockerTagClient) Download(name string, dst io.Writer) error {
	path, err := c.path(name)
	if err != nil {
		return err
	}
	return c.client.download(path, dst)
}

// Upload uploads src as the value of name. name should be in the format "repo:tag".
func (c *DockerTagClient) Upload(name string, src io.Reader) error {
	path, err := c.path(name)
	if err != nil {
		return err
	}
	return c.client.upload(path, src)
}