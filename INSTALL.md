# Tools intalling

1. air `asdf plugin add air` and `asdf install air 1.61.7`, then set it for use locally `asdf set air 1.61.7`
2. create director and setup go.mod `go mod init`
3. Could set GOPROXY= . to your companies artifact repo if needed
4. the main.go file in the step_01 branch will have gin framework and a simple html file 
5. Do an pnpm init
6. install tailwind for css tooling `pnpm add -D tailwindcss postcss autoprefixer`
7. install vite for javascript bundling `pnpm install -D vite` 
8. install concurrently to help with doing vite builds along with tailwind builds `pnpm install -D concurrently`
9. make sure the src/ directory contains the index.css and the htmx.js.  The index.css is used by the tailwind css processor to ensure css classes are generated, the htmx.js is critical for pulling down th htmx javascrit and loading into the ./static dir for use in the live app
10. Setup a "dev" script in the package.json that uses concurrently to run tailwind and vite ` "dev": "concurrently \"vite build --watch\" \"tailwindcss -i ./src/index.css -o ./static/index.css --watch\""`
11. At this point you have enough data in place to run `air init` which will generate a comprable .air.toml file to control the air relaoder
12. Add htmx to the the node modules `pnpm install htmx.org`
13. templ : install from https://templ.guide/quick-start/installation
14. Once you have templ installed, add it to your air config in .air.toml `cmd = "templ generate && go build -o ./tmp/main ."` .  
15.  Start up the two processes to keep your dev servers refreshed in two separate terminal windows `pnpm dev` and `air`
16.  Navigate to http://localhost:8080/

Makefiles are recommended here




## Trickyness
1. Newer go versions, starting in 1.23.10? , they try to check the x509 certificates for validity at a higher stantand during go module downloads a, you will get a error like this `failed to parse certificate, 509, negative serial number`.   You will want to a flag into godebug `export GODEBUG="x509negativeserial=1"`
