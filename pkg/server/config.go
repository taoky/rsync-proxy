package server

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"

	"github.com/ustclug/rsync-proxy/pkg/log"
)

type Upstream struct {
	Host    string   `toml:"host"`
	Port    int      `toml:"port"`
	Modules []string `toml:"modules"`
}

type Config struct {
	Upstreams map[string]*Upstream `toml:"upstreams"`
}

func (s *Server) LoadConfig(r io.Reader) error {
	log.V(3).Infof("[INFO] loading config")

	dec := toml.NewDecoder(r)
	var c Config
	err := dec.Decode(&c)
	if err != nil {
		return err
	}

	s.Upstreams = c.Upstreams
	return s.complete()
}

func (s *Server) LoadConfigFromFile() error {
	f, err := os.Open(s.ConfigPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.LoadConfig(f)
}
