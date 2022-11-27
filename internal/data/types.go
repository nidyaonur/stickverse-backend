package data

import (
	"fmt"

	"github.com/nidyaonur/stickverse-backend/entities"
)

func (d *data) initResources() {
	entries := []*entities.Resource{
		{
			Name: "coal",
			NameLocal: entities.JSONMap{
				"en": "coal",
				"tr": "kömür",
			},
			Description: entities.JSONMap{
				"en": "It is obtained from nibs and lead pens. Basically, the coal is the lifeblood of the stick figures.",
				"tr": "Uç ve kurşun kalemlerden elde edilir. Temel olarak kömür, çöp adamların can damarıdır.",
			},
		},
		{
			Name: "paper",
			NameLocal: entities.JSONMap{
				"en": "paper",
				"tr": "kağıt",
			},
			Description: entities.JSONMap{
				"en": "Paper is the basic construction material. Since the coal alone is not sufficient to build large structures, it is used as the skeleton of the coal structures or only in paper structures.",
				"tr": "Kağıt, temel yapı malzemesidir. Kömür tek başına büyük yapılar inşa etmek için yeterli olmadığı için kömürlü yapıların iskeleti olarak veya sadece kağıt yapılarda kullanılmaktadır.",
			},
		},
		{
			Name: "eraser",
			NameLocal: entities.JSONMap{
				"en": "eraser",
				"tr": "silgi",
			},
			Description: entities.JSONMap{
				"en": "Eraser is the most basic material of war against the enemies created with coal.",
				"tr": "Kömürle yaratılan düşmanlara karşı savaşın en temel malzemesi silgidir.",
			},
		},
		{
			Name: "ink",
			NameLocal: entities.JSONMap{
				"en": "ink",
				"tr": "mürekkep",
			},
			Description: entities.JSONMap{
				"en": "Ink is the coal of the new age. Since it is not carbon-based, living units cannot be produced with it alone, but the durability against the eraser can be increased as a coating. Its persistence is also very important for science.",
				"tr": "Mürekkep, yeni çağın kömürüdür. Karbon bazlı olmadığı için tek başına canlı birimler üretilemez ancak kaplama olarak silgiye karşı dayanıklılığı arttırılabilir. Kalıcılığı da bilimsel gelişmeler için çok önemlidir.",
			},
		},
	}
	d.repo.Upsert(&entries, "resources_name_key", true, []string{"name_local", "description"})

}

func (d *data) initLocationTypes() {
	entries := []*entities.LocationType{
		{
			Name: "town",
			NameLocal: entities.JSONMap{
				"en": "town",
				"tr": "köy",
			},
			Description: entities.JSONMap{
				"en": "Stick figures came alive in this paper, and they made this place their town.",
				"tr": "Çöp adamlar bu kağıtta canlandı, ve burayı onların şehri yaptılar.",
			},
		},
		{
			Name: "unoccupied",
			NameLocal: entities.JSONMap{
				"en": "unoccupied",
				"tr": "boş",
			},
			Description: entities.JSONMap{
				"en": "This is an unoccupied location. It is not used by any stick figure.",
				"tr": "Bu boş bir yerdir. Herhangi bir çöp adam tarafından kullanılmaz.",
			},
		},
	}
	d.repo.Upsert(&entries, "location_types_name_key", true, []string{"name_local", "description"})
}

