package logic

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type CodeGenerator struct {
	rng    *rand.Rand
	length int
}

func NewCodeGenerator(length int) *CodeGenerator {
	source := rand.NewSource(time.Now().UnixNano())

	return &CodeGenerator{
		rng:    rand.New(source),
		length: length,
	}
}

func (g *CodeGenerator) Generate() string {
	b := make([]byte, g.length)

	for i := range b {
		b[i] = charset[g.rng.Intn(len(charset))]
	}
	return string(b)
}
