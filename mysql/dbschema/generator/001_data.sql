use `famiphoto_db`;

INSERT INTO `oauth_clients` (`oauth_client_id`, `name`, `client_secret`, `scope`, `client_type`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web', 'Famiphoto Web', '', 'Admin', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
