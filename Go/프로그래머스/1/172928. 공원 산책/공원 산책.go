import "strconv"
import "strings"

type Coordinate struct {
    X int
    Y int
}

func (c Coordinate) toArray() []int {
    return []int{c.Y, c.X}
}

type ParkMap struct {
    DogPosition Coordinate
    Hurdles     []Coordinate
    MaxDistance Coordinate
}

func createParkMap(park []string) *ParkMap{
    m := ParkMap{MaxDistance: Coordinate{X: (len(park[0])-1), Y: (len(park)-1)}}
    
    for i, row := range park {
        for j, char := range row {
            switch char {
            case 'S':
                m.DogPosition = Coordinate{X: j, Y: i}
            case 'X':
                newHurdle := Coordinate{X: j, Y: i}
                m.Hurdles = append(m.Hurdles, newHurdle)
            }
        }
    }
    
    return &m
}

func (m *ParkMap) validateCheck(cor Coordinate) bool {
    for _, h := range m.Hurdles {
        if h == cor {
            return false
        }
    }
    return true
}

func (m *ParkMap) isInBound(cor Coordinate) bool {
    return cor.X >= 0 && cor.X <= m.MaxDistance.X && cor.Y >= 0 && cor.Y <= m.MaxDistance.Y
}

func (m *ParkMap) move(routes []string) {
    for _, route := range routes {
        char := strings.Split(route, " ")
        direction := char[0]
        distance, _ := strconv.Atoi(char[1])
        
        toUpdate := m.DogPosition
        isValid := true
        
        for i := 1; i <= distance; i++ {
            switch direction {
            case "N":
                toUpdate.Y -= 1
            case "S":
                toUpdate.Y += 1
            case "W":
                toUpdate.X -= 1
            case "E":
                toUpdate.X += 1
            }
            if !m.isInBound(toUpdate) || !m.validateCheck(toUpdate) {
                isValid = false
                break
            }
        }
        
        if isValid {
            m.DogPosition = toUpdate
        }
    }
}

func solution(park []string, routes []string) []int {
    solutionMap := createParkMap(park)
    solutionMap.move(routes)
    return solutionMap.DogPosition.toArray()
}