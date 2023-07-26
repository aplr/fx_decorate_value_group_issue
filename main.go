package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// MARK: -- APP

func main() {
	log := zap.Must(zap.NewProduction())

	app := fx.New(
		fx.Supply(log),
		fx.Module("my_service",
			fx.Decorate(func(log *zap.Logger) *zap.Logger {
				return log.Named("my_service")
			}),
			fx.Provide(NewDummyHandler),
			fx.Provide(NewMyService),
			fx.Invoke(func(srv *MyService) {
				srv.Run()
			}),
			// fx.Invoke(func(dot fx.DotGraph) {
			// 	fmt.Print(dot)
			// }),
		),
	)

	app.Run()
}

// MARK: -- HANDLER

type ServiceHandler struct {
	Name    string
	Handler func() error
}

type DummyHandlerParams struct {
	fx.In

	Log *zap.Logger
}

type HandlerResult struct {
	fx.Out

	Handler *ServiceHandler `group:"handlers"`
}

func NewDummyHandler(params DummyHandlerParams) HandlerResult {
	return HandlerResult{
		Handler: &ServiceHandler{
			Name: "dummy",
			Handler: func() error {
				params.Log.Named("dummy").Info("name should be `my_service.dummy`")
				return nil
			},
		},
	}
}

// MARK: -- SERVICE

type MyService struct {
	log      *zap.Logger
	handlers []*ServiceHandler
}

type ServiceParams struct {
	fx.In

	Log      *zap.Logger
	Handlers []*ServiceHandler `group:"handlers"`
}

func NewMyService(params ServiceParams) *MyService {
	return &MyService{
		log:      params.Log.Named("service"),
		handlers: params.Handlers,
	}
}

func (s *MyService) Run() error {
	s.log.Info("name should be `my_service.service`")

	for _, handler := range s.handlers {
		if err := handler.Handler(); err != nil {
			return err
		}
	}

	return nil
}
