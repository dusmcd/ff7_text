package engine

type Character struct {
	Name        string
	Description string
	Weapon      Weapon
	Armor       Armor
	Stats       Stats
	Health
	DefenseStatus bool
	Materias      []string
	Accessory     Accessory
}

type Materia struct {
	Name       string
	Spells     []string
	BaseDamage byte
}

type Enemy struct {
	Name        string
	Description string
	Stats       Stats
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
			Name:        character.Name,
			Description: character.Description,
			Health: Health{
				CurrentHP: character.CurrentHP,
				MaxHP:     character.MaxHP,
				CurrentMP: character.CurrentMP,
				MaxMP:     character.MaxMP,
			},
			Armor: Armor{
				Name:         character.ArmorName,
				Defense:      byte(character.ArmorDefense),
				MagicDefense: byte(character.ArmorMagicDefense),
			},
			Weapon: Weapon{
				Name:     character.WeaponName,
				Strength: byte(character.WeaponStrength),
			},
			Stats: Stats{
				"Luck":             character.Luck,
				"Dexterity":        character.Dexterity,
				"BaseStrength":     character.BaseStrength,
				"BaseDefense":      character.BaseDefense,
				"BaseMagicDefense": character.BaseMagicDefense,
			},
			Materias: character.Materia,
		})
	}
	return result
}

func MapJsonToEnemies(enemies []EnemyJson) []Enemy {
	result := []Enemy{}
	for _, enemy := range enemies {
		result = append(result, Enemy{
			Name:        enemy.Name,
			Description: enemy.Description,
			Health: Health{
				CurrentHP: enemy.CurrentHP,
				CurrentMP: enemy.CurrentMP,
				MaxHP:     enemy.MaxHP,
				MaxMP:     enemy.MaxMP,
			},
		})
	}

	return result
}
