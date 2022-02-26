package model

type Pokemon struct {
	PokemonID      string `csv:"#" bson:"pokemon_id"`
	Name           string `csv:"Name" bson:"name"`
	Type1          string `csv:"Type 1" bson:"type_1"`
	Type2          string `csv:"Type 2" bson:"type_2"`
	Total          int32  `csv:"Total" bson:"total"`
	Hp             int32  `csv:"HP" bson:"hp"`
	Attack         int32  `csv:"Attack" bson:"attack"`
	Defense        int32  `csv:"Defense" bson:"defense"`
	SpecialAttack  int32  `csv:"Sp. Atk" bson:"special_attack"`
	SpecialDefense int32  `csv:"Sp. Def" bson:"special_defense"`
	Speed          int32  `csv:"Speed" bson:"speed"`
	Generation     int32  `csv:"Generation" bson:"generation"`
	Legendary      bool   `csv:"Legendary" bson:"legendary"`
}
