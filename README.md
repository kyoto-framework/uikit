
# Temporarily archived

Will be used as a basement for a newer version, compatible with latest kyoto changes.

---

<p align="center">
    <img src="https://raw.githubusercontent.com/kyoto-framework/uikit/master/docs/assets/uikit.svg" height="200" />
</p>

<h1 align="center">kyoto uikit</h1>
<p align="center">
    UIKit for rapid development
</p>

![License](https://img.shields.io/github/license/kyoto-framework/uikit)
[![Go Reference](https://pkg.go.dev/badge/github.com/kyoto-framework/uikit.svg)](https://pkg.go.dev/github.com/kyoto-framework/kyoto)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyoto-framework/uikit)](https://goreportcard.com/report/github.com/kyoto-framework/uikit)


## Requirements

- `kyoto` page
- configured SSA
- basic knowledge of `kyoto`
- (twui only) configured tailwindcss

## Installation

### Git submodule

The easiest way to install for now, because it handles templates as well

- Add repo as git submodule to your app: `git submodule add https://github.com/kyoto-framework/uikit.git uikit`
- (twui only) Add replace directive to your `go.mod` file: `github.com/kyoto-framework/uikit/twui => ./uikit/twui`
- (twui only) Add twui path to tailwind config file for JIT mode

### Go package + templates

`<kit>` - component library you want to use

- Install Go package into your project with `go get github.com/kyoto-framework/uikit/<kit>`
- Manually copy all `<kit>/*.html` files to your templates directory

## Example

Simple example using tailwind ui stats component

![stats](https://imgur.com/TjBmJ26.png)

```go
package main

import (
    "github.com/kyoto-framework/kyoto"
    "github.com/kyoto-framework/uikit/twui"
)


type PageExample struct {
    Stats smode.Component
}

func (p *PageExample) Template() *template.Template {
    return template.Must(template.New("page.example.html").Funcs(smode.FuncMap(p)).ParseGlob("templates/*.html"))
}

func (p *PageExample) Init() {
    p.Stats = smode.UseC(p, &twui.AppUIDataStats{
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

- `uikit` project: https://github.com/kyoto-framework/uikit
- `kyoto` project: https://github.com/kyoto-framework/kyoto 
- Documentation: not ready yet 🤷‍♂️  
- Demo project: not ready yet 🤷‍♂️

## Contribution

Development setup is not ready yet, but feel free to use test files for testing components. Just check how existing components are used in tests.

## Support

<a target="_blank" href="https://www.buymeacoffee.com/yuriizinets"><img alt="Buy me a Coffee" src="https://github.com/egonelbre/gophers/blob/master/.thumb/animation/buy-morning-coffee-3x.gif?raw=true"></a>


Or directly with Bitcoin: `bc1qgxe4u799f8pdyzk65sqpq28xj0yc6g05ckhvkk`
