Архитектура и данные (Первый чек‑пойнт)

1. Диаграмма архитектуры
   
Компоненты:
```
Frontend (Vite + React): публичные страницы «О ЦОДД», «Новости», «Статистика», «Услуги», «Команда», «Проекты», экраны логина admin/editor. Общается с API через HTTPS JSON, axios.

Backend API (Go + Gin): REST, JWT‑аутентификация, ролевой доступ admin/editor, CORS, сбор и агрегация данных.

PostgreSQL: централизованная БД (штрафы, эвакуации, маршруты эвакуации, реестр светофоров, контент сайта, пользователи ролей). [attached_file:7eceee85-b83b-4bbf-a7d2-46669f640f9f]

Docker Compose: локальный стенд (postgres + pgadmin), миграции SQL применяются на старте.
```

Потоки данных:
```
Гости → Frontend → GET /api/* (публичные данные, без токена).

Admin/Editor → Frontend → POST /api/auth/* (получение JWT) → Protected /api/admin|/api/editor CRUD (с токеном). 

Backend ↔ PostgreSQL (CRUD, агрегации /api/stats, /api/traffic). [attached_file:7eceee85-b83b-4bbf-a7d2-46669f640f9f]
```

Развёртывание:
```

Dev: Vite (5173) ↔ Go API (8080) ↔ Postgres (5432), CORS для локальных доменов.

Prod: статический фронт (nginx/Cloud) + Go API, строгий CORS по белому списку доменов.
```

2. ER‑диаграмма (сущности и связи)

```
users(id, email, password_hash, role['admin'|'editor'], is_active, created_at, updated_at). 

news(id, title, content, tag, date, created_at, updated_at).

services(id, title, description, price, category, icon_url, created_at, updated_at).

team(id, name, position, experience, photo_url, created_at, updated_at).

projects(id, title, description, category, status, created_at, updated_at).

fines(id, date, violations_total, orders_total, fines_amount_total, collected_amount_total, created_at, updated_at)

evacuations(id, date, evacuators_count, trips_count, evacuations_count, fine_lot_income, created_at, updated_at)

evacuation_routes(id, year, month, route, created_at, updated_at)

traffic_lights(id, address, light_type, install_year, status, created_at, updated_at)
```



<!doctype html> <html lang="ru"> <head> <meta charset="utf-8"> <title>Локальный запуск — инструкция</title> <meta name="viewport" content="width=device-width, initial-scale=1"> <style> :root{ --bg:#0f172a; --card:#111827; --accent:#62a744; --muted:#9aa4b2; --text:#e5e7eb; --code:#0b1220; --border:#1f2937; } html,body{margin:0;padding:0;background:var(--bg);color:var(--text);font:16px/1.6 system-ui,-apple-system,Segoe UI,Roboto,Ubuntu,Cantarell,"Noto Sans","Helvetica Neue",Arial,"Apple Color Emoji","Segoe UI Emoji";} .wrap{max-width:920px;margin:48px auto;padding:0 20px;} .title{font-weight:800;font-size:28px;margin:0 0 20px;letter-spacing:.2px} .subtitle{color:var(--muted);margin:-6px 0 28px} .card{background:var(--card);border:1px solid var(--border);border-radius:14px;padding:22px 22px 6px;margin-bottom:18px;box-shadow:0 8px 28px rgba(0,0,0,.3)} h2{margin:0 0 14px;font-size:20px} ol{margin:0 0 8px 20px;padding:0} li{margin:8px 0} .block{margin:8px 0 18px} .badge{display:inline-block;font-weight:700;text-transform:uppercase;font-size:12px;letter-spacing:.06em;background:rgba(98,167,68,.15);color:var(--accent);border:1px solid rgba(98,167,68,.35);padding:5px 10px;border-radius:999px} pre{background:linear-gradient(180deg,rgba(17,24,39,.9),rgba(11,18,32,.95));border:1px solid var(--border);border-radius:12px;padding:14px 16px;overflow:auto;margin:10px 0 18px} code{font-family:ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,"Liberation Mono","Courier New",monospace;font-size:13.5px;color:#d1e7ff} .hint{color:var(--muted);font-size:14px;margin:8px 0 18px} .ok{color:#10b981} .warn{color:#f59e0b} .footer{color:var(--muted);font-size:13px;margin-top:26px} .kbd{font:12px ui-monospace,Menlo,Consolas,monospace;border:1px solid var(--border);background:#0b1220;padding:2px 6px;border-radius:6px;color:#cbd5e1} a{color:var(--accent);text-decoration:none} a:hover{text-decoration:underline} </style> </head> <body> <div class="wrap"> <h1 class="title">Локальный запуск проекта</h1> <p class="subtitle">Быстрые шаги для поднятия базы, бэкенда и фронтенда в режиме разработки.</p>
text
<div class="card">
  <span class="badge">Шаг 1</span>
  <h2>База данных (PostgreSQL + миграции)</h2>
  <div class="block">
    <pre><code>cd backend
docker-compose up -d postgres</code></pre>
</div>
<p class="hint">Будет поднят контейнер <span class="kbd">postgres</span>, а миграции из <span class="kbd">migrations/init.sql</span> применятся автоматически. Дефолт: host <span class="kbd">localhost</span>, порт <span class="kbd">5432</span>, пользователь <span class="kbd">postgres</span>, пароль <span class="kbd">password</span>, база <span class="kbd">smolathon_db</span>.</p>
</div>

text
<div class="card">
  <span class="badge">Шаг 2</span>
  <h2>Бэкенд (Go + Gin)</h2>
  <div class="block">
    <pre><code>cd backend
go mod tidy
go run cmd/backend/main.go</code></pre>
</div>
<p class="hint ok">API поднимется на <span class="kbd">http://localhost:8080</span>. Конфигурация берётся из <span class="kbd">backend/.env</span> (CORS, БД, JWT и т.п.).</p>
</div>

text
<div class="card">
  <span class="badge">Шаг 3</span>
  <h2>Фронтенд (Vite + React)</h2>
  <div class="block">
    <pre><code>cd frontend/Smolathon_nextgen
echo "VITE_API_BASE=http://localhost:8080" > .env.local
npm install
npm run dev</code></pre>
</div>
<p class="hint ok">Dev‑сервер доступен на <span class="kbd">http://localhost:5173</span>. Переменная <span class="kbd">VITE_API_BASE</span> указывает адрес бэкенда.</p>
<p class="hint warn">Если браузер показывает CORS‑ошибку — добавьте домен фронтенда в <span class="kbd">ALLOWED_ORIGINS</span> внутри <span class="kbd">backend/.env</span> и перезапустите бэкенд.</p>
</div>

text
<p class="footer">Подсказка: для продакшн‑сборки фронта используйте <span class="kbd">npm run build</span>, статические файлы появятся в директории <span class="kbd">dist/</span>.</p>
</div> </body> </html>
