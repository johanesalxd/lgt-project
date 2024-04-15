go build -o http_server main/*

curl http://localhost:8498/players/(player_name)
curl -X POST http://localhost:8498/players/(player_name)