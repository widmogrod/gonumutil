# gonumutil
[GoNum](https://github.com/gonum) utility functions

## NewConstantNumTicker
The example demonstrates how to change the scale of the X and Y axis to have constant distance between markers
```go
p, err := plot.New()
if err != nil {
    panic(err)
}
p.Title.Text = "Average CPU utilization in percentage"
p.X.Tick.Marker = gonumutil.NewConstantNumTicker(1)
p.Y.Tick.Marker = gonumutil.NewConstantNumTicker(5)
```

## NewConstantTicker
Equivalent function to `NewConstantNumTicker` but with more degree of freedom, that enables label customization.
```go
p, err := plot.New()
if err != nil {
    panic(err)
}
p.Title.Text = "Average CPU utilization in percentage"
p.X.Tick.Marker = NewConstantTicker("%1.0f", 1.0)
p.Y.Tick.Marker = NewConstantTicker("%1.0f", 5.0)
```