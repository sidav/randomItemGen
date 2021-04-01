package main

var (
	weaponRuleset = ruleset{
		possibleNames: []string{
			"меч", "молот", "посох", "лук", "кинжал", "серп", "резак", "резольвиндер", "арбалет", "катар", "фламберг",
			// дальше неадекватные
			"носовой платок", "карандаш", "дробовик", "пузырёк зелёнки", "подозритель", "базука", "плазмотрон", "котик",
			"огурец", "пердимонокль", "пердёж", "окурок", "миниган", "спичкострел", "гвоздодёр", "танк",
			"огрызок яблока", "огнемёт", "водомёт", "кухонный нож", "боевой робот", "туалетный утёнок", "игольник",
			"шприц", "гранатомёт", "психо-усилитель", "индуктор", "экзоскелет", "имплантат",
		},

		possbleAffixes: []affix{
			{
				lines: &affixLines{
					prefixForItemName: "вампирский",
					suffixForItemName: "вампиризма",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"вампиризм"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "самый обычный",
					suffixForItemName: "",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пламенный",
					suffixForItemName: "",
				},
				listOfEffects: []string{"огонь"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ледяной",
					suffixForItemName: "",
				},
				listOfEffects: []string{"мороз"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "ревнителя",
				},
				listOfEffects: []string{"святой"},
			},
			// неадекватные
			{
				lines: &affixLines{
					prefixForItemName: "резиновый",
					suffixForItemName: "из резины",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"пружинит"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "всратый",
					suffixForItemName: "",
				},
				listOfEffects: []string{"некрасивый"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "ближнего боя",
				},
				listOfEffects: []string{"дальность 1м"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "на солярке",
				},
				listOfEffects: []string{"без заправки не работает"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "с бесконечным боезапасом",
				},
				listOfEffects: []string{"больше патронов!"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "недоделанный",
					suffixForItemName: "который скоро сломается",
				},
				listOfEffects: []string{"ой"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "некачественный",
					suffixForItemName: "",
				},
				listOfEffects: []string{"верните деньги"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "бесплатный",
					suffixForItemName: "зато бесплатный",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"скоро сломается"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "самонаводящийся",
					suffixForItemName: "самонаведения",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"с системой распознавания цели"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "опустошающий",
					suffixForItemName: "опустошения",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"накладывает опустошение"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "свинцовый",
					suffixForItemName: "целиком вылитый из свинца",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"почти защищает от радиации"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "картографирующий",
					suffixForItemName: "для составления карт",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"заполняет карту"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "встроенный",
					suffixForItemName: "мультитул",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"выполняет несколько функций сразу, но фигово"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "терминаторский",
					suffixForItemName: "терминатора",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"I'll be back"},
			},
			{
				lines: &affixLines{
					prefixForItemName: "вместительный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "скорострельный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "точный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "проникающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "прочный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "автозаряжающийся",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "лечащий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ремонтирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "самоуничтожающийся",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "светящийся",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "радиоактивный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "помповый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "обморочный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нокаутирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "усыпляющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "опьяняющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "галлюциногенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "канцерогенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "травмирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кровопускающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "отравленный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "горящий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "химический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "базовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разгонный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "наэлектризованный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "полированный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "заземлённый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разжижающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "термический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "анионный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "лечебный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "старый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "древний",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "магический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "сумрачный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "техномагический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "искристый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "технологичный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нестабильный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мясной",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "поглощающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "спектральный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "усиливающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ослабляющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нормирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "солнечный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "отражающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "преломляющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "возвращающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "наводящийся",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "оптимизированный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "калиброванный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "расщепляющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "стабильный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пехотный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "классический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "диагностический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "симметричный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "адекватный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "синергетический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "стратегический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "тактический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "добрый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "злой",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "уродливый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "хрустящий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "треснутый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "хлипкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "твёрдый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "жидкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "густой",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "плазменный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "усреднённый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "элегантный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "неуклюжий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "грубый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "солнцезащитный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "миниатюрный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "посредственный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "хрупкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "прочный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "потрёпанный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "потёртый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "гибкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "изворотливый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ломкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "хитиновый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "костяной",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "несокрушимый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "панцирный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "расширенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "исполинский",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "рваный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "вздувшийся",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "плотный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "несгибаемый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "щелочной",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нейтрализующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "химический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "асбестовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "перевязанный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "патентованный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "заклеенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "гладкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "шершавый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "глянцевый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "матовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "морозный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мотивирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "комплексный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "комплектный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "инертный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "давящий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пылкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пассивный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "активный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "свинцовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разгонный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "быстродействующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нагнетающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "термальный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "эндотермический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "закалённый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "инвазивный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "рецессивный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "модернизированный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "улучшенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "экстренный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "расторопный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "громоздкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "компактный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ускоряющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "выжимающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мобильный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "изолированный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "красочный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "терапевтический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "заградительный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "разрушительный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "опустошающий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "отборный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "вибрационный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "рециркулирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "картографирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "визуализирющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "вспышкоподавляющий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "инфракрасный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ультрафиолетовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "квантовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кохлеарный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мембранный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "садохистический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "синоптический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "социальный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "церебральный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пищеварительный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "встроенный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "агрессивный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "адреналиновый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "электромагнитный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мономолекулярный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "механический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "амортизирующий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пахнущий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "маскировочный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "интегрированный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "гравитационный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кожаный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "мышечный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "нервный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "деревянный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "бумажный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "текстильный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "антисептический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "масляный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "биоразлагаемый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "асбестовый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "алюминиевый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "железный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "стальной",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "керамический",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "медный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "серебряный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "золотой",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "стеклянный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "клейкий",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "углеродный",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "пластиковый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "резиновый",
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "силиконовый",
				},
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
					"Damage":   {3, 6, 10, 15},
					"Accuracy": {2, 4, 6, 8},
				},
			},
			// дальше неадекватные
			{
				lines: &affixLines{
					prefixForItemName: "",
					suffixForItemName: "точный что аж пиздец",
				},
				listOfEffects: []string{"имба"},
				modifiers: map[string][]int{
					"Accuracy": {1000},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "кристально чистый",
					suffixForItemName: "из рекламы",
					selectOnlyOne: true,
				},
				listOfEffects: []string{"\"о, я видел это по телеку\""},
				modifiers: map[string][]int{
					"Accuracy": {2, 5, 7},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "ФЕНОМЕНАЛЬНО всратый",
					suffixForItemName: "зато с хорошим уроном",
				},
				listOfEffects: []string{"кто вообще сделал эту срань"},
				modifiers: map[string][]int{
					"Charisma": {-5},
					"Damage":   {5, 6, 7, 8},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "говорящий",
					suffixForItemName: "по имени Брумгильда",
				},
				listOfEffects: []string{"\"да когда оно уже заткнётся?..\""},
				modifiers: map[string][]int{
					"Accuracy": {2, 5, 7},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "очень стрёмный",
					suffixForItemName: "с щупальцами",
					selectOnlyOne: false,
				},
				listOfEffects: []string{"\"буэээээ\""},
				modifiers: map[string][]int{
					"Damage": {-4, 1, 6},
					"Accuracy": {2, 5, 7},
					"Количество щупалец": {1, 2, 3, 4, 5, 6, 7, 8},
				},
			},
			{
				lines: &affixLines{
					prefixForItemName: "самый хуёвый в истории человечества",
					suffixForItemName: "",
				},
				listOfEffects: []string{"\"м-да.\""},
				modifiers: map[string][]int{
					"Damage":   {-11, -14, -18},
					"Accuracy": {-10, -15, -17},
				},
			},
		},
	}
)

