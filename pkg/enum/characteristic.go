package enum

type Characteristic int

const (
	Attack Characteristic = iota
	Defense
	Health
	Evasion
	Targetability
	Cariability
	CanFly
	CanCarryUnits
	CanCarryResources
	TargetOrder
	TargetCount
	Volume
	Capacity
)

var CharacteristicNames = []string{"attack", "defense", "health", "evasion", "targetability", "cariability", "canFly", "canCarryUnits", "canCarryResources", "targetOrder", "targetCount", "volume", "capacity"}

var CharacteristicMap = map[string]Characteristic{
	CharacteristicNames[0]:  Attack,
	CharacteristicNames[1]:  Defense,
	CharacteristicNames[2]:  Health,
	CharacteristicNames[3]:  Evasion,
	CharacteristicNames[4]:  Targetability,
	CharacteristicNames[5]:  Cariability,
	CharacteristicNames[6]:  CanFly,
	CharacteristicNames[7]:  CanCarryUnits,
	CharacteristicNames[8]:  CanCarryResources,
	CharacteristicNames[9]:  TargetOrder,
	CharacteristicNames[10]: TargetCount,
	CharacteristicNames[11]: Volume,
	CharacteristicNames[12]: Capacity,
}

func (c Characteristic) String() string {
	return CharacteristicNames[c]
}
