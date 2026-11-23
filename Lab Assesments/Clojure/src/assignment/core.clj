(ns assignment.core
  (:require [ring.adapter.jetty :refer [run-jetty]]
            [clojure.data.json :as json])
  (:import [java.sql DriverManager]))

(def database "jdbc:sqlite:College.db")

(def products
  [{:id 1 :name "Motherboard" :quantity 2 :price 50000}
   {:id 2 :name "CPU" :quantity 5 :price 15000}
   {:id 3 :name "GPU" :quantity 3 :price 150000}
   {:id 4 :name "Power Supply" :quantity 2 :price 12000}
   {:id 5 :name "Cabinet" :quantity 4 :price 10000}])

(defn process-data []
  (let [processed (map
                    (fn [item]
                      (assoc item :total
                             (* (:quantity item) (:price item))))
                    products)

        filtered (filter
                   #(>= (:total %) 1000)
                   processed)

        summary (reduce
                  (fn [result item]
                    {:total-sales (+ (:total-sales result)
                                     (:total item))
                     :total-products (inc (:total-products result))})
                  {:total-sales 0 :total-products 0}
                  filtered)]

    {:rows (vec filtered)
     :summary summary}))



(defn create-table []
  (try
    (with-open [connection (DriverManager/getConnection database)
                statement (.createStatement connection)]
      (.executeUpdate statement
        "CREATE TABLE IF NOT EXISTS Products
         (id INTEGER PRIMARY KEY,
          name TEXT,
          quantity INTEGER,
          price INTEGER,
          total INTEGER)"))
    (println "Database table created")
    (catch Exception e
      (println "Database error:" (.getMessage e)))))

(defn save-data [items]
  (try
    (with-open [connection (DriverManager/getConnection database)
                statement (.prepareStatement connection
                  "INSERT OR REPLACE INTO Products
                   VALUES (?, ?, ?, ?, ?)")]

      (doseq [item items]
        (.setInt statement 1 (:id item))
        (.setString statement 2 (:name item))
        (.setInt statement 3 (:quantity item))
        (.setInt statement 4 (:price item))
        (.setInt statement 5 (:total item))
        (.executeUpdate statement)))

      (println "Data saved successfully")

    (catch Exception e
      (println "Save error:" (.getMessage e)))))

(defn get-data []
  (try
    (with-open [connection (DriverManager/getConnection database)
                statement (.createStatement connection)
                result (.executeQuery statement "SELECT * FROM Products")]

      (loop [items []]
        (if (.next result)
          (recur (conj items
                   {:id (.getInt result "id")
                    :name (.getString result "name")
                    :quantity (.getInt result "quantity")
                    :price (.getInt result "price")
                    :total (.getInt result "total")}))
          items)))

    (catch Exception e
      (println "Read error:" (.getMessage e))
      [])))

(defn handler [request]
  (try
    (case (:uri request)
      "/results"
      {:status 200
       :headers {"Content-Type" "application/json"}
       :body (json/write-str (get-data))}

      "/"
      {:status 200
       :headers {"Content-Type" "text/plain"}
       :body "Clojure Data Pipeline is running!"}

      {:status 404
       :headers {"Content-Type" "text/plain"}
       :body "Page not found"})

    (catch Exception e
      {:status 500
       :headers {"Content-Type" "text/plain"}
       :body "Something went wrong"})))

(defn -main [& args]
  (let [data (process-data)]
    (create-table)
    (save-data (:rows data))
    (println "Summary:" (:summary data))
    (println "Starting web server on port 3000")
    (run-jetty handler {:port 3000 :join? true})))
