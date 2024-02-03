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
	s.Widgets = append(s.Widgets, widget)
	return &s.Widgets[len(s.Widgets)-1]
}

func (w *Widget) AddButtonList() *ButtonList {
	var buttonList ButtonList
	w.ButtonList = &buttonList
	return &buttonList
}

func (b *ButtonList) AddButton() *Button {
	var button Button
	b.Buttons = append(b.Buttons, button)
	return &button
}

func (b *Button) AddOnClick(url string) *OnClick {
	var onClick OnClick
	b.Text = "Someone is outside"
	b.OnClick = &onClick
	return &onClick
}

func (o *OnClick) AddOpenLink(url string) *OpenLink {
	var openLink OpenLink
	openLink.Url = url
	o.OpenLink = &openLink
	return &openLink
}

func CreateChatCard(title string) *ChatCard {
	var chatCard ChatCard
	var cardsv2 CardsV2
	card := CreateCard(title)
	section := card.AddSection("Someone pressed the doorbell")
	section.AddWidget().
		AddButtonList().
		AddButton().
		AddOnClick("https://google.com").
		AddOpenLink("https://google.com")
	cardsv2.CardId = "Doorbell"
	cardsv2.Card = card
	chatCard.CardsV2 = append(chatCard.CardsV2, cardsv2)
	return &chatCard
}
