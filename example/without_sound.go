package main

import (
	YtChat "github.com/akami-channel/youtube-live-chat-downloader/v2"
	"fmt"
	"log"
	
	//"os"
	//"time"
	// "github.com/faiface/beep"
	// "github.com/faiface/beep/mp3"
	// "github.com/faiface/beep/speaker"
)

func main() {
	continuation, cfg, error := YtChat.ParseInitialData("https://www.youtube.com/watch?v=86YLFOog4GM")
		//play_sound()
	if error != nil {
		log.Fatal(error)
	}
	for {
		chat, newContinuation, error := YtChat.FetchContinuationChat(continuation, cfg)
		if error != nil {
			log.Print(error)
			continue
		}
		continuation = newContinuation
		for _, msg := range chat {
			fmt.Print(msg.Timestamp, " | ")
			fmt.Println(msg.AuthorName, ": ", msg.Message)
			// uncomment if you want it to play a sound on each message
			//play_sound()
		}
	}
}

