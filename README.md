# Vicyber api
This is a web server deployed to google cloud that I made for my robotics team, Vicyber

## endpoints

POST /article - submit article without id, return article with id \
GET /article - get all articles \
GET /article/{id} - get article by id \
DELETE /article/{id}

POST /image - submit image obj with image data and width \
GET /image/{id} - get image by id \
DELETE /image/{id}

## unused endpoints on client:

PUT /article/{id} - submit article to update

## notes:
Server expects API key in header and API key env var \
Server expects postgres database URL env var \
Server expects port env var \
Server uses google cloud bucket