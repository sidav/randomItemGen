package main

import (
	"fmt"
	"math/rand"
)

var (
	jokeSexes = []string{
		"мужской", "женский", "мужской", "женский", "неизвестен", "не поддаётся классификации", "[ЗАСЕКРЕЧЕНО]",
	}

	jokeClassPrefixes = []string {
		"супер", "недо", "контр", "недопере", "кибер", "робо", "био", "техно", "зомбо", "убер", "нано", "гига",
		"стелс-", "эрзац-",
		"великанский ", "тайный ", "незаменимый ", "крайне опытный ", "опытный ", "начинающий ",  "собачий ", "спящий ",
		"свежеиспечённый ", "тёмный ", "высший ", "вовсе не ", "безбашенный ", "неудержимый ", "неустрашимый ",
		"гнусный ", "подлый ", "поганый ", "вонючий ", "незаменимый ", "опасный ", "многократный ", "бессмертный ",
		"непревзойдённый ", "непобедимый ", "самый надёжный ", "всемирно известный ", "генномодифицированный ",
		"бесшумный ", "элитный ", "единственный в своём роде ", "высокооплачиваемый ", "безработный ", "упоротый ",
		"юный ", "судебный ", "карликовый ", "рукопашный ", "блатной ", "ссыкливый ", "невидимый ", "воображаемый ",
		"любопытный ", "переоценённый ", "недооценённый ", "дешёвый ", "пафосный ", "добрый ", "милосердный ",
		"главный ", "просветлённый", "павший ", "недоверчивый ", "параноидальный ", "непонятный ", "неизвестный ",
		"безмозглый ", "прокачанный ", "подкачанный ", "накачанный ", "вечно пьяный ", "забывчивый ", "рассеянный ",
		"камуфляжный ",
	}
	jokeClasses = []string {
		"ниндзя", "алкаш", "паладин", "дуралей", "шизофреник", "программист", "боксёр", "задрот", "ассассин",
		"терминатор", "громоотвод", "каратист", "эксгибиционист", "спецназовец", "турист", "археолог", "икскомовец",
		"пришелец", "призрак", "морпех", "синхрофазотрон", "словоблуд", "демагог", "вуайерист", "сабзиро", "самовар",
		"испытатель", "мизантроп", "филателист", "прокрастинатор", "ублюдок", "ассенизатор", "ведьмак", "шахматист",
		"лохотронщик", "фетишист", "психопат", "таксист", "пердун", "шантажист", "порноактёр", "импотент",
		"лупоглазик", "недоучка", "бизнесмен", "проститут", "лётчик", "гамадрил", "рогоносец", "козёл",
		"террорист", "энтузиаст", "геронтофил", "сутенёр", "отшельник", "дегустатор", "аристократ", "хакер",
		"киберспортсмен", "чемпион", "литрболист", "блевотрон", "пиздабол", "попрошайка", "пародист", "юрист",
		"рецидивист", "матершинник", "психопат", "потаскун", "комментатор", "спортсмен", "инвалид", "супергерой",
		"герой", "ментат", "бэтмен", "агент", "спецагент", "боец", "пловец", "жнец", "пахарь", "тракторист", "робот",
		"робокоп", "чупакабра", "перевозчик", "дальнобойщик", "дайвер", "кровопийца", "протосс", "командир",
		"человек-молекула", "пиромант", "пироманьяк", "криомант", "некромант", "оборотень", "вампир", "камикадзе",
	}
	jokeClassSuffixes = []string {
		"атор 3000",
		", человек и пароход", ", но это не точно", ", последняя надежда человечества", ", человек из стали",
		" с кучей бабла", " без бабла", " с венерическими болезнями", " который не лажает", " по вызову",
		" за копейки", ", готовый работать за бесплатно", ", готовый работать за еду", " без мозгов",
		" без страха и упрёка", " на пособии", " в отставке", " с боевым опытом", " с серьёзными намерениями",
		" с манией величия", " с синдромом самозванца", " со страхом газонокосилок", " с боязнью зеркал",
		" и победитель чемпионата Сыктывкара по шахматам в 2019г.", " с добрыми глазами", " с синдромом Дауна",
	}

	jokeInvWeaponUsages = []string{
		"Основное", "Запасное", "Не умеет пользоваться", "Носит, но боится использовать", "Постоянно проёбывает",
		"Прячет в труднодоступных местах", "Достаёт, когда дела совсем плохи", "Сам изготовил", "Украл в детском саду",
		"Продаёт", "Главный трофей", "Купил за бесценок", "Использует в качестве костыля", "Сломал", "Секретное оружие",
	}
	totalWeaponUsages = []int{1, 1, 1, 2, 2, 3, 4}

	jokeInvArmorUsages = []string{
		"Постоянно", "Для жары", "Для холодов", "Для выступлений", "Любимое", "Для тяжёлых перестрелок", "Для флирта",
		"Для конспирации", "Носит среди своих", "Хочет снять, но приклеился", "Сам сшил", "Украл в секонд-хенде",
		"Выиграл в карты", "Снимает в помещении", "Боится показать", "Ни за что не наденет",
	}
	totalArmorUsages = []int{1, 1, 2, 2, 3, 4}

	jokeInvBookUsages = []string{
		"Любимое", "Самое пригождающееся", "Выучено наизусть", "Всем рассказывает про", "Мечтает прочесть", "Фанат",
		"Всех задолбал разговорами о", "Стал экспертом по", "Прочёл в детстве", "Причина психической травмы в детстве",
		"Ненавидит", "Постоянно носит с собой, но не осилил", "Считает слишком сложной", "Сам написал",
		"Единственный, кто понял", "Единственное, что прочёл",
	}
	totalBookUsages = []int{1, 1, 2, 2, 2, 3, 4}

	jokeInvDusposableUsages = []string{
		"Постоянно использует", "Носит, но не умеет пользоваться", "Никак не может избавиться", "Обожает",
		"Всё отдаст за", "Всё забывает вытащить", "Забывает в других штанах", "Больше всего нуждается", "Умрёт без",
		"Зависим от", "Продаёт", "Курит", "Ест", "Нуждается по медицинским показаниям",
		"Использует не по прямому назначению", "Не помнит, где нашёл",
	}
	totalDisposablesUsages = []int{1, 1, 2, 2, 2, 3, 3, 4}
)

