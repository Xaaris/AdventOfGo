package main

// Pos represents a position on the grid
type Pos struct {
	x int
	y int
}

func (p *Pos) move(dx, dy int) {
	p.x += dx
	p.y += dy
}

func (p *Pos) isAdjacent(o *Pos) bool {
	if absDiffInt(p.x, o.x) <= 1 && absDiffInt(p.y, o.y) <= 1 {
		return true
	}
	return false
}

func (p *Pos) isInSameRow(o *Pos) bool {
	if absDiffInt(p.y, o.y) == 0 {
		return true
	}
	return false
}

func (p *Pos) isInSameCol(o *Pos) bool {
	if absDiffInt(p.x, o.x) == 0 {
		return true
	}
	return false
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
