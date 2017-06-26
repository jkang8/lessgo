package main

import (
  "github.com/nsf/termbox-go"
  "bufio"
  //"io"
  //"io/ioutil"
  "os"
  "log"
  //"reflect"
  //"fmt"
)

func drawLine(text string, line_number int) {
  //const coldef = termbox.ColorDefault
  //termbox.Clear(coldef, coldef)
  //w, h := termbox.Size()
  for pos, char := range text {
      termbox.SetCell(pos, line_number, char, 'b', 'a')
  }
  termbox.Flush()
}

func initial(pages []string, scanner *bufio.Scanner, term_height int, line_number int) (int) {
        for line_number < term_height && scanner.Scan() {
            text := scanner.Text()
						pages = append(pages, text)
            //pages[line_number] = text
            drawLine(text, line_number)
						line_number++
        }
				return line_number
}

func scrollDown(pages []string, next_line string, line_number int) {
		drawLine(next_line, line_number)
}

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }

    args := os.Args[1:]
    filename := args[0]
    //pages := make(map[int]string)
		var pages []string
		_, term_height := termbox.Size()

    if file, err := os.Open(filename); err == nil {
        scanner := bufio.NewScanner(file)
				line_number := initial(pages, scanner, term_height, 0)

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
														line_number++
														scrollDown(pages, next_line, line_number)
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

