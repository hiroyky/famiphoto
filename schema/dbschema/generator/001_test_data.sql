use `famiphoto_db`;

INSERT INTO `oauth_clients` (`oauth_client_id`, `name`, `client_secret`, `scope`, `client_type`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web1', 'Famiphoto Web', 'JDJhJDEwJFo2VkRBUlBrZkFUSjN2SmtzcFI0bnVyWDJlUWdLWVNQam5RMUcyMnVjdVVyOThHLk9KSGJL', 'Admin', 2, '2022-03-31 16:46:33', '2022-03-31 16:46:33'),
    ('famiphoto_web2', 'Famiphoto Web', 'JDJhJDEwJFphQU5OcFZWMGc2THJUemtNaTVUbC4zZXBhU3Nsa2RoUTNvQTdGWEpEN1BzLmwzbS5XejMu', 'Admin', 1, '2022-03-31 16:48:27', '2022-03-31 16:48:27');

INSERT INTO `oauth_client_redirect_urls` (`oauth_client_id`, `redirect_url`, `oauth_client_redirect_url_id`, `created_at`, `updated_at`)
VALUES
    ('famiphoto_web2', 'http://localhost:3000/auth/login', 2, '2022-03-31 16:48:27', '2022-03-31 16:48:27');
