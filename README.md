# Vicyber api

# endpoints

POST /article - submit article without id, return article with id \
GET /article - get all articles \
GET /article/{id} - get article by id \
DELETE /article/{id}

# unused endpoints on client:

PUT /article/{id} - submit article to update

# notes:
Server expects API key in header
Server expects postgres database URL env var
Server expects port env var
Server uses google cloud bucket