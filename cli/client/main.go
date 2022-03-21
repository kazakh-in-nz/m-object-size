package main

import (
	"flag"
	"time"

	pbgameengine "github.com/kazakh-in-nz/m-game-engine/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60052", "address to connect")

	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to close connection")
		}
	}()

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-object-size gRPC service")
	}

	c := pbgameengine.NewGameEngineClient(conn)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.GetSize(timeoutCtx, &pbgameengine.GetSizeRequest{})

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("object size", r.GetSize()).Msg("Highscore from m-object-size microservice")
	} else {
		log.Error().Msg("Couldn't get object size")
	}
}
