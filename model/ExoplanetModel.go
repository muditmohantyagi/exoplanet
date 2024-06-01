package model

import (
	"errors"

	"gorm.io/gorm"
	"planet.com/dto"
)

type Expplanet struct {
	ID                uint
	Name              string  `gorm:"type:varchar(250);not null"`
	Description       string  `gorm:"type:varchar(250);not null"`
	DistanceFromEarth int     `gorm:"not null"`
	Radius            float64 `gorm:"not null"`
	Mass              float64
	Type              string `gorm:"type:varchar(250);not null"`
	Date              `gorm:"embedded"`
}

func FinAll(InputDTO dto.SortAndFilter) ([]Expplanet, error) {
	var Expplanet []Expplanet
	queryDB := DB
	if InputDTO.FilterBymass != 0 {
		queryDB = queryDB.Where("mass = ?", InputDTO.FilterBymass).Session(&gorm.Session{})

	}
	if InputDTO.SortByRadius != "" {
		queryDB = queryDB.Order("radius " + InputDTO.SortByRadius).Session(&gorm.Session{})

	}
	if result := queryDB.Find(&Expplanet); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return Expplanet, nil
		}
		return nil, result.Error
	}
	return Expplanet, nil

}

func FindById(id int) (*Expplanet, error) {
	var Expplanet *Expplanet = new(Expplanet)

	if result := DB.Where("id", id).Take(&Expplanet); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, nil
		}
		return nil, result.Error
	}
	return Expplanet, nil

}
func UpdateById(data Expplanet, id int) (bool, error) {
	if result := DB.Model(&Expplanet{}).Where("id=?", id).Updates(&data); result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		err := errors.New("invalid operation")

		return false, err
	}
	return true, nil
}

func DeteteByID(id int) (bool, error) {

	var Expplanet Expplanet
	if result := DB.Where("id = ?", id).Delete(&Expplanet); result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		err := errors.New("invalid operation")

		return false, err
	}
	return true, nil
}
