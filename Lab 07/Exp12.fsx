open System

let firstList = [12; 15; 18; 21; 24]
let secondList = [31; 34; 37; 40; 43]

printfn "Original Lists:"
printfn "%A" firstList
printfn "%A" secondList

let firstSquare =
    firstList |> List.map (fun num -> num * num)

let secondSquare =
    secondList |> List.map (fun num -> num * num)

printfn "\n1. List.map - Square of Numbers:"
printfn "%A" firstSquare
printfn "%A" secondSquare

let firstEven =
    firstList |> List.filter (fun value -> value % 2 = 0)

let secondEven =
    secondList |> List.filter (fun value -> value % 2 = 0)

printfn "\n2. List.filter - Even Numbers:"
printfn "%A" firstEven
printfn "%A" secondEven

let firstSum =
    firstList |> List.fold (fun result item -> result + item) 0

let secondSum =
    secondList |> List.fold (fun result item -> result + item) 0

printfn "\n3. List.fold - Sum of Numbers:"
printfn "Sum of First List = %d" firstSum
printfn "Sum of Second List = %d" secondSum