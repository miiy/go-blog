package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/service/web/internal/app/article"
	"goblog.com/service/web/internal/app/book"
	"goblog.com/service/web/internal/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Config  *config.Config
	Router  *gin.Engine
	Logger  *zap.Logger
	Article *article.Article
	Book    *book.Book
}

func NewApplication(config *config.Config, router *gin.Engine, logger *zap.Logger, article *article.Article, book *book.Book) *Application {
	return &Application{
		Config: config,
		Logger: logger,
		Router: router,
		Article: article,
		Book: book,
	}
}

var ProviderSet = wire.NewSet(NewApplication)

func (app *Application) RegisterRouter(funcs ...func(r *gin.Engine)) {
	for _, f := range funcs {
		f(app.Router)
	}
}

// Run
// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
func (app *Application) Run(addr ...string) {
	o := app.Config.App
	address := resolveAddress(addr, app.Config.App)
	if o.Debug {
		fmt.Printf( "[App-debug] Listening and serving HTTP on %s\n", address)
	}

	srv := &http.Server{
		Addr:    address,
		Handler: app.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func resolveAddress(addr []string, c config.App) string {
	switch len(addr) {
	case 0:
		return c.Addr
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
