package main

import "github.com/himanshudogra/gadget"

func PlayList(device gadget.Player, songs []string) {

	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var device gadget.Player = gadget.TapePlayer{}

	mixtape := []string{"tere naam", "kal ho na ho"}

	PlayList(device, mixtape)

	device = gadget.TapeRecorder{}
	PlayList(device, mixtape)
}
