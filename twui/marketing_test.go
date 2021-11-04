package twui

import (
	"html/template"
	"testing"

	"github.com/yuriizinets/kyoto"
)

type TestMarketingPage struct {
	MarketingSectionPricing []kyoto.Component
	MarketingSectionCTA     []kyoto.Component
	MarketingSectionHero    kyoto.Component
	MarketingSectionHero2   kyoto.Component
	MarketingSectionHero3   kyoto.Component
	MarketingSectionFooter  kyoto.Component
	MarketingSectionFooter2 kyoto.Component
	MarketingSectionFooter3 kyoto.Component
	MarketingSectionFooter4 kyoto.Component
	MarketingSectionFooter5 kyoto.Component
}

func (p *TestMarketingPage) Template() *template.Template {
	return template.Must(template.New("marketing_test.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *TestMarketingPage) Init() {
	p.MarketingSectionPricing = append(p.MarketingSectionPricing, kyoto.RegC(p, &MarketingSectionPricing{
		Title:       "Pricing Plans",
		Description: "Start building for free, then add a site plan to go live. Account plans unlock additional features.",
		CardBlock: []MarketingSectionPricingCardBlock{
			{
				Title:         "Monthly billing",
				FeaturesTitle: "What's included",
				Cards: []MarketingSectionPricingCard{
					{
						Title:       "Hobby",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    12,
							Currency: "$",
							Suffix:   "/mo",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Hobby",
							Href: "/",
						},
					},
					{
						Title:       "Freelancer",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    24,
							Currency: "$",
							Suffix:   "/mo",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
							"Donec mauris sit in eu tincidunt etiam.",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Freelancer",
							Href: "/",
						},
					},
					{
						Title:       "Startup",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    32,
							Currency: "$",
							Suffix:   "/mo",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
							"Donec mauris sit in eu tincidunt etiam.",
							"Faucibus volutpat magna.",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Startup",
							Href: "/",
						},
					},
					{
						Title:       "Enterprise",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    48,
							Currency: "$",
							Suffix:   "/mo",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
							"Donec mauris sit in eu tincidunt etiam.",
							"Faucibus volutpat magna.",
							"Id sed tellus in varius quisque.",
							"Risus egestas faucibus",
							"Risus cursus ullamcorper",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Enterprise",
							Href: "/",
						},
					},
				},
			},
			{
				Title:         "Yearly billing",
				Inline:        true,
				FeaturesTitle: "What's included",
				Cards: []MarketingSectionPricingCard{
					{
						Title:       "Startup",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    335,
							Currency: "$",
							Suffix:   "/year",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
							"Donec mauris sit in eu tincidunt etiam.",
							"Faucibus volutpat magna.",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Startup",
							Href: "/",
						},
					},
					{
						Title:       "Enterprise",
						Description: "All the basics for starting a new business",
						Price: MarketingSectionPricingCardPrice{
							Value:    500,
							Currency: "$",
							Suffix:   "/year",
						},
						Features: []string{
							"Potenti felis, in cras at at ligula nunc.",
							"Orci neque eget pellentesque",
							"Donec mauris sit in eu tincidunt etiam.",
							"Faucibus volutpat magna.",
							"Id sed tellus in varius quisque.",
							"Risus egestas faucibus",
							"Risus cursus ullamcorper",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Buy Enterprise",
							Href: "/",
						},
					},
				},
			},
		},
	}))
	p.MarketingSectionPricing = append(p.MarketingSectionPricing, kyoto.RegC(p, &MarketingSectionPricing{
		Title:       "Simple no-tricks pricing",
		Description: "If you're not satisfied, contact us within the first 14 days and we'll send you a full refund",
		CardBlock: []MarketingSectionPricingCardBlock{
			{
				Title:         "Yearly billing",
				FeaturesTitle: "What's included",
				Cards: []MarketingSectionPricingCard{
					{
						Title:       "Lifetime Membership",
						Description: "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Voluptate officia, reprehenderit reiciendis quas alias doloribus, sed nobis provident.",
						PriceTitle:  "Pay once, own it forever",
						Price: MarketingSectionPricingCardPrice{
							Value:    349,
							Currency: "$",
							Suffix:   "USD",
						},
						AdditionalInformation: MarketingSectionPricingLink{
							Text: "Learn about our membership policy",
							Href: "/",
						},
						CTALink: MarketingSectionPricingLink{
							Text: "Get Access",
							Href: "/",
						},
						Features: []string{
							"Private forum access",
							"Member resources",
							"Entry to annual conference",
							"Official member t-shirt",
						},
						FreeSample: MarketingSectionPricingCardSample{
							Enabled: true,
							Sample:  template.HTML(`<a href="/" class="text-sm">Get a free sample <span class="text-gray-500">(20 MB)</span></a>`),
						},
					},
				},
			},
		},
	}))
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Content: MarketingSectionCTAContent{
			Title: MarketingSectionCTAContentText{
				Text: "Ready to dive in?",
			},
			Subtitle: MarketingSectionCTAContentText{
				Color: "text-indigo-500",
				Text:  "Start your free trial today.",
			},
		},
		Links: []MarketingSectionCTALink{
			{
				Text:    "Get started",
				Primary: true,
				Href:    "/",
			},
			{
				Text: "Learn more",
				Href: "/",
			},
		},
	})
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Background: "bg-indigo-700",
		Content: MarketingSectionCTAContent{
			Centered: true,
			Title: MarketingSectionCTAContentText{
				Text:     "Boost your productivity.",
				Color:    "text-white",
				Centered: true,
			},
			Subtitle: MarketingSectionCTAContentText{
				Text:     "Start using TWUI today.",
				Color:    "text-white",
				Centered: true,
			},
			Description: MarketingSectionCTAContentText{
				Text:     "Ac eusismod vel sit maecenas id pellentesque eu sed consectecur. Malesuada adipiscing sagittis vel nulla nec.",
				Color:    "text-gray-300",
				Centered: true,
			},
			LinksCentered: true,
			Links: []MarketingSectionCTALink{
				{
					Text: "Sign up for free",
					Href: "/",
				},
			},
		},
	})
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Content: MarketingSectionCTAContent{
			Centered: true,
			Title: MarketingSectionCTAContentText{
				Text:     "Ready to dive in?",
				Centered: true,
			},
			Subtitle: MarketingSectionCTAContentText{
				Text:     "Start your free trial today.",
				Centered: true,
			},
			LinksCentered: true,
			Links: []MarketingSectionCTALink{
				{
					Text:    "Get started",
					Primary: true,
					Href:    "/",
				},
				{
					Text: "Learn more",
					Href: "/",
				},
			},
		},
	})
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Background: "bg-indigo-50",
		Content: MarketingSectionCTAContent{
			Title: MarketingSectionCTAContentText{
				Text: "Ready to dive in?",
			},
			Subtitle: MarketingSectionCTAContentText{
				Color: "text-indigo-500",
				Text:  "Start your free trial today.",
			},
		},
		Links: []MarketingSectionCTALink{
			{
				Text: "Get started",
				Href: "/",
			},
		},
	})
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Content: MarketingSectionCTAContent{
			Title: MarketingSectionCTAContentText{
				Text: "Ready to dive in?",
			},
			Subtitle: MarketingSectionCTAContentText{
				Text:  "Start your free trial today.",
				Color: "text-indigo-500",
			},
			Links: []MarketingSectionCTALink{
				{
					Text:    "Get started",
					Primary: true,
					Href:    "/",
				},
				{
					Text: "Learn more",
					Href: "/",
				},
			},
		},
	})
	p.MarketingSectionCTA = append(p.MarketingSectionCTA, &MarketingSectionCTA{
		Background: "bg-indigo-700",
		Content: MarketingSectionCTAContent{
			Title: MarketingSectionCTAContentText{
				Text:  "Ready to dive in?",
				Color: "text-white",
			},
			Subtitle: MarketingSectionCTAContentText{
				Text:  "Start your free trial today.",
				Color: "text-white",
			},
			Description: MarketingSectionCTAContentText{
				Text:  "Ac eusismod vel sit maecenas id pellentesque eu sed consectecur. Malesuada adipiscing sagittis vel nulla nec.",
				Color: "text-gray-300",
			},
			Links: []MarketingSectionCTALink{
				{
					Text: "Sign up for free",
					Href: "/",
				},
			},
		},
		Image: "https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80",
	})
	p.MarketingSectionHero = kyoto.RegC(p, &MarketingSectionHero{
		Navbar: MarketingSectionHeroNavbar{
			Enabled: true,
			Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/yuriizinets/kyoto/master/.docs/.vuepress/public/kyoto.svg" class="h-8 w-8 scale-150" />
		</a>`),
			Links: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Product",
				},
				{
					Href: "#",
					Text: "Features",
				},
				{
					Href: "#",
					Text: "Marketplace",
				},
				{
					Href: "#",
					Text: "Company",
				},
			},
			AutorizationLinks: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Log in",
				},
			},
		},
		Body: MarketingSectionHeroBody{
			Image: MarketingSectionHeroBodyImage{
				Src:    "https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80",
				Cover:  true,
				Angled: true,
			},
			Title:    "Data to enrich your",
			Subtitle: "online business",
			Text:     "Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.",
			Actions: []MarketingSectionHeroBodyActions{
				{
					Text:    "Get started",
					Href:    "#",
					Primary: true,
				},
				{
					Text: "Live demo",
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionHero2 = kyoto.RegC(p, &MarketingSectionHero{
		BackgroundImage: "https://wallpaperaccess.com/full/1093402.jpg",
		Navbar: MarketingSectionHeroNavbar{
			Centered: true,
			Enabled:  true,
			Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/yuriizinets/kyoto/master/.docs/.vuepress/public/kyoto.svg" class="h-12 w-12 scale-150" />
		</a>`),
			Links: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Product",
				},
				{
					Href: "#",
					Text: "Features",
				},
				{
					Href: "#",
					Text: "Marketplace",
				},
				{
					Href: "#",
					Text: "Company",
				},
			},
			AutorizationLinks: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Log in",
				},
			},
		},
		Body: MarketingSectionHeroBody{
			Title:    "Data to enrich your",
			Subtitle: "online business",
			Text:     "Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.",
			Centered: true,
			Actions: []MarketingSectionHeroBodyActions{
				{
					Text:    "Get started",
					Href:    "#",
					Primary: true,
				},
				{
					Text: "Live demo",
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionHero3 = kyoto.RegC(p, &MarketingSectionHero{
		Navbar: MarketingSectionHeroNavbar{
			Enabled: true,
			Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/yuriizinets/kyoto/master/.docs/.vuepress/public/kyoto.svg" class="h-8 w-8 scale-150" />
		</a>`),
			Links: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Solutions",
					Links: []MarketingSectionHeroNavbarLink{
						{
							Text: "Marketing",
							Href: "#",
						},
						{
							Text: "Analytics",
							Href: "#",
						},
						{
							Text: "Commerce",
							Href: "#",
						},
						{
							Text: "Insight",
							Href: "#",
						},
					},
				},
				{
					Href: "#",
					Text: "Pricing",
				},
				{
					Href: "#",
					Text: "Docs",
				},
				{
					Href: "#",
					Text: "More",
					Links: []MarketingSectionHeroNavbarLink{
						{
							Text: "About",
							Href: "#",
						},
						{
							Text: "Blog",
							Href: "#",
						},
						{
							Text: "Jobs",
							Href: "#",
						},
						{
							Text: "Press",
							Href: "#",
						},
						{
							Text: "Partners",
							Href: "#",
						},
					},
				},
			},
			AutorizationLinks: []MarketingSectionHeroNavbarLink{
				{
					Href: "#",
					Text: "Log in",
				},
			},
		},
		Body: MarketingSectionHeroBody{
			Image: MarketingSectionHeroBodyImage{
				Src:   "https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80",
				Cover: true,
			},
			Title:    "Data to enrich your",
			Subtitle: "online business",
			Text:     "Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.",
			Actions: []MarketingSectionHeroBodyActions{
				{
					Text:    "Get started",
					Href:    "#",
					Primary: true,
				},
				{
					Text: "Live demo",
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionFooter = kyoto.RegC(p, &MarketingSectionFooter{
		Main: MarketingSectionFooterMain{
			Enabled: true,
			Blocks: []MarketingSectionFooterMainBlock{
				{
					Title: "Solutions",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Marketing",
							Href: "#",
						},
						{
							Text: "Analytics",
							Href: "#",
						},
						{
							Text: "Commerce",
							Href: "#",
						},
						{
							Text: "Insight",
							Href: "#",
						},
					},
				},
				{
					Title: "Support",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Pricing",
							Href: "#",
						},
						{
							Text: "Documentation",
							Href: "#",
						},
						{
							Text: "Guides",
							Href: "#",
						},
						{
							Text: "API Status",
							Href: "#",
						},
					},
				},
				{
					Title: "Company",
					Links: []MarketingSectionFooterLink{
						{
							Text: "About",
							Href: "#",
						},
						{
							Text: "Blog",
							Href: "#",
						},
						{
							Text: "Jobs",
							Href: "#",
						},
						{
							Text: "Press",
							Href: "#",
						},
						{
							Text: "Partners",
							Href: "#",
						},
					},
				},
				{
					Title: "Legal",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Claim",
							Href: "#",
						},
						{
							Text: "Privacy",
							Href: "#",
						},
						{
							Text: "Terms",
							Href: "#",
						},
					},
				},
			},
			Company: MarketingSectionFooterMainCompany{
				Enabled: true,
				Logo:    template.HTML(`<a href="/" class="text-2xl font-bold">TWUI</a>`),
				Text:    "Making the world a better place through construction elegant hierarchies",
				Socials: []MarketingSectionFooterSocial{
					{
						Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 12.067C0 18.033 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067z" fill="currentColor"/></g></svg>`,
						Href: "#",
					},
					{
						Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511zm8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379zm-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986zM8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996zm10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89z" fill="currentColor"/></g></svg>`,
						Href: "#",
					},
					{
						Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41l.002-.003z" fill="currentColor"/></g></svg>`,
						Href: "#",
					},
					{
						Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" fill="currentColor"/></g></svg>`,
						Href: "#",
					},
					{
						Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M512 0q-52 0-101.5 10T315 40l-6 3Q171 103 85.5 230T0 512q0 104 40.5 198.5T150 874t163.5 109.5T512 1024t198.5-40.5T874 874t109.5-163.5T1024 512q0-83-26-161t-73.5-141T814 99.5T673 26T512 0zm448 512q0 2-.5 6t-.5 5q-155-42-326-4q-5-12-10.5-23.5t-11-23.5t-11.5-24t-13-24q0-1-1.5-3.5T583 416q110-42 172.5-96T846 214q113 128 114 298zM798 167q-6 14-14.5 28t-28 36t-44.5 42.5t-66.5 44T553 359Q459 188 378 85q65-21 134-21q53 0 104.5 12.5T713 112t85 55zm-482-57q18 21 39 50t59 89.5T491 380q-70 21-140 33.5t-125.5 15t-90 2T72 428q20-105 85-188.5T316 110zM64 512q0-9 1-21q40 5 88 5q183 0 369-59q2 4 4.5 8.5t4.5 8.5q10 20 20 40.5t19 41.5q-8 2-16.5 5l-17 6l-16.5 6.5l-16.5 7L486 568q-64 29-118 65.5t-89.5 73t-56 62.5t-34.5 51q-59-61-91.5-140.5T64 512zm172 352q10-18 25.5-39.5T309 767t85.5-74T513 626q13-6 27-11.5t27.5-10.5t27.5-9q67 172 88 331q-40 16-83 25t-88 9q-154 0-276-96zm508 31q-23-154-86-316q51-10 102.5-12t91 2.5t62 9T953 587q-11 65-40 123t-72 105t-97 80z" fill="currentColor"/></svg>`,
						Href: "#",
					},
				},
			},
		},
		Simple: MarketingSectionFooterSimple{
			Enabled:  true,
			Text:     "© 2020 Workflow, Inc. All rights reserved",
			Centered: true,
		},
	})
	p.MarketingSectionFooter2 = kyoto.RegC(p, &MarketingSectionFooter{
		Dark: true,
		Main: MarketingSectionFooterMain{
			Enabled: true,
			Reverse: true,
			Blocks: []MarketingSectionFooterMainBlock{
				{
					Title: "Solutions",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Marketing",
							Href: "#",
						},
						{
							Text: "Analytics",
							Href: "#",
						},
						{
							Text: "Commerce",
							Href: "#",
						},
						{
							Text: "Insight",
							Href: "#",
						},
					},
				},
				{
					Title: "Support",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Pricing",
							Href: "#",
						},
						{
							Text: "Documentation",
							Href: "#",
						},
						{
							Text: "Guides",
							Href: "#",
						},
						{
							Text: "API Status",
							Href: "#",
						},
					},
				},
				{
					Title: "Company",
					Links: []MarketingSectionFooterLink{
						{
							Text: "About",
							Href: "#",
						},
						{
							Text: "Blog",
							Href: "#",
						},
						{
							Text: "Jobs",
							Href: "#",
						},
						{
							Text: "Press",
							Href: "#",
						},
						{
							Text: "Partners",
							Href: "#",
						},
					},
				},
				{
					Title: "Legal",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Claim",
							Href: "#",
						},
						{
							Text: "Privacy",
							Href: "#",
						},
						{
							Text: "Terms",
							Href: "#",
						},
					},
				},
			},
			Localization: MarketingSectionFooterMainLocalization{
				Enabled: true,
				Title:   "Language & Currency",
				Fields: []MarketingSectionFooterSelect{
					{
						Name:  "Language",
						Label: "Language",
						Options: []AppUIFormLayoutSelectOption{
							{Label: "English", Value: "US"},
							{Label: "Russian", Value: "RU"},
							{Label: "Spain", Value: "ES"},
							{Label: "Chinese", Value: "ZH"},
						},
					},
					{
						Name:  "Currency",
						Label: "Currency",
						Options: []AppUIFormLayoutSelectOption{
							{Label: "AUD", Value: "AUD"},
							{Label: "USD", Value: "USD"},
							{Label: "ZL", Value: "ZL"},
							{Label: "UAH", Value: "UAH"},
						},
					},
				},
			},
		},
		Newsletter: MarketingSecionsFooterNewsletter{
			Enabled: true,
			Action:  "/",
			Method:  "GET",
			Title:   "Subscribe to our newsletter",
			Text:    "The latest news, articles, and resources, sent to your inbox weekly.",
			Input: MarketingSectionFooterInput{
				Name:        "Email",
				Label:       "Email address",
				Placeholder: "Enter your email",
			},
			Button: MarketingSectionFooterAction{
				Text:    "Subscribe",
				Value:   "Subscribe",
				Primary: true,
			},
		},
		Simple: MarketingSectionFooterSimple{
			Enabled: true,
			Text:    "© 2020 Workflow, Inc. All rights reserved",
			Socials: []MarketingSectionFooterSocial{
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 12.067C0 18.033 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511zm8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379zm-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986zM8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996zm10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41l.002-.003z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M512 0q-52 0-101.5 10T315 40l-6 3Q171 103 85.5 230T0 512q0 104 40.5 198.5T150 874t163.5 109.5T512 1024t198.5-40.5T874 874t109.5-163.5T1024 512q0-83-26-161t-73.5-141T814 99.5T673 26T512 0zm448 512q0 2-.5 6t-.5 5q-155-42-326-4q-5-12-10.5-23.5t-11-23.5t-11.5-24t-13-24q0-1-1.5-3.5T583 416q110-42 172.5-96T846 214q113 128 114 298zM798 167q-6 14-14.5 28t-28 36t-44.5 42.5t-66.5 44T553 359Q459 188 378 85q65-21 134-21q53 0 104.5 12.5T713 112t85 55zm-482-57q18 21 39 50t59 89.5T491 380q-70 21-140 33.5t-125.5 15t-90 2T72 428q20-105 85-188.5T316 110zM64 512q0-9 1-21q40 5 88 5q183 0 369-59q2 4 4.5 8.5t4.5 8.5q10 20 20 40.5t19 41.5q-8 2-16.5 5l-17 6l-16.5 6.5l-16.5 7L486 568q-64 29-118 65.5t-89.5 73t-56 62.5t-34.5 51q-59-61-91.5-140.5T64 512zm172 352q10-18 25.5-39.5T309 767t85.5-74T513 626q13-6 27-11.5t27.5-10.5t27.5-9q67 172 88 331q-40 16-83 25t-88 9q-154 0-276-96zm508 31q-23-154-86-316q51-10 102.5-12t91 2.5t62 9T953 587q-11 65-40 123t-72 105t-97 80z" fill="currentColor"/></svg>`,
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionFooter3 = kyoto.RegC(p, &MarketingSectionFooter{
		Main: MarketingSectionFooterMain{
			Enabled: true,
			Reverse: true,
			Blocks: []MarketingSectionFooterMainBlock{
				{
					Title: "Solutions",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Marketing",
							Href: "#",
						},
						{
							Text: "Analytics",
							Href: "#",
						},
						{
							Text: "Commerce",
							Href: "#",
						},
						{
							Text: "Insight",
							Href: "#",
						},
					},
				},
				{
					Title: "Support",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Pricing",
							Href: "#",
						},
						{
							Text: "Documentation",
							Href: "#",
						},
						{
							Text: "Guides",
							Href: "#",
						},
						{
							Text: "API Status",
							Href: "#",
						},
					},
				},
				{
					Title: "Company",
					Links: []MarketingSectionFooterLink{
						{
							Text: "About",
							Href: "#",
						},
						{
							Text: "Blog",
							Href: "#",
						},
						{
							Text: "Jobs",
							Href: "#",
						},
						{
							Text: "Press",
							Href: "#",
						},
						{
							Text: "Partners",
							Href: "#",
						},
					},
				},
				{
					Title: "Legal",
					Links: []MarketingSectionFooterLink{
						{
							Text: "Claim",
							Href: "#",
						},
						{
							Text: "Privacy",
							Href: "#",
						},
						{
							Text: "Terms",
							Href: "#",
						},
					},
				},
			},
			Newsletter: MarketingSecionsFooterNewsletter{
				Enabled: true,
				Action:  "/",
				Method:  "GET",
				Title:   "Subscribe to our newsletter",
				Text:    "The latest news, articles, and resources, sent to your inbox weekly.",
				Input: MarketingSectionFooterInput{
					Name:        "Email",
					Label:       "Email address",
					Placeholder: "Enter your email",
				},
				Button: MarketingSectionFooterAction{
					Text:    "Subscribe",
					Value:   "Subscribe",
					Primary: true,
				},
			},
		},
		Simple: MarketingSectionFooterSimple{
			Enabled: true,
			Text:    "© 2020 Workflow, Inc. All rights reserved",
			Socials: []MarketingSectionFooterSocial{
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 12.067C0 18.033 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511zm8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379zm-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986zM8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996zm10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41l.002-.003z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M512 0q-52 0-101.5 10T315 40l-6 3Q171 103 85.5 230T0 512q0 104 40.5 198.5T150 874t163.5 109.5T512 1024t198.5-40.5T874 874t109.5-163.5T1024 512q0-83-26-161t-73.5-141T814 99.5T673 26T512 0zm448 512q0 2-.5 6t-.5 5q-155-42-326-4q-5-12-10.5-23.5t-11-23.5t-11.5-24t-13-24q0-1-1.5-3.5T583 416q110-42 172.5-96T846 214q113 128 114 298zM798 167q-6 14-14.5 28t-28 36t-44.5 42.5t-66.5 44T553 359Q459 188 378 85q65-21 134-21q53 0 104.5 12.5T713 112t85 55zm-482-57q18 21 39 50t59 89.5T491 380q-70 21-140 33.5t-125.5 15t-90 2T72 428q20-105 85-188.5T316 110zM64 512q0-9 1-21q40 5 88 5q183 0 369-59q2 4 4.5 8.5t4.5 8.5q10 20 20 40.5t19 41.5q-8 2-16.5 5l-17 6l-16.5 6.5l-16.5 7L486 568q-64 29-118 65.5t-89.5 73t-56 62.5t-34.5 51q-59-61-91.5-140.5T64 512zm172 352q10-18 25.5-39.5T309 767t85.5-74T513 626q13-6 27-11.5t27.5-10.5t27.5-9q67 172 88 331q-40 16-83 25t-88 9q-154 0-276-96zm508 31q-23-154-86-316q51-10 102.5-12t91 2.5t62 9T953 587q-11 65-40 123t-72 105t-97 80z" fill="currentColor"/></svg>`,
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionFooter4 = kyoto.RegC(p, &MarketingSectionFooter{
		Simple: MarketingSectionFooterSimple{
			Enabled:  true,
			Text:     "© 2020 Workflow, Inc. All rights reserved",
			Centered: true,
			Socials: []MarketingSectionFooterSocial{
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 12.067C0 18.033 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511zm8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379zm-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986zM8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996zm10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41l.002-.003z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M512 0q-52 0-101.5 10T315 40l-6 3Q171 103 85.5 230T0 512q0 104 40.5 198.5T150 874t163.5 109.5T512 1024t198.5-40.5T874 874t109.5-163.5T1024 512q0-83-26-161t-73.5-141T814 99.5T673 26T512 0zm448 512q0 2-.5 6t-.5 5q-155-42-326-4q-5-12-10.5-23.5t-11-23.5t-11.5-24t-13-24q0-1-1.5-3.5T583 416q110-42 172.5-96T846 214q113 128 114 298zM798 167q-6 14-14.5 28t-28 36t-44.5 42.5t-66.5 44T553 359Q459 188 378 85q65-21 134-21q53 0 104.5 12.5T713 112t85 55zm-482-57q18 21 39 50t59 89.5T491 380q-70 21-140 33.5t-125.5 15t-90 2T72 428q20-105 85-188.5T316 110zM64 512q0-9 1-21q40 5 88 5q183 0 369-59q2 4 4.5 8.5t4.5 8.5q10 20 20 40.5t19 41.5q-8 2-16.5 5l-17 6l-16.5 6.5l-16.5 7L486 568q-64 29-118 65.5t-89.5 73t-56 62.5t-34.5 51q-59-61-91.5-140.5T64 512zm172 352q10-18 25.5-39.5T309 767t85.5-74T513 626q13-6 27-11.5t27.5-10.5t27.5-9q67 172 88 331q-40 16-83 25t-88 9q-154 0-276-96zm508 31q-23-154-86-316q51-10 102.5-12t91 2.5t62 9T953 587q-11 65-40 123t-72 105t-97 80z" fill="currentColor"/></svg>`,
					Href: "#",
				},
			},
			Links: []MarketingSectionFooterLink{
				{
					Text: "About",
					Href: "#",
				},
				{
					Text: "Blog",
					Href: "#",
				},
				{
					Text: "Jobs",
					Href: "#",
				},
				{
					Text: "Press",
					Href: "#",
				},
				{
					Text: "Partners",
					Href: "#",
				},
			},
		},
	})
	p.MarketingSectionFooter5 = kyoto.RegC(p, &MarketingSectionFooter{
		Simple: MarketingSectionFooterSimple{
			Enabled: true,
			Text:    "© 2020 Workflow, Inc. All rights reserved",
			Socials: []MarketingSectionFooterSocial{
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 12.067C0 18.033 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511zm8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379zm-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986zM8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996zm10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41l.002-.003z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"><g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" fill="currentColor"/></g></svg>`,
					Href: "#",
				},
				{
					Icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="20px" height="20px" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M512 0q-52 0-101.5 10T315 40l-6 3Q171 103 85.5 230T0 512q0 104 40.5 198.5T150 874t163.5 109.5T512 1024t198.5-40.5T874 874t109.5-163.5T1024 512q0-83-26-161t-73.5-141T814 99.5T673 26T512 0zm448 512q0 2-.5 6t-.5 5q-155-42-326-4q-5-12-10.5-23.5t-11-23.5t-11.5-24t-13-24q0-1-1.5-3.5T583 416q110-42 172.5-96T846 214q113 128 114 298zM798 167q-6 14-14.5 28t-28 36t-44.5 42.5t-66.5 44T553 359Q459 188 378 85q65-21 134-21q53 0 104.5 12.5T713 112t85 55zm-482-57q18 21 39 50t59 89.5T491 380q-70 21-140 33.5t-125.5 15t-90 2T72 428q20-105 85-188.5T316 110zM64 512q0-9 1-21q40 5 88 5q183 0 369-59q2 4 4.5 8.5t4.5 8.5q10 20 20 40.5t19 41.5q-8 2-16.5 5l-17 6l-16.5 6.5l-16.5 7L486 568q-64 29-118 65.5t-89.5 73t-56 62.5t-34.5 51q-59-61-91.5-140.5T64 512zm172 352q10-18 25.5-39.5T309 767t85.5-74T513 626q13-6 27-11.5t27.5-10.5t27.5-9q67 172 88 331q-40 16-83 25t-88 9q-154 0-276-96zm508 31q-23-154-86-316q51-10 102.5-12t91 2.5t62 9T953 587q-11 65-40 123t-72 105t-97 80z" fill="currentColor"/></svg>`,
					Href: "#",
				},
			},
		},
	})
}

func TestMarketing(t *testing.T) {
	serve(&TestMarketingPage{})
}
