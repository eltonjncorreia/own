package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// application é a estrutura que armazena as
// dependências e usada dentro dos manipuladores e funções.
type application struct {
	Logger *slog.Logger
}

func main() {
	// define porta do servidor, seu valor default
	// vem do valor da variável de ambiente PORT
	port := flag.String("port", fmt.Sprintf(":%v", os.Getenv("PORT")), "server port")

	flag.Parse()

	// instancia um log estruturado
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Instancia uma aplication e injeta logger como dependência.
	app := application{Logger: logger}

	app.Logger.Info("starting server on ", "port:", *port)

	// Inicia servidor
	err := http.ListenAndServe(*port, app.routes())
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}

// routes define a url e manipuladores
func (app application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.shortLink)

	return mux
}

// ShortLink exibe o formulário para criar um link encurtado
func (a *application) shortLink(w http.ResponseWriter, r *http.Request) {
	a.Logger.Info("Home")
}
