package indexer

import (
	"github.com/antholord/poe-ML-indexer/api"
	"github.com/antholord/poe-ML-indexer/db"
	"strings"
	"regexp"
)

var ModReg *regexp.Regexp = regexp.MustCompile("(?:\\-?\\d*\\.)?\\-?\\d+")

type Parser struct {
	conn *db.DB
}

func NewParser() *Parser {
	x := db.Connect()
	return &Parser{
		conn : x,
	}
}

func (p *Parser) processStash(stash *api.Stash, itemSlice *[]*api.Item) {

	for _, item := range stash.Items {
		 p.processItem(&item, itemSlice)
	}
}

func (p *Parser) processItem(item *api.Item, itemSlice *[]*api.Item){
	if item.League == "Legacy" && item.FrameType == 2 && item.IsIdentified && strings.Contains(item.Note, "chaos") {
		item.Type = p.removeSuperior(item.Type)
		//log.Println("Adding" + item.Name)
		*itemSlice = append(*itemSlice, item)
	}

}

func (p *Parser) removeSuperior(s string) string{
	if strings.Contains(s, "Superior "){
		return strings.Replace(s, "Superior ", "", 2)
	}else {
		return s
	}
}

