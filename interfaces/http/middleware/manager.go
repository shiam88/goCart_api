package middleware

import "net/http"

type Manager struct {
	global []Middleware
}

func NewManager() *Manager {
	return &Manager{global: make([]Middleware, 0)}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.global = append(m.global, middlewares...)
}

func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func (m *Manager) Wrap(handler http.Handler) http.Handler {
	h := handler
	for i := len(m.global) - 1; i >= 0; i-- {
		h = m.global[i](h)
	}
	return h
}
