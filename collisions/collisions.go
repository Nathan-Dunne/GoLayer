package collisions

import (
	"math"

	"github.com/Nathan-Dunne/GoLayer/element"
)

func collides(c1, c2 element.Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) +
		math.Pow(c2.Center.Y-c1.Center.Y, 2))

	return dist <= c1.Radius+c2.Radius
}

func CheckCollisions() error {
	for i := 0; i < len(element.Elements)-1; i++ {
		for j := i + 1; j < len(element.Elements); j++ {
			for _, c1 := range element.Elements[i].Collisions {
				for _, c2 := range element.Elements[j].Collisions {
					if collides(c1, c2) && element.Elements[i].Active && element.Elements[j].Active {
						err := element.Elements[i].Collision(element.Elements[j])
						if err != nil {
							return err
						}

						err = element.Elements[j].Collision(element.Elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
