# Go MySQL User CRUD API
Initially adapted from https://github.com/icodestuff-io/golang-docker-tutorial for Dockerfile

## Endpoints
### /users
queries mysql with "SELECT * FROM users"
returns
#### id
#### name

### /addUser
Adds a user row based on the body of a GET request
accepts
#### name

returns
#### id


### /homepage
prints welcome message
