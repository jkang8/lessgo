package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
  "os"
  "log"
  //"fmt"
)

// Lines struct
type Lines struct {
  lines []string
  lines_focus []string
  min int
  max int
}

func NewLines(min int, max int) *Lines {
  l := new(Lines)
  l.lines = make([]string, 0)
  l.lines_focus = make([]string, 0)
  l.min = min
  l.max = max
  return l
}

func (l *Lines) scrollDown() []string {
  l.min++
  l.max++
  l.lines_focus = l.lines[l.min:l.max]
  return l.lines_focus
}

func (l *Lines) scrollUp() []string {
  l.min--
  l.max--
  l.lines_focus = l.lines[l.min:l.max]
  return l.lines_focus
}

func (l *Lines) append(text string) {
    l.lines = append(l.lines, text)
}

// Main
func drawLine(text string, line_number int) {
  for pos, char := range text {
      termbox.SetCell(pos, line_number, char, 'b', 'a')
  }
  termbox.Flush()
}

func initial(linesObject *Lines, scanner *bufio.Scanner, term_height int, line_number int) (*Lines) {
        for line_number < term_height && scanner.Scan() {
            text := scanner.Text()
            linesObject.append(text)
            drawLine(text, line_number)
						line_number++
        }
				return linesObject
}

func handleScrollDown(lines *Lines, next_line string) {
    //termbox.Flush()
    termbox.Clear('b', 'a')
    counter := 0
    //current_lines := lines[1:line_number+1]
    current_lines := lines.scrollDown()
    for _, str := range current_lines {
      drawLine(string(str), counter)
      counter++
    }
}

func handleScrollUp(lines *Lines) {
    termbox.Clear('b', 'a')
    counter := 0
    current_lines := lines.scrollUp()
    for _, str := range current_lines {
      drawLine(string(str), counter)
      counter++
    }
}

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }

    args := os.Args[1:]
    filename := args[0]
		_, term_height := termbox.Size()
		//var lines []string
    //lines := Lines{lines: make([]string, 0), min: 0, max: term_height}
    lines := NewLines(0, term_height)

    if file, err := os.Open(filename); err == nil {
        scanner := bufio.NewScanner(file)
				lines := initial(lines, scanner, term_height, 0)

				// main loop
				listener:
				for {
						switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                switch ev.Key {
                    case termbox.KeyEsc:
                        break listener
                    case termbox.KeyArrowDown:
                        scanner.Scan()
                        next_line := scanner.Text()
                        lines.append(next_line)
                        handleScrollDown(lines, next_line)
                    case termbox.KeyArrowUp:
                        handleScrollUp(lines)
                }
            case termbox.EventResize:
            }
				}
        defer file.Close()
    } else {
        log.Fatal(err)
    }

    defer termbox.Close()
}

