package a2a

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"

	einoA2A "github.com/cloudwego/eino-ext/a2a/extension/eino"
	"github.com/cloudwego/eino-ext/a2a/transport/jsonrpc"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/hertz/pkg/app"
	hertzServer "github.com/cloudwego/hertz/pkg/app/server"
)

// Server represents an A2A server instance that can register and run a single agent.
// This design moves away from the singleton pattern to provide better isolation and flexibility.
type Server struct {
	agent  adk.Agent          // The registered agent instance
	opts   *agentOption       // Configuration options for the agent
	mu     sync.RWMutex       // Mutex for thread-safe operations
	server *hertzServer.Hertz // The underlying HTTP server
}

// agentOption holds configuration options for the registered agent
type agentOption struct{}

// AgentOptionFn is a function type for configuring agent options using the functional options pattern
type AgentOptionFn func(*agentOption)

// runOption holds configuration options for running the server
type runOption struct {
	Host string // Server host address (e.g., "0.0.0.0", "localhost")
	Port int    // Server port number (e.g., 8080)

	Middlewares []app.HandlerFunc
}

// RunOptionFn is a function type for configuring run options using the functional options pattern
type RunOptionFn func(*runOption)

// WithHost sets the server host address
func WithHost(host string) RunOptionFn {
	return func(o *runOption) {
		o.Host = host
	}
}

// WithPort sets the server port number
func WithPort(port int) RunOptionFn {
	return func(o *runOption) {
		o.Port = port
	}
}

// WithMiddlewares sets the server middlewares
func WithMiddlewares(middlewares ...app.HandlerFunc) RunOptionFn {
	return func(o *runOption) {
		o.Middlewares = middlewares
	}
}

// New creates a new Server instance with default configuration
func New() *Server {
	return &Server{}
}

// RegisterAgent registers a single agent with the server.
// Only one agent can be registered per server instance to maintain simplicity and clarity.
func (s *Server) RegisterAgent(ctx context.Context, agent adk.Agent, opts ...AgentOptionFn) error {
	if agent == nil {
		return fmt.Errorf("agent cannot be nil")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.agent != nil {
		return fmt.Errorf("agent already registered, only one agent per server is allowed")
	}

	// Build agent options using the functional options pattern
	agentOpts := &agentOption{}

	for _, opt := range opts {
		opt(agentOpts)
	}

	s.agent = agent
	s.opts = agentOpts

	return nil
}

// Run starts the server and blocks until the context is cancelled or an error occurs.
// This method combines server startup and graceful shutdown in a single call.
func (s *Server) Run(ctx context.Context, opts ...RunOptionFn) error {
	// Ensure an agent is registered before starting
	s.mu.RLock()
	if s.agent == nil {
		s.mu.RUnlock()
		return fmt.Errorf("no agent registered")
	}
	agent := s.agent
	s.mu.RUnlock()

	// Build run options with defaults
	runOpts := &runOption{
		Host: "0.0.0.0", // Default to all interfaces
		Port: 8000,      // Default HTTP port
	}
	for _, opt := range opts {
		opt(runOpts)
	}

	// Create Hertz HTTP server instance
	h := hertzServer.Default(
		hertzServer.WithHostPorts(net.JoinHostPort(runOpts.Host, strconv.Itoa(runOpts.Port))),
	)
	s.server = h

	if len(runOpts.Middlewares) > 0 {
		h.Use(runOpts.Middlewares...)
	}

	// Create JSON-RPC registrar for handling agent communication
	r, err := jsonrpc.NewRegistrar(ctx, &jsonrpc.ServerConfig{
		Router:      h,
		HandlerPath: "/", // Handle requests at root path
	})
	if err != nil {
		return fmt.Errorf("failed to create registrar: %w", err)
	}

	// Register agent handlers with the A2A framework
	err = einoA2A.RegisterServerHandlers(ctx, agent, &einoA2A.ServerConfig{
		Registrar: r,
	})
	if err != nil {
		return fmt.Errorf("failed to register server handlers: %w", err)
	}

	return h.Run()
}