func stringArrHasValue(strArr *[]string, value string) bool {
	for _, s := range *strArr {
		if s == value {
			return true
		}
	}
	return false
}

func appendRandomNonUsedValueFromTo(from *[]string, to *[]string) {
	i := rand.Intn(len(*from))
	for stringArrHasValue(to, (*from)[i]) {
		i = rand.Intn(len(*from))
	}
	*to = append((*to), (*from)[i])
}

func getRandomFromIntArray(ia *[]int) int {
	return (*ia)[rand.Intn(len(*ia))]
}
func getRandomFromStringArray(ia *[]string) string {
	return (*ia)[rand.Intn(len(*ia))]
}

func createJokeCraracter() {
	fmt.Printf("===== СУБЪЕКТ №%d =====\n", rand.Intn(8888888) + 1111111)
	fmt.Println("ИМЯ: [ЗАСЕКРЕЧЕНО]")
	fmt.Printf("ВОЗРАСТ: %d\n", 16 + rand.Intn(15) + rand.Intn(15))
	fmt.Printf("ПОЛ: %s\n", getRandomFromStringArray(&jokeSexes))

	class := getRandomFromStringArray(&jokeClasses)
	if rand.Intn(100) < 45 {
		class = getRandomFromStringArray(&jokeClassPrefixes) + class
	}
	if rand.Intn(100) < 45 {
		class = getRandomFromStringArray(&jokeClassPrefixes) + class
	}
	if rand.Intn(100) < 50 {
		class += "-"+getRandomFromStringArray(&jokeClasses)
	}
	if rand.Intn(100) < 40 {
		class += getRandomFromStringArray(&jokeClassSuffixes)
	}
	fmt.Printf("СПЕЦИАЛИЗАЦИЯ: %s\n", class)

	fmt.Println("=== ОРУЖИЕ: ===")
	uses := []string{}
	totalUses := getRandomFromIntArray(&totalWeaponUsages)
	for i := 0; i < totalUses; i++ {
		appendRandomNonUsedValueFromTo(&jokeInvWeaponUsages, &uses)
	}
	for _, use := range uses {
		fmt.Println(use + ": " + gen.createItemByRule(weaponRuleset.getWeightedRandomRule()).getFullName(ADD_STATS_LINE, true, true))
	}

	fmt.Println("=== ОДЕЖДА: ===")
	uses = []string{}
	totalUses = getRandomFromIntArray(&totalArmorUsages)
	for i := 0; i < totalUses; i++ {
		appendRandomNonUsedValueFromTo(&jokeInvArmorUsages, &uses)
	}
	for _, use := range uses {
		fmt.Println(use + ": " + gen.createItemByRule(armorRuleset.getWeightedRandomRule()).getFullName(ADD_STATS_LINE, true, true))
	}

	fmt.Println("=== ЧТИВО: ===")
	uses = []string{}
	totalUses = getRandomFromIntArray(&totalBookUsages)
	for i := 0; i < totalUses; i++ {
		appendRandomNonUsedValueFromTo(&jokeInvBookUsages, &uses)
	}
	for _, use := range uses {
		fmt.Println(use + ": " + gen.createItemByRule(bookRuleset.getWeightedRandomRule()).getFullName(ADD_STATS_LINE, true, true))
	}

	fmt.Println("=== РЕСУРСЫ: ===")
	uses = []string{}
	totalUses = getRandomFromIntArray(&totalDisposablesUsages)
	for i := 0; i < totalUses; i++ {
		appendRandomNonUsedValueFromTo(&jokeInvDusposableUsages, &uses)
	}
	for _, use := range uses {
		fmt.Println(use + ": " + gen.createItemByRule(disposablesRuleset.getWeightedRandomRule()).getFullName(ADD_STATS_LINE, true, true))
	}
}
