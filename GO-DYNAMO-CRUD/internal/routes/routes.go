package routes

import (
	ServerConfig "githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/config"
	HealthHandler "githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/internal/handlers/health"
	ProductHandler "githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/internal/handlers/product"
	"githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/internal/repository/adapter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repo adapter.Interface) *chi.Mux {
	r.setConfigsRouters()
	r.RouterHealth(repo)
	r.RouterProduct(repo)

	return r.router
}

func (r *Router) setConfigsRouters() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repo adapter.Interface) {
	handler := HealthHandler.NewHandler(repo)

	r.router.Route("/health", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repo adapter.Interface) {
	handler := ProductHandler.NewHandler(repo)

	r.router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}
