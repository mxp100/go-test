# Binary Search Tree Simple Application



## Run
```shell script
go run visiongotest.go 
```
## Build
```shell script
go build visiongotest.go
```

## API
### Search
Method: GET

|Params|Type|Description|Example|
|---|---|---|---|
|val|integer|Searching value|36|

Example:
```shell script
curl --request GET --url 'http://localhost:8080/search?val=36'
```

Output:
```json
{
  "Found": true,
  "Tree": {
    "Value": 36,
    "Left": {
      "Value": 34,
      "Left": {
        "Value": 32,
        "Left": null,
        "Right": null
      },
      "Right": null
    },
    "Right": null
  }
}
```

### Delete
Method: DELETE

|Params|Type|Example|
|---|---|---|
|val|integer|412|

Example:
```shell script
curl --request DELETE --url 'http://localhost:8080/delete?val=412'
```

Output:
```
{
  "Deleted": [
    412
  ],
  "Tree": ... Full tree
```

### Insert
Method: POST

Content-Type: application/json

Body: int[]

Example:
```shell script
curl --request POST \
  --url http://localhost:8080/insert \
  --header 'content-type: application/json' \
  --data '[2]'
```

Output:
```
{
  "Inserted": [
    2
  ],
  "Tree": ... Full tree
```