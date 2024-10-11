# OzScraper (Read Adviser Bot)

Данный бот умеет сохранять ссылки, которые ему скидывают собеседники, и по запросу отправлять
случайную ссылку из сохраненных.

Это полезно для тех людей, которые часто сохраняют много статей, но забывают их читать :)

Код написан таким образом, чтобы его легко было расширять. К примеру, к нему без труда можно
добавить реализацию для любого другого мессенджера, добавив соотствующий клиент. Остальная логика
останется без изменений.



### Запустить проект:

go run . -link "https://www.ozon.ru/product/interaktivnyy-igrovoy-nabor-syurpriz-mega-dom-uzhasov-smashers-ot-zuru-kollektsionnaya-igrushka-dlya-1662934903/"



###   Как запустить

1) Билд

go build -o read-adviser-bot


2) Запуск

./read-adviser-bot -tg-bot-token '6231102487:AAEJ_e0G9gKF9qbwyAjQ03ZiyJTtwoVAoMc'