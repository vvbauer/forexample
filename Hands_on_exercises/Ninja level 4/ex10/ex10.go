package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m["man_iron"] = []string{`Miss Potts`, `Fast cars`, `Engineering`}

	fmt.Println(m)

	delete(m, `no_dr`)

	for i, v := range m {
		fmt.Println("Record for", i)
		for k, s := range v {
			fmt.Println("\t", k, s)
		}
	}
}
