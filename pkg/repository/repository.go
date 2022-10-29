package repository

import (
	"fmt"
	"log"

	"github.com/nidyaonur/stickverse-backend/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CloseDBConnection()
	Insert(i interface{}) (interface{}, error)
	Update(i interface{}) (interface{}, error)
	Delete(i interface{}) error
	DeleteOnCondition(i interface{}, condition string) error
	FindFirstWithCondition(i interface{}, condition string, values ...interface{}) (interface{}, error)
	FindFirstWithCustomPreloads(i interface{}, preloads []string, condition string, values ...interface{}) (interface{}, error)
	FindAll(a interface{}, pageSize, offset *int, condition string) (interface{}, error)
	FindAllWithConditionValues(a interface{}, pageSize, offset *int, condition, search string, values ...interface{}) (interface{}, error)
	FindAllWithCustomPreloads(a interface{}, pageSize, offset *int, preloads []string, condition string, values ...interface{}) (interface{}, error)
	LoadUser(user *entities.User) error
	LoadUsers(users []*entities.User) error
	LoadWithUsername(username string) (*entities.User, error)
	GetCapacities(maxCapacity int) []uint
	GetMinGridWithCapacity(wantedCapacity, maxCapacity int) (gridIDs []uint)
}

type Repo struct {
	db *gorm.DB
}

func NewRepository(databaseUrl string) Repository {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database", err)
		return nil
	}
	db.AutoMigrate(
		// user
		&entities.User{},
		&entities.LoginHistory{},
		// Alliance
		&entities.Alliance{},
		&entities.AllianceMember{},
		&entities.MembershipAction{},
		&entities.MembershipHistory{},
		&entities.MembershipType{},
		&entities.AllowedAction{},
		&entities.Location{},
		&entities.LocationType{},
		&entities.LocationUnit{},
		&entities.LocationResource{},
		&entities.Grid{},
		&entities.GridType{},
		&entities.Structure{},
		&entities.StructureResource{},
		&entities.StructureBuilt{},
		&entities.Prerequisite{},
		&entities.Resource{},
		&entities.Research{},
		&entities.ResearchLevel{},
		&entities.ResearchResource{},
		&entities.Unit{},
		&entities.UnitCharacteristic{},
		&entities.UnitCost{},
		&entities.Characteristic{},
		&entities.GroupMovement{},
		&entities.GroupMovementUnit{},
		&entities.MovementType{},
		&entities.GroupMovement{},
		&entities.Unit{},
	)
	return &Repo{db}
}

func (r *Repo) CloseDBConnection() {
	con, err := r.db.DB()
	if err != nil {
		log.Fatal("Could not close database connection", err)
		return
	}
	if err := con.Close(); err != nil {
		log.Fatal("Could not close database connection", err)
	}
}

func (r *Repo) Insert(i interface{}) (interface{}, error) {
	err := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(i).Error
	return i, err
}

