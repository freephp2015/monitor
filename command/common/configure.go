package common

import (
    "strings"
    "path/filepath"
)

type Server struct {
    Daemon bool
    Pid    string
    Unix   string
    Log    string
    Addr   string
}

type Client struct {
    Unix string
}

type Database struct {
    Host     string `json:"host"`
    Port     int32  `json:"port"`
    Auth     string `json:"auth"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Configure struct {
    Server  Server
    Client  Client
    MongoDB Database
}

func (c *Command) ReadConf() error {
    Dir, File := filepath.Split(c.Flags.Main.Config)
    Ext := filepath.Ext(File)
    
    c.Viper.SetConfigName(strings.Replace(File, Ext, "", 1))
    c.Viper.SetConfigType(strings.Trim(Ext, "."))
    c.Viper.AddConfigPath(Dir)
    
    if err := c.Viper.ReadInConfig(); err != nil {
        return err
    }
    
    c.Configure = &Configure{
        Server: Server{
            Daemon: c.Viper.GetBool("server.daemon"),
            Pid:  c.Viper.GetString("server.pid"),
            Log: c.Viper.GetString("server.log"),
            Unix: c.Viper.GetString("server.unix"),
        },
        Client: Client{
            Unix: c.Viper.GetString("client.unix"),
        },
        MongoDB: Database{
            Host: c.Viper.GetString("mongodb.host"),
            Port: int32(c.Viper.GetInt("mongodb.port")),
            Auth: c.Viper.GetString("mongodb.auth"),
            Username: c.Viper.GetString("mongodb.username"),
            Password: c.Viper.GetString("mongodb.password"),
        },
    }
    
    return nil
}

