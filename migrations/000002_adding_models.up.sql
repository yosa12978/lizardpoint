BEGIN;

CREATE TABLE roles (
    name VARCHAR NOT NULL, 
    PRIMARY KEY (name)
);

CREATE TABLE accounts (
    id uuid DEFAULT gen_random_uuid(),
    username VARCHAR(32) NOT NULL,
    password_hash VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_accounts_username ON accounts(username);

CREATE TABLE accounts_roles (
    account_id uuid NOT NULL,
    role_name VARCHAR NOT NULL,
    FOREIGN KEY (role_name) REFERENCES roles(name) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE,
    PRIMARY KEY (account_id, role_name)
);

CREATE TABLE channels (
    id uuid DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_channels_name ON channels(name);

CREATE TABLE messages (
    id uuid DEFAULT gen_random_uuid(),
    content VARCHAR NOT NULL,
    account_id uuid NOT NULL,
    channel_id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE NO ACTION, 
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_messages_channel_id ON messages(channel_id);
CREATE INDEX IF NOT EXISTS idx_messages_account_id ON messages(account_id);

CREATE TABLE write_permissions (
    channel_id uuid NOT NULL,
    role_name VARCHAR NOT NULL,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (role_name) REFERENCES roles(name) ON DELETE CASCADE,
    PRIMARY KEY (channel_id, role_name) 
);

CREATE TABLE read_permissions (
    channel_id uuid NOT NULL,
    role_name VARCHAR NOT NULL,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (role_name) REFERENCES roles(name) ON DELETE CASCADE,
    PRIMARY KEY (channel_id, role_name) 
);

COMMIT;
