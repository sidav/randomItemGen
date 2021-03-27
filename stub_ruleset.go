package main

var (
	stubRuleset = ruleset{
		possibleNames: []string{
			"sword", "hammer",
		},

		possbleAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "flaming",
					suffixForItemName: "of flame",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"flame"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "of ice",
				},
				listOfEffects: []string{"freeze"},
			},
		},

		possibleUniqueAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "of great warrior",
				},
				listOfEffects: []string{"smash"},
				modifiers: map[string][]int{
					"Damage":   {3, 6, 10, 15},
					"Accuracy": {2, 4, 6, 8},
				},
			},
		},
	}
)

func initStubRuleset() {
	stubRuleset.rules = []*rule {
		&rule{
			rarityName:        "Common",
			possibleItemNames: &stubRuleset.possibleNames,
			possibleAffixes:   nil,
			totalAffixes:      0,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Damage":   {-2, -1, 0, 0, 1, 2},
				"Accuracy": {-2, -1, 0, 0, 1, 2},
			},
		},
		&rule{
			rarityName:        "Uncommon",
			possibleItemNames: &stubRuleset.possibleNames,
			possibleAffixes:   stubRuleset.possbleAffixes,
			totalAffixes:      1,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Damage":   {-2, -1, 0, 0, 1, 2, 3},
				"Accuracy": {-2, -1, 0, 0, 1, 2, 3},
			},
		},
		&rule{
			rarityName:        "Rare",
			possibleItemNames: &stubRuleset.possibleNames,
			possibleAffixes:   stubRuleset.possbleAffixes,
			totalAffixes:      2,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Damage":   {-2, -1, -1, 0, 1, 1, 2, 3},
				"Accuracy": {-2, -1, 0, -1, 0, 1, 1, 2, 3},
			},
		},
		&rule{
			rarityName:            "Epic",
			possibleItemNames:     &stubRuleset.possibleNames,
			possibleAffixes:       stubRuleset.possbleAffixes,
			totalAffixes:          2,
			possibleUniqueAffixes: stubRuleset.possibleUniqueAffixes,
			totalUniqueAffixes:    1,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Damage":   {-3, -2, -1, 0, 1, 2, 3, 4, 5},
				"Accuracy": {-3, -2, -1, 0, 1, 2, 3, 4, 5},
			},
		},
		&rule{
			rarityName:            "LEGENDARY",
			possibleItemNames:     &stubRuleset.possibleNames,
			possibleAffixes:       stubRuleset.possbleAffixes,
			totalAffixes:          3,
			possibleUniqueAffixes: stubRuleset.possibleUniqueAffixes,
			totalUniqueAffixes:    2,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Damage":   {-3, -2, -1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5},
				"Accuracy": {-3, -2, -1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5},
			},
			aliasPreSyllables:  []string{"Qu", "Li", "Non-"},
			aliasSyllables:     []string{"kir", "qua", "li"},
			aliasPostSyllables: []string{"zum", "kem", "ral"},
		},
	}
}


