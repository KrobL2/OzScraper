# Scrappy

Данный бот умеет сохранять ссылки, которые ему скидывают собеседники, и по запросу отправлять
случайную ссылку из сохраненных.

Это полезно для тех людей, которые часто сохраняют много статей, но забывают их читать :)


### Запустить проект(old):

go run . -link "https://www.ozon.ru/product/interaktivnyy-igrovoy-nabor-syurpriz-mega-dom-uzhasov-smashers-ot-zuru-kollektsionnaya-igrushka-dlya-1662934903/"



###   Как запустить

1) Билд

go build -o scrappy


2) Запуск

./scrappy -tg-bot-token '6231102487:AAEJ_e0G9gKF9qbwyAjQ03ZiyJTtwoVAoMc'