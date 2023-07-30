# tune-bot api
The REST API that handles managing users and their playlists for tune-bot. Requires [database](https://github.com/tune-bot/database) credentials and a connection to the Tunebot database.

To run locally:
```
git clone https://github.com/tune-bot/api
cd api
go run .
```

## Docker
```
docker build -t tune-bot-api .
docker run -p 80:80 -td tune-bot-api
```
