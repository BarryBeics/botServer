package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/barrybeics/botServer/graph"
	"github.com/rs/zerolog/log"

	"flag"
	"fmt"

	"github.com/rs/zerolog"
)

const defaultPort = "8080"

func main() {
	// Initialize logger
	SetupLogger()

	// Iterate over environment variables and print them
	for _, envVar := range os.Environ() {
		fmt.Println(envVar)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info().Str("Port", port).Msg("connect to http://localhost: for GraphQL playground on:")

	// Use a channel to block the main goroutine
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Run the server in a goroutine
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Block the main goroutine until a signal is received
	<-stop
}

// SetupLogger is a function that configures logging for the application. It
// sets the output format to include caller information, parses command-line
// arguments, sets the log level based on the LOG_LEVEL environment variable
// or the -debug flag (default: Info), and prints the log level to the console
// for debugging purposes. The function uses the zerolog logging library and
// the standard flag package to handle command-line arguments.
func SetupLogger() {
	// Set output format to include caller information
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	// Parse command-line arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-debug] [-trace]\n", os.Args[0])
		flag.PrintDefaults()
	}
	debug := flag.Bool("debug", false, "enable debug logging")
	trace := flag.Bool("trace", false, "enable trace logging")
	flag.Parse()

	// Set log level based on environment variable (default: Info)
	logLevel := zerolog.InfoLevel
	if envLogLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if lvl, err := zerolog.ParseLevel(envLogLevel); err == nil {
			logLevel = lvl
		}
	} else if *debug {
		logLevel = zerolog.DebugLevel
	} else if *trace {
		logLevel = zerolog.TraceLevel
	}

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

}
