package main

import (
	"fmt"
	"os"

	"github.com/parquet-go/parquet-go"
)

func main() {
	for _, p := range os.Args[1:] {
		f, _ := os.Open(p)
		st, _ := f.Stat()
		pf, err := parquet.OpenFile(f, st.Size())
		if err != nil {
			panic(err)
		}
		fmt.Println(p, pf.NumRows())
		fmt.Println(pf.Schema())
		r := parquet.NewGenericReader[map[string]any](pf)
		rows := make([]map[string]any, 1)
		n, err := r.Read(rows)
		fmt.Println(n, err)
		s := fmt.Sprint(rows[0])
		if len(s) > 600 {
			s = s[:600]
		}
		fmt.Println(s)
	}
}
