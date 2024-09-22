BEGIN;

DROP TABLE IF EXISTS read_permissions;
DROP TABLE IF EXISTS write_permissions;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS channels;
DROP TABLE IF EXISTS accounts_roles;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS roles;

DROP INDEX IF EXISTS idx_accounts_username;
DROP INDEX IF EXISTS idx_channels_name;
DROP INDEX IF EXISTS idx_messages_account_id;
DROP INDEX IF EXISTS idx_messages_channel_id;

COMMIT;
