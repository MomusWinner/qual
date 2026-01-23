# Проект

Бэкенд-сервис с базой данных и метриками Prometheus.

## Быстрый старт

1. Скопируйте файл с переменными окружения:
```bash
cp example.env .env
```
2. Настройте переменные в .env файле

## Доступные команды
- Генерация SQL кода: make sql
- Генерация Swagger документации: make swagger
- Запуск интеграционных тестов: make test-integration

## Метрики Prometheus
- http_requests_total - количество HTTP запросов
- http_response_time_seconds - время обработки HTTP запросов
- db_response_time_seconds - время ответа от базы данных
