package gopractising


/*
Your function takes two arguments:

current father's age (years)
current age of his son (years)
Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
*/

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	var age = 0
	if dadYearsOld < sonYearsOld * 2 {
		for {
			dadYearsOld++
			age++
			if sonYearsOld * 2 == dadYearsOld {
				break
			}
		}
	} else if dadYearsOld > sonYearsOld * 2 {
		age = dadYearsOld - sonYearsOld * 2

	}
	return age
}