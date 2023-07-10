package pdfbox

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
)

// TestRealPDF open an actual file ...
func TestRealPDF(t *testing.T) {
	ctx := context.Background()

	p, nerr := New(ctx, "pdfbox://")
	if nerr != nil {
		t.Fatal(nerr)
	}
	// Extract text from PDF
	r, err := os.Open("pdf/murray.pdf")
	if err != nil {
		t.Fatalf("Failed to open %s, %v", "murray.pdf", err)
	}
	defer r.Close()

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = p.ExecuteWithReaderAndWriter(ctx, r, wr, "ExtractText", READER, WRITER)

	if err != nil {
		t.Fatalf("Failed to extract text, %v", err)
	}

	wr.Flush()

	body := buf.Bytes()
	//body = bytes.TrimSpace(body)
	fmt.Print(string(body))

	os.WriteFile("data/murray.txt", body, 0644)

	err = p.Close()
	if err != nil {
		t.Fatal(err)
	}
}
