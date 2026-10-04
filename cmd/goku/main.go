package main

import (
	"fmt"
	"os"

	"github.com/sirojmu/Goku/internal/scaffold"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "new" {
		fmt.Fprintln(os.Stderr, "Penggunaan: goku new <nama-proyek>")
		os.Exit(2)
	}
	if err := scaffold.Create(".", os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, "Gagal membuat proyek:", err)
		os.Exit(1)
	}
	fmt.Printf("Proyek %q dibuat. Jalankan: cd %s && go mod tidy && go run .\n", os.Args[2], os.Args[2])
}
