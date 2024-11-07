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
	AfterHide  func()
	AfterShow  func()
	BeforeHide func()
	BeforeShow func()
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
	if h.BeforeHide != nil {
		h.BeforeHide()
	}

	if len(h.dependants) > 0 {
		for _, dependant := range h.dependants {
			dependant.Show()
		}
	}

	h.canRender = false

	if h.AfterHide != nil {
		h.AfterHide()
	}
}

func (h *HideMixin) Show() {
	if h.BeforeShow != nil {
		h.BeforeShow()
	}

	if len(h.dependants) > 0 {
		for _, dependant := range h.dependants {
			dependant.Show()
		}
	}

	h.canRender = true

	if h.AfterShow != nil {
		h.AfterShow()
	}
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
