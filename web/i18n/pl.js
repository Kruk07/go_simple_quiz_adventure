window.quizLocaleData = window.quizLocaleData || {};

(function () {
  const categories = [
    { id: 'cat-programming', name: 'Programowanie' },
    { id: 'cat-history', name: 'Historia' },
    { id: 'cat-science', name: 'Nauka' },
    { id: 'cat-sports', name: 'Sport' },
    { id: 'cat-geography', name: 'Geografia' },
    { id: 'cat-literature', name: 'Literatura' },
    { id: 'cat-movies', name: 'Filmy' },
    { id: 'cat-music', name: 'Muzyka' },
    { id: 'cat-math', name: 'Matematyka' },
    { id: 'cat-nature', name: 'Natura' },
    { id: 'cat-technology', name: 'Technologia' },
    { id: 'cat-art', name: 'Sztuka' },
    { id: 'cat-food', name: 'Jedzenie' },
    { id: 'cat-animals', name: 'Zwierzęta' },
    { id: 'cat-space', name: 'Kosmos' },
    { id: 'cat-health', name: 'Zdrowie' },
    { id: 'cat-language', name: 'Język' },
    { id: 'cat-psychology', name: 'Psychologia' },
    { id: 'cat-religion', name: 'Religia' },
    { id: 'cat-business', name: 'Biznes' },
    { id: 'cat-gaming', name: 'Gry' },
    { id: 'cat-cars', name: 'Samochody' },
    { id: 'cat-astronomy', name: 'Astronomia' },
    { id: 'cat-philosophy', name: 'Filozofia' },
    { id: 'cat-design', name: 'Design' }
  ];

  const categoryQuestionBanks = {
    'cat-programming': [
      { text: 'Które słowo kluczowe deklaruje funkcję w Go?', options: ['func', 'let', 'class', 'def'], correctOption: 'A' },
      { text: 'Do czego Go używa do równoległego wykonywania kodu?', options: ['gorutyn', 'wątków', 'klas', 'modułów'], correctOption: 'A' },
      { text: 'Która struktura danych przechowuje pary klucz-wartość w Go?', options: ['map', 'stack', 'queue', 'array'], correctOption: 'A' },
      { text: 'Które polecenie uruchamia program Go z katalogu modułu?', options: ['go run .', 'npm start', 'cargo run', 'python app.py'], correctOption: 'A' },
      { text: 'Które słowo kluczowe służy do importowania pakietu w Go?', options: ['import', 'include', 'require', 'using'], correctOption: 'B' },
      { text: 'Jaki typ jest najczęściej używany do tekstu w Go?', options: ['string', 'int', 'bool', 'char'], correctOption: 'A' },
      { text: 'Które narzędzie Go kompiluje pakiety i zależności?', options: ['go build', 'gcc', 'javac', 'tsc'], correctOption: 'A' },
      { text: 'Czym jest slice w Go?', options: ['dynamiczna sekwencja wartości', 'pojedyncza liczba całkowita', 'skompilowany plik', 'rekord bazy danych'], correctOption: 'A' },
      { text: 'Który typ w Go przechowuje pojedynczą wartość prawda/fałsz?', options: ['bool', 'string', 'int', 'byte'], correctOption: 'A' },
      { text: 'Które słowo kluczowe Go tworzy nową zmienną z wywnioskowanym typem?', options: ['var', 'class', 'def', 'new'], correctOption: 'A' }
    ],
    'cat-history': [
      { text: 'Kto był pierwszym prezydentem Stanów Zjednoczonych?', options: ['George Washington', 'Thomas Jefferson', 'John Adams', 'Abraham Lincoln'], correctOption: 'A' },
      { text: 'W którym roku runął Mur Berliński?', options: ['1989', '1987', '1991', '1979'], correctOption: 'A' },
      { text: 'Które imperium było rządzone przez Czyngis-chana?', options: ['Imperium mongolskie', 'Imperium rzymskie', 'Imperium osmańskie', 'Imperium bizantyjskie'], correctOption: 'A' },
      { text: 'Magna Charta została podpisana w którym roku?', options: ['1215', '1315', '1415', '1115'], correctOption: 'A' },
      { text: 'Która starożytna cywilizacja zbudowała piramidy w Gizie?', options: ['Egipcjanie', 'Rzymianie', 'Grecy', 'Wikingowie'], correctOption: 'A' },
      { text: 'Który konflikt toczył się pomiędzy Północą a Południem w Stanach Zjednoczonych?', options: ['Amerykańska wojna secesyjna', 'I wojna światowa', 'Wojna rewolucyjna', 'Wojna z 1812 roku'], correctOption: 'A' },
      { text: 'Kto odkrył Amerykę w 1492 roku?', options: ['Krzysztof Kolumb', 'Ferdynand Magellan', 'Marco Polo', 'Amerigo Vespucci'], correctOption: 'A' },
      { text: 'Które miasto zostało zniszczone przez wybuch Wezuwiusza?', options: ['Pompeje', 'Rzym', 'Ateny', 'Neapol'], correctOption: 'A' },
      { text: 'Który dokument ustanowił ramy rządów w USA?', options: ['Konstytucja', 'Deklaracja niepodległości', 'Artykuły Konfederacji', 'Prawa do wstępu'], correctOption: 'A' },
      { text: 'Który król został zmuszony do podpisania Magna Charta?', options: ['Jan bez Ziemi', 'Henryk VIII', 'Edward I', 'Ryszard III'], correctOption: 'A' }
    ],
    'cat-science': [
      { text: 'Która planeta jest znana jako Czerwona Planeta?', options: ['Mars', 'Wenus', 'Merkury', 'Jowisz'], correctOption: 'A' },
      { text: 'Jaki jest wzór chemiczny wody?', options: ['H2O', 'CO2', 'NaCl', 'O2'], correctOption: 'A' },
      { text: 'Jaka siła utrzymuje planety na orbitach wokół Słońca?', options: ['Grawitacja', 'Magnetyzm', 'Elektryczność', 'Tarcie'], correctOption: 'A' },
      { text: 'Która część komórki zawiera materiał genetyczny?', options: ['Jądro', 'Błona komórkowa', 'Rybozom', 'Cytoplazma'], correctOption: 'A' },
      { text: 'Który gaz rośliny pobierają z atmosfery?', options: ['Dwutlenek węgla', 'Tlen', 'Azot', 'Wodór'], correctOption: 'A' },
      { text: 'Jak nazywa się środek atomu?', options: ['Jądro', 'Elektron', 'Proton', 'Neutron'], correctOption: 'A' },
      { text: 'Jaka jest temperatura wrzenia wody na poziomie morza?', options: ['100°C', '90°C', '80°C', '110°C'], correctOption: 'A' },
      { text: 'Który narząd pompuje krew po organizmie?', options: ['Serce', 'Wątroba', 'Płuco', 'Nerka'], correctOption: 'A' },
      { text: 'Jaki gaz ludzie wydychają w dużych ilościach?', options: ['Dwutlenek węgla', 'Tlen', 'Hel', 'Argon'], correctOption: 'A' },
      { text: 'Który ptak słynie z naśladowania dźwięków?', options: ['Papuga', 'Orzeł', 'Pingwin', 'Sowa'], correctOption: 'A' }
    ],
    'cat-sports': [
      { text: 'Ilu zawodników jest na boisku w jednej drużynie piłki nożnej?', options: ['11', '9', '10', '12'], correctOption: 'A' },
      { text: 'W baseballu czym jest home run?', options: ['Biegi zdobyte po uderzeniu poza boisko', 'Strikeout', 'Foul ball', 'Double play'], correctOption: 'A' },
      { text: 'Ile minut trwa standardowy mecz piłki nożnej?', options: ['90', '60', '75', '120'], correctOption: 'A' },
      { text: 'Który sport używa lotki?', options: ['Badminton', 'Tenis', 'Squash', 'Krykiet'], correctOption: 'A' },
      { text: 'Ile kół jest na fladze olimpijskiej?', options: ['5', '4', '6', '7'], correctOption: 'A' },
      { text: 'Który sport jest grany pięcioma zawodnikami na boisku?', options: ['Koszykówka', 'Siatkówka', 'Baseball', 'Rugby'], correctOption: 'A' },
      { text: 'Który kraj wygrał Mistrzostwa Świata FIFA w 2018 roku?', options: ['Francja', 'Niemcy', 'Brazylia', 'Argentyna'], correctOption: 'A' },
      { text: 'W tenisie wynik zero nazywa się?', options: ['Love', 'Nil', 'Blank', 'Zero'], correctOption: 'A' },
      { text: 'Jaki jest maksymalny wynik w jednej kolejce ten-pin bowling?', options: ['300', '250', '200', '275'], correctOption: 'A' },
      { text: 'Który sport używa krążka?', options: ['Hokej na lodzie', 'Golf', 'Krykiet', 'Tenis'], correctOption: 'A' }
    ],
    'cat-geography': [
      { text: 'Który kontynent ma największą powierzchnię?', options: ['Azja', 'Afryka', 'Europa', 'Ameryka Północna'], correctOption: 'A' },
      { text: 'Jaka jest stolica Francji?', options: ['Paryż', 'Berlin', 'Madryt', 'Rzym'], correctOption: 'A' },
      { text: 'Która rzeka jest najdłuższa na świecie?', options: ['Nil', 'Amazonka', 'Jangcy', 'Missisipi'], correctOption: 'A' },
      { text: 'Który ocean leży na wschód od Stanów Zjednoczonych?', options: ['Ocean Atlantycki', 'Ocean Spokojny', 'Ocean Indyjski', 'Ocean Arktyczny'], correctOption: 'A' },
      { text: 'Który kraj ma największą liczbę ludności?', options: ['Indie', 'Chiny', 'Stany Zjednoczone', 'Brazylia'], correctOption: 'A' },
      { text: 'Który łańcuch górski zawiera Mount Everest?', options: ['Himalaje', 'Andy', 'Alpy', 'Rocky Mountains'], correctOption: 'A' },
      { text: 'Jaka jest stolica Japonii?', options: ['Tokio', 'Seul', 'Pekin', 'Bangkok'], correctOption: 'A' },
      { text: 'Która pustynia jest największą gorącą pustynią na Ziemi?', options: ['Sahara', 'Gobi', 'Kalahari', 'Mojave'], correctOption: 'A' },
      { text: 'Który stan USA ma przydomek Sunshine State?', options: ['Floryda', 'Kalifornia', 'Teksas', 'Arizona'], correctOption: 'A' },
      { text: 'Który kraj jest najmniejszy na świecie pod względem powierzchni?', options: ['Watykan', 'Monako', 'San Marino', 'Liechtenstein'], correctOption: 'A' }
    ],
    'cat-literature': [
      { text: 'Kto napisał Romeo i Julię?', options: ['William Shakespeare', 'Charles Dickens', 'Jane Austen', 'Mark Twain'], correctOption: 'A' },
      { text: 'Jaki jest tytuł antyutopijnego powieści Orwella?', options: ['1984', 'Nowy wspaniały świat', 'Fahrenheit 451', 'Dawca'], correctOption: 'A' },
      { text: 'Kto napisał Duma i uprzedzenie?', options: ['Jane Austen', 'Emily Brontë', 'Charlotte Brontë', 'Louisa May Alcott'], correctOption: 'A' },
      { text: 'Która powieść zaczyna się od słów „Zawołaj mnie Izmaelem”?', options: ['Moby Dick', 'Odyseja', 'Wyspa skarbów', 'Iliada'], correctOption: 'A' },
      { text: 'Która szkoła jest centralna w Harrym Potterze?', options: ['Hogwart', 'Beauxbatons', 'Durmstrang', 'Ilvermorny'], correctOption: 'A' },
      { text: 'Kto napisał Hobbit, czyli tam i z powrotem?', options: ['J.R.R. Tolkien', 'C.S. Lewis', 'J.K. Rowling', 'George R.R. Martin'], correctOption: 'A' },
      { text: 'Kto jest pomocnikiem Sherlocka Holmesa?', options: ['Dr Watson', 'Inspektor Lestrade', 'Pani Hudson', 'Profesor Moriarty'], correctOption: 'A' },
      { text: 'Która powieść przedstawia Atticusa Fincha?', options: ['Zabić drozdka', 'Wielki Gatsby', 'Na południu', '1984'], correctOption: 'A' },
      { text: 'Jaki typ wiersza ma 14 wersów?', options: ['Sonet', 'Haiku', 'Limerick', 'Epos'], correctOption: 'A' },
      { text: 'Która książka zaczyna się od słów „To były najlepsze czasy, to były najgorsze czasy”?', options: ['A Tale of Two Cities', 'Great Expectations', 'Oliver Twist', 'David Copperfield'], correctOption: 'A' }
    ],
    'cat-movies': [
      { text: 'Który film przedstawia postać Indiany Jonesa?', options: ['Poszukiwacze zaginionej Arki', 'Matrix', 'Park Jurajski', 'Powrót do przyszłości'], correctOption: 'A' },
      { text: 'Kto wyreżyserował Park Jurajski?', options: ['Steven Spielberg', 'James Cameron', 'Christopher Nolan', 'Ridley Scott'], correctOption: 'A' },
      { text: 'Który film zdobył Oscara za najlepszy film w 2020 roku?', options: ['Parasite', '1917', 'Joker', 'Pewnego razu... w Hollywood'], correctOption: 'A' },
      { text: 'Który bohater pojawia się w filmie The Dark Knight?', options: ['Batman', 'Spider-Man', 'Iron Man', 'Superman'], correctOption: 'A' },
      { text: 'Który film przedstawia statek Titanic?', options: ['Titanic', 'The Poseidon Adventure', 'Piraci z Karaibów', 'Życie Pi'], correctOption: 'A' },
      { text: 'Który film animowany przedstawia Olafa, śnieżynkę?', options: ['Kraina lodu', 'Toy Story', 'Shrek', 'Moana'], correctOption: 'A' },
      { text: 'Jaką tabletkę bierze Neo w filmie Matrix?', options: ['Czerwoną', 'Niebieską', 'Zieloną', 'Żółtą'], correctOption: 'A' },
      { text: 'Który film Toma Hanksa rozgrywa się na bezludnej wyspie?', options: ['Cast Away', 'The Beach', 'Życie Pi', 'Robin Crusoe'], correctOption: 'A' },
      { text: 'Która seria filmów obejmuje Darth Vadera?', options: ['Gwiezdne wojny', 'Star Trek', 'Avatar', 'Matrix'], correctOption: 'A' },
      { text: 'Który film zawiera utwór „My Heart Will Go On”?', options: ['Titanic', 'The Bodyguard', 'Ghost', 'Dirty Dancing'], correctOption: 'A' }
    ],
    'cat-music': [
      { text: 'Która grupa wydała Hey Jude?', options: ['The Beatles', 'The Rolling Stones', 'Queen', 'Pink Floyd'], correctOption: 'A' },
      { text: 'Który instrument ma 88 klawiszy?', options: ['Pianino', 'Gitara', 'Skrzypce', 'Flet'], correctOption: 'A' },
      { text: 'Kto jest znany jako Król Popu?', options: ['Michael Jackson', 'Elvis Presley', 'Prince', 'Madonna'], correctOption: 'A' },
      { text: 'Która piosenkarka wydała Rolling in the Deep?', options: ['Adele', 'Beyoncé', 'Taylor Swift', 'Rihanna'], correctOption: 'A' },
      { text: 'Jaki instrument zwykle gra perkusista?', options: ['Perkusja', 'Skrzypce', 'Flet', 'Wiolonczela'], correctOption: 'A' },
      { text: 'Która grupa jest znana z utworu Bohemian Rhapsody?', options: ['Queen', 'The Beatles', 'Nirvana', 'U2'], correctOption: 'A' },
      { text: 'Jakiego gatunku muzycznego Taylor Swift jest najbardziej znana?', options: ['Pop i country', 'Klasyczna', 'Jazz', 'Hip hop'], correctOption: 'A' },
      { text: 'Ile strun ma standardowa gitara?', options: ['6', '4', '7', '8'], correctOption: 'A' },
      { text: 'Które miasto jest blisko związane z muzyką country w Stanach Zjednoczonych?', options: ['Nashville', 'Austin', 'Nowy Orlean', 'Los Angeles'], correctOption: 'A' },
      { text: 'Który artysta wydał album Thriller?', options: ['Michael Jackson', 'Prince', 'Madonna', 'Whitney Houston'], correctOption: 'A' }
    ],
    'cat-math': [
      { text: 'Ile to 8 × 7?', options: ['56', '54', '58', '60'], correctOption: 'A' },
      { text: 'Jaka jest wartość π z dokładnością do dwóch miejsc po przecinku?', options: ['3,14', '3,15', '3,13', '3,12'], correctOption: 'A' },
      { text: 'Jaki jest pierwiastek kwadratowy z 81?', options: ['9', '7', '8', '10'], correctOption: 'A' },
      { text: 'Ile to 15% z 200?', options: ['30', '20', '25', '35'], correctOption: 'A' },
      { text: 'Jaka jest następna liczba pierwsza po 7?', options: ['11', '9', '10', '13'], correctOption: 'A' },
      { text: 'Ile to 2 + 2 × 2?', options: ['6', '8', '4', '10'], correctOption: 'A' },
      { text: 'Ile stopni ma kąt prosty?', options: ['90', '45', '180', '360'], correctOption: 'A' },
      { text: 'Jaki jest obwód kwadratu o boku 5?', options: ['20', '25', '10', '15'], correctOption: 'A' },
      { text: 'Jaka jest jedyna parzysta liczba pierwsza?', options: ['2', '0', '4', '6'], correctOption: 'A' },
      { text: 'Jaka jest suma kątów wewnętrznych trójkąta?', options: ['180°', '90°', '360°', '270°'], correctOption: 'A' }
    ],
    'cat-nature': [
      { text: 'Jaką proces wykorzystują rośliny do zamiany światła słonecznego w energię?', options: ['Fotosynteza', 'Oddychanie', 'Trawienie', 'Transpiracja'], correctOption: 'A' },
      { text: 'Które zwierzę jest największe na Ziemi?', options: ['Słoń afrykański', 'Żyrafa', 'Hipopotam', 'Nosorożec'], correctOption: 'A' },
      { text: 'Które zwierzę jest często nazywane królem dżungli?', options: ['Lew', 'Tygrys', 'Słoń', 'Goryl'], correctOption: 'A' },
      { text: 'Co pszczoły zbierają z kwiatów?', options: ['Nektar', 'Wodę', 'Piasek', 'Sok z liści'], correctOption: 'A' },
      { text: 'Jaka jest chemiczna nazwa soli stołowej?', options: ['Chlorek sodu', 'Węglan wapnia', 'Azotan potasu', 'Siarczan magnezu'], correctOption: 'A' },
      { text: 'Jaką warstwę Ziemi zamieszkujemy?', options: ['Skorupa', 'Płaszcz', 'Jądro', 'Jądro zewnętrzne'], correctOption: 'A' },
      { text: 'Który ocean jest największy na Ziemi?', options: ['Ocean Spokojny', 'Ocean Atlantycki', 'Ocean Indyjski', 'Ocean Arktyczny'], correctOption: 'A' },
      { text: 'Który typ drzewa zachowuje liście przez cały rok?', options: ['Wiecznie zielone', 'Liściaste', 'Brzoza', 'Klon'], correctOption: 'A' },
      { text: 'Jaki gaz ludzie wdychają i wykorzystują do przetrwania?', options: ['Tlen', 'Dwutlenek węgla', 'Azot', 'Hel'], correctOption: 'A' },
      { text: 'Który ptak jest znany z naśladowania ludzkiej mowy?', options: ['Papuga', 'Sowa', 'Orzeł', 'Gołąb'], correctOption: 'A' }
    ],
    'cat-technology': [
      { text: 'Która firma stworzyła iPhone?', options: ['Apple', 'Samsung', 'Nokia', 'Microsoft'], correctOption: 'A' },
      { text: 'Co oznacza CPU?', options: ['Central Processing Unit', 'Computer Power Unit', 'Central Program Utility', 'Core Processing Utility'], correctOption: 'A' },
      { text: 'Jaki język jest najczęściej używany do stron internetowych?', options: ['HTML', 'C++', 'Go', 'Rust'], correctOption: 'A' },
      { text: 'Które urządzenie przechowuje dane trwale?', options: ['Dysk twardy', 'RAM', 'CPU', 'Monitor'], correctOption: 'A' },
      { text: 'Co oznacza Wi‑Fi?', options: ['Wireless Fidelity', 'Wide Fiber', 'Web Internet Frequency', 'Wireless File Interface'], correctOption: 'A' },
      { text: 'Jaki język programowania został stworzony przez Sun Microsystems?', options: ['Java', 'Python', 'Ruby', 'Swift'], correctOption: 'A' },
      { text: 'Która technologia służy do wysyłania wiadomości przez internet?', options: ['Pakiety', 'Kawałki dźwięku', 'Tylko kable', 'Tylko Bluetooth'], correctOption: 'A' },
      { text: 'Która firma stworzyła Windows?', options: ['Microsoft', 'Apple', 'Google', 'IBM'], correctOption: 'A' },
      { text: 'Jaki jest główny cel systemu operacyjnego?', options: ['Zarządzanie zasobami sprzętowymi i programowymi', 'Łączenie drukarek', 'Przechowywanie energii elektrycznej', 'Wyświetlanie tylko tekstu'], correctOption: 'A' },
      { text: 'Który protokół sieciowy jest używany do sieci web?', options: ['HTTP', 'FTP', 'SMTP', 'SSH'], correctOption: 'A' }
    ],
    'cat-art': [
      { text: 'Który artysta namalował Mona Lisę?', options: ['Leonardo da Vinci', 'Vincent van Gogh', 'Pablo Picasso', 'Claude Monet'], correctOption: 'A' },
      { text: 'Jaki typ sztuki powstaje z pigmentów na płótnie?', options: ['Malarstwo', 'Rzeźba', 'Fotografia', 'Architektura'], correctOption: 'A' },
      { text: 'Który kolor powstaje z połączenia niebieskiego i żółtego?', options: ['Zielony', 'Pomarańczowy', 'Fioletowy', 'Czerwony'], correctOption: 'A' },
      { text: 'Jaka forma sztuki wykorzystuje glinę i narzędzia do formowania?', options: ['Ceramika', 'Malarstwo', 'Grafika', 'Rysunek'], correctOption: 'A' },
      { text: 'Który budynek słynie z wejścia w postaci szklanej piramidy?', options: ['Luwr', 'Wieża Eiffla', 'Tadź Mahal', 'Big Ben'], correctOption: 'A' },
      { text: 'Który artysta jest znany z „Gwiaździstej nocy”?', options: ['Vincent van Gogh', 'Leonardo da Vinci', 'Michał Anioł', 'Salvador Dalí'], correctOption: 'A' },
      { text: 'Czym jest rzeźba?', options: ['Trójwymiarowy obiekt artystyczny', 'Płaski rysunek', 'Fotografia', 'Poemat'], correctOption: 'A' },
      { text: 'Jaki materiał używa się do rysunku węglem?', options: ['Węgiel', 'Olej', 'Akwarela', 'Marmur'], correctOption: 'A' },
      { text: 'Jaka sztuka powstaje z układania kolorowych kawałków?', options: ['Mozaika', 'Drewno', 'Kaligrafia', 'Ceramika'], correctOption: 'A' },
      { text: 'Który styl artystyczny używa mocnych, uproszczonych form i wyrazistych kolorów?', options: ['Sztuka nowoczesna', 'Renesans', 'Barok', 'Klasyka'], correctOption: 'A' }
    ],
    'cat-food': [
      { text: 'Które owoce są zwykle używane do przygotowania guacamole?', options: ['Awokado', 'Mango', 'Jagoda', 'Pomarańcza'], correctOption: 'A' },
      { text: 'Które ziarno jest używane do pieczenia chleba?', options: ['Pszenica', 'Ryż', 'Kukurydza', 'Jęczmień'], correctOption: 'A' },
      { text: 'Która kuchnia słynie z sushi?', options: ['Japońska', 'Włoska', 'Meksykańska', 'Francuska'], correctOption: 'A' },
      { text: 'Który witaminę znajdujemy w owocach cytrusowych?', options: ['Witamina C', 'Witamina D', 'Witamina A', 'Witamina K'], correctOption: 'A' },
      { text: 'Jaki jest główny składnik hummusu?', options: ['Ciecierzyca', 'Pomidory', 'Fasola', 'Ryż'], correctOption: 'A' },
      { text: 'Która metoda gotowania wykorzystuje suchą temperaturę w piekarniku?', options: ['Pieczenie', 'Gotowanie', 'Parzenie', 'Fermentacja'], correctOption: 'A' },
      { text: 'Które owoce są znane z czerwonego i słodkiego smaku?', options: ['Truskawka', 'Cytryna', 'Gruszka', 'Banan'], correctOption: 'A' },
      { text: 'Jaki jest główny składnik sera?', options: ['Mleko', 'Mąka', 'Jajka', 'Ryż'], correctOption: 'A' },
      { text: 'Który warzywo jest często używane do zupy pomidorowej?', options: ['Pomidory', 'Ziemniaki', 'Marchew', 'Kapusta'], correctOption: 'A' },
      { text: 'Który przyprawa jest zwykle używana do przyprawiania pizzy?', options: ['Oregano', 'Cynamon', 'Gałka muszkatołowa', 'Kardamon'], correctOption: 'A' }
    ],
    'cat-animals': [
      { text: 'Które ssaki są znane z umiejętności latania?', options: ['Nietoperz', 'Mysz', 'Wiewiórka', 'Królik'], correctOption: 'A' },
      { text: 'Które zwierzę jest największe na Ziemi?', options: ['Niebieski wieloryb', 'Słoń', 'Żyrafa', 'Hipopotam'], correctOption: 'A' },
      { text: 'Które zwierzę jest znane z noszenia swojego domu na plecach?', options: ['Ślimak', 'Żółw', 'Krab', 'Homar'], correctOption: 'A' },
      { text: 'Które zwierzę jest najszybszym ssakiem lądowym?', options: ['Gepard', 'Koń', 'Lew', 'Gazela'], correctOption: 'A' },
      { text: 'Które zwierzę jest znane z hibernacji zimą?', options: ['Niedźwiedź', 'Królik', 'Lis', 'Wilk'], correctOption: 'A' },
      { text: 'Które zwierzę morskie używa muszli do ochrony?', options: ['Żółw', 'Delfin', 'Rekin', 'Foka'], correctOption: 'A' },
      { text: 'Które zwierzę jest znane z czarno-białych pasków?', options: ['Zebra', 'Tygrys', 'Lampart', 'Panda'], correctOption: 'A' },
      { text: 'Które zwierzę jest symbolem mądrości w wielu kulturach?', options: ['Sowa', 'Orzeł', 'Sokół', 'Wróbel'], correctOption: 'A' },
      { text: 'Które zwierzę jest często nazywane królem dżungli?', options: ['Lew', 'Słoń', 'Goryl', 'Tygrys'], correctOption: 'A' },
      { text: 'Który duży kot jest słynny z pomarańczowej sierści i czarnych pasków?', options: ['Tygrys', 'Lampart', 'Gepard', 'Ryś'], correctOption: 'A' }
    ],
    'cat-space': [
      { text: 'Która planeta jest najbliżej Słońca?', options: ['Merkury', 'Wenus', 'Mars', 'Ziemia'], correctOption: 'A' },
      { text: 'Jak nazywa się naturalny satelita Ziemi?', options: ['Księżyc', 'Tytan', 'Europa', 'Fobos'], correctOption: 'A' },
      { text: 'Która planeta jest znana z pierścieni?', options: ['Saturn', 'Jowisz', 'Neptun', 'Mars'], correctOption: 'A' },
      { text: 'Jaka siła utrzymuje planety na orbitach wokół Słońca?', options: ['Grawitacja', 'Magnetyzm', 'Tarcie', 'Promieniowanie'], correctOption: 'A' },
      { text: 'Która galaktyka zawiera Układ Słoneczny?', options: ['Droga Mleczna', 'Andromeda', 'Whirlpool', 'Sombrero'], correctOption: 'A' },
      { text: 'Z czego głównie składa się kometa?', options: ['Lód i pył', 'Metal i lawa', 'Szkło i gaz', 'Kamień i woda'], correctOption: 'A' },
      { text: 'Która planeta jest znana jako Czerwona Planeta?', options: ['Mars', 'Wenus', 'Merkury', 'Jowisz'], correctOption: 'A' },
      { text: 'Który statek kosmiczny wylądował ludzi na Księżycu?', options: ['Apollo', 'Voyager', 'Gemini', 'Sputnik'], correctOption: 'A' },
      { text: 'Jak nazywa się gwiazda w centrum naszego Układu Słonecznego?', options: ['Słońce', 'Proxima Centauri', 'Sirius', 'Betelgeuse'], correctOption: 'A' },
      { text: 'Która planeta jest największa w Układzie Słonecznym?', options: ['Jowisz', 'Saturn', 'Ziemia', 'Neptun'], correctOption: 'A' }
    ],
    'cat-health': [
      { text: 'Która aktywność poprawia zdrowie serca?', options: ['Regularny trening', 'Spanie przez cały dzień', 'Pomijanie posiłków', 'Palenie'], correctOption: 'A' },
      { text: 'Jaki jest główny cel snu?', options: ['Odpoczynek i regeneracja', 'Trawienie jedzenia', 'Produkcja krwi', 'Schładzanie ciała'], correctOption: 'A' },
      { text: 'Która grupa pokarmów jest niezbędna dla mocnych kości?', options: ['Nabiał i produkty bogate w wapń', 'Słodycze', 'Soda', 'Fast food'], correctOption: 'A' },
      { text: 'Który nawyk jest najbardziej szkodliwy dla płuc?', options: ['Palenie', 'Picie wody', 'Spacerowanie', 'Rozciąganie'], correctOption: 'A' },
      { text: 'Jaki jest zdrowy sposób radzenia sobie z stresem?', options: ['Ćwiczenia i odpoczynek', 'Pomijanie snu', 'Palenie', 'Jedzenie fast foodów'], correctOption: 'A' },
      { text: 'Który narząd jest najbezpośredniej odpowiedzialny za filtrowanie krwi?', options: ['Nerki', 'Żołądek', 'Płuca', 'Wątroba'], correctOption: 'A' },
      { text: 'Co powinieneś zrobić przed ćwiczeniami?', options: ['Rozgrzać się', 'Pić sodę', 'Pomijać nawodnienie', 'Położyć się'], correctOption: 'A' },
      { text: 'Który składnik odżywczy pomaga budować mięśnie?', options: ['Białko', 'Cukier', 'Sól', 'Kofeina'], correctOption: 'A' },
      { text: 'Jak możesz utrzymywać odporność na wysokim poziomie?', options: ['Dobrze jeść i spać wystarczająco', 'Jeść tylko słodycze', 'Unikać ćwiczeń', 'Siedzieć do późna'], correctOption: 'A' },
      { text: 'Jaki jest najlepszy sposób na nawodnienie?', options: ['Regularnie pić wodę', 'Pić tylko sodę', 'Pomijać płyny', 'Pić tylko kawę'], correctOption: 'A' }
    ],
    'cat-language': [
      { text: 'Jaki jest główny cel gramatyki?', options: ['Strukturyzowanie zdań w jasny sposób', 'Dodawanie kolorów', 'Liczenie liczb', 'Losowa zmiana pisowni'], correctOption: 'A' },
      { text: 'Które słowo jest rzeczownikiem?', options: ['Stół', 'Szybko', 'Pięknie', 'Biegać'], correctOption: 'A' },
      { text: 'Jaki jest antonim słowa „zacząć”?', options: ['Zakończyć', 'Otworzyć', 'Rozpocząć', 'Skonczyć'], correctOption: 'A' },
      { text: 'Które zdanie jest napisane poprawnie?', options: ['Ona chodzi do szkoły codziennie.', 'Ona chodzi do szkoły każdego dnia.', 'Ona chodzą do szkoły codziennie.', 'Ona chodził do szkoły codziennie.'], correctOption: 'A' },
      { text: 'Czym jest synonim słowa „szczęśliwy”?', options: ['Radosny', 'Smutny', 'Zły', 'Cichy'], correctOption: 'A' },
      { text: 'Jaki jest antonim słowa „gorący”?', options: ['Zimny', 'Ciepły', 'Jasny', 'Szybki'], correctOption: 'A' },
      { text: 'Która część mowy opisuje czasownik, przymiotnik albo inny przysłówek?', options: ['Przysłówek', 'Rzeczownik', 'Zaimek', 'Przyimek'], correctOption: 'A' },
      { text: 'Czym jest akapit?', options: ['Grupa powiązanych zdań', 'Pojedyncze słowo', 'Znacznik interpunkcyjny', 'Liczba'], correctOption: 'A' },
      { text: 'Który znak interpunkcyjny kończy zdanie?', options: ['Kropka', 'Przecinek', 'Znak zapytania', 'Dwukropek'], correctOption: 'A' },
      { text: 'Do czego służy słownik?', options: ['Znajdowanie znaczeń słów', 'Mierzenie wzrostu', 'Gotowanie jedzenia', 'Budowanie krzesła'], correctOption: 'A' }
    ],
    'cat-psychology': [
      { text: 'Czym jest pamięć?', options: ['Zdolność do zapamiętywania i przywoływania informacji', 'Poczucie głodu', 'Rodzaj rośliny', 'Rodzaj światła'], correctOption: 'A' },
      { text: 'Który emocja jest często związana z lękiem?', options: ['Niepokój', 'Radość', 'Spokój', 'Ekscytacja'], correctOption: 'A' },
      { text: 'Czym jest empatia?', options: ['Rozumienie i współdzielenie uczuć innych', 'Ignorowanie innych', 'Poczucie złości', 'Bycie cichym'], correctOption: 'A' },
      { text: 'Jakim jest wspólny cel uważności?', options: ['Bycie obecnym i skupionym', 'Ignorowanie myśli', 'Spanie przez cały czas', 'Unikanie wszystkich emocji'], correctOption: 'A' },
      { text: 'Które zachowanie jest zwykle związane z dobrym zdrowiem psychicznym?', options: ['Zdrowe strategie radzenia sobie', 'Unikanie wszystkich obowiązków', 'Stały niepokój', 'Brak emocji'], correctOption: 'A' },
      { text: 'Co oznacza percepcja?', options: ['Jak interpretujemy informacje zmysłowe', 'Rodzaj pamięci', 'Norma społeczna', 'Nastrój'], correctOption: 'A' },
      { text: 'Który czynnik często wpływa na podejmowanie decyzji?', options: ['Emocje i doświadczenie', 'Tylko przypadkowy los', 'Tylko pogoda', 'Tylko daty kalendarzowe'], correctOption: 'A' },
      { text: 'Czym jest motywacja?', options: ['Pęd, który prowadzi do działania', 'Rodzaj jedzenia', 'Kształt pudełka', 'Limit prędkości'], correctOption: 'A' },
      { text: 'Co stres zwykle robi z koncentracją?', options: ['Może ją obniżać', 'Zawsze ją poprawia', 'Czyni cię niewidzialnym', 'Usuwa wszystkie myśli'], correctOption: 'A' },
      { text: 'Czym jest nawyk?', options: ['Powtarzalne zachowanie', 'Losowy pomysł', 'Rodzaj owocu', 'Obiekt fizyczny'], correctOption: 'A' }
    ],
    'cat-religion': [
      { text: 'Która książka jest święta w chrześcijaństwie?', options: ['Biblia', 'Tora', 'Koran', 'Wedy'], correctOption: 'A' },
      { text: 'Jakie jest święte miasto islamu?', options: ['Mekka', 'Jerozolima', 'Rzym', 'Kair'], correctOption: 'A' },
      { text: 'Która religia przestrzega Pięciu Filarów?', options: ['Islam', 'Chrześcijaństwo', 'Hinduizm', 'Judaizm'], correctOption: 'A' },
      { text: 'Jaki jest święty tekst judaizmu?', options: ['Tora', 'Biblia', 'Koran', 'Tripitaka'], correctOption: 'A' },
      { text: 'Która religia jest związana z karmą i dharmą?', options: ['Hinduizm', 'Islam', 'Buddyzm', 'Chrześcijaństwo'], correctOption: 'A' },
      { text: 'Czym jest świątynia?', options: ['Miejsce kultu', 'Szkoła', 'Farma', 'Targowisko'], correctOption: 'A' },
      { text: 'Która wiara skupia się na naukach Buddy?', options: ['Buddyzm', 'Chrześcijaństwo', 'Judaizm', 'Shinto'], correctOption: 'A' },
      { text: 'Które miasto jest święte dla chrześcijan, Żydów i muzułmanów?', options: ['Jerozolima', 'Kair', 'Mekka', 'Rzym'], correctOption: 'A' },
      { text: 'Która religia opiera się na życiu i naukach Jezusa?', options: ['Chrześcijaństwo', 'Islam', 'Hinduizm', 'Sikhizm'], correctOption: 'A' },
      { text: 'Czym jest sanktuarium?', options: ['Święte lub uświęcone miejsce', 'Biblioteka', 'Stadion sportowy', 'Szpital'], correctOption: 'A' }
    ],
    'cat-business': [
      { text: 'Jaki jest główny cel biznesu?', options: ['Tworzenie wartości i zarabianie zysków', 'Tylko wydawanie pieniędzy', 'Unikanie klientów', 'Ignorowanie innowacji'], correctOption: 'A' },
      { text: 'Czym jest budżet?', options: ['Plan dochodów i wydatków', 'Rodzaj maszyny', 'Przepis prawny', 'Projekt produktu'], correctOption: 'A' },
      { text: 'Który termin oznacza kwotę pieniędzy, którą firma zarabia?', options: ['Przychód', 'Dług', 'Zapasy', 'Koszt'], correctOption: 'A' },
      { text: 'Czym jest marketing?', options: ['Promocja i sprzedaż produktów', 'Tylko liczenie zapasów', 'Budowanie dróg', 'Zmiana prawa'], correctOption: 'A' },
      { text: 'Czym jest branding?', options: ['Tworzenie rozpoznawalnej tożsamości', 'Tworzenie umowy prawnej', 'Pisanie kodu', 'Wysyłka produktów'], correctOption: 'A' },
      { text: 'Który dokument opisuje pomysł i plan biznesowy?', options: ['Plan biznesowy', 'Paragon', 'Faktura', 'Wyciąg bankowy'], correctOption: 'A' },
      { text: 'Czym jest interesariusz?', options: ['Osoba dotknięta decyzją biznesową', 'Część maszyny', 'Kampania marketingowa', 'Funkcja produktu'], correctOption: 'A' },
      { text: 'Czego mierzy ROI?', options: ['Zwrot z inwestycji', 'Uruchomienie operacyjnych zapasów', 'Przychód od odsetek', 'Tempo dochodu'], correctOption: 'A' },
      { text: 'O co chodzi w obsłudze klienta?', options: ['Pomaganie klientom przed i po zakupie', 'Ignorowanie problemów klientów', 'Tylko obniżanie cen', 'Produkcja surowców'], correctOption: 'A' },
      { text: 'Który dokument finansowy pokazuje zysk i stratę?', options: ['Rachunek zysków i strat', 'Bilans', 'Kasa', 'Forma płacowa'], correctOption: 'A' }
    ],
    'cat-gaming': [
      { text: 'Jaki jest cel większości gier wyścigowych?', options: ['Dojazd do mety pierwszy', 'Rozwiązywanie zagadek', 'Zbieranie roślin', 'Pisanie kodu'], correctOption: 'A' },
      { text: 'Co oznacza FPS w grach?', options: ['First-person shooter', 'Fast play score', 'Full power system', 'Frame priority setting'], correctOption: 'A' },
      { text: 'Które działanie zwykle wygrywa w grze strategicznej?', options: ['Planowanie z wyprzedzeniem', 'Losowe klikanie', 'Ignorowanie zasobów', 'Pomijanie tur'], correctOption: 'A' },
      { text: 'Czym jest loot w RPG?', options: ['Nagrody, takie jak ekwipunek lub przedmioty', 'Rodzaj mapy', 'Dźwięk broni', 'Poziom trudności'], correctOption: 'A' },
      { text: 'Do czego służy kontroler gry?', options: ['Sterowanie wejściem w grach', 'Słuchanie muzyki', 'Czytanie książek', 'Robienie kawy'], correctOption: 'A' },
      { text: 'Co oznacza multiplayer?', options: ['W grze uczestniczy więcej niż jeden gracz', 'Tryb solo', 'Ekran o wysokiej rozdzielczości', 'Nowy silnik gry'], correctOption: 'A' },
      { text: 'Który gatunek skupia się na rozwiązywaniu logicznych zagadek?', options: ['Gry logiczne', 'Gry wyścigowe', 'Gry sportowe', 'Platformery'], correctOption: 'A' },
      { text: 'Czym jest punkt zapisu?', options: ['Miejsce, w którym zapisywany jest postęp', 'Kod oszustwa', 'Tytuł gry', 'Ikona mapy'], correctOption: 'A' },
      { text: 'Która funkcja zwykle poprawia wydajność gry?', options: ['Lepsze ustawienia grafiki', 'Wolniejsze ładowanie', 'Więcej przypadkowych awarii', 'Brak dźwięku'], correctOption: 'A' },
      { text: 'Czym jest mechanika gry?', options: ['Zasada lub system definiujący rozgrywkę', 'Nazwa gracza', 'Dysk twardy', 'Filtr zdjęć'], correctOption: 'A' }
    ],
    'cat-cars': [
      { text: 'Jaki jest główny cel silnika samochodu?', options: ['Generowanie mocy do poruszania pojazdu', 'Tylko magazynowanie paliwa', 'Odtwarzanie muzyki', 'Kontrola świateł'], correctOption: 'A' },
      { text: 'Do czego pomaga ABS w samochodzie?', options: ['Zapobiega blokowaniu kół podczas hamowania', 'Zwiększa ciśnienie w oponach', 'Obraca kierownicę', 'Chłodzi kabinę'], correctOption: 'A' },
      { text: 'Które paliwo jest najczęściej używane w większości samochodów benzynowych?', options: ['Benzyna', 'Tylko olej napędowy', 'Węgiel', 'Gaz ziemny'], correctOption: 'A' },
      { text: 'Czym jest skrzynia biegów?', options: ['System przekazujący moc na koła', 'Radio samochodowe', 'Poduszka siedzenia', 'Kierownica'], correctOption: 'A' },
      { text: 'Co mierzy obrotomierz?', options: ['Obroty silnika na minutę', 'Poziom paliwa', 'Temperaturę na zewnątrz', 'Pozycję siedzenia'], correctOption: 'A' },
      { text: 'Która część odpowiada za kierowanie samochodem?', options: ['Kierownica i mechanizm', 'Deska rozdzielcza', 'Reflektory', 'Bagażnik dachowy'], correctOption: 'A' },
      { text: 'Jaki jest cel akumulatora samochodowego?', options: ['Uruchamianie silnika i zasilanie elektroniki', 'Bezpośrednio porusza koła', 'Mierzy ciśnienie w oponach', 'Chłodzi hamulce'], correctOption: 'A' },
      { text: 'Który typ pojazdu jest zaprojektowany głównie do jazdy w terenie?', options: ['SUV', 'Sedan', 'Hatchback', 'Coupe'], correctOption: 'A' },
      { text: 'Co mierzy MPG?', options: ['Mile na galon paliwa', 'Mile na godzinę', 'Miligramy na gram', 'Metry na galon'], correctOption: 'A' },
      { text: 'Która część samochodu odpowiada za hamowanie?', options: ['Klocki hamulcowe', 'Reflektory', 'Dach', 'Spojler'], correctOption: 'A' }
    ],
    'cat-astronomy': [
      { text: 'Czym jest galaktyka?', options: ['Ogromne skupisko gwiazd, gazu i pyłu', 'Pojedyncza gwiazda', 'Planeta', 'Księżyc'], correctOption: 'A' },
      { text: 'Która planeta jest znana z wyraźnych pierścieni?', options: ['Saturn', 'Mars', 'Merkury', 'Wenus'], correctOption: 'A' },
      { text: 'Z czego głównie składa się kometa?', options: ['Lód, pył i kamień', 'Gorący metal', 'Tylko para wodna', 'Miedź i stal'], correctOption: 'A' },
      { text: 'Do czego astronomowie używają do badania odległych gwiazd?', options: ['Teleskopów', 'Mikroskopów', 'Magnetów', 'Kompasów'], correctOption: 'A' },
      { text: 'Która planeta jest najbliżej Słońca?', options: ['Merkury', 'Wenus', 'Ziemia', 'Mars'], correctOption: 'A' },
      { text: 'Czym jest zaćmienie Księżyca?', options: ['Ziemia przechodzi między Słońcem a Księżycem', 'Księżyc przechodzi między Ziemią a Słońcem', 'Planeta zderza się z Księżycem', 'Słońce znika'], correctOption: 'A' },
      { text: 'Czym jest czarna dziura?', options: ['Obszar przestrzeni o bardzo silnym grawitacji', 'Gigantyczna gwiazda', 'Planeta bez atmosfery', 'Pusta mgławica'], correctOption: 'A' },
      { text: 'Która planeta ma Wielką Czerwoną Plamę?', options: ['Jowisz', 'Mars', 'Neptun', 'Wenus'], correctOption: 'A' },
      { text: 'Jak nazywa się gwiazda Ziemi?', options: ['Słońce', 'Sirius', 'Polaris', 'Wega'], correctOption: 'A' },
      { text: 'Co opisuje orbita?', options: ['Ścieżka, po której obiekt porusza się wokół innego obiektu', 'Błysk światła', 'Pogoda', 'Kształt góry'], correctOption: 'A' }
    ],
    'cat-philosophy': [
      { text: 'Czym głównie zajmuje się etyka?', options: ['Dobro i zło w zachowaniu', 'Reakcje chemiczne', 'Pogoda', 'Daty historyczne'], correctOption: 'A' },
      { text: 'Kto jest znany z twierdzenia „Myślę, więc jestem”?', options: ['Descartes', 'Platon', 'Arystoteles', 'Sokrates'], correctOption: 'A' },
      { text: 'Czym jest logika?', options: ['Rozumowanie zgodnie z prawami logicznymi', 'Eksperyment fizyczny', 'Kształt budynku', 'Forma medycyny'], correctOption: 'A' },
      { text: 'Czego dotyczy metafizyka?', options: ['Natury rzeczywistości', 'Tylko liczb', 'Systemów pogodowych', 'Anatomii człowieka'], correctOption: 'A' },
      { text: 'Czym jest prawda?', options: ['Koncepcja tego, co jest poprawne i realne', 'Rodzaj kamienia', 'Typ muzyki', 'Forma światła'], correctOption: 'A' },
      { text: 'Czym jest argument w filozofii?', options: ['Zestaw powodów wspierających wniosek', 'Głośna kłótnia', 'Typ budynku', 'Wydarzenie polityczne'], correctOption: 'A' },
      { text: 'Który filozof nauczał w Atenach i zadawał wiele pytań?', options: ['Sokrates', 'Descartes', 'Hume', 'Kant'], correctOption: 'A' },
      { text: 'Czym jest wolna wola?', options: ['Możliwość wyboru własnych działań', 'Prawo fizyki', 'Instynkt zwierzęcia', 'Stan nastroju'], correctOption: 'A' },
      { text: 'Czym jest epistemologia?', options: ['Badanie wiedzy', 'Badanie liczb', 'Badanie komórek', 'Badanie języka'], correctOption: 'A' },
      { text: 'Na czym skupia się egzystencjalizm?', options: ['Ludzkim istnieniu i wolnym wyborze', 'Dowodach matematycznych', 'Biologii roślin', 'Projektowaniu mechaniki'], correctOption: 'A' }
    ],
    'cat-design': [
      { text: 'Na czym głównie koncentruje się projektowanie doświadczeń użytkownika?', options: ['Jak osoba czuje się podczas korzystania z produktu', 'Tylko kolorystyka', 'Pisanie długich dokumentów', 'Druk na papierze'], correctOption: 'A' },
      { text: 'Czym jest wireframe?', options: ['Prosty plan układu', 'Kompilator kodu', 'Zapytanie do bazy danych', 'Raport marketingowy'], correctOption: 'A' },
      { text: 'Dlaczego kontrast jest ważny w projektowaniu?', options: ['Pomaga elementom wyróżniać się i poprawia czytelność', 'Zmniejsza kreatywność', 'Zawsze spowalnia użytkowników', 'Usuwa kolor'], correctOption: 'A' },
      { text: 'Co oznacza UI?', options: ['User Interface', 'Universal Input', 'User Information', 'Unified Internet'], correctOption: 'A' },
      { text: 'Czym jest hierarchia w projektowaniu?', options: ['Ułożenie elementów według ważności', 'Zestaw kolorów', 'Metoda przechowywania', 'Tylko siatka układu'], correctOption: 'A' },
      { text: 'Który zasada pomaga użytkownikom zrozumieć, na co patrzeć najpierw?', options: ['Wizualna hierarchia', 'Losowe rozmieszczenie', 'Przepełnienie', 'Brak odstępów'], correctOption: 'A' },
      { text: 'Czym jest prototyp?', options: ['Wczesny model projektu', 'Finalny produkt', 'Hasło marketingowe', 'Biblioteka kodu'], correctOption: 'A' },
      { text: 'Dlaczego białe miejsce jest przydatne w projektowaniu?', options: ['Poprawia czytelność i skupienie', 'Czyni tekst niewidocznym', 'Usuwa znaczenie', 'Dodaje hałas'], correctOption: 'A' },
      { text: 'Czym jest dostępność w projektowaniu?', options: ['Projektowanie dla osób o różnych potrzebach i możliwościach', 'Tylko skracanie treści', 'Używanie tylko jednego koloru', 'Ignorowanie użytkowników'], correctOption: 'A' },
      { text: 'Które narzędzie projektowe jest powszechnie używane do interfejsów?', options: ['Figma', 'Excel', 'Photoshop bez narzędzi projektowych', 'Linia poleceń'], correctOption: 'A' }
    ]
  };

  function buildQuestionPoolFromBank(questionBank) {
    const questions = {};
    let questionNumber = 1;

    categories.forEach(({ id }) => {
      const pool = questionBank[id] || [];
      for (let index = 0; index < 42; index++) {
        const question = pool[index % pool.length] || {
          text: `Które odpowiedzi najlepiej pasują do ${id}?`,
          options: ['Nauka', 'Losowość', 'Zamieszanie', 'Opóźnienie'],
          correctOption: 'A'
        };

        questions[`q-${questionNumber}`] = {
          text: question.text,
          options: {
            A: question.options[0],
            B: question.options[1],
            C: question.options[2],
            D: question.options[3]
          },
          correctOption: question.correctOption
        };

        questionNumber += 1;
      }
    });

    return questions;
  }

  const translatedCategories = {};
  categories.forEach(({ id, name }) => {
    translatedCategories[id] = name;
  });

  window.quizLocaleData.pl = {
    categories: translatedCategories,
    questions: buildQuestionPoolFromBank(categoryQuestionBanks)
  };
})();
