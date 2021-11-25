
# With starter

The quickest way to get started with `kyoto` and `uikit` is a prepared starter project.  

## Installation

First, you need to clone starter as your new project.  
`<app name>` - your new project name

```bash
git clone --recursive https://github.com/kyoto-framework/starter <app name>
```

Then, we need to install node dependencies and build css.  
All statics and assets for building are placed in `static` directory.  

```bash
(cd static; npm i; npm run build)
```

In the final step, you'll need to set git origin URL to your repository.  
`<repo>` - your's project repository URL

```bash
git remote set-url origin <repo>
```

## What's included?

- `kyoto` - core library
- `uikit` - [UI Kit](https://github.com/kyoto-framework/uikit), built on top of `kyoto`
- `tailwindcss` - [Tailwind CSS](https://imgur.com/RN4YbvR.png) library
