package twui

import (
	"html/template"
	"testing"

	"github.com/kyoto-framework/kyoto"
)

type TestAppUIPage struct {
	AppUINavNavbar       kyoto.Component
	AppUINavSidebar      kyoto.Component
	AppUINavBreadcrumbs  []kyoto.Component
	AppUINavPagination   kyoto.Component
	AppUIHeadPage        []kyoto.Component
	AppUIOverlayModal    kyoto.Component
	AppUIOverlaySideover kyoto.Component
	AppUIDataStats       kyoto.Component
	AppUIFormLayout      kyoto.Component
	AppUIFormSignIn      kyoto.Component
	AppUIListTable       kyoto.Component
	AppUIListFeed        kyoto.Component
}

func (p *TestAppUIPage) Template() *template.Template {
	return template.Must(template.New("appui_test.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *TestAppUIPage) Init() {
	p.AppUINavNavbar = kyoto.RegC(p, &AppUINavNavbar{
		Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/kyoto-framework/kyoto/master/.docs/.vuepress/public/kyoto.svg" class="h-8 w-8 scale-150" />
		</a>`),
		Links: []AppUINavNavbarLink{
			{Text: "Dashboard", Href: "#"},
			{Text: "Team", Href: "#"},
			{Text: "Projects", Href: "#"},
			{Text: "Calendar", Href: "#"},
		},
		Search: AppUINavNavbarSearch{
			Enabled: true,
			Action:  "/",
			Method:  "GET",
		},
		Notifications: AppUINavNavbarNotifications{
			Enabled: true,
			Href:    "#",
			Counter: 4,
		},
		Profile: AppUINavNavbarProfile{
			Enabled: true,
			Avatar:  `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
			// Avatar: `<div class="p-1 rounded-full bg-gray-900"><svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg></div>`,
			Links: []AppUINavNavbarLink{
				{Text: "Your Profile", Href: "#"},
				{Text: "Settings", Href: "#"},
				{Text: "Sign Out", Href: "#"},
			},
		},
	})
	p.AppUINavSidebar = kyoto.RegC(p, &AppUINavSidebar{
		Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/kyoto-framework/kyoto/master/.docs/.vuepress/public/kyoto.svg" class="mx-auto h-16 w-16 scale-150" />
		</a>`),
		Search: AppUINavSidebarSearch{
			Enabled: true,
		},
		Links: []AppUINavSidebarLink{
			{
				Text: "Dashboard",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
			},
			{
				Separator: true,
			},
			{
				Text: "Team",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path></svg>`,
			},
			{
				Text: "Projects",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v-1m4 1v-3m4 3V8M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path></svg>`,
			},
			{
				Text:    "Calendar",
				Href:    "#",
				Counter: 2,
				Icon:    `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
			},
		},
		Profile: AppUINavSidebarProfile{
			Enabled:     true,
			Avatar:      `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
			Username:    "Tom Cook",
			Text:        "View profile",
			SignoutHref: "/signout",
		},
	})
	bclinks := []AppUINavBreadcrumbsLink{
		{
			Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path></svg>`,
			Href: "/",
		},
		{
			Text: "Projects",
			Href: "#",
		},
		{
			Text: "UIKit",
			Href: "#",
		},
		{
			Text: "TWUI",
			Href: "#",
		},
	}
	p.AppUINavBreadcrumbs = append(p.AppUINavBreadcrumbs, kyoto.RegC(p, &AppUINavBreadcrumbs{
		Links:     bclinks,
		Separator: "slash",
	}))
	p.AppUINavBreadcrumbs = append(p.AppUINavBreadcrumbs, kyoto.RegC(p, &AppUINavBreadcrumbs{
		Links: bclinks,
	}))
	p.AppUINavBreadcrumbs = append(p.AppUINavBreadcrumbs, kyoto.RegC(p, &AppUINavBreadcrumbs{
		Links:     bclinks,
		Contained: true,
		Rounded:   true,
	}))
	p.AppUINavPagination = kyoto.RegC(p, &AppUINavPagination{
		Page:       97,
		TotalPages: 100,
	})
	p.AppUIHeadPage = append(p.AppUIHeadPage, kyoto.RegC(p, &AppUIHeadPage{
		Breadcrumbs: kyoto.RegC(p, &AppUINavBreadcrumbs{
			Links: []AppUINavBreadcrumbsLink{
				{Text: "Jobs", Href: "#"},
				{Text: "Engineering", Href: "#"},
				{Text: "Back End Developer", Href: "#"},
			},
		}),
		Title:    "Back End Developer",
		Subtitle: "Applied for Front End Developer on August 25, 2020",
		Actions: []AppUIHeadPageAction{
			{
				Icon: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>`,
				Text: "Edit",
				Href: "#",
			},
			{
				Icon: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path></svg>`,
				Text: "View",
				Href: "#",
			},
			{
				Icon:    `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>`,
				Text:    "Publish",
				Href:    "#",
				Primary: true,
			},
		},
		Meta: []AppUIHeadPageMeta{
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path></svg>`,
				Text: "Full-time",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>`,
				Text: "Remote",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`,
				Text: "$120k - $140k",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
				Text: "Closing on January 9, 2020",
			},
		},
	}))
	p.AppUIHeadPage = append(p.AppUIHeadPage, kyoto.RegC(p, &AppUIHeadPage{
		Banner: "https://source.unsplash.com/random/1024x400",
		Avatar: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
		Breadcrumbs: kyoto.RegC(p, &AppUINavBreadcrumbs{
			Links: []AppUINavBreadcrumbsLink{
				{Text: "Jobs", Href: "#"},
				{Text: "Engineering", Href: "#"},
				{Text: "Back End Developer", Href: "#"},
			},
		}),
		Title:    "Back End Developer",
		Subtitle: "Applied for Front End Developer on August 25, 2020",
		Actions: []AppUIHeadPageAction{
			{
				Icon: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>`,
				Text: "Edit",
				Href: "#",
			},
			{
				Icon: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path></svg>`,
				Text: "View",
				Href: "#",
			},
			{
				Icon:    `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>`,
				Text:    "Publish",
				Href:    "#",
				Primary: true,
			},
		},
		Meta: []AppUIHeadPageMeta{
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path></svg>`,
				Text: "Full-time",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>`,
				Text: "Remote",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`,
				Text: "$120k - $140k",
			},
			{
				Icon: `<svg class="flex-shrink-0 mr-1.5 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
				Text: "Closing on January 9, 2020",
			},
		},
	}))
	p.AppUIOverlayModal = kyoto.RegC(p, &AppUIOverlayModal{
		ID: "TestModal",
		Content: `
			<div class="p-4">
				<div class="text-xl">Test Modal</div>
				<button onclick="Action(this, 'Close')" class="mt-2 py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Close</button>
			</div>`,
		MaxWidth: "max-w-screen-sm",
	})
	p.AppUIOverlaySideover = kyoto.RegC(p, &AppUIOverlaySideover{
		ID: "TestSideover",
		Content: `
			<div class="p-4">
				<div class="text-xl">Test Side-over</div>
				<button onclick="Action(this, 'Close')" class="mt-2 py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Close</button>
			</div>`,
		MaxWidth: "max-w-md",
	})
	p.AppUIDataStats = kyoto.RegC(p, &AppUIDataStats{
		Items: []AppUIDataStatsItem{
			{
				Image: `
					<div class="p-2 bg-indigo-500 text-white rounded-lg">
						<svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg>
					</div>
				`,
				Title:  "Total Subscribers",
				Value:  71897,
				Change: 122,
			},
			{
				Title:        "Avg. Open Rate",
				Value:        58.16,
				ValueSymbol:  "%",
				Change:       5.4,
				ChangeSymbol: "%",
			},
			{
				Title:        "Avg. Click Rate",
				Value:        24.57,
				ValueSymbol:  "%",
				Change:       -3.2,
				ChangeSymbol: "%",
			},
		},
	})
	p.AppUIFormLayout = kyoto.RegC(p, &AppUIFormLayout{
		Action: "/",
		Method: "GET",
		Sections: []AppUIFormLayoutSection{
			{
				Title:       "Profile",
				Description: "This information will be displayed publicly",
				Fields: []AppUIFormLayoutField{
					{
						Type:        FORM_INPUT,
						Name:        "Website",
						Label:       "Website",
						Preffix:     "https://",
						Placeholder: "example.com",
					},
					{
						Type:        FORM_TEXTAREA,
						Name:        "About",
						Label:       "About",
						Placeholder: "Here is some information about me ...",
						Description: "Brief description for your profile",
					},
					{
						Type:    FORM_RADIO,
						Name:    "UnderEighteen",
						Label:   "I'm under 18",
						Value:   "true",
						Checked: true,
					},
					{
						Type:  FORM_RADIO,
						Name:  "UnderEighteen",
						Label: "I'm above 18",
						Value: "false",
					},
					{
						Type:  FORM_CHECKBOX,
						Name:  "MarketingEnabled",
						Label: "Enable marketing emails",
					},
				},
			},
			{
				Title:       "Personal Information",
				Description: "Use permanent address where you can receive mail",
				Fields: []AppUIFormLayoutField{
					{
						Type:  FORM_INPUT,
						Name:  "FirstName",
						Label: "First name",
						Width: "w-1/2",
					},
					{
						Type:  FORM_INPUT,
						Name:  "LastName",
						Label: "Last name",
						Width: "w-1/2",
					},
					{
						Type:  FORM_INPUT,
						Name:  "Email",
						Label: "Email address",
					},
					{
						Type:  FORM_SELECT,
						Name:  "Country",
						Label: "Country",
						Options: []AppUIFormLayoutSelectOption{
							{Label: "United States", Value: "US"},
							{Label: "Canada", Value: "CA"},
							{Label: "Mexico", Value: "MX"},
						},
					},
					{
						Type:  FORM_INPUT,
						Name:  "City",
						Label: "City",
						Width: "w-1/3",
					},
					{
						Type:  FORM_INPUT,
						Name:  "State",
						Label: "State / Province",
						Width: "w-1/3",
					},
					{
						Type:  FORM_INPUT,
						Name:  "Zip",
						Label: "Zip / Postal code",
						Width: "w-1/3",
					},
				},
			},
		},
		Actions: []AppUIFormLayoutAction{
			{Text: "Reset", Value: "Reset"},
			{Text: "Submit", Value: "Submit", Primary: true},
		},
	})
	rows := []map[string]template.HTML{}
	for i := 0; i < 100; i++ {
		rows = append(rows, map[string]template.HTML{
			"Name":   "Jane Cooper",
			"Title":  "Regional Paradigm Technician",
			"Status": `<span class="px-2 py-1 rounded-full bg-green-200 text-green-800">active</span>`,
			"Role":   "Admin",
			"Actions": `
				<div class="w-full flex justify-end gap-4">
					<a href="#" class="text-blue-500 hover:underline">Edit</a>
					<a href="#" class="text-red-500 hover:underline">Delete</a>
				</div>
			`,
		})
	}
	p.AppUIListTable = kyoto.RegC(p, &AppUIListTable{
		Columns: []AppUIListTableColumn{
			{Label: "Name", Key: "Name"},
			{Label: "Title", Key: "Title"},
			{Label: "Status", Key: "Status"},
			{Label: "Role", Key: "Role"},
			{Label: "", Key: "Actions"},
		},
		Rows: rows,
	})
	p.AppUIListFeed = kyoto.RegC(p, &AppUIListFeed{
		Entries: []AppUIListFeedEntry{
			{
				Icon:        `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
				Title:       "Eduardo Benz",
				Subtitle:    "Commented",
				Time:        "1d ago",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			},
			{
				Icon: `
				<div class="p-1 rounded-full bg-blue-500 text-white">
					<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path></svg>
				</div>
			`,
				Title:    "Eduardo Benz",
				Time:     "2d ago",
				Subtitle: "Friend request",
			},
			{
				Icon: `
				<div class="p-1 rounded-full bg-gray-500 text-white">
					<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path></svg>
				</div>
			`,
				RAW: `
				<div>
					<span class="font-semibold">Oleg Nazarov</span>
					added tags
					<span class="px-2 border rounded-full font-semibold">
						<span class="text-red-500">•</span>
						Bug
					</span>
					<span class="ml-1 px-2 border rounded-full font-semibold">
						<span class="text-indigo-500">•</span>
						Accessibility
					</span>
					<span class="ml-1">3d ago</span>
				</div>
			`,
			},
		},
	})
}

func TestAppUI(t *testing.T) {
	serve(&TestAppUIPage{})
}
