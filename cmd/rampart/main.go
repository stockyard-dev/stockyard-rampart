package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-rampart/internal/server";"github.com/stockyard-dev/stockyard-rampart/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="10030"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./rampart-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("rampart: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Rampart — IP blocklist manager\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("rampart: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
