package marsrover

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReceiveSingleCommandShouldMove(t *testing.T) {
	plateau := Plateau{maxX: 10, maxY: 10}
	startingPosition := Coordinates{1,8}
	marsRover := &MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{F}
	marsRover.acceptCommands(commands)

	expectedPosition := Coordinates{1,9}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
}
