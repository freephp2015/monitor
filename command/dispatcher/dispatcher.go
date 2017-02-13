package dispatcher

import (
    "net"
    "fmt"
    "encoding/json"
    "monitor/command/common"
    "monitor/command/protocol"
    "monitor/monitor"
)

type Dispatcher struct {
    Message protocol.SocketMsg
    Conn    *net.UnixConn
}

func (dis *Dispatcher) Res(Code int, Msg string) (int, error) {
    Res := protocol.Response{
        Code: Code,
        Body: []byte(Msg),
    }
    Byte, _ := json.Marshal(Res)
    return dis.Conn.Write(Byte)
}

func Run(Msg protocol.SocketMsg, Conn *net.UnixConn, Monitor *monitor.Monitor) {
    Dis := &Dispatcher{
        Conn: Conn,
        Message: Msg,
    }
    
    switch Dis.Message.Command {
    case common.CMD_ROLE:
    
    case common.CMD_SERVER_INIT:
        err := Monitor.SInit(Dis.Message);
        if err != nil {
            Dis.Res(-1, fmt.Sprintf("%v", err))
            return
        }
        Dis.Res(-1, "success")
        return
    case common.CMD_SERVER_TOKEN:
    
    case common.CMD_JOIN:
    
    default:
        
    }
}
