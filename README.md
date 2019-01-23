# go-lambda-invoke

Small package to allow you to invoke your Go AWS lambda locally.

This _might_ be useful for:

* You want to run an integration test, maybe in conjunction with [LocalStack](https://github.com/atlassian/localstack)?
  * Unit testing is probably better in most cases
* You want to validate your CI has built a valid `linux` binary of your application before deploying

I wrote a blog post about it here https://djhworld.github.io/post/2018/01/27/running-go-aws-lambda-functions-locally/

## Installing

```
go get -u github.com/djhworld/go-lambda-invoke/golambdainvoke
```

## Example usage

Run the example lambda [toupperlambda.go](/toupperlambda.go) on port 8001

```
_LAMBDA_SERVER_PORT=8001 go run ./toupperlambda.go
```

Then use this library in tests or wherever you need it, by calling 

```
response, err := golambdainvoke.Run(Input{
    Port:    8001,
    Payload: "payload",
})
```

Note that `Payload` can be any structure that can be encoded by the `encoding/json` package. Your lambda function will need to use this structure in its type signature.
