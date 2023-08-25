package main

import (
	"fmt"
	"github.com/mindwingx/abstraction"
	"github.com/mindwingx/go-cache-wrapper"
	"github.com/mindwingx/go-http-wrapper"
	"github.com/mindwingx/go-locale-wrapper"
	"github.com/mindwingx/go-registry-wrapper"
	"github.com/mindwingx/go-rpc-wrapper"
	"github.com/mindwingx/go-sql-wrapper"
	"github.com/mindwingx/go-wrappers/rpcEntities"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		registry abstraction.Registry
		locale   abstraction.Locale
		cache    abstraction.Cache
		sql      abstraction.Sql
		rpcSrv   abstraction.RpcService
		http     abstraction.ApiService
	)

	// registry
	registry = registrywrapper.New()
	_ = registry.InitRegistry(
		"yml",
		fmt.Sprintf("%s/config.yaml", BasePath()),
	)

	// locale
	locale = localewrapper.NewLocale(registry.ValueOf("locale"))
	locale.InitLocaleJson([]string{
		fmt.Sprintf("%s/locale/%s", BasePath(), "en-US.json"),
	})

	fmt.Println(locale.Get("hello"))

	// cache
	cache = cachewrapper.New(registry.ValueOf("cache"), locale)
	cache.InitCache()

	// sql
	sql = sqlwrapper.NewSql(registry.ValueOf("locale"), locale)
	sql.InitSql()

	// rpc
	rpcSrv = rpcwrapper.New(registry.ValueOf("rpc"), locale)
	rpcSrv.InitRpcService([]interface{}{
		&rpcentities.UserRpc{},
	})

	go func() {
		rpcSrv.StartRpc()
	}()

	// http
	http = httpwrapper.NewGin(registry.ValueOf("http"), locale)
	go func() {
		http.StartHttp()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
