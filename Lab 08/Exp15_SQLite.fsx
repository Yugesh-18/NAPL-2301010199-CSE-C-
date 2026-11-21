#r "nuget: Microsoft.Data.Sqlite"

open System
open Microsoft.Data.Sqlite

let dbConnection =
    new SqliteConnection("Data Source=College.db")

dbConnection.Open()

let createTable = dbConnection.CreateCommand()

createTable.CommandText <-
    """CREATE TABLE IF NOT EXISTS Students
       (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)"""

createTable.ExecuteNonQuery() |> ignore

let clearRecords = dbConnection.CreateCommand()
clearRecords.CommandText <- "DELETE FROM Students"
clearRecords.ExecuteNonQuery() |> ignore

let showRecords () =
    let readCommand = dbConnection.CreateCommand()
    readCommand.CommandText <- "SELECT * FROM Students"

    use reader = readCommand.ExecuteReader()

    let mutable found = false

    while reader.Read() do
        found <- true

        printfn "ID: %d, Name: %s, Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

    if not found then
        printfn "No records found."

printfn "1. INSERT RECORD"

let insertRecord = dbConnection.CreateCommand()

insertRecord.CommandText <-
    "INSERT INTO Students VALUES (201, 'Amit', 21)"

insertRecord.ExecuteNonQuery() |> ignore

printfn "Student record inserted successfully."

printfn "\n2. DISPLAY RECORD"

showRecords()

printfn "\n3. UPDATE RECORD"

let updateRecord = dbConnection.CreateCommand()

updateRecord.CommandText <-
    "UPDATE Students SET Age = 24 WHERE Id = 201"

updateRecord.ExecuteNonQuery() |> ignore

printfn "Student record updated successfully."

showRecords()

printfn "\n4. DELETE RECORD"

let deleteRecord = dbConnection.CreateCommand()

deleteRecord.CommandText <-
    "DELETE FROM Students WHERE Id = 201"

deleteRecord.ExecuteNonQuery() |> ignore

printfn "Student record deleted successfully."

printfn "\nFINAL RECORDS:"

showRecords()

dbConnection.Dispose()