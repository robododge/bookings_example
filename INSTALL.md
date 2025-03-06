# Tools intalling

1. air `asdf plugin add air` and `asdf install air 1.61.7`, then set it for use locally `asdf set air 1.61.7`
2. create director and setup go.mod `go mod init`
3. Could set GOPROXY= . to your companies artifact repo if needed
4. the main.go file in the step_01 branch will have gin framework and a simmple html file 
5. Do an pnpm init
6. install tailwind for css tooling `pnpm add -D tailwindcss postcss autoprefixer`
7. install vite for javascript bundling `pnpm install -D vite` 
8. install concurrently to help with doing vite builds along with tailwind builds `pnpm install -D concurrently`
9. make sure the src/ directory contains the index.css and the htmx.js.  The index.css is used by the tailwind css processor to ensure css classes are generated, the htmx.js is critical for pulling down th htmx javascrit and loading into the ./static dir for use in the live app
10. At this point you have enough data in place to run `air init` which will generate a comprable .air.toml file to control the air relaoder
11. Add htmx to the the node modules
12. Setup a "dev" script in the package.json
