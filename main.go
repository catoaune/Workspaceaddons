package WorkspaceAddOns

func (ra *RenderAction) CreateAction() *NavigationAction {
	var action NavigationAction
	action.Link = nil
	ra.Action = &action
	return ra.Action
}

func (na *NavigationAction) AddNavigation() *Navigation {
	navigation := new(Navigation)
	na.Navigations = append(na.Navigations, *navigation)
	return &na.Navigations[len(na.Navigations)-1]
}

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
	return &c.Sections[len(c.Sections)-1]
}

func (s *Section) AddWidget() *Widget {
	var widget Widget
	s.Widgets = append(s.Widgets, widget)
	return &s.Widgets[len(s.Widgets)-1]
}

func (w *Widget) AddTextInput(name, label, value string) *TextInput {
	var textInput TextInput
	textInput.Name = name
	textInput.Label = &label
	textInput.Value = &value
	w.TextInput = &textInput
	return w.TextInput
}

func (w *Widget) AddTextParagraph(text string) *TextParagraph {
	var textParagraph TextParagraph
	textParagraph.Text = text
	w.TextParagraph = &textParagraph
	return w.TextParagraph
}

func (w *Widget) AddImage(altText, url string) *Image {
	var image Image
	image.AltText = &altText
	image.ImageUrl = url
	w.Image = &image
	return w.Image
}

func (w *Widget) AddButtonList() *ButtonList {
	var buttonList ButtonList
	w.ButtonList = &buttonList
	return &buttonList
}

func (b *ButtonList) AddButton() *Button {
	var button Button
	b.Buttons = append(b.Buttons, button)
	return &b.Buttons[len(b.Buttons)-1]
}

func (b *Button) AddOnClick(label string) *OnClick {
	var onClick OnClick
	b.Text = label
	b.OnClick = &onClick
	return &onClick
}

func (o *OnClick) AddOpenLink(url string) *OpenLink {
	var openLink OpenLink
	openLink.Url = url
	o.OpenLink = &openLink
	return &openLink
}

// Not generated automatically since Action is used in several places with different meaning
type FormAction struct {
	Function string `json:"function,omitempty"`
}

// Wrapper since the payload must start with renderAction when it is a response to a form submit
type RenderActionWrapper struct {
	RenderAction *RenderAction `json:"renderActions,omitempty"`
}

func (ra *RenderActionWrapper) AddRenderAction() *RenderAction {
	var renderAction RenderAction
	ra.RenderAction = &renderAction
	return &renderAction
}

// Helper functions
/*
func (r *RenderAction) CreateAction() *Action {
	var action Action
	r.Action = &action
	return &action
}
*/
/*
func (a *Action) AddNavigation() *Navigation {
	navigation := new(Navigation)
	a.Navigations = append(a.Navigations, *navigation)
	return &a.Navigations[len(a.Navigations)-1]
}
*/

func (a *NavigationAction) AddNotification(text string) {
	var notification Notification
	notification.Text = &text
	a.Notification = &notification
}

func (n *Navigation) AddCard() *Card {
	var card Card
	n.PushCard = &card
	return &card
}

func (c *Card) CreateHeader(imageType string) {
	var h CardHeader
	h.ImageType = (*CardHeaderImageType)(&imageType)
	c.Header = &h
}

func (b *ButtonList) AddSubmitButton(text, url string) {
	var button Button
	button.Text = text
	formAction := new(Action)
	formAction.Function = url
	var onClick OnClick
	onClick.Action = formAction
	button.OnClick = &onClick
	//	button.OnClick.Action = formAction
	b.Buttons = append(b.Buttons, button)
}
