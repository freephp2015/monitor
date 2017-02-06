package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "runtime"
    "encoding/json"
    "monitor/daemon"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

const (
    READ_LENGTH = 2048
)

type DBAuth struct {
    Host     string `json:"host"`
    Port     int `json:"port"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func Scheduler(Con *net.UnixConn) {
    for {
        Buffer := make([]byte, READ_LENGTH)
        Len, err := Con.Read(Buffer);
        if err != nil {
            Con.Close()
            return
        }
        var Message protocols.Socket
        json.Unmarshal(Buffer[0: Len], &Message)
        // todo 接收到cli信息,然后处理
        log.Println(Message)
    }
}

func main() {
    var (
        // 配置文件
        ConfFile string
        
        PidFile string
        LogFile string
    )
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    Viper := viper.GetViper()
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor daemon",
        Long: "######to do...",
        RunE:func(cmd *cobra.Command, args []string) error {
            
            Conf := configures.Initialize(Viper, ConfFile)
    
            fmt.Println(*Conf)
            // TODO run ...
            Daemon := &daemon.Daemon{
                PidFile: "/var/run/monitord.pid",
                UnixFile: "/var/run/monitord.sock",
                LogFile: "/var/log/monitord.log",
            }
            
            Daemon.Daemon(func(Unix *net.UnixListener) {
                defer Unix.Close()
                
                for {
                    if Fd, err := Unix.AcceptUnix(); err != nil {
                        log.Printf("%v\r\n", err)
                    } else {
                        go Scheduler(Fd)
                    }
                }
            })
            return nil
        },
    }
    
    Flags := RootCmd.Flags()
    
    // 配置文件路径
    Flags.StringVarP(&ConfFile, "config", "c", "/etc/monitor.yaml", "monitor config file path")
    
    Flags.StringVarP(&PidFile, "pid", "", "/var/run/monitord.pid", "monitor daemon pid file")
    Flags.StringVarP(&LogFile, "log", "l", "/var/log/monitord.log", "monitor daemon log file")
    // todo socket file and mongodb config
    
    Viper.BindPFlag("server.pid_file", Flags.Lookup("pid"))
    Viper.BindPFlag("server.log_file", Flags.Lookup("log"))
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}