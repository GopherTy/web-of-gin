{
    DB: {
        Driver: "mysql",
        Source: "ty:123@tcp(localhost:3306)/gopherty?charset=utf8",
        ShowSQL: true,
        UserManageDisable: true,
        MaxOpenConns: 100,
        MaxIdleConns: 5,
        Cached: 200
    },
    Server: {
        Address: ":8080",
        Release: false,
        CertFile: "",
        KeyFile: "",
        LogsPath: ""
    },
    Logger: {
        Level: "debug",
        Encoding: "console",
        Development: false,
        LogsPath: ""
    }
}