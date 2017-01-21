package option

import (
    "fmt"
    "time"
    "strconv"
    "math/rand"
    "crypto/md5"
    "encoding/hex"
    "monitor/collector/collection"
    "monitor/collector/model"
)

func (o *Options) Init() {
    
    DBNames := []string{collection.GATHER, collection.COLLOCATE}
    
    for _, Name := range DBNames {
        o.DBHandler.C(Name).DropCollection()
    }
    
    rand.Seed(time.Now().UnixNano())
    Md5 := md5.New()
    Md5.Write([]byte(strconv.Itoa(rand.Intn(10))))
    Token := hex.EncodeToString(Md5.Sum(nil))
    
    if nil == o.DBHandler.C(collection.COLLOCATE).Insert(model.Collocate{
        Name: "token",
        Value: Token,
    }) {
        fmt.Println("")
        return
    }
    
    fmt.Println("")
    return
}