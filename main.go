package WorkspaceAddOns

func createCard(title string) *Card {
	var card Card
	var header CardHeader
	header.Title = title
	card.Header = &header
	return &card
}

func (c *Card) addSection(title string) *Section {
	var section Section
	section.Header = &title
	c.Sections = append(c.Sections, section)
	return &section
}

func (s *Section) addWidget() *Widget {
	var widget Widget

	return &widget
}
