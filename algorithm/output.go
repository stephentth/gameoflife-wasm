package algorithm

import "fmt"

const LIVECELL = "<td class='cell live'></td>"
const DEADCELL = "<td class='cell dead'></td>"

// Print World into std.out
func (w World) Print() {
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			if w.board[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// RenderHTML ..
func (w World) RenderHTML() string {
	html := ""
	html += "<table>"
	for i := 0; i < w.height; i++ {
		html += "<tr>"
		for j := 0; j < w.width; j++ {
			if w.board[i][j] {
				html += LIVECELL
			} else {
				html += DEADCELL
			}
		}
		html += "</tr>\n"
	}
	html += "</table>\n"
	return html
}
