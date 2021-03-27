package main

var (
	armorRuleset = ruleset{
		possibleNames: []string{
			"бригантина",
			"пара штанов",
			"накидка",
			"зимняя куртка",
			"пара перчаток",
			// дальше неадекватные
			"бронежилетка",
			"сисечная броня",
			"пара трусов",
			"пара носков",
			"шляпка",
			"пара берцев",
		},

		possbleAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "шерстяная",
					suffixForItemName: "",
				},
				listOfEffects: []string{"колется"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "синтетическая",
					suffixForItemName: "из синтетики",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"неприятно"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "невидимая",
					suffixForItemName: "",
				},
				listOfEffects: []string{"невидима только одежда"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "слишком длинная",
					suffixForItemName: "",
				},
				listOfEffects: []string{"споткнуться можно"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "гламурная",
					suffixForItemName: "",
				},
				listOfEffects: []string{"зато со стразиками"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кожаная",
					suffixForItemName: "в стиле БДСМ",
					selectOnlyOne: true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "узкая",
					suffixForItemName: "на два размера меньше",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"для похудения"},
			},
		},

		possibleUniqueAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "великого воина",
				},
				listOfEffects: []string{},
				modifiers: map[string][]int{
					"Защита":   {3, 6, 10, 15},
				},
			},
			// дальше неадекватные
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "из свинцовой ткани",
				},
				listOfEffects: []string{"защищает от радиации"},
				modifiers: map[string][]int{
					"Защита":   {-2, -1, 0, 1},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "говорящая",
					suffixForItemName: "по имени Добрыня Святополкович Зильберштейн",
				},
				listOfEffects: []string{"\"да когда оно уже заткнётся?..\""},
				modifiers: map[string][]int{
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "паршивая",
					suffixForItemName: "",
				},
				listOfEffects: []string{"\"уёбищная, да и нитки вылезают\""},
				modifiers: map[string][]int{
					"Защита":   {-11, -14, -18},
				},
			},
		},
	}
)

func initArmorRuleset() {
	armorRuleset.rules = []*rule {
		&rule{
			rarityName:        "Common",
			possibleItemNames: &armorRuleset.possibleNames,
			possibleAffixes:   nil,
			totalAffixes:      0,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Защита":   {-2, -1, 0, 0, 1, 2},
			},
		},
		&rule{
			rarityName:        "Uncommon",
			possibleItemNames: &armorRuleset.possibleNames,
			possibleAffixes:   append(append(armorRuleset.possbleAffixes, generalAffixes...), generalAffixes...),
			totalAffixes:      1,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Защита":   {-2, -1, 0, 0, 1, 2, 3},
			},
		},
		&rule{
			rarityName:        "Rare",
			possibleItemNames: &armorRuleset.possibleNames,
			possibleAffixes:   append(armorRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:      2,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Защита":   {-2, -1, -1, 0, 1, 1, 2, 3},
			},
		},
		&rule{
			rarityName:            "Epic",
			possibleItemNames:     &armorRuleset.possibleNames,
			possibleAffixes:       append(armorRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          2,
			possibleUniqueAffixes: append(armorRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    1,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Защита":   {-3, -2, -1, 0, 1, 2, 3, 4, 5},
			},
		},
		&rule{
			rarityName:            "LEGENDARY",
			possibleItemNames:     &armorRuleset.possibleNames,
			possibleAffixes:       append(armorRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          3,
			possibleUniqueAffixes: append(armorRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    2,
			modifiersNamesAndPossibleValues: map[string][]int{
				"Защита":   {-3, -2, -1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5},
			},
			uniqueNamePrefixes:  []string{"пре", "пере", "до", "недо", "блево", "броне", "тряп", "брен"},
			uniqueNameSyllables: []string{"ли", "ка", "ша", "шаш", "зюн", "бав", "брю", "хыыых"},
			uniqueNameSuffixes:  []string{"алка", "ень", "ый", "-кладенец"},
		},
	}
}

