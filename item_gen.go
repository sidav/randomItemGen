package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type ruleset struct {
	possibleNames         []string
	possbleAffixes        []affix
	possibleUniqueAffixes []affix
	rules                 []*rule
}

type affixLines struct {
	prefixForItemName string // _flaming_ sword
	suffixForItemName string // sword _of flame_
	selectOnlyOne     bool
}

type affix struct {
	lines         *affixLines
	listOfEffects []string // flaming sword {_flame_}
	modifiers     map[string][]int
}

type rule struct {
	rarityName            string
	possibleItemNames     *[]string
	possibleAffixes       []affix
	possibleUniqueAffixes []affix
	totalAffixes          int
	totalUniqueAffixes    int

	uniqueNamePrefixes  []string
	uniqueNameSyllables []string
	uniqueNameSuffixes  []string

	modifiersNamesAndPossibleValues map[string][]int
}

type resultingItem struct {
	modifiers  map[string]int
	affixes    []*affix
	uniqueName string
	name       string
}

type itemGenerator struct {
}

func (ig *itemGenerator) createItemByRule(r *rule) *resultingItem {
	ri := resultingItem{
		modifiers: map[string]int{},
		affixes:   []*affix{},
		name:      (*r.possibleItemNames)[rand.Intn(len(*r.possibleItemNames))],
	}
	for k, v := range r.modifiersNamesAndPossibleValues {
		ri.modifiers[k] = 0 + v[rand.Intn(len(v))]
	}

	for i := 0; i < r.totalAffixes; i++ {
		if len(ri.affixes) == r.totalAffixes {
			break
		}
		rndPrefix := rand.Intn(len(r.possibleAffixes))
		// check if that affixLines is already used
		for ri.hasTheAffix(r.possibleAffixes[rndPrefix].lines.prefixForItemName,
			r.possibleAffixes[rndPrefix].lines.suffixForItemName) {
			rndPrefix = rand.Intn(len(r.possibleAffixes))
		}
		ri.affixes = append(ri.affixes, &r.possibleAffixes[rndPrefix])
	}
	if r.totalUniqueAffixes > 0 {
		for i := 0; i < r.totalUniqueAffixes; i++ {
			uaffix := r.possibleUniqueAffixes[rand.Intn(len(r.possibleUniqueAffixes))]
			for k, v := range uaffix.modifiers {
				ri.addModifier(k, v[rand.Intn(len(v))])
			}
			if !ri.hasTheAffix(uaffix.lines.prefixForItemName, uaffix.lines.suffixForItemName) {
				ri.affixes = append(ri.affixes, &uaffix)
			}
		}
	}
	if len(r.uniqueNameSyllables) > 0 {
		ri.generateUniqueName(r)
	}
	return &ri
}

func (ri *resultingItem) addModifier(name string, value int) {
	ri.modifiers[name] += value
}

func (ri *resultingItem) hasTheAffix(prefix, suffix string) bool {
	for _, p := range ri.affixes {
		if p.lines.prefixForItemName == prefix && p.lines.suffixForItemName == suffix {
			return true
		}
	}
	return false
}

func (ri *resultingItem) generateUniqueName(r *rule) {
	hasPrefix := rand.Intn(2) == 1
	hasSuffix := rand.Intn(2) == 1
	syllables := rand.Intn(3) + 1
	if hasPrefix && hasSuffix {
		syllables -= 1
	}
	for i := 0; i < syllables; i++ {
		ri.uniqueName += r.uniqueNameSyllables[rand.Intn(len(r.uniqueNameSyllables))]
	}
	if hasPrefix {
		ri.uniqueName = r.uniqueNamePrefixes[rand.Intn(len(r.uniqueNamePrefixes))] + ri.uniqueName
	}
	if hasSuffix {
		ri.uniqueName += r.uniqueNameSuffixes[rand.Intn(len(r.uniqueNameSuffixes))]
	}
}

func (ri *resultingItem) getFullName(addStatsLine, forceAffixes bool) string {
	name := ri.name
	prefixes := ""
	suffixes := ""
	statsLine := ""
	for k, val := range ri.modifiers {
		if val == 0 {
			continue
		}
		valString := strconv.Itoa(val)
		if !(val < 0) {
			valString = "+" + valString
		}
		statsLine += k + " " + valString + ", "
	}

	for _, pref := range ri.affixes {
		if pref.lines.selectOnlyOne {
			prefixNotSuffix := rand.Intn(2) == 0
			if prefixNotSuffix {
				prefixes += pref.lines.prefixForItemName + " "
			} else {
				if suffixes != "" {
					suffixes = suffixes + ", "
				}
				suffixes += pref.lines.suffixForItemName
			}
		} else {
			if pref.lines.prefixForItemName != "" {
				prefixes += pref.lines.prefixForItemName + " "
			}
			if pref.lines.suffixForItemName != "" {
				if suffixes != "" {
					suffixes = suffixes + ", "
				}
				suffixes += pref.lines.suffixForItemName
			}
		}
		for _, e := range pref.listOfEffects {
			if e != "" {
				statsLine += e + ", "
			}
		}
	}
	if len(ri.uniqueName) == 0 || forceAffixes {
		name = fmt.Sprintf("%s%s %s", prefixes, name, suffixes)
	} else {
		name = fmt.Sprintf("%s \"%s\" ", ri.name, ri.uniqueName)
	}
	if statsLine != "" && addStatsLine {
		name += "{" + statsLine + "}"
	}
	name = strings.Replace(name, ", }", "}", 1)
	return name
}
