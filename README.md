# bs-cli

binarysearch.com in terminal

#### first add binarysearch config file in your home dir

```
touch ~/.bs/conf.toml

accessToken = "your token"
lang = "cpp"
devDir = "path-to-your-source-dir"
```

Now you can query and fetch question with `bs_cli`
```
Usage: ./bs_cli [-ls] params...
  -l string
        Input Problem name (default "Add")
  -s int
        Input Problem number
```