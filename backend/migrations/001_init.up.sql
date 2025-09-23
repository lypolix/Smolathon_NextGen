-- Создание таблиц для базы данных

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица штрафов (из Excel)
CREATE TABLE IF NOT EXISTS fines (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    violations_total INTEGER NOT NULL,
    orders_total INTEGER NOT NULL,
    fines_amount_total BIGINT NOT NULL,
    collected_amount_total BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица эвакуации (из Excel)
CREATE TABLE IF NOT EXISTS evacuations (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    evacuators_count INTEGER NOT NULL,
    trips_count INTEGER NOT NULL,
    evacuations_count INTEGER NOT NULL,
    fine_lot_income BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица маршрутов эвакуации (из Excel)
CREATE TABLE IF NOT EXISTS evacuation_routes (
    id SERIAL PRIMARY KEY,
    year INTEGER NOT NULL,
    month VARCHAR(50) NOT NULL,
    route TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица светофоров (из Excel)
CREATE TABLE IF NOT EXISTS traffic_lights (
    id SERIAL PRIMARY KEY,
    address VARCHAR(500) NOT NULL,
    light_type VARCHAR(50) NOT NULL,
    install_year INTEGER NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Остальные таблицы (существующие)
CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    tag VARCHAR(50) NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price INTEGER NOT NULL,
    category VARCHAR(100) NOT NULL,
    icon_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS team (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    experience TEXT NOT NULL,
    photo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Вставка тестовых данных
INSERT INTO users (username, email, password, role) VALUES 
('admin', 'admin@smolensk.ru', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin'),
('editor', 'editor@smolensk.ru', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'editor');

-- Вставка данных из Excel файла (примеры)
INSERT INTO fines (date, violations_total, orders_total, fines_amount_total, collected_amount_total) VALUES
('2024-01-01', 919, 832, 608816, 0),
('2024-01-02', 1733, 1576, 1153298, 122300),
('2024-01-03', 2800, 2581, 2041916, 339416);

INSERT INTO evacuations (date, evacuators_count, trips_count, evacuations_count, fine_lot_income) VALUES
('2024-01-01', 1, 3, 2, 10000),
('2024-01-02', 2, 5, 4, 20000),
('2024-01-03', 2, 7, 6, 30000);

INSERT INTO evacuation_routes (year, month, route) VALUES
(2024, 'Январь', 'ул. Большая Советская (д.1-25) → ул. Ленина (д.10-40) → пр-т Гагарина (д.5-35) → ул. Николаева (д.3-22) → ул. Багратиона (д.7-18)'),
(2024, 'Февраль', 'ул. Дзержинского (д.1-15) → ул. Маршала Конева (д.4-28) → ул. Ново-Ленинградская (д.6-32) → ул. Рыленкова (д.2-19) → ул. Кашена (д.8-24)');

INSERT INTO traffic_lights (address, light_type, install_year, status) VALUES
('ул. Большая Советская / ул. Ленина', 'Т.1', 2018, 'active'),
('ул. Николаева / ул. Кашена', 'Т.2', 2015, 'active'),
('пр-т Гагарина / ул. Багратиона', 'Т.3', 2020, 'active');

INSERT INTO news (title, content, tag) VALUES
('Новые камеры контроля скорости установлены', 'Больше камер — меньше аварий', 'Камеры'),
('Центр запускает онлайн-сервис проверки штрафов', 'Штрафы можно проверить за минуту', 'Штрафы'),
('Обновлена схема движения в центре города', 'Новые маршруты снизят пробки', 'Движения'),
('Стартовала программа безопасных пешеходных переходов', 'Освещение и разметка — безопасность людей', 'Пешеходы');

INSERT INTO services (title, description, price, category) VALUES
('Платная справка "чистоты автомобиля"', 'Проверка авто перед покупкой на наличие штрафов, ограничений, долгов, арестов', 10000, 'auto_check'),
('Сервис "Автопомощь"', 'Услуги эвакуатора по вызову через центр', 4000, 'evacuation'),
('Оформление страховок', 'Подбор подходящей страховой компании', 5000, 'insurance'),
('Автодокументы под ключ', 'Заполнение заявлений для МФЦ и ГИБДД', 10000, 'documents');

INSERT INTO team (name, position, experience) VALUES
('Иван Иванов', 'Начальник отдела', 'Стаж: 15 лет'),
('Петр Петров', 'Инженер', 'Стаж: 8 лет'),
('Сидор Сидоров', 'Аналитик', 'Стаж: 5 лет');

INSERT INTO projects (title, description, category) VALUES
('Чистый автомобиль', 'Команда ЦОДД проверяет автомобили через официальные базы и формирует прозрачные отчёты', 'auto_check'),
('Штрафы без очередей', 'Специалисты ЦОДД создали удобный сервис, где жители могут сразу узнать о штрафах и оплатить их онлайн', 'fines'),
('Помощь на дороге', 'Сотрудники ЦОДД обеспечивают круглосуточную поддержку водителей: эвакуатор, консультация или помощь при ДТП', 'roadside'),
('Умные парковки', 'Команда ЦОДД внедряет технологии, которые показывают свободные парковочные места и помогают жителям меньше тратить времени в поисках и избегать штрафов', 'parking');
