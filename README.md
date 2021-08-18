# API-REST-GO
An API you can register, update, list and delete people.
Code written in golang

Usage: 

- POST
```
http://localhost:3000/api/person/

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
http://localhost:3000/api/person/

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