func initWeaponRuleset() {
	weaponRuleset.rules = []*rule {
			&rule{
				rarityName:        "Common",
				ruleWeight: 3,
				possibleItemNames: &weaponRuleset.possibleNames,
				possibleAffixes:   nil,
				totalAffixes:      0,
				modifiersNamesAndPossibleValues: map[string][]int{
					"Damage":   {-2, -1, 0, 0, 1, 2},
					"Accuracy": {-2, -1, 0, 0, 1, 2},
				},
			},
			&rule{
				rarityName:        "Uncommon",
				ruleWeight: 4,
				possibleItemNames: &weaponRuleset.possibleNames,
				possibleAffixes:   append(weaponRuleset.possbleAffixes, generalAffixes...),
				totalAffixes:      1,
				modifiersNamesAndPossibleValues: map[string][]int{
					"Damage":   {-2, -1, 0, 0, 1, 2, 3},
					"Accuracy": {-2, -1, 0, 0, 1, 2, 3},
				},
			},
			&rule{
				rarityName:        "Rare",
				ruleWeight: 3,
				possibleItemNames: &weaponRuleset.possibleNames,
				possibleAffixes:   append(weaponRuleset.possbleAffixes, generalAffixes...),
				totalAffixes:      2,
				modifiersNamesAndPossibleValues: map[string][]int{
					"Damage":   {-2, -1, -1, 0, 1, 1, 2, 3},
					"Accuracy": {-2, -1, 0, -1, 0, 1, 1, 2, 3},
				},
			},
			&rule{
				rarityName:            "Epic",
				ruleWeight: 2,
				possibleItemNames:     &weaponRuleset.possibleNames,
				possibleAffixes:       append(append(weaponRuleset.possbleAffixes, generalAffixes...), generalAffixes...),
				totalAffixes:          2,
				possibleUniqueAffixes: append(weaponRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
				totalUniqueAffixes:    1,
				modifiersNamesAndPossibleValues: map[string][]int{
					"Damage":   {-3, -2, -1, 0, 1, 2, 3, 4, 5},
					"Accuracy": {-3, -2, -1, 0, 1, 2, 3, 4, 5},
				},
			},
			&rule{
				rarityName:            "LEGENDARY",
				ruleWeight: 1,
				possibleItemNames:     &weaponRuleset.possibleNames,
				possibleAffixes:       append(weaponRuleset.possbleAffixes, generalAffixes...),
				totalAffixes:          3,
				possibleUniqueAffixes: append(weaponRuleset.possibleUniqueAffixes, generalUniqueAffixes...),
				totalUniqueAffixes:    2,
				modifiersNamesAndPossibleValues: map[string][]int{
					"Damage":   {-3, -2, -1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5},
					"Accuracy": {-3, -2, -1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5},
				},
				aliasPreSyllables:  []string{"пре", "пере", "до", "недо", "блево", "бюбюка", "псевдо", "супер"},
				aliasSyllables:     []string{"ли", "ка", "ша", "шаш", "зюн", "бав", "брю", "хым", "дроб", "бах",
					"рас", "пыщь", "тель", "зам", "брым", "кольт", "шот", "ган", "стрел"},
				aliasPostSyllables: []string{"алка", "ень", "ый", "-кладенец", "бабах", " бабахович", "тель"},
			},
		}
}
