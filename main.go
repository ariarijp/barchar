package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

func main() {
	re := regexp.MustCompile(`^ *(\d+) +(.*)$`)
	lw, tw := 20, 60

	scanner := bufio.NewScanner(os.Stdin)
	max, r := -1, float64(-1)
	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())

		n, err := strconv.Atoi(m[1])
		if err != nil {
			log.Println(err)
		}
		l := fmt.Sprintf(m[2])

		if max < 0 {
			max = n
			r = float64(tw) / float64(max)
		}

		w := int(float64(n) * float64(r))
		ts := runewidth.Truncate(l, lw, "...")
		tsw := runewidth.StringWidth(ts)
		if tsw != lw {
			ts += strings.Repeat(" ", lw-tsw)
		}

		fmt.Println(ts, strings.Repeat("#", w), n)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
