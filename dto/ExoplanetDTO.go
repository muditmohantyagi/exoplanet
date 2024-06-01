package dto

type AddExpplanet struct {
	Name              string  `binding:"required"`
	Description       string  `binding:"required"`
	DistanceFromEarth int     `binding:"required,min=10,max=1000"`
	Radius            float64 `binding:"required,min=0.1,max=10"`
	Mass              float64 `binding:"required_if=Type Terrestrial"`
	Type              string  `binding:"required,oneof=GasGiant Terrestrial"`
}
type UpdateExpplanet struct {
	Id                int     `binding:"required"`
	Name              string  `binding:"required"`
	Description       string  `binding:"required"`
	DistanceFromEarth int     `binding:"required,min=10,max=1000"`
	Radius            float64 `binding:"required,min=0.1,max=10"`
	Mass              float64 `binding:"required_if=Type Terrestrial"`
	Type              string  `binding:"required,oneof=GasGiant Terrestrial"`
}
type GetId struct {
	Id int `uri:"id" binding:"required"`
}
type FuleCalculation struct {
	ExoPlanetId  int `binding:"required"`
	CrewCapacity int `binding:"required"`
}
type SortAndFilter struct {
	SortByRadius string  `binding:"omitempty,required,oneof=asc desc"`
	FilterBymass float64 `binding:"omitempty,required,min=0.1,max=10"`
}
