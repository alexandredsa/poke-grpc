# POKE-GRPC

This project is a PoC(Proof of Concept) of a gRPC server implemented using Go.


## Data
To enrich this implementation a [Pokemon](https://www.kaggle.com/abcsds/pokemon/version/2) dataset was downloaded from Kaggle. The application seed the database (Mongo) using the csv file.

## Methods

### List

**Request**

| Parameter | Type      | Allowed Values                                                                   | 
| :-------- | :-------  | :-------------------------                                                       | 
| `filterKey` | string  | any attribute from pokemon object's response.                                    |
| `filterType` | string | could be `regex` (has a SQL's `LIKE` behavior) or `exact`. Default is `regex`.   | 
| `filterValue` | string| N/A                                                                              | 

**Response**

```json
{
  "pokemons": [
    {
      "id": "7",
      "name": "Squirtle",
      "type1": "Water",
      "type2": "",
      "total": 314,
      "hp": 44,
      "attack": 48,
      "defense": 65,
      "specialAttack": 50,
      "specialDefense": 64,
      "speed": 43,
      "generation": 1,
      "legendary": false
    }
  ]
}
```

---

## How can I run it?

---
**NOTE:**

You must have an environment with: `Go`, `Docker`, `Docker Compose`, and maybe some grpc client to check it out.
I personally recommend [grpcui](https://github.com/fullstorydev/grpcui) (at least for now since I don't have so much experience with gRPC).

---

You can run the application with `development` configuration (that is the project's default using `.env` file's located at project's root).

1. Run `docker-compose up -d` to provision local MongoDB.
2. To seed the bank with `pokemon.csv` run: `make seed-db`
3. And finally, to run the app: `make run`
