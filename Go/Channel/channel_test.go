package Channel

import (
	"fmt"
	"testing"
)

func TestChannelOperate(t *testing.T) {
	ChannelOperate()
}

func TestChannelFor(t *testing.T) {
	ChannelFor()
}

func TestNoBufferChannel(t *testing.T) {
	NoBufferChannel()
}

func TestBufferChannel(t *testing.T) {
	BufferChannel()
}

func TestChannelGroutineNumCtl(t *testing.T) {
	ChannelGroutineNumCtl()
}

func TestChannelDirection(t *testing.T) {
	ChannelDirection()
}

func TestSelectTest(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "----------------")
		SelectTest()
	}
}

func TestForSelect(t *testing.T) {
	ForSelect()
}

func TestNilChannel(t *testing.T) {
	NilChannel()
}

func TestGuessChannel(t *testing.T) {
	GuessChannel()
}

func TestSelectNonBlock(t *testing.T) {
	SelectNonBlock()
}

func TestRaceTest(t *testing.T) {
	RaceTest()
}
