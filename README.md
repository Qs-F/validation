# package validation

## installation

`go get github.com/Qs-F/validation`

## usage

1. import `github.com/Qs-F/validation`
2. use like `validation.Required(VALUE)`

## future design is complete! 

```go
v := validation.SetValue(VALUE)

// you can set err message last arg.
v.Required("this field is required.")

// or if you want to use default error message, all you have to do is DONOT set last string arg
v.MaxSize(2)

// you can get all error message like this:
for _, value := range v.Errors {
  println(value.Error())
}
```

