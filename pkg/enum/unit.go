package enum

type Unit uint64

const (
	CargoPlane Unit = iota
	Transporter
	FighterJet
	Bomber
	UAV
	AttackHelicopter
	Pikeman
	Knight
	Samurai
	Cavalry
	Undead
	Ghoul
	Archer
	Rifleman
	Sniper
	Wizard
	WitchHunter
	Medic
	Engineer
	Necromancer
	Vampire
)

var UnitNames = []string{"cargoPlane", "transporter", "fighterJet", "bomber", "uav", "attackHelicopter", "pikeman", "knight", "samurai", "cavalry", "undead", "ghoul", "archer", "rifleman", "sniper", "wizard", "witchHunter", "medic", "engineer", "necromancer", "vampire"}

var UnitMap = map[string]Unit{
	UnitNames[0]:  CargoPlane,
	UnitNames[1]:  Transporter,
	UnitNames[2]:  FighterJet,
	UnitNames[3]:  Bomber,
	UnitNames[4]:  UAV,
	UnitNames[5]:  AttackHelicopter,
	UnitNames[6]:  Pikeman,
	UnitNames[7]:  Knight,
	UnitNames[8]:  Samurai,
	UnitNames[9]:  Cavalry,
	UnitNames[10]: Undead,
	UnitNames[11]: Ghoul,
	UnitNames[12]: Archer,
	UnitNames[13]: Rifleman,
	UnitNames[14]: Sniper,
	UnitNames[15]: Wizard,
	UnitNames[16]: WitchHunter,
	UnitNames[17]: Medic,
	UnitNames[18]: Engineer,
	UnitNames[19]: Necromancer,
	UnitNames[20]: Vampire,
}

var UnitToValue = map[Unit]string{
	CargoPlane:       UnitNames[0],
	Transporter:      UnitNames[1],
	FighterJet:       UnitNames[2],
	Bomber:           UnitNames[3],
	UAV:              UnitNames[4],
	AttackHelicopter: UnitNames[5],
	Pikeman:          UnitNames[6],
	Knight:           UnitNames[7],
	Samurai:          UnitNames[8],
	Cavalry:          UnitNames[9],
	Undead:           UnitNames[10],
	Ghoul:            UnitNames[11],
	Archer:           UnitNames[12],
	Rifleman:         UnitNames[13],
	Sniper:           UnitNames[14],
	Wizard:           UnitNames[15],
	WitchHunter:      UnitNames[16],
	Medic:            UnitNames[17],
	Engineer:         UnitNames[18],
	Necromancer:      UnitNames[19],
	Vampire:          UnitNames[20],
}

func (u Unit) String() string {
	return UnitNames[u]
}
