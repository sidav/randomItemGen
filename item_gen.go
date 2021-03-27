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
	rarityName string
	ruleWeight int // the higher, the more probable the rule is to apply

	possibleItemNames     *[]string
	possibleAffixes       []affix
	possibleUniqueAffixes []affix
	totalAffixes          int
	totalUniqueAffixes    int

	aliasPreSyllables  []string
	aliasSyllables     []string
	aliasPostSyllables []string

	modifiersNamesAndPossibleValues map[string][]int
}

type resultingItem struct {
	modifiers map[string]int
	affixes   []*affix
	alias     string
	name      string
	rarityName string
}

type itemGenerator struct {
}

func (rs *ruleset) getWeightedRandomRule() *rule {
	totalWeight := 0
	for i := range rs.rules {
		if rs.rules[i].ruleWeight == 0 {
			rs.rules[i].ruleWeight = 1
		}
		totalWeight += rs.rules[i].ruleWeight
	}
	currWeight := rand.Intn(totalWeight)
	for _, r := range rs.rules {
		if r.ruleWeight > currWeight {
			return r
		} else {
			currWeight -= r.ruleWeight
		}
	}
	panic("Oh noes!")
}

func (ig *itemGenerator) createItemByRule(r *rule) *resultingItem {
	ri := resultingItem{
		modifiers: map[string]int{},
		affixes:   []*affix{},
		name:      (*r.possibleItemNames)[rand.Intn(len(*r.possibleItemNames))],
		rarityName: r.rarityName,
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
	if len(r.aliasSyllables) > 0 {
		ri.generateAlias(r)
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

func (ri *resultingItem) generateAlias(r *rule) {
	hasPrefix := rand.Intn(2) == 1
	hasSuffix := rand.Intn(2) == 1
	syllables := rand.Intn(3) + 1
	if hasPrefix && hasSuffix {
		syllables -= 1
	}
	for i := 0; i < syllables; i++ {
		ri.alias += r.aliasSyllables[rand.Intn(len(r.aliasSyllables))]
	}
	if hasPrefix {
		ri.alias = r.aliasPreSyllables[rand.Intn(len(r.aliasPreSyllables))] + ri.alias
	}
	if hasSuffix {
		ri.alias += r.aliasPostSyllables[rand.Intn(len(r.aliasPostSyllables))]
	}
}

func (ri *resultingItem) getFullName(addStatsLine, forceAffixes, addRarityName bool) string {
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

	startingAffixIndex := 0
	if len(ri.affixes) > 0 {
		startingAffixIndex = rand.Intn(len(ri.affixes))
	}
	for aliasIndex := range ri.affixes {
		currAffix := ri.affixes[(aliasIndex+startingAffixIndex)%len(ri.affixes)]
		if currAffix.lines.selectOnlyOne {
			prefixNotSuffix := rand.Intn(2) == 0
			if prefixNotSuffix {
				prefixes += currAffix.lines.prefixForItemName + " "
			} else {
				if suffixes != "" {
					if rand.Intn(2) == 0 {
						suffixes = suffixes + ", "
					} else {
						suffixes = suffixes + " "
					}
				}
				suffixes += currAffix.lines.suffixForItemName
			}
		} else {
			if currAffix.lines.prefixForItemName != "" {
				prefixes += currAffix.lines.prefixForItemName + " "
			}
			if currAffix.lines.suffixForItemName != "" {
				if suffixes != "" {
					if rand.Intn(2) == 0 {
						suffixes = suffixes + ", "
					} else {
						suffixes = suffixes + " "
					}
				}
				suffixes += currAffix.lines.suffixForItemName
			}
		}
		for _, e := range currAffix.listOfEffects {
			if e != "" {
				statsLine += e + ", "
			}
		}
	}
	if len(ri.alias) == 0 || forceAffixes {
		name = fmt.Sprintf("%s%s %s", prefixes, name, suffixes)
	} else {
		name = fmt.Sprintf("%s \"%s\" ", ri.name, ri.alias)
	}
	if statsLine != "" && addStatsLine {
		name += "{" + statsLine + "}"
	}
	name = strings.Replace(name, ", }", "}", 1)
	if addRarityName {
		name += fmt.Sprintf(" (%s)", ri.rarityName)
	}
	return name
}
