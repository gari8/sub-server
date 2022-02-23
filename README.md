# sub-server

# How to install
```bash
$ go install github.com/gari8/sub-server@latest
```

# How to use
```bash
$ mkdir mock-server && cd mock-server

# creating json files
$ mkdir 1 && touch index.json post.json 1/index.json

# creating config file
$ sub-server init
```

```bash
$ vim index.json
# ... 
```

# JSON example
https://github.com/gari8/mock-server-by-sub-server

### index.json
```json: index.json
[
  {
    "id": 1,
    "title": "Who I am"
  },
  {
    "id": 2,
    "title": "Who I am"
  },
  {
    "id": 3,
    "title": "Who I am"
  },
  {
    "id": 4,
    "title": "Who I am"
  }
]
```

### post.json
```json: post.json
{
  "id": 5,
  "title": "new one"
}
```

### 1/index.json
```json: index.json
{
  "id": 1,
  "title": "Who I am"
}
```

# Running
```bash
$ sub-server serve


        ░██████╗██╗░░░██╗██████╗░  ░██████╗███████╗██████╗░██╗░░░██╗███████╗██████╗░
        ██╔════╝██║░░░██║██╔══██╗  ██╔════╝██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗
        ╚█████╗░██║░░░██║██████╦╝  ╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝█████╗░░██████╔╝
        ░╚═══██╗██║░░░██║██╔══██╗  ░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██╔══╝░░██╔══██╗
        ██████╔╝╚██████╔╝██████╦╝  ██████╔╝███████╗██║░░██║░░╚██╔╝░░███████╗██║░░██║
        ╚═════╝░░╚═════╝░╚═════╝░  ╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝


        GET >> http://localhost:8080/app/v1/start/1
        GET >> http://localhost:8080/app/v1/start
        POST >> http://localhost:8080/app/v1/start

```

# How to write config.toml
```toml
[app] # App resource
server_name = "server name"
port = "8080" # server port

[routing]
root_path = "." # JSON root path at local
origin_root = "/app/v1" # server url root

# routing resources
[[routing.origins]]
uri = "/start" # server url after origin_root
file_path = "index.json" # JSON path at local after root_path
method = "GET"

[[routing.origins]]
uri = "/start/1"
file_path = "1/index.json"
method = "GET"

[[routing.origins]]
uri = "/start"
file_path = "post.json"
method = "POST"
```
