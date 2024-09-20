package partials

import "github.com/maddalax/htmgo/framework/h"

type NavItem struct {
	Name string
	Url  string
}

func NavBar() *h.Element {

	star := h.Raw(`
		<a 
		class="github-button" 
		href="https://github.com/maddalax/htmgo" 
		data-color-scheme="no-preference: light; light: light; dark: dark;" 
		data-icon="octicon-star" 
		data-size="large" 
		data-show-count="true"
		aria-label="Star maddalax/htmgo on GitHub">Star</a>
	`)

	navItems := []NavItem{
		{Name: "Docs", Url: "/docs"},
		{Name: "Examples", Url: "/examples"},
	}

	prelease := h.Div(h.Class("bg-yellow-200 text-yellow-800 text-center p-2"),
		h.Text("This is a prerelease version and generally should not be used at this time. Watch on github for the stable release!"),
	)

	return h.Div(
		prelease,
		h.Nav(
			h.Class("bg-neutral-100 border border-b-slate-300 p-4 md:p-3"),
			h.Div(
				h.Class("max-w-[95%] md:max-w-prose mx-auto"),
				h.Div(
					h.Class("flex justify-between items-center"),
					h.Div(
						h.Class("flex items-center"),
						h.A(
							h.Class("text-2xl"),
							h.Href("/"),
							h.Text("htmgo"),
						)),
					h.Div(
						h.Class("flex gap-4 items-center"),
						h.List(navItems, func(item NavItem, index int) *h.Element {
							return h.Div(
								h.Class("flex items-center"),
								h.A(
									h.Class(""),
									h.Href(item.Url),
									h.Text(item.Name),
								),
							)
						}),
						h.Div(h.Class("ml-2"), star),
					),
				),
			),
		),
	)
}
