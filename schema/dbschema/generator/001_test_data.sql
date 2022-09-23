use `famiphoto_db`;

INSERT INTO `oauth_clients` (`oauth_client_id`, `name`, `client_secret`, `scope`, `client_type`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web', 'Famiphoto Web', 'JDJhJDEwJFBjWGZXeFQ5REQ5aHBvY3ZhV2RBcGVRbXFGT1NLMksxMDltY3B4bHl6bUJMRHZybWJqUTBt', 'Admin', 1, '2022-04-10 08:37:43', '2022-04-10 08:37:43');

INSERT INTO `oauth_client_redirect_urls` (`oauth_client_id`, `redirect_url`, `oauth_client_redirect_url_id`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web', 'http://localhost:3000/auth/redirect', 3, '2022-04-10 08:37:43', '2022-04-10 08:37:43');


INSERT INTO `users` (`user_id`, `name`, `status`, `created_at`, `updated_at`)
VALUES
    ('hiro', 'ヒロ', 1, '2022-04-10 09:17:22', '2022-04-10 09:17:22');

-- password: password1
INSERT INTO `user_passwords` (`user_id`, `password`, `last_modified_at`, `is_initialized`, `created_at`, `updated_at`)
VALUES
    ('hiro', 'JDJhJDEwJEoyWVNFcFptNkticjlmNEE0Qmc1Mk9CdjFUY1pFTWRDSUd2cEhmN1VoclhsTDZxcWFQaUUu', '2022-04-10 09:17:22', 1, '2022-04-10 09:17:22', '2022-04-10 09:17:22');

INSERT INTO `groups` (`group_id`, `name`, `created_at`, `updated_at`)
VALUES
    ('yokoyama', 'よこやま', '2022-04-21 18:22:30', '2022-04-21 18:22:30');

INSERT INTO `group_users` (`group_id`, `user_id`, `created_at`, `updated_at`)
VALUES
    ('yokoyama', 'hiro', '2022-04-21 18:22:48', '2022-04-21 18:22:48');
