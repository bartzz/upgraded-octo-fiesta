# Currency API

Cześć! Dzięki za poświęcenie czasu na sprawdzenie mojego zadania.

**Uruchamianie**: `make start-api` - zbuduje aplikację, uruchomi za pomocą `docker compose` w trybie detached (-d) i zacznie tailować logi `docker compose logs -f`.

**Wyłączanie**:

`make stop-api` - po prostu alias do `docker compose down`

## Proces myślowy, notatki

- Pominąłem opisywanie tego projektu w tym README, każdy wie o co chodzi.
- Na produkcyjnym repo nie zostawiłbym ważnego klucza API do płatnego serwisu, przeważnie trzyma się to w AWS secrets, Kubernetes secrets itd. Dla ułatwienia Wam sprawdzenia pozostawiłem swój klucz z "free tier".
- Obecny serwis `rates_service.go` działa poprawnie, jednak w przypadku dużego RPS warto byłoby ograniczyć liczbę zapytań do płatnego API poprzez cache’owanie danych np. w Redisie. Zakładając horyzontalne skalowanie aplikacji (np. przez K8s), trzymanie cache’u lokalnie przy pomocy `sync.RWMutex` mogłoby prowadzić do niespójności przy load balancerze np typu round robin, dlatego współdzielony cache byłby bezpieczniejszym rozwiązaniem. Tutaj jest sporo przemyśleń, to zależy od wielu czynników.
- Dodałem `type ExchangeRatesProvider interface` żeby nie uzależniać całego kodu od jednego providera kursów, można sobie napisać nowy i podmieniać za pomocą DI.
- Nie wspominaliście nic o loggerze, acz w przypadku napotkania błędów, np podczas callowania zew. API fajnie byłoby to logować np logrusem żeby wiedzieć że coś się wykrzaczyło. Także to pozostawiam w mistycznej sekcji // TODO :) A gdyby to była ważna apka produkcyjna to idealnie byłoby wystawić `/metrics`, liczyć każdą awarię zapytania do API, podpiąć Prometheusa, do tego Grafanę i od razu byłoby widać że provider do kursów którego używamy zaczyna słabo działać. (https://github.com/prometheus/client_golang)
- W kontekście drugiego "opcjonalnego" endpointa `/exchange` dorzuciłem [mema](./gratisowy_meme.png) który od razu mi się skojarzył.
- Można by dorzucić jeszcze więcej testów, szczególnie np. klienta `openexchangerates.go`

Dzięki i pozdrawiam,

Bartłomiej Szabłowski.