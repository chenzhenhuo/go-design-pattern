package decorator

// Picture 画
type Picture interface {
	Do() string
}

type Paper struct {
}

func (p *Paper) Do() string {
	return "纸"
}

type ColorDo struct {
	paper Picture
	color string
}

func (receiver *ColorDo) Do() string {
	return "这是" + receiver.color + "的" + receiver.paper.Do()
}

func NewPicture(paper Picture, color string) *ColorDo {
	return &ColorDo{
		paper: paper,
		color: color,
	}
}
