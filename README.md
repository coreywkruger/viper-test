# viper-test

Tests UnmarshalKey() from https://github.com/spf13/viper

Simply execute `go run main.go` and watch output.

The json file that is being loaded is of the format:
```json
{
  "data":{
	  "normal": "Normal stuff",
      "snake_case": "Snake case stuff that is skipped :("
	}
}
```

The struct it is being loaded into is defined here:
```go
type configJSON struct {
	Normal    string `json:"normal"`
	SnakeCase string `json:"snake_case"`
}
```

The json file is loaded like so:
```go
// Unmarshalling from viper config:
config := viper.New()
config.SetConfigFile("./config.json")
err := config.ReadInConfig()
if err != nil {
	log.Fatal(err)
}
```
... then un-marshaled  and printed ...
```go
c := configJSON{}
config.UnmarshalKey("data", &c)

// Print out the values of the fields
fmt.Printf("normal: `%s`\n", c.Normal)
fmt.Printf("snake_case: `%s`\n", c.SnakeCase) // <-- Notice how this is empty
```
The console output looks like:
```
normal: `Normal stuff`
snake_case: ``
```
`snake_case` is an empty string (but shouldn't be).
