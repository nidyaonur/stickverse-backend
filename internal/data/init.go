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
	gridTypeMap := d.getGridTypeMap()
	if len(gridTypeMap) == 0 {
		panic("grid type map is empty")
	}
	resourceMap := d.getResourceMap()
	if len(resourceMap) == 0 {
		panic("resource map is empty")
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
				reosurceID := resourceMap["ink"]
				grid.ResourceID = &reosurceID
				fmt.Println("printer")
			case modX == 2 && modY == 7:
				grid.GridTypeID = gridTypeMap["pencilBox"]
				reosurceID := resourceMap["coal"]
				grid.ResourceID = &reosurceID
				fmt.Println("pencilBox")
			case modX == 7 && modY == 2:
				grid.GridTypeID = gridTypeMap["recycleBin"]
				reosurceID := resourceMap["paper"]
				grid.ResourceID = &reosurceID
				fmt.Println("recycleBin")
			case modX == 7 && modY == 7:
				grid.GridTypeID = gridTypeMap["dustBin"]
				reosurceID := resourceMap["eraser"]
				grid.ResourceID = &reosurceID
				fmt.Println("dustBin")
			case modX == 1 && modY == 1 || // first column
				modX == 1 && modY == 3 ||
				modX == 3 && modY == 1 || // second column
				modX == 3 && modY == 3:
				grid.GridTypeID = gridTypeMap["table"]
				grid.Capacity = 10
				reosurceID := resourceMap["ink"]
				grid.ResourceID = &reosurceID
			case modX == 1 && modY == 6 ||
				modX == 1 && modY == 8 ||
				modX == 3 && modY == 6 ||
				modX == 3 && modY == 8:
				grid.GridTypeID = gridTypeMap["table"]
				grid.Capacity = 10
				reosurceID := resourceMap["coal"]
				grid.ResourceID = &reosurceID

			case modX == 6 && modY == 1 || // third column
				modX == 6 && modY == 3 ||
				modX == 8 && modY == 1 || // fourth column
				modX == 8 && modY == 3:
				grid.GridTypeID = gridTypeMap["table"]
				grid.Capacity = 10
				reosurceID := resourceMap["paper"]
				grid.ResourceID = &reosurceID

			case modX == 6 && modY == 6 ||
				modX == 6 && modY == 8 ||
				modX == 8 && modY == 6 ||
				modX == 8 && modY == 8:
				grid.GridTypeID = gridTypeMap["table"]
				grid.Capacity = 10
				reosurceID := resourceMap["eraser"]
				grid.ResourceID = &reosurceID
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
	if len(locationTypeMap) == 0 {
		panic("location type map is empty")
	}
	structureMap := d.getStructureMap()
	if len(structureMap) == 0 {
		panic("structure map is empty")
	}
	resourceMap := d.getResourceMap()
	if len(resourceMap) == 0 {
		panic("resource map is empty")
	}

	grids := []*entities.Grid{}
	d.repo.FindAll(&grids, nil, nil, "")
	for _, grid := range grids {
		if grid.GridType.Name != "table" {
			continue
		}

		occupied := grid.Occupied

		for i := 0; i < int(grid.Capacity-occupied); i++ {
			// create location
			location := &entities.Location{
				Name: "emptyPage",
				NameLocal: entities.JSONMap{
					"en": "empty page",
					"tr": "boş sayfa",
				},
				Workers:        100,
				GridID:         grid.ID,
				GridIndex:      int(grid.Occupied) + i,
				LocationTypeID: locationTypeMap["unoccupied"],
			}
			d.repo.Insert(location)
			grid.Occupied++
			d.repo.Update(grid)

			// create structure
			structureBuilts := []*entities.StructureBuilt{}
			for name, id := range structureMap {
				structureBuilt := &entities.StructureBuilt{
					StructureID: id,
					LocationID:  location.ID,
				}
				if name == "townhall" {
					structureBuilt.Level = 1
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
				if id == *grid.ResourceID {
					locationResource.Multiplier = 0.3
				}

				locationResources = append(locationResources, locationResource)
			}
			d.repo.Insert(locationResources)
		}
	}
}

func (d *data) createPrerequisites() {
	preqs := []*entities.StructurePrerequisite{}
	structureTypes := []*entities.Structure{}
	d.repo.FindAll(&structureTypes, nil, nil, "")
	for _, structureType := range structureTypes {
		switch structureType.Name {
		case "townhall":
			preqs = append(preqs, &entities.StructurePrerequisite{
				StructureID: structureType.ID,
				UniqueID:    "townhall:resource:coal",
				Type:        "resource",
				SubType:     "coal",
				Formula:     "exp:2:50",
			}, &entities.StructurePrerequisite{
				StructureID: structureType.ID,
				UniqueID:    "townhall:resource:paper",
				Type:        "resource",
				SubType:     "paper",
				Formula:     "exp:1.5:50",
			})
		case "airport":
			preqs = append(preqs, &entities.StructurePrerequisite{
				StructureID: structureType.ID,
				UniqueID:    "airport:resource:coal",
				Type:        "resource",
				SubType:     "coal",
				Formula:     "exp:1.5:50",
			}, &entities.StructurePrerequisite{
				StructureID: structureType.ID,
				UniqueID:    "airport:resource:paper",
				Type:        "resource",
				SubType:     "paper",
				Formula:     "exp:2:50",
			})
		}

	}
	d.repo.Upsert(&preqs, "structure_prerequisites_unique_id_key", true, []string{"formula", "amount"})

}

func (d *data) InitDataValues() {
	d.initResources()
	d.initLocationTypes()
	d.initStructures()
	d.initGridTypes()
	d.initGridMap()
	d.fillGridWithLocations()
	d.createPrerequisites()
	d.initAirCraftUnits()
	d.initMeleeUnits()
	d.initRangedUnits()
	d.initUtiltyUnits()
	d.initCharacteristics()

}
