package game

type Difficulty struct {
	Name    string
	Chances int
	Level   int
}

var difficulties = map[int]Difficulty{
	1: {"Easy", 10, 1},
	2: {"Medium", 5, 2},
	3: {"Hard", 3, 3},
}

func GetDifficulty(choice int) (Difficulty, bool) {
	diff, exists := difficulties[choice]
	return diff, exists
}

func GetAvailableDifficulties() map[int]Difficulty {
	return difficulties
}
