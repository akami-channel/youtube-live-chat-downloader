package main

import (
	YtChat "github.com/akami-channel/youtube-live-chat-downloader/v2"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	continuation, cfg, error := YtChat.ParseInitialData("https://www.youtube.com/watch?v=ydYDqZQpim8")
		play_sound()
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
			play_sound()
		}
	}
}

func play_sound() {
	f, err := os.Open("test.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
