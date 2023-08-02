package trim

import (
	"strings"
)

// TrimNewline adalah fungsi yang menghapus karakter newline (baris baru) dari string input.
// Fungsi ini menggunakan strings.TrimSpace() untuk membuang spasi ekstra di awal dan akhir string.
func TrimNewline(s string) string {
	return strings.TrimSpace(s)
}
