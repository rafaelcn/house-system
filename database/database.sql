
CREATE TABLE IF NOT EXISTS TB_USER_TYPE (
    USER_TYPE_ID SERIAL PRIMARY KEY,
    USER_TYPE_NAME VARCHAR(255) NOT NULL,
    UNIQUE (USER_TYPE_NAME)
);

CREATE TABLE IF NOT EXISTS TB_USER (
    USER_ID SERIAL PRIMARY KEY,
    USER_NAME CHAR(50) NOT NULL,
    USER_EMAIL VARCHAR(100) NOT NULL,
    USER_USERNAME VARCHAR(20) DEFAULT '', 
    USER_PASSWORD VARCHAR(255) NOT NULL,
    USER_PHONE VARCHAR(12) NOT NULL,
    USER_BIRTH DATE NOT NULL,
    USER_TYPE INTEGER NOT NULL,
    CONSTRAINT FK_USER_TYPE
        FOREIGN KEY (USER_TYPE)
        REFERENCES TB_USER_TYPE(USER_TYPE_ID)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS TB_HOUSE (
    HOUSE_ID SERIAL PRIMARY KEY,
    HOUSE_NAME VARCHAR(255) NOT NULL,
    HOUSE_OWNER INTEGER NOT NULL,
    CONSTRAINT FK_HOUSE_OWNER
        FOREIGN KEY (HOUSE_OWNER)
        REFERENCES TB_USER(USER_ID)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS TB_OBJECT_TYPE (
    OBJECT_TYPE_ID SERIAL PRIMARY KEY,
    OBJECT_TYPE_NAME VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS TB_OBJECT (
    OBJECT_ID VARCHAR(255) PRIMARY KEY,
    OBJECT_NAME VARCHAR(255) NOT NULL,
    OBJECT_STATUS BOOLEAN NOT NULL,
    OBJECT_TYPE INTEGER NOT NULL,
    OBJECT_HOUSE INTEGER NOT NULL,
    OBJECT_ATTR_INTENSITY REAL DEFAULT 0.0,
    OBJECT_ATTR_DISTANCE REAL DEFAULT 0.0,
    OBJECT_ATTR_VOLUME REAL DEFAULT 0.0,
    OBJECT_ATTR_TEMPERATURE REAL DEFAULT 0.0,
    CONSTRAINT FK_OBJECT_HOUSE
        FOREIGN KEY (OBJECT_HOUSE)
        REFERENCES TB_HOUSE(HOUSE_ID)
        ON DELETE CASCADE,
    CONSTRAINT FK_OBJECT_TYPE
        FOREIGN KEY (OBJECT_TYPE)
        REFERENCES TB_OBJECT_TYPE(OBJECT_TYPE_ID)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS TB_INVITE (
    INVITE_ID SERIAL PRIMARY KEY,
    INVITE_SENDER INTEGER,
    INVITE_RECEIVER VARCHAR(255) NOT NULL,
    CONSTRAINT FK_INVITE_SENDER
        FOREIGN KEY (INVITE_SENDER)
        REFERENCES TB_USER(USER_ID)
        ON DELETE CASCADE,
    CONSTRAINT FK_INVITE_RECEIVER
        FOREIGN KEY (INVITE_RECEIVER)
        REFERENCES TB_USER(USER_ID)
        ON DELETE CASCADE
)