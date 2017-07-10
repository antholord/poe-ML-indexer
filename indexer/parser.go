package indexer

import (
	"github.com/antholord/poe-micro/api"
	"log"
	"github.com/antholord/poe-micro/db"
	"strings"
	"regexp"
)

var ModReg *regexp.Regexp = regexp.MustCompile("(?:\\-?\\d*\\.)?\\-?\\d+")

type Parser struct {
	ModsMap map[string]string

	TypesMap map[string]bool
	SubCategories map[string]bool
	TopCategories map[string]bool

	ItemsMap map[string]bool

	IconsMap map[string]string

	db *db.DB


}

func NewParser() *Parser {
	x := db.Connect()
	types := x.InitializeStringMap("Types")
	items := x.InitializeStringMap("Items")
	icons := x.InitializeIconsMap()
	mods := x.InitializeModsMap()
	return &Parser{
		ModsMap : mods,
		TypesMap : types,
		SubCategories : make(map[string]bool),
		TopCategories : make(map[string]bool),
		ItemsMap : items,
		IconsMap : icons,
		db : x,
	}

}

func (p *Parser) initData() {

}

func (p *Parser) processStash(stash *api.Stash) {

	for _, item := range stash.Items {
		 p.processItem(&item)
	}

}

func (p *Parser) processItem(item *api.Item){
	if item.FrameType != 1 && item.IsIdentified {
		_type := p.removeSuperior(item.Type)
		p.parseType(_type)
		p.parseIcon(item.Icon, _type, item.DescriptionText)
		p.parseFullName(item.Name, _type, item.FrameType)
	}
	if (item.FrameType <3 && item.IsIdentified && !strings.Contains(item.Type, "Leaguestone")){
		p.ParseMods(item.ImplicitMods, "implicit")
		p.ParseMods(item.ExplicitMods, "explicit")
		p.ParseMods(item.CraftedMods, "crafted")
		p.ParseMods(item.EnchantMods, "enchant")

	}
}

func (p *Parser) ParseMods(mods []string, t string) {
	for _, mod := range mods {
		if mod != "" {
			str := strings.Replace(ModReg.ReplaceAllString(mod, "#"), "# to #", "#", 2)
			_, ok := p.ModsMap[str] ; if !ok{
				//Add to map
				p.ModsMap[str]=t
				go p.db.UpsertMods(str, t)
			}
		}
	}

}


func (p *Parser) parseType(s string){
	if strings.Contains(s, "Superior "){
		s = strings.Replace(s, "Superior ", "", 2)
	}

	_, ok := p.TypesMap[s] ; if !ok{
		go p.db.UpsertString("Types", s)

		//Update new type
		p.TypesMap[s] = true
	}

}

func (p *Parser) parseIcon(icon string, t string, desc string){
	temp := strings.Split(icon, "?")
	if (len(temp) == 1){
		//This should be a flask. We can match with descText
		_, ok := p.IconsMap[desc] ; if !ok{
			go p.db.UpsertIcon(desc,t )
			p.IconsMap[desc] = t
		}
	}else if (len(temp) > 0){
		icon = temp[0]
		_, ok := p.IconsMap[icon] ; if !ok{
			go p.db.UpsertIcon(icon,t )
			p.IconsMap[icon] = t
		}
	}
}

func (p *Parser) parseFullName(n string, t string, frameType api.FrameType){
	var FName string
	switch frameType{
	case 3, 4, 6, 7, 8, 9 :
		if (t != "" && n != ""){
		FName = n + " " + t
		}else {
			FName = n + t
		}
	}
	if (FName != ""){
		_, ok := p.ItemsMap[FName] ; if !ok{
			go p.db.UpsertString("Items", FName)
			//Update new type
			p.ItemsMap[FName] = true
		}

	}


}

func (p *Parser) removeSuperior(s string) string{
	if strings.Contains(s, "Superior "){
		return strings.Replace(s, "Superior ", "", 2)
	}else {
		return s
	}

}

func (p *Parser) UpsertStringDB(coll string, s string){

}

func (p *Parser) updateDB(){
	log.Println("Updating DB")
	p.db.UpdateCollection("Types", p.TypesMap)

}



