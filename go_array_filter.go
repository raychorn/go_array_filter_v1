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
	canusethis := 0
	n := len(items)
	for i:= 0; i < n; i++ {
		if (j % 2) == 1 {
			if (opt.smalllarge) {
				if (i < (n-1)) {
					nextitem := items[i+1]
					if (opt.verbose) {
						msg := fmt.Sprintf("(S) %d - (L) %d", items[i], nextitem)
						fmt.Print(msg + " ")
					}
					if (nextitem < items[i]) {
						canusethis--
						if (opt.verbose) {
							fmt.Println("[-]")
						}
						continue
					}
				}
			} else {
				if (i < (n-1)) {
					nextitem := items[i+1]
					if (opt.verbose) {
						msg := fmt.Sprintf("(L) %d - (S) %d", items[i], nextitem)
						fmt.Print(msg + " ")
					}
					if (nextitem > items[i]) {
						canusethis--
						if (opt.verbose) {
							fmt.Println("[-]")
						}
						continue
					}
				}
			}
		} else {
			if (opt.smalllarge) {
				if (i < (n-1)) {
					nextitem := items[i+1]
					if (opt.verbose) {
						msg := fmt.Sprintf("(L) %d - (S) %d", items[i], nextitem)
						fmt.Print(msg + " ")
					}
					if (nextitem > items[i]) {
						canusethis--
						if (opt.verbose) {
							fmt.Println("[-]")
						}
						continue
					}
				}
			} else {
				if (i < (n-1)) {
					nextitem := items[i+1]
					if (opt.verbose) {
						msg := fmt.Sprintf("(S) %d - (L) %d", items[i], nextitem)
						fmt.Print(msg + " ")
					}
					if (nextitem < items[i]) {
						canusethis--
						if (opt.verbose) {
							fmt.Println("[-]")
						}
						continue
					}
				}
			}
		}
		j++
		if (opt.verbose) {
			fmt.Println("[+]")
		}
		canusethis++
	}
	if (opt.verbose) {
		//fmt.Println(fmt.Sprintf("canusethis -> %d\n", canusethis))
		fmt.Println("")
	}
	return canusethis == 4
}

func prettyPrintTheGuess(items []int) string {
	output := "{"
	delim := ""
	n := len(items)
	for i, value := range items {
		delim = ""
		if (i < n-1) {
			delim = ","
		}
		output += fmt.Sprintf("%d%s", value, delim)
	}
	output += "}"
	return output
}

func compareArrays(a, b []int) bool {
	isok := true
	if (len(a) != len(b)) {
		isok = false
	}
	if (isok) {
		for i, val := range a {
			if (val != b[i]) {
				isok = false
				break
			}
		}
	}
	return isok
}

func performTest(orig []int) [][]int {
	opts := options{smalllarge: true, verbose: false}

	theguess := []int{1, 7, 3, 9}
	useit := false

	responses := [][]int{}
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		theguess = getPerm(orig, p)
		useit = usethis(theguess, opts)
		if (useit) {
			responses = append(responses, theguess)
		}
	}
	return responses
}

func test1() {
	// Unit test mock-up.
    type test struct {
        input      []int
        expected   [][]int
    }

    tests := []test{
        {input: []int{1, 7, 3, 9}, expected: [][]int{{1, 7, 3, 9}, {1, 9, 3, 7}, {7, 9, 1, 3}, {3, 7, 1, 9}, {3, 9, 1, 7}} }, 
    } 

    for _, tc := range tests {
        // perform setUp before each test here
		actual := performTest(tc.input)
		for i, tcExpected := range tc.expected {
			fmt.Print(fmt.Sprintf("(%d) expected '%s', got '%s'", i, prettyPrintTheGuess(tcExpected), prettyPrintTheGuess(actual[i])))
			if (!compareArrays(tcExpected, actual[i])) {
				fmt.Println(fmt.Sprintf("(FAILED)"))
			} else {
				fmt.Println(fmt.Sprintf("(SUCCESS)"))
			}
		}
    }
}

func main2() {
	opts := options{smalllarge: true, verbose: false}

	orig := []int{1, 7, 3, 9}
	theguess := []int{1, 7, 3, 9}
	useit := false
	useitcount := 0
	fmt.Print(orig)
	fmt.Print(" --> ")
	if (opts.verbose) {
		fmt.Println()
	}
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		theguess = getPerm(orig, p)
		useit = usethis(theguess, opts)
		if (useit) {
			fmt.Print(prettyPrintTheGuess(theguess))
			fmt.Print(" ")
			useitcount++
		}
	}
	msg := fmt.Sprintf(" (There are %d)", useitcount)
	fmt.Print(msg)
	fmt.Println(" !!!")
}

func main() {
	test1()
}