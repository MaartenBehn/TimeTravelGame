package math

type CardPos struct {
	X float64
	Y float64
}

func (c *CardPos) Add(b CardPos) CardPos {
	return CardPos{c.X + b.X, c.Y + b.Y}
}

func (c *CardPos) AddFloat(b float64) CardPos {
	return CardPos{c.X + b, c.Y + b}
}

func (c *CardPos) Sub(b CardPos) CardPos {
	return CardPos{c.X - b.X, c.Y - b.Y}
}

func (c *CardPos) SubFloat(b float64) CardPos {
	return CardPos{c.X - b, c.Y - b}
}

func (c *CardPos) Mul(b CardPos) CardPos {
	return CardPos{c.X * b.X, c.Y * b.Y}
}

func (c *CardPos) MulFloat(b float64) CardPos {
	return CardPos{c.X * b, c.Y * b}
}

func (c *CardPos) Div(b CardPos) CardPos {
	return CardPos{c.X / b.X, c.Y / b.Y}
}

func (c *CardPos) DivFloat(b float64) CardPos {
	return CardPos{c.X / b, c.Y / b}
}

func (c *CardPos) ToAxial() AxialPos {
	return AxialPos{c.Y - (2.0/3.0)*c.X, (4.0 / 3.0) * c.X}
}
