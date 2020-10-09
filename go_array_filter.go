package main

import "fmt"

type options struct {
	smalllarge bool		// pick small then large or vice-versa.
	verbose bool
}

func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}

// rules :: 1st number is smaller than next

func usethis(items []int, opt options) bool {
	j := 1
	symbol := ""
	nextsymbol := ""
	canusethis := 0
	n := len(items)
	for i:= 0; i < n; i++ {
		if (i < (n-1)) {
			nextitem := items[i+1]
			if (nextsymbol == "S") {
				if (nextitem > items[i]) {
					canusethis--
					continue
				}
			} else {
				if (nextitem < items[i]) {
					canusethis--
					continue
				}
			}
		}
		if (j % 2) == 1 {
			if (opt.smalllarge) {
				symbol = "S"
				nextsymbol = "L"
			} else {
				symbol = "L"
				nextsymbol = "S"
			}
		} else {
			if (opt.smalllarge) {
				symbol = "L"
				nextsymbol = "S"
			} else {
				symbol = "S"
				nextsymbol = "L"
			}
		}
		if (opt.verbose) {
			fmt.Print(fmt.Sprint(items[i]) + symbol + " ")
		}
		j++
		canusethis++
	}
	if (opt.verbose) {
		//fmt.Println(fmt.Sprintf("canusethis -> %d\n", canusethis))
		fmt.Println("")
	}
	return canusethis == 4
}

func main() {
	opts := options{smalllarge: true, verbose: false}

	orig := []int{1, 7, 3, 9}
	theguess := []int{1, 7, 3, 9}
	useit := false
	useitcount := 0
	fmt.Print(orig)
	fmt.Print(" --> ")
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		theguess = getPerm(orig, p)
		useit = usethis(theguess, opts)
		if (useit) {
			fmt.Print(theguess)
			fmt.Print(" ")
			useitcount++
		}
	}
	msg := fmt.Sprintf(" (There are %d)", useitcount)
	fmt.Print(msg)
	fmt.Println(" !!!")
}
