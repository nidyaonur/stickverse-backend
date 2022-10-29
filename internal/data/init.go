package data

import (
	"fmt"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/pkg/utils"
)

var (
	gridTypeNames = []string{"floor", "table", "pencilBox", "recycleBin", "printer", "dustBin"}
)

func (d *data) initGridMap() {
	gridTypeMap := make(map[string]uint)
	for _, gridTypeName := range gridTypeNames {
		gridType := entities.GridType{Name: gridTypeName}
		gridTypeInterface, err := d.repo.FindFirstWithCondition(&gridType, "name = ?", gridTypeName)
		if err != nil {
			panic(err)
		}
		gridType = *gridTypeInterface.(*entities.GridType)
		gridTypeMap[gridTypeName] = gridType.ID
	}

	grids := []*entities.Grid{}
	for x := -10; x < 10; x++ {
		for y := -10; y < 10; y++ {
			grid := &entities.Grid{
				X: x,
				Y: y,
			}
			modX := utils.PositiveMod(float64(x), 10)
			modY := utils.PositiveMod(float64(y), 10)

			fmt.Println(x, y, modX, modY)
			grid.Capacity = 0
			switch {
			case modX == 2 && modY == 2:
				grid.GridTypeID = gridTypeMap["printer"]
				fmt.Println("printer")
			case modX == 2 && modY == 7:
				grid.GridTypeID = gridTypeMap["pencilBox"]
				fmt.Println("pencilBox")
			case modX == 7 && modY == 2:
				grid.GridTypeID = gridTypeMap["recycleBin"]
				fmt.Println("recycleBin")
			case modX == 7 && modY == 7:
				grid.GridTypeID = gridTypeMap["dustBin"]
				fmt.Println("dustBin")
			case modX == 1 && modY == 1 || // first column
				modX == 1 && modY == 3 ||
				modX == 1 && modY == 6 ||
				modX == 1 && modY == 8 ||
				modX == 3 && modY == 1 || // second column
				modX == 3 && modY == 3 ||
				modX == 3 && modY == 6 ||
				modX == 3 && modY == 8 ||
				modX == 6 && modY == 1 || // third column
				modX == 6 && modY == 3 ||
				modX == 6 && modY == 6 ||
				modX == 6 && modY == 8 ||
				modX == 8 && modY == 1 || // fourth column
				modX == 8 && modY == 3 ||
				modX == 8 && modY == 6 ||
				modX == 8 && modY == 8:
				grid.GridTypeID = gridTypeMap["table"]
				fmt.Println("table")
				grid.Capacity = 10
			default:
				grid.GridTypeID = gridTypeMap["floor"]
				fmt.Println("floor")

			}
			grids = append(grids, grid)

		}
	}

	d.repo.Insert(grids)
}

func (d *data) fillGridWithLocations() {
	locationTypeMap := d.getLocationTypeMap()
	structureMap := d.getStructureMap()
	resourceMap := d.getResourceMap()

	grids := []*entities.Grid{}
	d.repo.FindAll(&grids, nil, nil, "")
	for _, grid := range grids {
		if grid.GridType.Name != "table" {
			continue
		}
		for i := 0; i < int(grid.Capacity-grid.Occupied); i++ {
			// create location
			location := &entities.Location{
				Name: "emptyPage",
				NameLocal: entities.JSONMap{
					"en": "empty page",
					"tr": "boş sayfa",
				},
				GridID:         grid.ID,
				LocationTypeID: locationTypeMap["unoccupied"],
			}
			d.repo.Insert(location)

			// create structure
			structureBuilts := []*entities.StructureBuilt{}
			for _, id := range structureMap {
				structureBuilt := &entities.StructureBuilt{
					StructureID: id,
					LocationID:  location.ID,
				}
				structureBuilts = append(structureBuilts, structureBuilt)
			}
			d.repo.Insert(structureBuilts)

			// create resource
			locationResources := []*entities.LocationResource{}
			for _, id := range resourceMap {
				locationResource := &entities.LocationResource{
					ResourceID: id,
					LocationID: location.ID,
				}
				locationResources = append(locationResources, locationResource)
			}
			d.repo.Insert(locationResources)
		}
	}
}

func (d *data) createPrerequisites() {
	preqs := []*entities.Prerequisite{}
	structureTypes := []*entities.Structure{}
	d.repo.FindAll(&structureTypes, nil, nil, "")
	for _, structureType := range structureTypes {
		switch structureType.Name {
		case "townhall":
			preqs = append(preqs, &entities.Prerequisite{
				StructureID:         &structureType.ID,
				Type:                "resource",
				PrerequisiteFormula: "coal:1:exp:1.2:50",
			}, &entities.Prerequisite{
				StructureID:         &structureType.ID,
				Type:                "resource",
				PrerequisiteFormula: "paper:3:exp:1.1:50",
			})
		case "airport":
			preqs = append(preqs, &entities.Prerequisite{
				StructureID:         &structureType.ID,
				Type:                "resource",
				PrerequisiteFormula: "coal:1:exp:1.2:50",
			}, &entities.Prerequisite{
				StructureID:         &structureType.ID,
				Type:                "resource",
				PrerequisiteFormula: "paper:1:exp:1.2:50",
			})
		}

	}
	d.repo.Insert(&preqs)

}

func (d *data) InitDataValues() {
	//d.initResources()
	//d.initLocationTypes()
	//d.initStructures()
	//d.initGridTypes()
	//d.initGridMap()
	//d.fillGridWithLocations()
	//d.createPrerequisites()

}
