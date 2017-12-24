package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
    "time"
    "encoding/json"
    "log"
    "net/http"
)

type Sleep struct {
    Sleep_date time.Time `json:"sleep_date"`
    Bedding_datetime time.Time `json:"bedding_datetime""`
    Sleeping_datetime time.Time `json:"sleeping_datetime"`
    Getting_up_datetime time.Time `json:"getting_up_datetime"`
    Time_of_sleeping time.Time `json:"time_of_sleeping"`
    Arousal_time time.Time `json:"arousal_time"`
    Sleep_latency time.Time `json:"sleep_latency"`
    Middle_awakening int `json:"middle_awakening"`
    Bedtime time.Time `json:"bedtime"`
} 

func main() {
    http.HandleFunc("/hypnos", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Origin", "*" )
        w.Header().Set("Access-Control-Allow-Methods","GET" )

        db, _ := sql.Open("postgres", "user=postgres dbname=hypnos password=t90thbntH sslmode=disable")
        defer db.Close()

        rows, _ := db.Query("select * from sleep;")
        sleeps:=[]Sleep{}
        for rows.Next() {
            var id int
            var sleep_date time.Time
            var bedding_datetime time.Time
            var sleeping_datetime time.Time
            var getting_up_datetime time.Time
            var time_of_sleeping time.Time
            var arousal_time time.Time
            var sleep_latency time.Time
            var middle_awakening int
            var bedtime time.Time

            rows.Scan(&id, &sleep_date, &bedding_datetime, &sleeping_datetime, &getting_up_datetime, &time_of_sleeping, &arousal_time, &sleep_latency, &middle_awakening, &bedtime)

            var sleep Sleep
            sleep.Sleep_date = sleep_date
            sleep.Bedding_datetime = bedding_datetime
            sleep.Sleeping_datetime = sleeping_datetime
            sleep.Getting_up_datetime = getting_up_datetime
            sleep.Time_of_sleeping = time_of_sleeping
            sleep.Arousal_time = arousal_time
            sleep.Sleep_latency = sleep_latency
            sleep.Middle_awakening = middle_awakening
            sleep.Bedtime = bedtime
            
            sleeps = append(sleeps, sleep)
        }

        jsonBytes, _ := json.Marshal(sleeps)

        fmt.Fprintf(w, string(jsonBytes))
    })

    log.Fatal(http.ListenAndServe(":3000", nil))
}