func (d *data) initStructures() {
	entries := []*entities.Structure{
		{
			Name: "academy",
			NameLocal: entities.JSONMap{
				"en": "academy",
				"tr": "akademi",
			},
			Description: entities.JSONMap{
				"en": "You can do  research for new technologies you need in this building.",
				"tr": "Bu binada ihtiyacınız olan yeni teknolojiler için araştırma yapabilirsiniz.",
			},
		},
		{
			Name: "airport",
			NameLocal: entities.JSONMap{
				"en": "airport",
				"tr": "havaalanı",
			},
			Description: entities.JSONMap{
				"en": "The building where you can craft new paper airplanes. Since the tables are too high for the stick figures, the only known way of transportation between these tables is paper planes.",
				"tr": "Yeni kağıt uçaklar üretebileceğiniz bina. Çöp adamlar için masalar çok yüksek olduğundan, bu masalar arasında bilinen tek ulaşım şekli kağıt uçaklardır.",
			},
		},
		{
			Name: "barracks",
			NameLocal: entities.JSONMap{
				"en": "barracks",
				"tr": "kışla",
			},
			Description: entities.JSONMap{
				"en": "The building where you can train new stick figures to use in your army.",
				"tr": "Ordunuzda kullanmak üzere yeni çöp adamlar yetiştirebileceğiniz bina.",
			},
		},
		{
			Name: "cemetery",
			NameLocal: entities.JSONMap{
				"en": "cemetery",
				"tr": "mezarlık",
			},
			Description: entities.JSONMap{
				"en": "There is no eternal rest in this world. You can recycle units here and get some of the materials back",
				"tr": "Birimleri burada geri dönüştürebilir ve bazı malzemeleri geri alabilirsiniz.",
			},
		},
		{
			Name: "embassy",
			NameLocal: entities.JSONMap{
				"en": "embassy",
				"tr": "elçilik",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "hospital",
			NameLocal: entities.JSONMap{
				"en": "hospital",
				"tr": "hastane",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "market",
			NameLocal: entities.JSONMap{
				"en": "market",
				"tr": "market",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "townhall",
			NameLocal: entities.JSONMap{
				"en": "townhall",
				"tr": "belediye binası",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "warehouse",
			NameLocal: entities.JSONMap{
				"en": "warehouse",
				"tr": "depo",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "workshop",
			NameLocal: entities.JSONMap{
				"en": "workshop",
				"tr": "atölye",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
	}
	d.repo.Upsert(&entries, "structures_name_key", true, []string{"name_local", "description"})

}

func (d *data) initAirCraftUnits() {
	entries := []*entities.Unit{
		{
			Name: "cargoPlane",
			NameLocal: entities.JSONMap{
				"en": "Cargo Plane",
				"tr": "Yük Uçağı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCargoPlane",
				"tr": "DescriptionCargoPlane",
			},
		},
		{
			Name: "transporter",
			NameLocal: entities.JSONMap{
				"en": "Transporter",
				"tr": "Taşıyıcı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionTransporter",
				"tr": "DescriptionTransporter",
			},
		},
		{
			Name: "fighterJet",
			NameLocal: entities.JSONMap{
				"en": "Fighter Jet",
				"tr": "Savaş Uçağı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionFighterJet",
				"tr": "DescriptionFighterJet",
			},
		},
		{
			Name: "bomber",
			NameLocal: entities.JSONMap{
				"en": "Bomber",
				"tr": "Bombardıman Zeplini",
			},
			Description: entities.JSONMap{
				"en": "DescriptionBomber",
				"tr": "DescriptionBomber",
			},
		},
		{
			Name: "uav",
			NameLocal: entities.JSONMap{
				"en": "UAV",
				"tr": "IHA",
			},
			Description: entities.JSONMap{
				"en": "DescriptionUAV",
				"tr": "DescriptionUAV",
			},
		},
		{
			Name: "attackHelicopter",
			NameLocal: entities.JSONMap{
				"en": "Attack Helicopter",
				"tr": "Saldırı Helikopteri",
			},
			Description: entities.JSONMap{
				"en": "DescriptionAttackHelicopter",
				"tr": "DescriptionAttackHelicopter",
			},
		},
	}
	d.repo.Upsert(&entries, "units_name_key", true, []string{"name_local", "description"})

}

func (d *data) initMeleeUnits() {
	entries := []*entities.Unit{
		{
			Name: "pikeman",
			NameLocal: entities.JSONMap{
				"en": "Pikeman",
				"tr": "Mızraklı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionPikeman",
				"tr": "DescriptionPikeman",
			},
		},
		{
			Name: "knight",
			NameLocal: entities.JSONMap{
				"en": "Knight",
				"tr": "Şövalye",
			},
			Description: entities.JSONMap{
				"en": "DescriptionKnight",
				"tr": "DescriptionKnight",
			},
		},
		{
			Name: "samurai",
			NameLocal: entities.JSONMap{
				"en": "Samurai",
				"tr": "Samuray",
			},
			Description: entities.JSONMap{
				"en": "DescriptionSamurai",
				"tr": "DescriptionSamurai",
			},
		},
		{
			Name: "cavalry",
			NameLocal: entities.JSONMap{
				"en": "Cavalry",
				"tr": "Atlı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCavalry",
				"tr": "DescriptionCavalry",
			},
		},
		{
			Name: "undead",
			NameLocal: entities.JSONMap{
				"en": "Undead",
				"tr": "Ölümsüz",
			},
			Description: entities.JSONMap{
				"en": "DescriptionUndead",
				"tr": "DescriptionUndead",
			},
		},
		{
			Name: "ghoul",
			NameLocal: entities.JSONMap{
				"en": "Ghoul",
				"tr": "Gulyabani",
			},
			Description: entities.JSONMap{
				"en": "DescriptionGhoul",
				"tr": "DescriptionGhoul",
			},
		},
	}
	d.repo.Upsert(&entries, "units_name_key", true, []string{"name_local", "description"})

}

func (d *data) initRangedUnits() {
	entries := []*entities.Unit{
		{
			Name: "archer",
			NameLocal: entities.JSONMap{
				"en": "Archer",
				"tr": "Okçu",
			},
			Description: entities.JSONMap{
				"en": "DescriptionArcher",
				"tr": "DescriptionArcher",
			},
		},
		{
			Name: "rifleman",
			NameLocal: entities.JSONMap{
				"en": "Rifleman",
				"tr": "Tüfekçi",
			},
			Description: entities.JSONMap{
				"en": "DescriptionRifleman",
				"tr": "DescriptionRifleman",
			},
		},
		{
			Name: "sniper",
			NameLocal: entities.JSONMap{
				"en": "Sniper",
				"tr": "Keskin Nişancı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionSniper",
				"tr": "DescriptionSniper",
			},
		},
		{
			Name: "wizard",
			NameLocal: entities.JSONMap{
				"en": "Wizard",
				"tr": "Büyücü",
			},
			Description: entities.JSONMap{
				"en": "DescriptionWizard",
				"tr": "DescriptionWizard",
			},
		},
		{
			Name: "witchHunter",
			NameLocal: entities.JSONMap{
				"en": "Witch Hunter",
				"tr": "Cadı Avcısı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionWitchHunter",
				"tr": "DescriptionWitchHunter",
			},
		},
	}
	d.repo.Upsert(&entries, "units_name_key", true, []string{"name_local", "description"})

}

func (d *data) initUtiltyUnits() {
	entries := []*entities.Unit{
		{
			Name: "Medic",
			NameLocal: entities.JSONMap{
				"en": "Medic",
				"tr": "Doktor",
			},
			Description: entities.JSONMap{
				"en": "DescriptionMedic",
				"tr": "DescriptionMedic",
			},
		},
		{
			Name: "engineer",
			NameLocal: entities.JSONMap{
				"en": "Engineer",
				"tr": "Mühendis",
			},
			Description: entities.JSONMap{
				"en": "DescriptionEngineer",
				"tr": "DescriptionEngineer",
			},
		},
		{
			Name: "necromancer",
			NameLocal: entities.JSONMap{
				"en": "Necromancer",
				"tr": "Ölüm Büyücüsü",
			},
			Description: entities.JSONMap{
				"en": "DescriptionNecromancer",
				"tr": "DescriptionNecromancer",
			},
		},
		{
			Name: "vampire",
			NameLocal: entities.JSONMap{
				"en": "Vampire",
				"tr": "Vampir",
			},
			Description: entities.JSONMap{
				"en": "DescriptionVampire",
				"tr": "DescriptionVampire",
			},
		},
	}
	d.repo.Upsert(&entries, "units_name_key", true, []string{"name_local", "description"})

}

func (d *data) initCharacteristics() {
	entries := []*entities.Characteristic{
		// Fundamental
		{
			Name: "attack",
			NameLocal: entities.JSONMap{
				"en": "Attack",
				"tr": "Saldırı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionAttack",
				"tr": "DescriptionAttack",
			},
		},
		{
			Name: "defense",
			NameLocal: entities.JSONMap{
				"en": "Defense",
				"tr": "Savunma",
			},
			Description: entities.JSONMap{
				"en": "DescriptionDefense",
				"tr": "DescriptionDefense",
			},
		},
		{
			Name: "health",
			NameLocal: entities.JSONMap{
				"en": "Health",
				"tr": "Sağlık",
			},
			Description: entities.JSONMap{
				"en": "DescriptionHealth",
				"tr": "DescriptionHealth",
			},
		},
		{
			Name: "evasion",
			NameLocal: entities.JSONMap{
				"en": "Evasion",
				"tr": "Kaçınma",
			},
			Description: entities.JSONMap{
				"en": "DescriptionEvade",
				"tr": "DescriptionEvade",
			},
		},
		{
			Name: "targetability",
			NameLocal: entities.JSONMap{
				"en": "Targetability",
				"tr": "Hedeflenebilirlik",
			},
			Description: entities.JSONMap{
				"en": "DescriptionTargetable",
				"tr": "DescriptionTargetable",
			},
		},
		{
			Name: "carriability",
			NameLocal: entities.JSONMap{
				"en": "Carriability",
				"tr": "Taşınabilirlik",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCarriable",
				"tr": "DescriptionCarriable",
			},
		},
		{
			Name: "canFly",
			NameLocal: entities.JSONMap{
				"en": "Can Fly",
				"tr": "Uçabilir",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCanFly",
				"tr": "DescriptionCanFly",
			},
		},
		{
			Name: "canCarryUnits",
			NameLocal: entities.JSONMap{
				"en": "Can Carry Units",
				"tr": "Birim Taşıyabilir",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCanCarryUnits",
				"tr": "DescriptionCanCarryUnits",
			},
		},
		{
			Name: "canCarryResources",
			NameLocal: entities.JSONMap{
				"en": "Can Carry Resources",
				"tr": "Kaynak Taşıyabilir",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCanCarryUnits",
				"tr": "DescriptionCanCarryUnits",
			},
		},
		{
			Name: "canAttack",
			NameLocal: entities.JSONMap{
				"en": "Cariability",
				"tr": "Taşınabilirlik",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCanAttack",
				"tr": "DescriptionCanAttack",
			},
		},
		{
			Name: "targetOrder",
			NameLocal: entities.JSONMap{
				"en": "Target Order",
				"tr": "Hedef Sırası",
			},
			Description: entities.JSONMap{
				"en": "DescriptionTargetPriority",
				"tr": "DescriptionTargetPriority",
			},
		},
		{
			Name: "targetCount",
			NameLocal: entities.JSONMap{
				"en": "Target Count",
				"tr": "Hedef Sayısı",
			},
			Description: entities.JSONMap{
				"en": "DescriptionTargetPriority",
				"tr": "DescriptionTargetPriority",
			},
		},
		{
			Name: "volume",
			NameLocal: entities.JSONMap{
				"en": "Volume",
				"tr": "Hacim",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCariability",
				"tr": "DescriptionCariability",
			},
		},
		{
			Name: "capacity",
			NameLocal: entities.JSONMap{
				"en": "Capacity",
				"tr": "Kapasite",
			},
			Description: entities.JSONMap{
				"en": "DescriptionCariability",
				"tr": "DescriptionCariability",
			},
		},
	}
	d.repo.Upsert(&entries, "characteristics_name_key", true, []string{"name_local", "description"})
}

func (d *data) initGridTypes() {
	entries := []*entities.GridType{
		{
			Name: "floor",
			NameLocal: entities.JSONMap{
				"en": "floor",
				"tr": "zemin",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "table",
			NameLocal: entities.JSONMap{
				"en": "table",
				"tr": "masa",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "pencilBox",
			NameLocal: entities.JSONMap{
				"en": "pencilBox",
				"tr": "kalem kutusu",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "recycleBin",
			NameLocal: entities.JSONMap{
				"en": "recycleBin",
				"tr": "geri dönüşüm kutusu",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "printer",
			NameLocal: entities.JSONMap{
				"en": "printer",
				"tr": "yazıcı",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
		{
			Name: "dustBin",
			NameLocal: entities.JSONMap{
				"en": "dustbin",
				"tr": "çöp kutusu",
			},
			Description: entities.JSONMap{
				"en": "",
				"tr": "",
			},
		},
	}
	d.repo.Upsert(&entries, "grid_types_name_key", true, []string{"name_local", "description"})
}

func (d *data) getResourceMap() map[string]uint {
	resourceMap := make(map[string]uint)
	resources := []*entities.Resource{}
	d.repo.FindAll(&resources, nil, nil, "")
	for _, structure := range resources {
		resourceMap[structure.Name] = structure.ID
	}
	return resourceMap
}

func (d *data) getLocationTypeMap() map[string]uint {
	locationTypeMap := make(map[string]uint)
	locationTypes := []*entities.LocationType{}
	d.repo.FindAll(&locationTypes, nil, nil, "")
	for _, locationType := range locationTypes {
		locationTypeMap[locationType.Name] = locationType.ID
	}
	fmt.Println("locationTypeMap", locationTypeMap)
	return locationTypeMap
}

func (d *data) getStructureMap() map[string]uint {
	structureMap := make(map[string]uint)
	structures := []*entities.Structure{}
	d.repo.FindAll(&structures, nil, nil, "")
	for _, structure := range structures {
		structureMap[structure.Name] = structure.ID
	}
	fmt.Println("structureMap", structureMap)
	return structureMap
}

func (d *data) getGridTypeMap() map[string]uint {
	gridTypeMap := make(map[string]uint)
	gridTypes := []*entities.GridType{}
	d.repo.FindAll(&gridTypes, nil, nil, "")
	for _, gridType := range gridTypes {
		gridTypeMap[gridType.Name] = gridType.ID
	}
	fmt.Println("gridTypeMap", gridTypeMap)
	return gridTypeMap
}
