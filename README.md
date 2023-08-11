# WBL0
# В БД:
Развернуть локально postgresql
Создать свою бд
Настроить своего пользователя.
Создать таблицы для хранения полученных данных.
В сервисе:
1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
их запроса по id
# Доп инфо:
• Данные статичны, исходя из этого подумайте насчет модели хранения
в Кеше и в pg. Модель в файле model.json
• В канал могут закинуть что угодно, подумайте как избежать проблем
из-за этого
• Чтобы проверить работает ли подписка онлайн, сделайте себе
отдельный скрипт, для публикации данных в канал
• Подумайте как не терять данные в случае ошибок или проблем с
сервисом
• Nats-streaming разверните локально ( не путать с Nats )
