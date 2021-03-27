package main

var (
	bookRuleset = ruleset{
		possibleNames: []string{
			"книга", "бумажка", "методичка", "инструкция", "рекламка", "брошюра", "газета", "тетрадь", "медкарта",
			"камасутра", "записка", "скрижаль", "живопись", "вывеска", "социальная реклама",
			"папка \"секретных материалов\"", "шпаргалка", "таблица", "объяснительная",
		},

		possbleAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "скучная",
					suffixForItemName: "",
				},
				listOfEffects: []string{"сон"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пустая",
					suffixForItemName: "",
				},
				listOfEffects: []string{"можно что-нибудь написать"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "написанная карандашом",
					suffixForItemName: "",
				},
				listOfEffects: []string{""},
			},
			{
				lines: &affixLines{
					prefixForItemName: "некрономическая",
					suffixForItemName: "демонического происхождения",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"сводит с ума"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "школьная",
					suffixForItemName: "для младших классов",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"написано с ошибками"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "медицинская",
					suffixForItemName: "с советами по медицине",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"самолечение +1"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "иллюстрированная",
					suffixForItemName: "в комиксах",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"интереснее обычной"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "иероглифическая",
					suffixForItemName: "на китайском",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"может, это японский?.."},
			},
			{
				lines: &affixLines{
					prefixForItemName: "испачканная",
					suffixForItemName: "в грязище",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"местами не разобрать"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "библиотечная",
					suffixForItemName: "украденная из библиотеки",
					selectOnlyOne:     true,
				},
				listOfEffects: []string{"вернуть до послезавтра"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "матерная",
					suffixForItemName: "с матерными частушками",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "сумасбродная",
					suffixForItemName: "пациента из психушки",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "безумно интересная",
					suffixForItemName: "с потрясающим сюжетом",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "детская",
					suffixForItemName: "для самых маленьких",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "потрёпанная",
					suffixForItemName: "с вырванными страницами",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "предсказуемая",
					suffixForItemName: "со спойлерами",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "написанная Брайлем",
					suffixForItemName: "для слепых",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "смешная",
					suffixForItemName: "в комиксах",
					selectOnlyOne:     true,
				},
			},
		},

		possibleUniqueAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по матанализу",
				},
				listOfEffects: []string{},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по демонологии",
				},
				listOfEffects: []string{"и чем это отличается от матана?"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "про программирование на php",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "про карательную кулинарию",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по зельеварению",
				},
				listOfEffects: []string{"сварим что угодно, Джесси"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по взлому жопы",
				},
				listOfEffects: []string{"1337"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по магии цветочков",
				},
				listOfEffects: []string{"боевая - как ни странно"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "полная брехни",
				},
				listOfEffects: []string{"и кому это надо?.."},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по дифференциальным уравнениям",
				},
				listOfEffects: []string{"убейте меня"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "по физкультуре",
				},
				listOfEffects: []string{"левой-правой"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "с читами для GTA",
				},
				listOfEffects: []string{"мне бы её лет десять назад..."},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "про Гарри Поттера",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "которую можно курить",
				},
			},
		},
	}
)

func initBookRuleset() {
	bookRuleset.rules = []*rule{
		&rule{
			rarityName:        "Common",
			ruleWeight:        5,
			possibleItemNames: &bookRuleset.possibleNames,
			possibleAffixes:   nil,
			totalAffixes:      0,
		},
		&rule{
			rarityName:        "Uncommon",
			ruleWeight:        4,
			possibleItemNames: &bookRuleset.possibleNames,
			possibleAffixes:   append(append(bookRuleset.possbleAffixes, generalAffixes...), generalAffixes...),
			totalAffixes:      1,
		},
		&rule{
			rarityName:        "Rare",
			ruleWeight:        3,
			possibleItemNames: &bookRuleset.possibleNames,
			possibleAffixes:   append(bookRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:      2,
		},
		&rule{
			rarityName:            "Epic",
			ruleWeight:            2,
			possibleItemNames:     &bookRuleset.possibleNames,
			possibleAffixes:       append(bookRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          2,
			possibleUniqueAffixes: append(bookRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    1,
		},
		&rule{
			rarityName:            "LEGENDARY",
			ruleWeight:            1,
			possibleItemNames:     &bookRuleset.possibleNames,
			possibleAffixes:       append(bookRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          3,
			possibleUniqueAffixes: append(bookRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    2,
			aliasPreSyllables:     []string{"Qu", "Li", "Non-"},
			aliasSyllables:        []string{"kir", "qua", "li"},
			aliasPostSyllables:    []string{"zum", "kem", "ral"},
		},
	}
}
