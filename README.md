# Vicyber api

# endpoints

POST /article - submit text, return article object (with id) \
GET /article \
GET /article/{id} \
DELETE /article/{id}

# unused endpoints on client:

PUT /article/{id} - submit article object

# notes:
Server expects API key in header
Server expects postgres database URL env var
Server expects port env var
Server uses google cloud bucket