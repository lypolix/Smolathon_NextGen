
# Архитектура и данные (Первый чек‑пойнт)

## 1. Диаграмма архитектуры

### Компоненты

```
Frontend (Vite + React): публичные страницы «О ЦОДД», «Новости», «Статистика», «Услуги», «Команда», «Проекты», экраны логина admin/editor. Общается с API через HTTPS JSON, axios.

Backend API (Go + Gin): REST, JWT‑аутентификация, ролевой доступ admin/editor, CORS, сбор и агрегация данных.

PostgreSQL: централизованная БД (штрафы, эвакуации, маршруты эвакуации, реестр светофоров, контент сайта, пользователи ролей).

Docker Compose: локальный стенд (postgres + pgadmin), миграции SQL применяются на старте.

```

### Потоки данных

```
Гости → Frontend → GET /api/* (публичные данные, без токена).

Admin/Editor → Frontend → POST /api/auth/* (получение JWT) → Protected /api/admin | /api/editor CRUD (с токеном).

Backend ↔ PostgreSQL (CRUD, агрегации /api/stats, /api/traffic).

```

### Развёртывание

```
Dev: Vite (5173) ↔ Go API (8080) ↔ Postgres (5432), CORS для локальных доменов.

Prod: статический фронт (nginx/Cloud) + Go API, строгий CORS по белому списку доменов.
```

## 2. ER‑диаграмма (сущности и связи)

```

users(id, email, password_hash, role['admin'|'editor'], is_active, created_at, updated_at)

news(id, title, content, tag, date, created_at, updated_at)

services(id, title, description, price, category, icon_url, created_at, updated_at)

team(id, name, position, experience, photo_url, created_at, updated_at)

projects(id, title, description, category, status, created_at, updated_at)

fines(id, date, violations_total, orders_total, fines_amount_total, collected_amount_total, created_at, updated_at)

evacuations(id, date, evacuators_count, trips_count, evacuations_count, fine_lot_income, created_at, updated_at)

evacuation_routes(id, year, month, route, created_at, updated_at)

traffic_lights(id, address, light_type, install_year, status, created_at, updated_at)

# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])



