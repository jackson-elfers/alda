package code_generator

import (
	"alda.io/client/model"
	"fmt"
	"io"
)

func Generate(score *model.Score, w io.Writer) {
	// TODO: Implement an Alda code generator
	// fmt.Fprintln(w, score.JSON().String())
	fmt.Fprintln(w, "...Some Alda Code...")
}
