#  GoJsonSchemaServer

This is an HTTP server designed to generate JsonSchema from Golang structs.

## ⚠️Warning:

**Deployment of this project should be done within Docker, as it relies on the Golang runtime and may pose risks or vulnerabilities otherwise.**
```shell
docker build -t go-json-schema-server . 

docker run --name go-json-schema-server -d -p 8080:8080 go-json-schema-server
```
### Run the current project
```shell

https://github.com/maocatooo/GoJsonSchemaServer.git

cd GoJsonSchemaServer

go mod test 

go run .

```


## Example



### input
```go
type html struct {
    Head struct {
        Title string `json:"title"` // this is title
    } `json:"head"` // this is head
    Body Body `json:"body"` // this is body
}

type Body struct {
    Div int `json:"div"` // this is div
}
```
### output
```json5
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "properties": {
    "head": {
      "properties": {
        "title": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "title"
      ],
      "description": "this is head"
    },
    "body": {
      "properties": {
        "div": {
          "type": "integer",
          "description": "this is div"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "div"
      ],
      "description": "this is body"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "head",
    "body"
  ]
}

```

