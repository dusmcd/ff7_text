package engine

type JsonCharacter struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	MaxHP             int      `json:"maxHP"`
	CurrentHP         int      `json:"currentHP"`
	MaxMP             int      `json:"maxMP"`
	CurrentMP         int      `json:"currentMP"`
	WeaponName        string   `json:"weapon_name"`
	WeaponStrength    int      `json:"weapon_strength"`
	ArmorName         string   `json:"armor_name"`
	ArmorDefense      int      `json:"armor_defense"`
	ArmorMagicDefense int      `json:"armor_magic_defense"`
	BaseStrength      int      `json:"base_strength"`
	BaseDefense       int      `json:"base_defense"`
	BaseMagicDefense  int      `json:"base_magic_defense"`
	Luck              int      `json:"luck"`
	Dexterity         int      `json:"dexterity"`
	Materia           []string `json:"materia"`
}

type EnemyJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MaxHP       int    `json:"maxHP"`
	CurrentHP   int    `json:"currentHP"`
	MaxMP       int    `json:"maxMP"`
	CurrentMP   int    `json:"currentMP"`
}

type JsonData struct {
	Characters []JsonCharacter `json:"characters"`
	Enemies    []EnemyJson     `json:"enemies"`
}
