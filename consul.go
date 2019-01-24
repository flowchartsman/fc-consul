package consul

import (
	"fmt"
	"strings"

	"github.com/flowchartsman/fc"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

// Source represents a config source pulling from a consul node
type Source struct {
	addr   string
	prefix string
	m      map[string][]string
}

// WithNode creates a new consul Source pulling from the consul node at addr,
// using the provided prefix
func WithNode(addr string, prefix string) *Source {
	return &Source{
		addr:   addr,
		prefix: prefix,
	}
}

// Get retrieves the values from the consul source at location <prefix>/<key>
func (s *Source) Get(key string) ([]string, error) {
	if s.m == nil {
		if err := s.init(); err != nil {
			return nil, err
		}
	}

	val, ok := s.m[key]
	if !ok {
		return nil, fc.ErrMissing
	}

	return val, nil
}

func (s *Source) init() error {
	clientConfig := api.DefaultConfig()
	clientConfig.Address = s.addr
	client, err := api.NewClient(clientConfig)

	kv, _, err := client.KV().List(s.prefix, nil)
	if err != nil {
		return err
	}

	if kv == nil || len(kv) == 0 {
		return errors.Wrapf(err, "prefix not found: %s", s.prefix)
	}

	s.m = make(map[string][]string)

	for _, v := range kv {
		key := strings.TrimPrefix(v.Key, s.prefix)
		s.m[key] = strings.Split(string(v.Value), ",")
	}

	return nil
}

// Name returns a useful name for the consul config source for usage
func (s *Source) Name() string {
	return fmt.Sprintf("consul node at %s with prefix %q", s.addr, s.prefix)
}

// Loc returns the object key where the value is expected to be found
func (s *Source) Loc(key string) string {
	return fmt.Sprintf("%s%s", s.prefix, key)
}
