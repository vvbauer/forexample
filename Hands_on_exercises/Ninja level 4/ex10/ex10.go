package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
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
