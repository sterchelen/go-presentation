type Abser interface {
  Abs() float
}

var a Abser
a = Point{3, 4}
print(a.Abs()}
a= Vector{1, 2, 3, 4}
print(a.Abs()}
