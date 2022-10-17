CREATE TABLE wagers (
  id bigserial primary key,
  total_wager_value integer NOT NULL,
  odds integer NOT NULL,
  selling_percentage integer NOT NULL,
  selling_price decimal(20,2) NOT NULL,
  current_selling_price decimal(20,2) NOT NULL,
  percentage_sold integer,
  amount_sold integer,
  placed_at bigint
);

CREATE TABLE wager_txn_logs (
  id bigserial primary key,
  wager_id bigint NOT NULL,
  user_id char(128),
  action integer NOT NULL,
  amount decimal(20,2) NOT NULL,
  post_selling_price decimal(20,2) NOT NULL,
  bought_at bigint
);