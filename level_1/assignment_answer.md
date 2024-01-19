# 1. Bagaimanakah dependency management dalam golang?
Dependency dalam Golang adalah modul-modul yang berisi packages yang kita import. Langkah-langkah untuk dependency management di antaranya adalah:
1. Mencari packages yang ingin digunakan di pkg.go.dev
2. Import packages yang diinginkan di code
3. Tambahkan kode ke dalam sebuah modul untuk dependency tracking, apabila belum berada dalam modul (go mod)
4. Tambah external packages sebagai dependency
5. Updgrade atau downgrade dependency versions jika diperlukan

# 2. Jelaskan kegunaan function fmt.Sprintln apa bedanya dengan fmt.Println? Beri contoh code, copy-paste outputnya.
`fmt.Sprintln` bisa dibilang mirip dengan `fmt.Println`, hanya saja `fmt.Sprintln` tidak menampilkan nilai, melainkan mengembalikan nilainya dalam bentuk string. Fungsi tersebut akan memberikan new line di akhir.

Kode:
```
package main

import "fmt"

func main() {

	const name = "Salsabila"
	s := fmt.Sprintln("Assalamu'alaikum! Saya", name)

	fmt.Println(s)

}
```

Output:
```
Assalamu'alaikum! Saya Salsabila

```

# 3. Jelaskan kegunaan function fmt.Errorf apa bedanya dengan errors.New? Beri contoh code, copy-paste outputnya.
Fungsi `errors.New` hanya akan mengembalikan error dengan pesan error statis, sementara fungsi `fmt.Errorf` memungkin formatting string dengan values, sama seperti `fmt.Printf` atau `fmt.Sprintf`

package main

import (
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}

