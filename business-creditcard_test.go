package creditcard

import (
	"testing"
)

type card struct {
	Number string
	Type string
	Valid bool
}

var cards = []*card{
	&card{"4242424242424242", "Visa", true}, // should pass
	&card{"4213729238347292", "Visa", false}, // should fail
	&card{"79927398713", "Unknown", true}, // should pass
	&card{"79927398710", "Unknown", false}, // should fail
	&card{"6011111111111117", "Discover", true}, // should pass
	&card{"6011342393482022", "Discover", false}, // should fail
	&card{"344347386473833", "AmericanExpress", false}, // should fail
	&card{"374347386473833", "AmericanExpress", false}, // should fail
	&card{"30569309025904", "DinersClub/Carteblanche", true}, // should pass
	&card{"30569309025905", "DinersClub/Carteblanche", false}, // should fail
	&card{"5555555555554444", "MasterCard", true}, // should pass
	&card{"5555555555554445", "MasterCard", false}, // should fail
	&card{"180034239348202", "JCB", false}, // should fail
}

func TestValidate(t *testing.T) {
	for _, card := range cards {
		valid := Validate(card.Number)
		
		if valid != card.Valid {
			t.Errorf("card validation [%v]; want [%v]", valid, card.Valid)
		}
	}
}

func TestCardype(t *testing.T) {
	for _, card := range cards {
		cardType := Cardtype(card.Number)
		
		if cardType != card.Type {
			t.Errorf("card type [%s]; want [%s]", cardType, card.Type)
		}
	}
}
