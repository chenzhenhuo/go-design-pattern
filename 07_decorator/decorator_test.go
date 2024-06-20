package decorator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}

func TestColorPaper(t *testing.T) {
	pp := &Paper{}
	cp := NewPicture(pp, "红色")
	fmt.Println(cp.Do())
}
