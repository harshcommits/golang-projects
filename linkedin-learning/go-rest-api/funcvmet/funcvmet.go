package funcvmet

type Dimensions struct {
	Length int
	Width  int
	Height int
}

// method definition
func (d Dimensions) Area() int {
	return d.Length * d.Width
}

func (d Dimensions) Volume() int {
	return d.Length * d.Width * d.Height
}

// function definition
func Values(length, width, height int) (area int, volume int) { // all the parameters of type int; this is the shorthand
	area = length * width
	volume = length * width * height

	return // no need to pass return values explicitly, since already defined at line 7
}
