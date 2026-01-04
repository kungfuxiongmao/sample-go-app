Check Models:
Update posts and comments as time goes

Add other models into automigrate
Check that api.go is fine

HANDLERS:

USERS:
1. Create a user: Check if username exists -> Store username and pw, return id of user
2. Login: POST rq of username & pw -> return id and login

Topics:
1. Create a post: Check that body is NOT NULL -> Create in db -> Refresh
2. Delete a post: Parse id of post -> Rmv from db

