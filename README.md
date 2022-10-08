# Memos_backend
This is the backend project of Memos.

# How to start
```bash
# 1. clone this repo
git clone https://github.com/ParadiseWitch/memos_backend/
# 2. go mod tidy
go mod tidy
# 3. settings
cp template-settings.yml settings.yml
# 4. edit your settings file
# 5. run
go run main.go server
```
You can also run it by adding an launch.json file
```
{
  "configurations": [
    {
      "name": "Launch file",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "main.go",
      "args": [
        "server"
      ],
    }
  ]
}
```

# Thinks
- [go-admin](https://github.com/go-admin-team/go-admin)
