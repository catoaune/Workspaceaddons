package WorkspaceAddOns

type ChatCard struct {
	CardsV2 []CardsV2 `json:"cardsV2"`
}

type CardsV2 struct {
	CardId string `json:"cardId"`
	Card   Card   `json:"card"`
}
