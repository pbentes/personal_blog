# Prerequisites

You must have [Templ](https://templ.guide/) installed to use this template. You can find more information on installing Templ [here](https://templ.guide/quick-start/installation).

Before running the project if you need a database create a .env file following the .env.example file. If you don't need a database you can just remove the calls to the db package from the main file.

# Running the project

## Using Air

Using air is recommended, it provides live reloads and pre/post build commands. You can find information on Air [here](https://github.com/cosmtrek/air).

We provide a pre configured air.toml file so once air is installed simply run `air` in the console to run the project.

## Manually

You can also run the project manually by running:

`npx tailwindcss -i ./assets/css/input.css -o ./assets/css/tailwind.css`

`templ generate`

`go run ./src/`
