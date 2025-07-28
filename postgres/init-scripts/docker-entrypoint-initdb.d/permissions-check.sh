#!/bin/sh
set -e

echo "Проверка прав доступа:"
ls -la /docker-entrypoint-initdb.d

echo "Тестовый файл создан" > /test_output.txt