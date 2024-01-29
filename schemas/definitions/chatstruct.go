package WorkspaceAddOns

type ChatCard struct {
	CardsV2 []CardsV2
}

type CardsV2 struct {
	CardId string
	Card   Card
}
