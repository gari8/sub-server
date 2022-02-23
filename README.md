# sub-server

# How to install
```bash
$ go install github.com/gari8/sub-server@latest
```

# How to use
```bash
$ mkdir mock-server && cd mock-server

# creating config file
$ sub-server init

# running mock-server
$ sub-server serve


        ░██████╗██╗░░░██╗██████╗░  ░██████╗███████╗██████╗░██╗░░░██╗███████╗██████╗░
        ██╔════╝██║░░░██║██╔══██╗  ██╔════╝██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗
        ╚█████╗░██║░░░██║██████╦╝  ╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝█████╗░░██████╔╝
        ░╚═══██╗██║░░░██║██╔══██╗  ░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██╔══╝░░██╔══██╗
        ██████╔╝╚██████╔╝██████╦╝  ██████╔╝███████╗██║░░██║░░╚██╔╝░░███████╗██║░░██║
        ╚═════╝░░╚═════╝░╚═════╝░  ╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝

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
