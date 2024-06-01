# exoplanet
this microservice is developed using
1) gin
2) gorm(MYSQL) and mysql is running in docker
3) golang
4) validation
5) Listing all exoplanet with  sorting or filter
6) test cases
7) docker

# EndPoints:
1) POST: localhost:8080/api/exoplanet/add_exoplanet
{
    "Name": "GasGiant 2",
    "Description": "i am GasGiant 2",
    "DistanceFromEarth": 400,
    "Radius":3,
    "Mass": 0,
    "Type": "GasGiant"
}
2) PUT: localhost:8080/api/exoplanet/update_exoplanet
{
    "Id": 3,
    "Name": "GasGiant update again",
    "Description": "i am GasGiant 2",
    "DistanceFromEarth": 400,
    "Radius": 3,
    "Mass": 0,
    "Type": "GasGiant"
}
3) with optional search and filter: GET: localhost:8080/api/exoplanet/list_all_exoplanet
{
    "SortByRadius":"desc",
    "FilterBymass":0
}

4) GET: localhost:8080/api/exoplanet/list_exoplanet_byid/1

5) DELET: localhost:8080/api/exoplanet/delete_exoplanet_byid/8

6) GET: localhost:8080/api/exoplanet/fuel_estimation
{
    "ExoPlanetId": 4,
    "CrewCapacity": 3
}
