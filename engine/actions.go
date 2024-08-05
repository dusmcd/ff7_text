package engine

type BattleMember[T any] interface {
	Attack(T)
	TakeDamage(int)
}

func (char Character) Attack(enemy *Enemy) {
	//damage := char.Stats["BaseStrength"] + int(char.Weapon.Strength) + 100
	damage := 100
	enemy.TakeDamage(damage)
}

func (char *Character) TakeDamage(amount int) {
	if char.DefenseStatus {
		amount /= 2
	}
	char.CurrentHP -= amount
}

func (enemy *Enemy) TakeDamage(amount int) {
	enemy.CurrentHP -= amount
}

func (enemy Enemy) Attack(char *Character) {
	//damage := enemy.Stats["BaseStrength"] + 100
	damage := 10
	char.TakeDamage(damage)
}

func (char Character) CastMagic(enemy *Enemy) {
	// check equipped materia
}

func (char *Character) Defend() {
	char.DefenseStatus = true
}
