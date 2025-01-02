# 全てコマンドはプロジェクトルートで実行すること
# 構築
build: 
	docker compose build
buildn:
	docker compose build --no-cache

# 起動
up:
	docker compose up
upd:
	docker compose up -d

# 再起動
re:
	docker compose restart

# シャットダウン
down:
	docker compose down --remove-orphans

# コンテナ内に入る
b:
	docker compose exec backend sh

n:
	docker compose exec nginx sh

d:
	docker compose exec db sh

# 開発環境構築
init:
	cp ./.env.example ./.env && \
	make buildn && \
	make upd

# フロントエンド
front:
	cd ./frontend && \
	npm start

# バックエンド
api:
	cd ./backend && \
	swag init
