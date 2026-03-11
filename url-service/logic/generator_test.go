package logic

import "testing"

func TestCodeGenerator_Generate(t *testing.T) {
	generator := NewCodeGenerator(6)

	code := generator.Generate()

	if len(code) != 6 {
		t.Fatalf("expected code length 6, got %d", len(code))
	}
}

func TestCodeGenerator_GenerateDifferentCodes(t *testing.T) {
	generator := NewCodeGenerator(6)

	code1 := generator.Generate()
	code2 := generator.Generate()

	if code1 == code2 {
		t.Fatal("expected different codes, got identical values")
	}
}