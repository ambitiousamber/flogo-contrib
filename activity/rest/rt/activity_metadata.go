package rest

var jsonMetadata = `{
  "name": "tibco-rest",
  "version": "0.0.1",
  "description": "REST activity",
  "inputs":[
    {
      "name": "method",
      "type": "string"
    },
    {
      "name": "uri",
      "type": "string"
    },
    {
      "name": "params",
      "type": "params"
    },
    {
      "name": "content",
      "type": "object"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "object"
    }
  ]
}`
