# Vicyber api
This is a web server deployed to google cloud that I made for my robotics team, Vicyber

## Links
The client repo is [here](https://github.com/CezarBaluta/viCyber). \
The live website is [here](https://vicyber.ro).

## Endpoints

POST /article - submit article without id, return article with id \
GET /article/{category} - get all articles in a category \
DELETE /article/{id}

## Notes
Server expects API key in header and API key env var \
Server expects postgres database URL env var \
Server expects port env var