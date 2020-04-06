package excel

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type SortType string

const (
	CreatedAtDesc SortType = "createdAtDesc"
	CreatedAtAsc  SortType = "createdAtAsc"
	UpdatedAtDesc SortType = "updatedAtDesc"
	UpdatedAtAsc  SortType = "updatedAtAsc"
	NameDesc      SortType = "nameDesc"
	NameAsc       SortType = "nameAsc"
	ItemCountDesc SortType = "itemCountDesc"
)

func Contains(arr *[]SortType, str SortType) bool {
	if arr == nil {
		return false
	}
	for _, v := range *arr {
		if v == str {
			return true
		}
	}
	return false
}

func SortString(name *string, sortType SortType) *string {
	var typeString string
	switch sortType {
	case CreatedAtDesc:
		typeString = "created_at desc"
	case CreatedAtAsc:
		typeString = "created_at asc"
	case UpdatedAtDesc:
		typeString = "created_at desc"
	case UpdatedAtAsc:
		typeString = "created_at asc"
	case NameDesc:
		typeString = "name desc"
	case NameAsc:
		typeString = "name asc"
	default:
	}
	if name != nil {
		typeString = *name + "." + typeString
	}
	return &typeString
}

func AppendSort(arr *[]SortType, db *gorm.DB) *gorm.DB {
	if arr == nil {
		return db
	}
	con := db
	if Contains(arr, CreatedAtDesc) {
		con = con.Order("created_at desc")
	}
	if Contains(arr, CreatedAtAsc) {
		con = con.Order("created_at asc")
	}
	if Contains(arr, UpdatedAtDesc) {
		con = con.Order("updated_at desc")
	}
	if Contains(arr, UpdatedAtAsc) {
		con = con.Order("updated_at asc")
	}
	if Contains(arr, NameDesc) {
		con = con.Order("name desc")
	}
	if Contains(arr, NameAsc) {
		con = con.Order("name asc")
	}
	return con
}

func Parse(stringType *string) *SortType {
	if stringType == nil {
		return nil
	}
	if *stringType == "created_at_desc" {
		sortType := CreatedAtDesc
		return &sortType
	}
	if *stringType == "created_at_asc" {
		sortType := CreatedAtAsc
		return &sortType
	}
	if *stringType == "updated_at_desc" {
		sortType := UpdatedAtDesc
		return &sortType
	}
	if *stringType == "updated_at_asc" {
		sortType := UpdatedAtAsc
		return &sortType
	}
	if *stringType == "item_count_desc" {
		sortType := ItemCountDesc
		return &sortType
	}
	return nil
}

type ExcelParams struct {
	UserIds    []int     `json:"userIds"`
	StarteDate time.Time `form:"start_date" time_format:"2006-01-02T15:04:05Z07:00"`
	SortType   SortType  `form:"sort_type"`
}

type User struct {
	ID                int    `gorm:"primary_key" json:"id"`
	Role              int    `json:"role"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
	ClientID          sql.NullInt32
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	Parent            sql.NullBool `sql:"-"`
}

type Client struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Current struct {
	User   *User
	Client *Client
}

type Excel struct {
	UserID           int
	UserName         string
	StartedDate      *time.Time
	StartedDatetime  *time.Time
	EndedDatetime    *time.Time
	ConstructionName string
	Memo             string
	Address          string
	SalesUserName    string
}

type JoinUser struct {
	ID        *int
	UserID    *int
	CreatedAt *time.Time
}

type excelPresenter struct{}

func NewExcelPresenter() ExcelPresenter {
	return &excelPresenter{}
}

type ExcelPresenter interface {
	ResponseExcelObject(current *Current, items []*Excel, joinUser []*JoinUser, params *ExcelParams) ([]byte, error)
}

func (excelPresenter *excelPresenter) CreateJsonParam(current *Current, excelDatas []*Excel, joinUser []*JoinUser, params *ExcelParams) (*string, error) {
	return nil, nil
}

func Sort(startJST time.Time, excelDatas []*Excel) map[int]map[int]map[int]*Excel {
	excelData := make(map[int]map[int]map[int]*Excel)
	weekDay := make(map[int]string)
	for i := 0; i < 7; i++ {
		weekDay[i] = startJST.AddDate(0, 0, i).Weekday().String()
	}
	page := 0
	for _, v := range excelDatas {
		day := 0
		if v.StartedDate != nil {
			s := v.StartedDate.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Weekday().String()
			for i, week := range weekDay {
				if week == s {
					day = i
					break
				}
			}
		} else {
			day = 7
		}
		if excelData[v.UserID] == nil {
			excelData[v.UserID] = make(map[int]map[int]*Excel)
		}
		for i := 0; i < 8; i++ {
			if excelData[v.UserID][i] == nil {
				excelData[v.UserID][i] = make(map[int]*Excel)
			}
		}
		count := len(excelData[v.UserID][day])
		count += 1
		excelData[v.UserID][day][count] = v
		page += 1
	}
	return excelData
}
