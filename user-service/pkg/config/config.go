package config

type Config struct {
    DBConnectionString string
}



func LoadConfig() (*Config, error) {
    // If the connection string is static, directly set it in the struct
    config := &Config{
        DBConnectionString: "postgres://u_tzt623sgywyqhd6:rekphmlozh3ce6g@02f7e6f1-1adb-4347-835a-02c74fcccb0e.db.cloud.postgresml.org:6432/pgml_m0rkbxzq1ewj2de",
    }

    return config, nil
}

func (c *Config) GetDBConnectionString() string {
    return c.DBConnectionString
}
