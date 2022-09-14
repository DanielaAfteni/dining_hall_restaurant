# dining_hall_restaurant

This is the first lab at PR. It is related to the kitchen_restaurant repository.

## To run the restaurant app with Docker

```bash
$ docker compose up --build
```
## To simply run the app

You need to change: `"kitchen_url": "http://localhost:8081"` in `config/scfg.json`.

```bash
$ go run .
```