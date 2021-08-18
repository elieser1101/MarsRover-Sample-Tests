package marsrover

import (
	"fmt"
    "strconv"
)

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

func (d Direction) ArrowString() string {
	return [...]string{"^", ">", "v", "<"}[d]
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

func (r *MarsRover) currentLocation() interface{} {
	return fmt.Sprintf("%d %d %s", r.position.x, r.position.y, r.heading.String())
}

func (r *MarsRover) acceptCommands(commands []Command) {
    fmt.Println("started at")
    r.currentLocation()
    for _, c := range commands {
        switch c {
            case L:
                r.turnLeft()
            case R:
                r.turnRight()
            case F:
                r.forward()
            case B:
                r.backward()
        }
        r.printTerrain()
	}
}

func (r *MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) forward() {
        switch r.heading {
            case N:
                if r.plateau.maxY > r.position.y {
                    r.position.y += 1
                }
            case S:
                if r.position.y > 0{
                    r.position.y -= 1
                }
            case E:
                if r.plateau.maxX > r.position.x {
                    r.position.x += 1
                }
            case W:
                if r.position.x > 0{
                    r.position.x -= 1
                }
        }

}

func (r *MarsRover) backward() {
        switch r.heading {
            case N:
                r.position.y -= 1
            case S:
                r.position.y += 1
            case E:
                r.position.x -= 1
            case W:
                r.position.x += 1
        }

}

func (r *MarsRover) turnRight() {
    if r.heading == 3 {
        r.heading = 0
    } else {
        r.heading += 1
    }

}


func (r *MarsRover) printTerrain() {
    xAxis := ""
    stepSeparator := ""
	for i := 0; i < r.plateau.maxX; i++ {
        line := ""
        line = line + strconv.Itoa(r.plateau.maxX - 1 - i)
        for j := 0; j < r.plateau.maxY; j++ {
            if i == 0 {
                xAxis = xAxis + strconv.Itoa(j) + "\t"
                stepSeparator = stepSeparator + "#######"
            }
            if i == r.position.y && j == r.position.x {
                line = line + "[" + r.heading.ArrowString() + "]\t"
            } else {
                line = line + "*\t"
            }
        }
	    fmt.Println(line)
    }
    fmt.Println(xAxis)
    fmt.Println(stepSeparator)
}
