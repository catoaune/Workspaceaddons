package WorkspaceAddOns

func CreateCard(title string) *Card {
	var card Card
	var header CardHeader
	header.Title = title
	card.Header = &header
	return &card
}

func (c *Card) AddSection(title string) *Section {
	var section Section
	section.Header = &title
	c.Sections = append(c.Sections, section)
	return &section
}

func (s *Section) AddWidget() *Widget {
	var widget Widget

	return &widget
}

func CreateChatCard(title string) *ChatCard {
	var chatCard ChatCard
	var cardsv2 CardsV2
	card := CreateCard(title)
	section := card.AddSection("Someone pressed the doorbell")
	cardsv2.CardId = "Doorbell"
	cardsv2.Card = card
	chatCard.CardsV2 = append(chatCard.CardsV2, cardsv2)
	return &chatCard
}
