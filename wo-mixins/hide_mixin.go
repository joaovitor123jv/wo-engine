package womixins

type Hideable interface {
	Hide()
	Show()
	ToggleVisibility()
	IsVisible() bool
	IsHidden() bool
}

type HideMixin struct {
	canRender  bool
	dependants []Hideable
}

func NewHideMixin() HideMixin {
	return HideMixin{
		canRender:  true,
		dependants: []Hideable{},
	}
}

func (h *HideMixin) AddDependant(dependant Hideable) {
	h.dependants = append(h.dependants, dependant)
}

func (h *HideMixin) Hide() {
	if len(h.dependants) > 0 {
		for _, dependant := range h.dependants {
			dependant.Show()
		}
	}

	h.canRender = false
}

func (h *HideMixin) Show() {
	if len(h.dependants) > 0 {
		for _, dependant := range h.dependants {
			dependant.Show()
		}
	}

	h.canRender = true
}

func (h *HideMixin) ToggleVisibility() {
	if h.canRender {
		h.Hide()
	} else {
		h.Show()
	}
}

func (h *HideMixin) IsVisible() bool {
	return h.canRender
}

func (h *HideMixin) IsHidden() bool {
	return !h.canRender
}
