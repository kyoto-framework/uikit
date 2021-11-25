
# From scratch

Installation from scratch is easy, but requires some additional steps. Unfortunately, native Go packaging systems not allows to include external resources into libraries, like templates, styles, etc.  

## As git submodule (recommended)

Installing UI Kit as git submodule is the easiest way for now, without any difficulties on keeping dependency up-to-date.  

- Add repo as git submodule to your app: `git submodule add https://github.com/kyoto-framework/uikit.git uikit`
- (twui only) Add replace directive to your `go.mod` file: `github.com/kyoto-framework/uikit/twui => ./uikit/twui`
- (twui only) Add twui path to tailwind config file for JIT mode

## Go package + templates

With this method you will install UI Kit as a regular library, but you'll need to manually copy all templates into your templates directory. Also, with this approach it's harder to keep sources up-to-date.  

`<kit>` - component library you want to use

- Install Go package into your project with `go get github.com/kyoto-framework/uikit/<kit>`
- Manually copy all `<kit>/*.html` files to your templates directory
