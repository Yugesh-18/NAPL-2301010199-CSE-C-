open System

[<AbstractClass>]
type Shape() =
    abstract member CalculateArea : unit -> float

type Circle(radius: float) =
    inherit Shape()

    override this.CalculateArea() =
        3.14 * radius * radius

type Rectangle(length: float, breadth: float) =
    inherit Shape()

    override this.CalculateArea() =
        length * breadth

let circle = Circle(7.0)
let rectangle = Rectangle(12.0, 8.0)

printfn "OBJECT ORIENTED PROGRAMMING"

printfn "\nCircle:"
printfn "Radius = 7"
printfn "Area = %.2f" (circle.CalculateArea())

printfn "\nRectangle:"
printfn "Length = 12"
printfn "Breadth = 8"
printfn "Area = %.2f" (rectangle.CalculateArea())

let shapeList: Shape list = [circle; rectangle]

printfn "\nDisplaying Areas Using Polymorphism:"

shapeList
|> List.iter (fun item ->
    printfn "Area = %.2f" (item.CalculateArea())
)