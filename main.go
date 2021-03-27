package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ADD_STATS_LINE = false
)

var (
	prefixesToCreate = ""
	suffixesToCreate = ""
)

func main() {
	createAffixesFromList()
	gen := itemGenerator{}
	rand.Seed(int64(time.Now().Unix()))

	initWeaponRuleset()
	initArmorRuleset()
	initBookRuleset()
	initDisposablesRuleset()
	for i := range weaponRuleset.rules {
		if i == 0 || i == -4 {
			continue
		}
		fmt.Println("===== " + weaponRuleset.rules[i].rarityName + ": =====")
		fmt.Println(" - " + gen.createItemByRule(weaponRuleset.rules[i]).getFullName(ADD_STATS_LINE, true))
		fmt.Println(" - " + gen.createItemByRule(armorRuleset.rules[i]).getFullName(ADD_STATS_LINE, true))
		fmt.Println(" - " + gen.createItemByRule(bookRuleset.rules[i]).getFullName(ADD_STATS_LINE, true))
		for j := 0; j < 1; j++ {
			fmt.Println(" - " + gen.createItemByRule(disposablesRuleset.rules[i]).getFullName(ADD_STATS_LINE, true))
		}
		fmt.Printf("\n")
	}
	fmt.Print()
}

func createAffixesFromList() {
	if prefixesToCreate != "" && suffixesToCreate != "" {
		fmt.Printf("Well, nothing will be done, both strs not empty")
		return
	}
	if prefixesToCreate != "" {
		splitted := strings.Split(prefixesToCreate, ", ")
		for _, s := range splitted {
			fmt.Printf("{\nlines: &affixLines{\nprefixForItemName: \"%s\",\n},\n},\n", strings.ToLower(s))
		}
	}
	if suffixesToCreate != "" {
		splitted := strings.Split(suffixesToCreate, ", ")
		for _, s := range splitted {
			fmt.Printf("{\nlines: &affixLines{\nsuffixForItemName: \"%s\",\n},\n},\n", strings.ToLower(s))
		}
	}
}
