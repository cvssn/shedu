package env

import (
	"os"
)

type Env interface {
	Get(key string) string
	Env() []string
}

type OsEnv struct{}

// get implementa a env
func (o *osEnv) Get(key string) string {
	return os.Getenv(key)
}

func (o *osEnv) Env() []string {
	return os.Environ()
}

func New() Env {
	return &osEnv{}
}

type mapEnv struct {
	m map[string]string
}

// get implementa a env
func (m *mapEnv) Get(key string) string {
	if value, ok := m.m[key]; ok {
		return value
	}

	return ""
}

// env implementa a env
func (m *mapEnv) Env() []string {
	env := make([]string, 0, len(m.m))

	for k, v := range m.m {
		env = append(env, k+"="+v)
	}

	return env
}

func NewFromMap(m map[string]string) Env {
	if m == nil {
		m = make(map[string]string)
	}
	
	return &mapEnv{m: m}
}