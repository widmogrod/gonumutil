# gonumutil
GoNum utility functions

## NewConstantNumTicker & NewConstantTicker
Example demonstrates how to change scale of `X` and `Y` axis to have constant distance between markers
```go
p, err := plot.New()
if err != nil {
    panic(err)
}
p.Title.Text = "Average CPU utilization in percentage"
p.X.Tick.Marker = gonumutil.NewConstantNumTicker(1)
p.Y.Tick.Marker = gonumutil.NewConstantNumTicker(5)

// equivalent function as above but help
p.X.Tick.Marker = NewConstantTicker("%1.0f", 1.0)
p.Y.Tick.Marker = NewConstantTicker("%1.0f", 5.0)

```

## NewConstantTicker
Equivalent function as above but with more degree of freedom, that enable label customization.
```go
p, err := plot.New()
if err != nil {
    panic(err)
}
p.Title.Text = "Average CPU utilization in percentage"
p.X.Tick.Marker = NewConstantTicker("%1.0f", 1.0)
p.Y.Tick.Marker = NewConstantTicker("%1.0f", 5.0)
```