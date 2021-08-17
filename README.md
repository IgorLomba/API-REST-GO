# API-REST-GO
 A API that you can register people write in golang

Usage: 

- POST
```
{
  "name": "name",
  "birth": "01-01-1991",
  "nif": "000000001",
  "address": {
    "city": "Portugal",
    "street": "Portugal"
	}
}
```
- PUT
```
{
  "id"  : 1,
  "name": "nameEdited",
  "birth": "28-06-1111",
  "nif": "909090909",
  "address": {
    "city": "Portugal",
    "street": "Portugalllll",
    "id" : 1   
    }
}
```
- DELETE:
```
    http://localhost:3000/api/person/:id
```

- GET (3 ways):
```
    http://localhost:3000/api/person/:id

    http://localhost:3000/api/person/name/:name
    
    http://localhost:3000/api/address/:StreetOrCity
```