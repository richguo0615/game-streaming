package controllers

import (
	"fmt"
	"github.com/richguo0615/game-streaming/proto"
	"time"
)

type GameController struct{}

func (GameController) Streaming(stream proto.Game_StreamingServer) (err error) {
	fmt.Println("streaming!")

	ticker := time.NewTicker(5 * time.Millisecond)
	for range ticker.C {
		packet, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(packet)
		fmt.Println("hearBeat...")
	}
	ticker.Stop()
	return
}
