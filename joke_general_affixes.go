package main

var (
	generalAffixes = []affix{
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "без особых примет",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "на гусеничном ходу",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией самоуничтожения",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с запасом прочности",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с эффектом зомбирования",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "на паровой тяге",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "кустарного производства",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с кустарной модификацией",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "со встроенным джетпаком",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией \"развидеть\"",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "из туалетной бумаги",
			},
			listOfEffects: []string{"хрупкая"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "инопланетного происхождения",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "со следами укусов",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "со следами эктоплазмы",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "прямиком из ада",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "из взрывоопасных веществ",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "апокалипсиса",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "белорусского производства",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "с острым наконечником",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "для особых случаев",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "на кодовом замке",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "с жидкокристаллическим дисплеем",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "особой секретности",
			},
		},
	}

	generalUniqueAffixes = []affix {
		{
			lines: &affixLines{
				suffixForItemName: "с пылу с жару",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: " - то, что доктор прописал",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "тётушки Совы",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "ниоткуда с любовью",
			},
		},
		{
			lines: &affixLines{
				suffixForItemName: "особой грубости",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для занятий астрологией",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "со встроенным вибратором",
			},
			listOfEffects: []string{"главное - не отвлекаться"},
			modifiers: map[string][]int{
				"Защита":   {-1, -2, -3},
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "из кожи жопы дракона",
			},
			listOfEffects: []string{"непробиваемо"},
			modifiers: map[string][]int{
				"Защита": {2, 3, 4, 5},
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с огоньком",
			},
			listOfEffects: []string{"жжёт жопу"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с дырками",
			},
			listOfEffects: []string{"продувает"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "всратости",
			},
			listOfEffects: []string{"некрасивый"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией добычи огня",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с искусственным интеллектом",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией доставки еды",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с выходом в Интернет",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с шипами",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с торчащими гвоздями",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "ликантропии",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией просмотра сериальчиков",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией массажа",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для повышения самооценки",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для настоящих засранцев",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для пришельцев",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с эффектом лоботомии",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с эффектом регенерации",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с лазерной гравировкой",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с установленным DOOM",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "саморазрушения",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с функцией заваривания чая",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с подогревом",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с вентиляцией",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "от фирмы Тошиба",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "от фирмы \"Туалетный утёнок\"",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "\"Каждый день\"",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "от фирмы \"Роскосмос\"",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для хардкорных геймерш",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "Чака Норриса",
			},
			listOfEffects: []string{"бьёт с разворота"},
			modifiers: map[string][]int{
				"Damage": {1000, 1001, 1002},
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "мистера Жопосранчика",
			},
			listOfEffects: []string{"интересный"},
			modifiers: map[string][]int{
				"Damage":   {-2},
				"Accuracy": {-1},
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "со списком ругательств в адрес владельца",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с порнографическими картинками",
			},
			listOfEffects: []string{"боевой дух +25"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "с советами по игре в покер",
			},
			listOfEffects: []string{"ва-банк"},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для захвата мира",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "без ГМО",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для воздействия на разум",
			},
		},
		{
			lines: &affixLines{
				prefixForItemName: "",
				suffixForItemName: "для уничтожения человечества",
			},
		},
	}
)
