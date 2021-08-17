package marsrover

type Coordinates struct {
	x int
	y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)
func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
	B Command = iota
	F
	L
	R
)

func (c Command) String() string {
	return [...]string{"B", "F", "L", "R"}[c]
}

type Obstacle struct {
	position Coordinates
}

type Plateau struct {
	maxX int
	maxY int
	obstacles []Obstacle
}

type Status int

const (
	OK Status = iota
	NOK
)
func (s Status) String() string {
	return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
	plateau Plateau
	heading Direction
	position Coordinates
	status Status
}

func (r *MarsRover) turnLeft() {
    if r.heading == 0 {
        r.heading = 3
    } else {
        r.heading -= 1
    }

}

func (r MarsRover) currentLocation() interface{} {
	return ""
}

func (r *MarsRover) acceptCommands(commands []Command) {
    for _, c := range commands {
        switch c {
            case L:
                r.turnLeft()
            case R:
                r.turnRight()
            case F:
                r.forward()
        }
	}
}

func (r *MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) forward() {
        switch r.heading {
            case N:
                r.position.y += 1
            case S:
                r.position.y -= 1
            case E:
                r.position.x += 1
            case W:
                r.position.x -= 1
        }

}

func (r MarsRover) backward() {

}

func (r MarsRover) turnRight() {
    if r.heading == 3 {
        r.heading = 0
    } else {
        r.heading += 1
    }

}
