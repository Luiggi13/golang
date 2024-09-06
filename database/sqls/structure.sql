CREATE TABLE IF NOT EXISTS qrs (
  id bigint primary key generated always as identity,
  qr_code text,
  userId text,
  url_text text,
  premium bool DEFAULT false null,
  created_at timestamp DEFAULT now()
);
