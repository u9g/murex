package readline

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	modeNormal = iota
	modeTabCompletion
)

var (
	tcPrefix      string
	tcSuggestions []string
	tcPosX        int
	tcPosY        int
	tcMaxX        int
	tcMaxY        int
	tcUsedY       int
	tcMaxLength   int
)

func tabCompletion() {
	if TabCompletion == nil {
		return
	}

	tcPrefix, tcSuggestions = TabCompletion(line, pos)
	if len(tcSuggestions) == 0 {
		return
	}

	if len(tcSuggestions) == 1 {
		if len(tcSuggestions[0]) == 0 || tcSuggestions[0] == " " || tcSuggestions[0] == "\t" {
			return
		}
		insert([]byte(tcSuggestions[0]))
		return
	}

	initTabGrid()
	renderSuggestions()
}

func initTabGrid() {
	getTermWidth()

	tcMaxLength := 1
	for i := range tcSuggestions {
		if len(tcPrefix+tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len([]rune(tcPrefix + tcSuggestions[i]))
		}
	}

	mode = modeTabCompletion
	tcPosX = 1
	tcPosY = 1
	tcMaxX = termWidth / (tcMaxLength + 2)
	tcMaxY = MaxTabCompletionRows
}

func moveTabHighlight(x, y int) {
	tcPosX += x
	tcPosY += y

	if tcPosX < 1 {
		tcPosX = tcMaxX
		tcPosY--
	}

	if tcPosX > tcMaxX {
		tcPosX = 1
		tcPosY++
	}

	if tcPosY < 1 {
		tcPosY = tcUsedY
	}

	if tcPosY > tcUsedY {
		tcPosY = 1
	}

	if tcPosY == tcUsedY && (tcMaxX*(tcPosY-1))+tcPosX > len(tcSuggestions) {
		if x < 0 {
			tcPosX = len(tcSuggestions) - (tcMaxX * (tcPosY - 1))
		}

		if x > 0 {
			tcPosX = 1
			tcPosY = 1
		}

		if y < 0 {
			tcPosY--
		}

		if y > 0 {
			tcPosY = 1
		}
	}

	renderSuggestions()
}

func renderSuggestions() {
	fmt.Print("\r\n")
	moveCursorDown(hintY)

	cellWidth := strconv.Itoa((termWidth / tcMaxX) - 2)
	x := 0
	y := 1

	for i := range tcSuggestions {
		x++
		if x > tcMaxX {
			x = 1
			y++
			if y > tcMaxY {
				y--
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		if x == tcPosX && y == tcPosY {
			fmt.Print(seqBgWhite + seqFgBlack)
		}
		fmt.Printf(" %-"+cellWidth+"s %s", tcPrefix+tcSuggestions[i], seqReset)
	}

	tcUsedY = y
	moveCursorUp(y + hintY)
	moveCursorBackwards(termWidth)
	moveCursorForwards(promptLen + pos)
}

func clearTabSuggestions() {
	move := termWidth * tcUsedY
	blank := strings.Repeat(" ", move)

	// It's a bit ugly this code but it helps minimize the likelihood of the console
	// bug for becoming visible.
	//if hintY > 0 {
	fmt.Print(seqPosSave + "\r\n" + blank + seqPosRestore)
	/*} else {
		fmt.Print("\r\n")
		moveCursorDown(hintY)
		fmt.Print(blank)
		moveCursorBackwards(termWidth)
		moveCursorUp(hintY + tcUsedY)
		moveCursorForwards(promptLen + pos + 1)
	}*/

	mode = modeNormal
}
