CREATE TABLE IF NOT EXISTS Servers(
    ServerID SERIAL NOT NULL PRIMARY KEY,
    Game VARCHAR(32) NOT NULL,
    Name VARCHAR(20) NOT NULL,
    Address VARCHAR(15) NOT NULL,
    RCONPort VARCHAR(5) NOT NULL,
    RCONPassword VARCHAR(128) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ModifiedAt TIMESTAMP
);

CREATE TRIGGER update_servers_modat BEFORE UPDATE ON Servers
    FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();