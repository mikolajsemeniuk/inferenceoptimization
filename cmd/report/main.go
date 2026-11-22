// Collects every output/<benchmark>/**/summary.json into one markdown file and one CSV per
// benchmark (output/report/). Distribution fields are reduced to their mean.
//
//	go run ./cmd/report
package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// flatten turns one summary object into column -> printable value.
func flatten(row map[string]any) map[string]string {
	out := map[string]string{}
	for k, v := range row {
		switch x := v.(type) {
		case map[string]any: // metrics.Summary
			if m, ok := x["mean"].(float64); ok {
				out[k] = fmt.Sprintf("%.4g", m)
			}
		case float64:
			out[k] = fmt.Sprintf("%.4g", x)
			if x == float64(int64(x)) {
				out[k] = fmt.Sprint(int64(x))
			}
		case []any:
			out[k] = strings.Trim(fmt.Sprint(x), "[]")
		default:
			out[k] = fmt.Sprint(x)
		}
	}
	return out
}

func main() {
	tables := map[string][]map[string]string{} // benchmark -> rows
	err := filepath.WalkDir("output", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || d.Name() != "summary.json" || strings.Contains(path, "/smoke/") || strings.Contains(path, "/plumbing/") {
			return err
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var rows []map[string]any
		if json.Unmarshal(raw, &rows) != nil {
			return nil
		}
		name := strings.Split(filepath.ToSlash(path), "/")[1]
		for _, r := range rows {
			f := flatten(r)
			f["source"] = filepath.ToSlash(filepath.Dir(path))
			tables[name] = append(tables[name], f)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("output/report", 0o755); err != nil {
		log.Fatal(err)
	}
	var names []string
	for n := range tables {
		names = append(names, n)
	}
	sort.Strings(names)
	var md strings.Builder
	for _, name := range names {
		rows := tables[name]
		seen := map[string]bool{}
		var cols []string
		for _, r := range rows {
			for k := range r {
				if !seen[k] {
					seen[k] = true
					cols = append(cols, k)
				}
			}
		}
		sort.Strings(cols)
		var csv strings.Builder
		csv.WriteString(strings.Join(cols, ",") + "\n")
		fmt.Fprintf(&md, "## %s\n\n| %s |\n|%s\n", name, strings.Join(cols, " | "), strings.Repeat("---|", len(cols)))
		for _, r := range rows {
			vals := make([]string, len(cols))
			for i, c := range cols {
				vals[i] = r[c]
			}
			fmt.Fprintf(&md, "| %s |\n", strings.Join(vals, " | "))
			csv.WriteString(strings.Join(vals, ",") + "\n")
		}
		md.WriteString("\n")
		if err := os.WriteFile("output/report/"+name+".csv", []byte(csv.String()), 0o644); err != nil {
			log.Fatal(err)
		}
	}
	if err := os.WriteFile("output/report/report.md", []byte(md.String()), 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote output/report/report.md for:", strings.Join(names, ", "))
}