func (r *Repo) Update(i interface{}) (interface{}, error) {
	err := r.db.Debug().Save(i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (r *Repo) Delete(i interface{}) error {
	return r.db.Delete(i).Error
}

func (r *Repo) DeleteOnCondition(i interface{}, condition string) error {
	return r.db.Where(condition).Delete(i).Error
}

func (r *Repo) FindFirstWithCondition(i interface{}, condition string, values ...interface{}) (interface{}, error) {
	dbQuery := r.db.Debug().Preload(clause.Associations).Order("updated_at desc").
		Where(condition, values...)
	err := dbQuery.First(i).Error
	return i, err
}

func (r *Repo) FindFirstWithCustomPreloads(i interface{}, preloads []string, condition string, values ...interface{}) (interface{}, error) {
	dbQuery := r.db.Debug().Order("updated_at desc").
		Where(condition, values...)
	for _, preload := range preloads {
		dbQuery = dbQuery.Preload(preload)
	}
	err := dbQuery.First(i).Error
	return i, err
}

func (r *Repo) FindAll(i interface{}, pageSize, offset *int, condition string) (interface{}, error) {
	dbQuery := r.db.Preload(clause.Associations).Order("updated_at desc")
	if condition != "" {
		dbQuery.Where(condition)
	}
	if pageSize != nil && offset != nil {
		dbQuery.Limit(*pageSize).Offset(*offset)
	}
	err := dbQuery.Find(i).Error
	return i, err
}

func (r *Repo) FindAllWithConditionValues(i interface{}, pageSize, offset *int, condition, search string, values ...interface{}) (interface{}, error) {
	dbQuery := r.db.Debug().Preload(clause.Associations).Order("updated_at desc").
		Where(condition, values...)
	if pageSize != nil && offset != nil {
		dbQuery.Limit(*pageSize).Offset(*offset)
	}
	if search != "" {
		search = "%" + search + "%"
		dbQuery = dbQuery.Where("description ILIKE ?", search)
	}
	err := dbQuery.Find(i).Error
	return i, err
}

func (r *Repo) FindAllWithCustomPreloads(i interface{}, pageSize, offset *int, preloads []string, condition string, values ...interface{}) (interface{}, error) {
	dbQuery := r.db.Debug().Order("updated_at desc").
		Where(condition, values...)
	for _, preload := range preloads {
		dbQuery = dbQuery.Preload(preload)
	}
	if pageSize != nil && offset != nil {
		dbQuery.Limit(*pageSize).Offset(*offset)
	}
	err := dbQuery.Find(i).Error
	return i, err
}

func (r *Repo) LoadUser(user *entities.User) error {
	err := r.db.Debug().Preload("Locations").Preload("Locations.Resources").Preload("Locations.Resources.Resource").
		Preload("Locations.Structures").Preload("Locations.Structures.Structure").
		Preload("Locations.Resources.Resource").Preload("Locations.Structures.Structure.Prerequisites").
		Preload("Locations.Structures.Location").Preload("Locations.Structures.Location.Resources").
		Preload("Locations.Structures.Location.User").Preload("Locations.Structures.Location.Structures").
		First(user).Error
	return err
}
func (r *Repo) LoadUsers(users []*entities.User) error {
	err := r.db.Debug().Preload("Locations").Preload("Locations.Resources").
		Preload("Locations.Structures").Preload("Locations.Structures.Structure").
		Preload("Locations.Resources.Resource").Preload("Locations.Structures.Structure.Prerequisites").
		Preload("Locations.Structures.Location").Preload("Locations.Structures.Location.Resources").
		Preload("Locations.Structures.Location.User").Preload("Locations.Structures.Location.Structures").
		Find(&users).Error
	return err
}

func (r *Repo) LoadWithUsername(username string) (*entities.User, error) {
	user := entities.User{}
	err := r.db.Preload(clause.Associations).Where("username = ?", username).Find(&user).Error
	return &user, err
}

func (r *Repo) GetCapacities(maxCapacity int) (capacities []uint) {
	//select grid capacity
	var capacityRes []map[string]interface{}
	//r.db.Table("grids").Select("occupied, count(id) as result1, sum(occupied) as result2, min(occupied) as result3, max(occupied) as result4, avg(occupied) as result5").
	//r.db.Table("grids").Select("occupied").
	//	Where("occupied < ?", maxCapacity).Group("occupied").Scan(&capacityRes)
	r.db.Raw(`SELECT occupied FROM grids WHERE occupied<? AND capacity=10 GROUP BY occupied HAVING count(*)>0;`, maxCapacity).Scan(&capacityRes)
	for _, capacity := range capacityRes {
		fmt.Println(capacity)
		capacities = append(capacities, uint(capacity["occupied"].(int64)))

	}
	fmt.Println(capacities)
	return capacities
}

func (r *Repo) GetMinGridWithCapacity(wantedCapacity, maxCapacity int) (gridIDs []uint) {
	//select grid capacity
	var capacityRes []map[string]interface{}
	//r.db.Table("grids").Select("occupied, count(id) as result1, sum(occupied) as result2, min(occupied) as result3, max(occupied) as result4, avg(occupied) as result5").
	//r.db.Table("grids").Select("occupied").Clauses(clause.OrderBy{Expression: "occupied", Direction: clause.Desc}).
	//	Where("occupied < ?", maxCapacity).Group("occupied").Scan(&capacityRes)
	r.db.Raw(`SELECT*,x*x+y*y AS"distance" FROM grids WHERE occupied=? AND capacity=? ORDER BY"distance" ASC`, wantedCapacity, maxCapacity).Scan(&capacityRes)
	var minDistance int64 = 999999999
	for _, capacity := range capacityRes {
		distance := capacity["distance"].(int64)
		if distance < minDistance {
			minDistance = distance
		} else if distance > minDistance {
			break
		}
		gridIDs = append(gridIDs, uint(capacity["id"].(int64)))
	}
	return
}
