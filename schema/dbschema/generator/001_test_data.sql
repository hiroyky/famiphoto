use `famiphoto_db`;

INSERT INTO `oauth_clients` (`oauth_client_id`, `name`, `client_secret`, `scope`, `client_type`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web', 'Famiphoto Web', 'JDJhJDEwJFBjWGZXeFQ5REQ5aHBvY3ZhV2RBcGVRbXFGT1NLMksxMDltY3B4bHl6bUJMRHZybWJqUTBt', 'Admin', 1, '2022-04-10 08:37:43', '2022-04-10 08:37:43');

INSERT INTO `oauth_client_redirect_urls` (`oauth_client_id`, `redirect_url`, `oauth_client_redirect_url_id`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web', 'http://localhost:3000/auth/redirect', 3, '2022-04-10 08:37:43', '2022-04-10 08:37:43');
