package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type BlockRule struct {
	ID string `json:"id"`
	CIDR string `json:"cidr"`
	Reason string `json:"reason"`
	Source string `json:"source"`
	Enabled int `json:"enabled"`
	HitCount int `json:"hit_count"`
	ExpiresAt string `json:"expires_at"`
	LastHitAt string `json:"last_hit_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"rampart.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS block_rules(id TEXT PRIMARY KEY,cidr TEXT NOT NULL,reason TEXT DEFAULT '',source TEXT DEFAULT '',enabled INTEGER DEFAULT 1,hit_count INTEGER DEFAULT 0,expires_at TEXT DEFAULT '',last_hit_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *BlockRule)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO block_rules(id,cidr,reason,source,enabled,hit_count,expires_at,last_hit_at,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.CIDR,e.Reason,e.Source,e.Enabled,e.HitCount,e.ExpiresAt,e.LastHitAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*BlockRule{var e BlockRule;if d.db.QueryRow(`SELECT id,cidr,reason,source,enabled,hit_count,expires_at,last_hit_at,created_at FROM block_rules WHERE id=?`,id).Scan(&e.ID,&e.CIDR,&e.Reason,&e.Source,&e.Enabled,&e.HitCount,&e.ExpiresAt,&e.LastHitAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]BlockRule{rows,_:=d.db.Query(`SELECT id,cidr,reason,source,enabled,hit_count,expires_at,last_hit_at,created_at FROM block_rules ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []BlockRule;for rows.Next(){var e BlockRule;rows.Scan(&e.ID,&e.CIDR,&e.Reason,&e.Source,&e.Enabled,&e.HitCount,&e.ExpiresAt,&e.LastHitAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *BlockRule)error{_,err:=d.db.Exec(`UPDATE block_rules SET cidr=?,reason=?,source=?,enabled=?,hit_count=?,expires_at=?,last_hit_at=? WHERE id=?`,e.CIDR,e.Reason,e.Source,e.Enabled,e.HitCount,e.ExpiresAt,e.LastHitAt,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM block_rules WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM block_rules`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]BlockRule{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (cidr LIKE ?)"
        args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["source"];ok&&v!=""{where+=" AND source=?";args=append(args,v)}
    if v,ok:=filters["enabled"];ok&&v!=""{where+=" AND enabled=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,cidr,reason,source,enabled,hit_count,expires_at,last_hit_at,created_at FROM block_rules WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []BlockRule;for rows.Next(){var e BlockRule;rows.Scan(&e.ID,&e.CIDR,&e.Reason,&e.Source,&e.Enabled,&e.HitCount,&e.ExpiresAt,&e.LastHitAt,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    return m
}
