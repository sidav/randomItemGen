package main

var (
	disposablesRuleset = ruleset{
		possibleNames: []string{
			"сигареты", "жвачки", "шнурки", "бинты", "медикаменты", "презервативы", "деньги", "фантики", "сосиски",
			"глазные капли", "ректальные свечи", "гирлянды", "жетоны для метро", "таблетки", "влажные салфетки",
			"зубочистки",
		},

		possbleAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "использованные",
					suffixForItemName: "б/у",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "подозрительные",
					suffixForItemName: "странной формы",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "невкусные",
					suffixForItemName: "с мерзким послевкусием",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "уникальные",
					suffixForItemName: "без аналогов",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "тактические",
					suffixForItemName: "в камуфляжной расцветке",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "свеженькие",
					suffixForItemName: "без следов использования",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "грязные",
					suffixForItemName: "в грязи",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "вонючие",
					suffixForItemName: "",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "полезные для здоровья",
					suffixForItemName: "исцеления",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "инопланетные",
					suffixForItemName: "с приветом землянам",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "быстрого приготовления",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "съедобные",
					suffixForItemName: "с кофеином",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "многоразовые",
					suffixForItemName: "",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "дорогие",
					suffixForItemName: "за кучу денег",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "собачьи",
					suffixForItemName: "для собак",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кошачьи",
					suffixForItemName: "для кошек",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "непромокаемые",
					suffixForItemName: "с эффектом влагостойкости",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "заплесневелые",
					suffixForItemName: "с плесенью",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "одушевлённые",
					suffixForItemName: "которые просто хотят жить",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разумные",
					suffixForItemName: "которые желают вам зла",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "непонятные",
					suffixForItemName: "но это не точно",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "странно притягательные",
					suffixForItemName: "которые нельзя забыть",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "незаконные",
					suffixForItemName: "которые нельзя развидеть",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "женские",
					suffixForItemName: "женственности",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мужские",
					suffixForItemName: "для настоящих мужчин",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "известные",
					suffixForItemName: "из рекламы",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "брендовые",
					suffixForItemName: "как в рекламе",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "светящиеся в темноте",
					suffixForItemName: "ослепительного света",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "скользкие",
					suffixForItemName: "смазанные солидолом",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разноцветные",
					suffixForItemName: "разных цветов",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "воображаемые",
					suffixForItemName: "которых не существует",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "убогие",
					suffixForItemName: "жуткой формы",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "скрюченные",
					suffixForItemName: "невозможной геометрии",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "высокотехнологичные",
					suffixForItemName: "с наночастицами",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "повышающие интеллект",
					suffixForItemName: "для повышения мозговой активности",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "клеёнчатые",
					suffixForItemName: "с крылышками",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "неудобные",
					suffixForItemName: "размера ХХХХХL",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ватные",
					suffixForItemName: "из ваты",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "бюджетные",
					suffixForItemName: "из фикспрайса",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "бессмертные",
					suffixForItemName: "от бессонницы",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "неразрушимые",
					suffixForItemName: "бракованные",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "космические",
					suffixForItemName: "для космонавтов",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "успокаивающие",
					suffixForItemName: "от тревоги и стресса",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "безумные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "музыкальные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "небанальные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "неправильные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "передовые",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "технологичные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нехуёвые",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "дааддалаовые",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "недоеденные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "сказочные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "радостные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "безэмоциональные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "тотальные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "оригинальные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кошерные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "шикарные",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "зашкварные",
				},
			},
		},

		possibleUniqueAffixes: []affix{
			{
				lines: &affixLines{
					suffixForItemName: "требующие наждачки",
				},
			},
			{
				lines: &affixLines{
					suffixForItemName: "страшные до усрачки",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "исполненные очей",
					suffixForItemName: "чей так светел взор незабываемый",
					selectOnlyOne:     true,
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "наследственные",
					suffixForItemName: "",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "бесконечные",
					suffixForItemName: "",
				},
			},
		},
	}
)

func initDisposablesRuleset() {
	disposablesRuleset.rules = []*rule{
		&rule{
			rarityName:        "Common",
			ruleWeight:        5,
			possibleItemNames: &disposablesRuleset.possibleNames,
			possibleAffixes:   nil,
			totalAffixes:      0,
		},
		&rule{
			rarityName:        "Uncommon",
			ruleWeight:        5,
			possibleItemNames: &disposablesRuleset.possibleNames,
			possibleAffixes:   append(disposablesRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:      1,
		},
		&rule{
			rarityName:        "Rare",
			ruleWeight:        4,
			possibleItemNames: &disposablesRuleset.possibleNames,
			possibleAffixes:   append(disposablesRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:      2,
		},
		&rule{
			rarityName:            "Epic",
			ruleWeight:            3,
			possibleItemNames:     &disposablesRuleset.possibleNames,
			possibleAffixes:       append(append(disposablesRuleset.possbleAffixes, generalAffixes...), generalAffixes...),
			totalAffixes:          2,
			possibleUniqueAffixes: append(disposablesRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    1,
		},
		&rule{
			rarityName:            "LEGENDARY",
			ruleWeight:            2,
			possibleItemNames:     &disposablesRuleset.possibleNames,
			possibleAffixes:       append(disposablesRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          3,
			possibleUniqueAffixes: append(disposablesRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    2,
		},
		&rule{
			rarityName:            "~ ETERNAL ~",
			ruleWeight:            1,
			possibleItemNames:     &disposablesRuleset.possibleNames,
			possibleAffixes:       append(disposablesRuleset.possbleAffixes, generalAffixes...),
			totalAffixes:          4,
			possibleUniqueAffixes: append(disposablesRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
			totalUniqueAffixes:    3,
		},
	}
}
