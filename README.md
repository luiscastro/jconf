JSON Configuration for Go
========

This is a simple package written in golang to load a JSON file and use it for configurations (like ini files).

How it works
--------------

This package use same ideas of INI configurations files but written in JSON, you can set your configuration
variables in different sections and inherit one or more sections. Like example below:

```
{
  "development" : {
    "+" : "production",
    "host" : "google.com",
    "port" : "80"
  },
  "production" : {
    "host" : "google.com",
    "port" : "80",
    "safe" : "true"
  },
  "testing" : {
    "+" : ["development", "production"],
    "host" : "google.com",
    "port" : "80",
    "safe" : "false"
  }
}
```

Now for use your configurations variables on your code just folow the next example:

````
func main() {
  config, err := jconf.New("my.conf", "production")
  if err != nil {
    panic("Error loading config file")
  }
  else {
    if host, err := config.Get("host"); err != nil {
      println("My host", host)
    }
  }
}
```

