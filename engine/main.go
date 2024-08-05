package engine

type Character struct {
	Name   string
	Weapon Weapon
	Armor  Armor
	Stats  Stats
	Health
	DefenseStatus bool
}

type Enemy struct {
	Name  string
	Stats Stats
	Health
}

type Health struct {
	CurrentHP int
	MaxHP     int
	CurrentMP int
	MaxMP     int
}

type Weapon struct {
	Name        string
	Description string
	Strength    byte
	Accuracy    float64
}

type Armor struct {
	Name         string
	Description  string
	Defense      byte
	MagicDefense byte
}

type Accessory struct {
	Name        string
	Description string
}

func EffectChange(stat string) {
}

type Stats map[string]int

func initializeStatFields() Stats {
	return Stats{
		"BaseStrength":     0,
		"BaseDefense":      0,
		"BaseMagicDefense": 0,
		"Luck":             0,
		"Dexterity":        0,
	}
}

func NewCharacter(name string) Character {
	return Character{
		Name:  name,
		Stats: initializeStatFields(),
	}
}

func MapJsonToCharacters(characters []JsonCharacter) []Character {
	result := []Character{}
	for _, character := range characters {
		result = append(result, Character{
			Name: character.Name,
		})
	}
	return result
}

func MapJsonToEnemies(enemies []EnemyJson) []Enemy {
	result := []Enemy{}
	for _, enemy := range enemies {
		result = append(result, Enemy{
			Name: enemy.Name,
		})
	}

	return result
}
