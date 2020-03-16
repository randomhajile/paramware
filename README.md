# paramware
Simple middlware for handling expected query parameters in Gin.

## Usage
Simply import this module and add the middleware before the handler that requires the
given parameter, e.g.
```go
router.GET(
    "/widgets",
    paramware.Int64Param("intParam", 11),
    paramware.StringParam("stringParam", "foo"),
    getWidgets,
)
```
