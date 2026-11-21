#r "nuget: Microsoft.Data.Sqlite"

open System
open Microsoft.Data.Sqlite

type Student = {
    Id: int
    Name: string
    Age: int
}

let db = new SqliteConnection("Data Source=College.db")
db.Open()

let cmd = db.CreateCommand()
cmd.CommandText <-
    "CREATE TABLE IF NOT EXISTS Students
     (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)"
cmd.ExecuteNonQuery() |> ignore

let addStudent name age =
    async {
        let c = db.CreateCommand()
        c.CommandText <- "INSERT INTO Students VALUES (NULL, @n, @a)"
        c.Parameters.AddWithValue("@n", name) |> ignore
        c.Parameters.AddWithValue("@a", age) |> ignore
        let! _ = c.ExecuteNonQueryAsync() |> Async.AwaitTask
        printfn "Student added!"
    }


let showStudents () =
    async {
        let c = db.CreateCommand()
        c.CommandText <- "SELECT Id, Name, Age FROM Students"
        let! r = c.ExecuteReaderAsync() |> Async.AwaitTask

        let mutable list = []
        while r.Read() do
            list <- { Id = r.GetInt32(0); Name = r.GetString(1); Age = r.GetInt32(2) } :: list

        r.Close()

        list |> List.map (fun s -> s.Name)
             |> List.iter (printfn "%s")

        list |> List.filter (fun s -> s.Age >= 18)
             |> List.iter (fun s -> printfn "%s is an adult" s.Name)
    }

let updateStudent id name =
    async {
        let c = db.CreateCommand()
        c.CommandText <- "UPDATE Students SET Name=@n WHERE Id=@id"
        c.Parameters.AddWithValue("@n", name) |> ignore
        c.Parameters.AddWithValue("@id", id) |> ignore
        let! _ = c.ExecuteNonQueryAsync() |> Async.AwaitTask
        printfn "Student updated!"
    }

let deleteStudent id =
    async {
        let c = db.CreateCommand()
        c.CommandText <- "DELETE FROM Students WHERE Id=@id"
        c.Parameters.AddWithValue("@id", id) |> ignore
        let! _ = c.ExecuteNonQueryAsync() |> Async.AwaitTask
        printfn "Student deleted!"
    }

let main () =
    async {
        match 1 with
        | 1 -> do! addStudent "Rahul" 20
        | _ -> printfn "Invalid choice"

        do! showStudents()
        do! updateStudent 1 "Aman"
        do! deleteStudent 1
    }

main() |> Async.RunSynchronously

db.Close()
