# The first laboratory work at PR dining_hall_restaurant

This is the dining_hall_restaurant. It is related to the kitchen_restaurant repository.

## Restaurant app with Docker (used here docker compose)

It is required to introduce in Terminal:

```bash
$ docker compose up --build
```
## Run the app in the Terminal

Firstly there is required to switch: `"kitchen_url": "http://localhost:8080/distribution"` in `main.go`.

Then to run in the Terminal:

```bash
$ go run .
```
## Try it by yourself

Pay attention at the order of running, because everytime the kitchen_restaurant is running first.