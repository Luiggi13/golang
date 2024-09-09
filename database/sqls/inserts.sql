INSERT INTO tags ("name", public) VALUES('tag1', false);
INSERT INTO tags ("name", public) VALUES('tag2', true);
INSERT INTO tags ("name", public) VALUES('tag3', false);

INSERT INTO qrs (qr_code, userid, url_text, premium, id_tag, created_at) VALUES('code1rrrr', 'userid1', 'www.google.es', false, 1, '2024-09-07 20:28:56.995');
INSERT INTO qrs (qr_code, userid, url_text, premium, id_tag, created_at) VALUES('qrcode2', 'userid2', 'www.code.dev', true, 2, '2024-09-07 20:29:31.543');
INSERT INTO qrs (qr_code, userid, url_text, premium, id_tag, created_at) VALUES('qrcode3', 'userid2', 'www.code.dev', true, 3, '2024-09-07 20:29:31.543');