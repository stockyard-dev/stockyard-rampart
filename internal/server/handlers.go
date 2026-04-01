package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-rampart/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list:=r.URL.Query().Get("list");entries,_:=s.db.List(list);if entries==nil{entries=[]store.Entry{}};writeJSON(w,200,entries)}
func(s *Server)handleAdd(w http.ResponseWriter,r *http.Request){var e store.Entry;json.NewDecoder(r.Body).Decode(&e);if e.IP==""||e.List==""{writeError(w,400,"ip and list required");return};s.db.Add(&e);writeJSON(w,201,e)}
func(s *Server)handleLookup(w http.ResponseWriter,r *http.Request){ip:=r.PathValue("ip");result,_:=s.db.Lookup(ip);writeJSON(w,200,result)}
func(s *Server)handleRemove(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Remove(id);writeJSON(w,200,map[string]string{"status":"removed"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
