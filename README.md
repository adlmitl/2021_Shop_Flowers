# Shop Flowers
First way run pet-project
1. For build project need install docker.
2. Run project command <code>'docker build -t go-docker .'</code>.
3. After created the image, check it in the list with the command <code>'docker images'</code>.
4. Create and run container command <code>'docker run --rm --name shopflowers -d -p 8000:8080 go-docker'</code>.

Second way run pet-project
1. Add file docker-compose for fast start pet-project. Start command <code>'docker-compose up -d'</code>





# FGW - Finish Goods Warehouse (Склад готовой продукции).

## Стиль именования коммитов в git

### Для именования коммитов можно воспользоваться следующими правилами:
1. Разделяй название коммита с описанием пустой стройкой:
   <br>Пример:
   <pre>Обновить документацию 

   Необходимо добавить новые комментарии к публичным методам класса Person.</pre>

2. Название коммита не должно превышать 50 символов;

3. Название коммита должно начинаться с большой буквы:
   <pre>Правильно: Accelerate to 88 miles per hour
   Не правильно: аccelerate to 88 miles per hour</pre>

4. Не нужно ставить знаки препинания в конце названия коммита:
   <pre>Правильно: Accelerate to 88 miles per hour
   Не правильно: аccelerate to 88 miles per hour.</pre>

5. Длина одной строки в описании не должна превышать 72 символа

6. В описании отвечайте на вопрос: <br> «Что сделать? и Зачем это делать?»

Пример коммита:

<pre>Создать проект какой-то

    1. создать архитектуру
    2. создать подключение к БД
    3. настроить передачу данных в json формате</pre>