package server

import (
	"backend/app/config"
	"backend/app/handler"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func Run() error {
	// db, err := infrastructure.NewDB(config.MySQLConfig().FormatDSN())
	// if err != nil {
	// 	return err
	// }

	//accountUsecase := usecase.NewAcocunt(db, dao.NewAccount(db))

	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)

	r := handler.NewRouter()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}
