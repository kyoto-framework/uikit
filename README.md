
# kyoto uikit

UIKit for rapid development

## Requirements

- `kyoto` page
- configured SSA
- basic knowledge of `kyoto`
- (twui) configured tailwindcss

## Installation

`<kit>` - component library you want to use

- Install Go package into your project with `go get github.com/yuriizinets/kyoto-uikit/<kit>`
- Copy all `<kit>/*.html` files to your templates directory

## Example

Simple example using tailwind ui stats component

![stats](https://imgur.com/TjBmJ26.png)

```go
package main

import (
    "github.com/yuriizinets/kyoto"
    "github.com/yuriizinets/kyoto-uikit/twui"
)


type PageExample struct {
    Stats kyoto.Component
}

func (p *PageExample) Template() *template.Template {
    return template.Must(template.New("page.example.html").Funcs(kyoto.TFuncMap()).ParseGlob("templates/*.html"))
}

func (p *PageExample) Init() {
    p.Stats = kyoto.RegC(p, &twui.AppUIDataStats{
        Items: []twui.AppUIDataStatsItem{
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
}
```

## Documentation

Sorry, but not ready yet 🤷‍♂️  
We are working on this

## Features

- Provides high-level blocks instead of primitives like buttons and labels
- Configurable enough for using in different situations
- Adaptive out of the box
- Don't like component UI? It's very easy to customize it!

## References

- `kyoto-uikit` project: https://github.com/yuriizinets/kyoto-uikit
- `kyoto` project: https://github.com/yuriizinets/kyoto 
- Documentation: not ready yet 🤷‍♂️  
- Demo project: not ready yet 🤷‍♂️

## Contribution

Development setup is not ready yet, but feel free to use test files for testing components. Just check how existing components are used in tests.

## Support

<a target="_blank" href="https://www.buymeacoffee.com/yuriizinets"><img alt="Buy me a Coffee" src="https://github.com/egonelbre/gophers/blob/master/.thumb/animation/buy-morning-coffee-3x.gif?raw=true"></a>


Or directly with Bitcoin: `18WY5KRWVKVxjWKFzJLqqALsRrsDh4snqg`
