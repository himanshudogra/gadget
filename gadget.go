package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}
type TapeRecorder struct {
	Microphones int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Playing stopped")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Playing stopped")
}

func (t TapeRecorder) Record(song string) {
	fmt.Println("Started Recording", song)
}

type Player interface {
	Play(string)
	Stop()
}
