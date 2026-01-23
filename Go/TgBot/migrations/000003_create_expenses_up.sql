CREATE table expenses (
id SERIAL PRIMARY key,
user_id BIGINT NOT NULL references users(user_id),
category VARCHAR(50) NOT NULL,
amount DECIMAL,
text VARCHAR(100),
created_at TIMESTAMP default CURRENT_TIMESTAMP
);