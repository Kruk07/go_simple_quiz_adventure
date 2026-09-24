window.quizLocaleData = window.quizLocaleData || {};

(function () {
  const categories = [
    { id: 'cat-programming', name: 'Programming' },
    { id: 'cat-history', name: 'History' },
    { id: 'cat-science', name: 'Science' },
    { id: 'cat-sports', name: 'Sports' },
    { id: 'cat-geography', name: 'Geography' },
    { id: 'cat-literature', name: 'Literature' },
    { id: 'cat-movies', name: 'Movies' },
    { id: 'cat-music', name: 'Music' },
    { id: 'cat-math', name: 'Math' },
    { id: 'cat-nature', name: 'Nature' },
    { id: 'cat-technology', name: 'Technology' },
    { id: 'cat-art', name: 'Art' },
    { id: 'cat-food', name: 'Food' },
    { id: 'cat-animals', name: 'Animals' },
    { id: 'cat-space', name: 'Space' },
    { id: 'cat-health', name: 'Health' },
    { id: 'cat-language', name: 'Language' },
    { id: 'cat-psychology', name: 'Psychology' },
    { id: 'cat-religion', name: 'Religion' },
    { id: 'cat-business', name: 'Business' },
    { id: 'cat-gaming', name: 'Gaming' },
    { id: 'cat-cars', name: 'Cars' },
    { id: 'cat-astronomy', name: 'Astronomy' },
    { id: 'cat-philosophy', name: 'Philosophy' },
    { id: 'cat-design', name: 'Design' }
  ];

  const categoryQuestionBanks = {
    'cat-programming': [
      { text: 'Which keyword declares a function in Go?', options: ['func', 'let', 'class', 'def'], correctOption: 'A' },
      { text: 'What does Go use to run code concurrently?', options: ['goroutines', 'threads', 'classes', 'modules'], correctOption: 'A' },
      { text: 'Which data structure keeps key-value pairs in Go?', options: ['map', 'stack', 'queue', 'array'], correctOption: 'A' },
      { text: 'Which command runs a Go program from the module root?', options: ['go run .', 'npm start', 'cargo run', 'python app.py'], correctOption: 'A' },
      { text: 'Which keyword is used to import a package in Go?', options: ['import', 'include', 'require', 'using'], correctOption: 'B' },
      { text: 'Which type is commonly used for text in Go?', options: ['string', 'int', 'bool', 'char'], correctOption: 'A' },
      { text: 'Which Go tool compiles packages and dependencies?', options: ['go build', 'gcc', 'javac', 'tsc'], correctOption: 'A' },
      { text: 'What is a slice in Go?', options: ['a dynamic sequence of values', 'a single integer', 'a compiled file', 'a database record'], correctOption: 'A' },
      { text: 'Which type in Go holds a single true/false value?', options: ['bool', 'string', 'int', 'byte'], correctOption: 'A' },
      { text: 'Which Go keyword creates a new variable with an inferred type?', options: ['var', 'class', 'def', 'new'], correctOption: 'A' }
    ],
    'cat-history': [
      { text: 'Who was the first President of the United States?', options: ['George Washington', 'Thomas Jefferson', 'John Adams', 'Abraham Lincoln'], correctOption: 'A' },
      { text: 'In which year did the Berlin Wall fall?', options: ['1989', '1987', '1991', '1979'], correctOption: 'A' },
      { text: 'Which empire was ruled by Genghis Khan?', options: ['Mongol Empire', 'Roman Empire', 'Ottoman Empire', 'Byzantine Empire'], correctOption: 'A' },
      { text: 'The Magna Carta was signed in which year?', options: ['1215', '1315', '1415', '1115'], correctOption: 'A' },
      { text: 'Which ancient civilization built the pyramids of Giza?', options: ['Egyptians', 'Romans', 'Greeks', 'Vikings'], correctOption: 'A' },
      { text: 'Which conflict was fought between the North and South in the United States?', options: ['American Civil War', 'World War I', 'Revolutionary War', 'War of 1812'], correctOption: 'A' },
      { text: 'Who discovered the Americas in 1492?', options: ['Christopher Columbus', 'Ferdinand Magellan', 'Marco Polo', 'Amerigo Vespucci'], correctOption: 'A' },
      { text: 'Which city was destroyed by the eruption of Mount Vesuvius?', options: ['Pompeii', 'Rome', 'Athens', 'Naples'], correctOption: 'A' },
      { text: 'Which document established the U.S. Constitution\'s framework of government?', options: ['The Constitution', 'The Declaration', 'The Articles of Confederation', 'The Bill of Rights'], correctOption: 'A' },
      { text: 'Which king was forced to sign the Magna Carta?', options: ['King John', 'Henry VIII', 'Edward I', 'Richard III'], correctOption: 'A' }
    ],
    'cat-science': [
      { text: 'What planet is known as the Red Planet?', options: ['Mars', 'Venus', 'Mercury', 'Jupiter'], correctOption: 'A' },
      { text: 'What is the chemical formula of water?', options: ['H2O', 'CO2', 'NaCl', 'O2'], correctOption: 'A' },
      { text: 'What force keeps planets in orbit around the Sun?', options: ['Gravity', 'Magnetism', 'Electricity', 'Friction'], correctOption: 'A' },
      { text: 'Which part of the cell contains genetic material?', options: ['Nucleus', 'Cell membrane', 'Ribosome', 'Cytoplasm'], correctOption: 'A' },
      { text: 'Which gas do plants absorb from the atmosphere?', options: ['Carbon dioxide', 'Oxygen', 'Nitrogen', 'Hydrogen'], correctOption: 'A' },
      { text: 'What is the center of an atom called?', options: ['Nucleus', 'Electron', 'Proton', 'Neutron'], correctOption: 'A' },
      { text: 'What is the boiling point of water at sea level?', options: ['100°C', '90°C', '80°C', '110°C'], correctOption: 'A' },
      { text: 'Which organ pumps blood through the body?', options: ['Heart', 'Liver', 'Lung', 'Kidney'], correctOption: 'A' },
      { text: 'What gas do humans exhale in large amounts?', options: ['Carbon dioxide', 'Oxygen', 'Helium', 'Argon'], correctOption: 'A' },
      { text: 'Which bird is famous for mimicking sounds?', options: ['Parrot', 'Eagle', 'Penguin', 'Owl'], correctOption: 'A' }
    ],
    'cat-sports': [
      { text: 'How many players are on the field for one soccer team?', options: ['11', '9', '10', '12'], correctOption: 'A' },
      { text: 'In baseball, what is a home run?', options: ['A run scored by hitting the ball out of play', 'A strikeout', 'A foul ball', 'A double play'], correctOption: 'A' },
      { text: 'How many minutes are in a standard soccer match?', options: ['90', '60', '75', '120'], correctOption: 'A' },
      { text: 'Which sport uses a shuttlecock?', options: ['Badminton', 'Tennis', 'Squash', 'Cricket'], correctOption: 'A' },
      { text: 'How many rings appear on the Olympic flag?', options: ['5', '4', '6', '7'], correctOption: 'A' },
      { text: 'Which sport is played with five players per team on a court?', options: ['Basketball', 'Volleyball', 'Baseball', 'Rugby'], correctOption: 'A' },
      { text: 'Which country won the FIFA World Cup in 2018?', options: ['France', 'Germany', 'Brazil', 'Argentina'], correctOption: 'A' },
      { text: 'In tennis, a score of zero is called what?', options: ['Love', 'Nil', 'Blank', 'Zero'], correctOption: 'A' },
      { text: 'What is the maximum score in a frame of ten-pin bowling?', options: ['300', '250', '200', '275'], correctOption: 'A' },
      { text: 'Which sport uses a puck?', options: ['Ice hockey', 'Golf', 'Cricket', 'Tennis'], correctOption: 'A' }
    ],
    'cat-geography': [
      { text: 'What is the largest continent by land area?', options: ['Asia', 'Africa', 'Europe', 'North America'], correctOption: 'A' },
      { text: 'What is the capital of France?', options: ['Paris', 'Berlin', 'Madrid', 'Rome'], correctOption: 'A' },
      { text: 'Which river is the longest in the world?', options: ['Nile', 'Amazon', 'Yangtze', 'Mississippi'], correctOption: 'A' },
      { text: 'Which ocean lies east of the United States?', options: ['Atlantic Ocean', 'Pacific Ocean', 'Indian Ocean', 'Arctic Ocean'], correctOption: 'A' },
      { text: 'Which country has the highest population?', options: ['India', 'China', 'United States', 'Brazil'], correctOption: 'A' },
      { text: 'Which mountain range contains Mount Everest?', options: ['Himalayas', 'Andes', 'Alps', 'Rockies'], correctOption: 'A' },
      { text: 'What is the capital of Japan?', options: ['Tokyo', 'Seoul', 'Beijing', 'Bangkok'], correctOption: 'A' },
      { text: 'Which desert is the largest hot desert on Earth?', options: ['Sahara', 'Gobi', 'Kalahari', 'Mojave'], correctOption: 'A' },
      { text: 'Which U.S. state is nicknamed the Sunshine State?', options: ['Florida', 'California', 'Texas', 'Arizona'], correctOption: 'A' },
      { text: 'What is the smallest country in the world by area?', options: ['Vatican City', 'Monaco', 'San Marino', 'Liechtenstein'], correctOption: 'A' }
    ],
    'cat-literature': [
      { text: 'Who wrote Romeo and Juliet?', options: ['William Shakespeare', 'Charles Dickens', 'Jane Austen', 'Mark Twain'], correctOption: 'A' },
      { text: 'What is the title of Orwell\'s dystopian novel?', options: ['1984', 'Brave New World', 'Fahrenheit 451', 'The Giver'], correctOption: 'A' },
      { text: 'Who wrote Pride and Prejudice?', options: ['Jane Austen', 'Emily Brontë', 'Charlotte Brontë', 'Louisa May Alcott'], correctOption: 'A' },
      { text: 'Which novel begins with “Call me Ishmael”?', options: ['Moby-Dick', 'The Odyssey', 'Treasure Island', 'The Iliad'], correctOption: 'A' },
      { text: 'Which school is central to Harry Potter?', options: ['Hogwarts', 'Beauxbatons', 'Durmstrang', 'Ilvermorny'], correctOption: 'A' },
      { text: 'Who wrote The Hobbit?', options: ['J.R.R. Tolkien', 'C.S. Lewis', 'J.K. Rowling', 'George R.R. Martin'], correctOption: 'A' },
      { text: 'Who is Sherlock Holmes\'s assistant?', options: ['Dr. Watson', 'Inspector Lestrade', 'Mrs. Hudson', 'Professor Moriarty'], correctOption: 'A' },
      { text: 'Which novel features Atticus Finch?', options: ['To Kill a Mockingbird', 'The Great Gatsby', 'Of Mice and Men', '1984'], correctOption: 'A' },
      { text: 'What type of poem has 14 lines?', options: ['Sonnet', 'Haiku', 'Limerick', 'Epic'], correctOption: 'A' },
      { text: 'Which book opens with “It was the best of times, it was the worst of times”?', options: ['A Tale of Two Cities', 'Great Expectations', 'Oliver Twist', 'David Copperfield'], correctOption: 'A' }
    ],
    'cat-movies': [
      { text: 'Which movie features the character Indiana Jones?', options: ['Raiders of the Lost Ark', 'The Matrix', 'Jurassic Park', 'Back to the Future'], correctOption: 'A' },
      { text: 'Who directed Jurassic Park?', options: ['Steven Spielberg', 'James Cameron', 'Christopher Nolan', 'Ridley Scott'], correctOption: 'A' },
      { text: 'Which film won Best Picture at the 2020 Academy Awards?', options: ['Parasite', '1917', 'Joker', 'Once Upon a Time in Hollywood'], correctOption: 'A' },
      { text: 'Which superhero appears in The Dark Knight?', options: ['Batman', 'Spider-Man', 'Iron Man', 'Superman'], correctOption: 'A' },
      { text: 'Which movie features the ship Titanic?', options: ['Titanic', 'The Poseidon Adventure', 'Pirates of the Caribbean', 'Life of Pi'], correctOption: 'A' },
      { text: 'Which animated movie features Olaf the snowman?', options: ['Frozen', 'Toy Story', 'Shrek', 'Moana'], correctOption: 'A' },
      { text: 'What color pill does Neo take in The Matrix?', options: ['Red', 'Blue', 'Green', 'Yellow'], correctOption: 'B' },
      { text: 'Which Tom Hanks movie is set on a deserted island?', options: ['Cast Away', 'The Beach', 'Life of Pi', 'Robin Crusoe'], correctOption: 'A' },
      { text: 'Which film series includes Darth Vader?', options: ['Star Wars', 'Star Trek', 'Avatar', 'The Matrix'], correctOption: 'A' },
      { text: 'Which film includes the song “My Heart Will Go On”?', options: ['Titanic', 'The Bodyguard', 'Ghost', 'Dirty Dancing'], correctOption: 'A' }
    ],
    'cat-music': [
      { text: 'Which band released Hey Jude?', options: ['The Beatles', 'The Rolling Stones', 'Queen', 'Pink Floyd'], correctOption: 'A' },
      { text: 'Which instrument has 88 keys?', options: ['Piano', 'Guitar', 'Violin', 'Flute'], correctOption: 'A' },
      { text: 'Who is known as the King of Pop?', options: ['Michael Jackson', 'Elvis Presley', 'Prince', 'Madonna'], correctOption: 'A' },
      { text: 'Which singer released Rolling in the Deep?', options: ['Adele', 'Beyoncé', 'Taylor Swift', 'Rihanna'], correctOption: 'A' },
      { text: 'What instrument does a drummer typically play?', options: ['Drums', 'Violin', 'Flute', 'Cello'], correctOption: 'A' },
      { text: 'Which band is known for Bohemian Rhapsody?', options: ['Queen', 'The Beatles', 'Nirvana', 'U2'], correctOption: 'A' },
      { text: 'Which genre is Taylor Swift best known for?', options: ['Pop and country', 'Classical', 'Jazz', 'Hip hop'], correctOption: 'A' },
      { text: 'How many strings does a standard guitar usually have?', options: ['6', '4', '7', '8'], correctOption: 'A' },
      { text: 'Which city is closely associated with Country music in the United States?', options: ['Nashville', 'Austin', 'New Orleans', 'Los Angeles'], correctOption: 'A' },
      { text: 'Which artist released the album Thriller?', options: ['Michael Jackson', 'Prince', 'Madonna', 'Whitney Houston'], correctOption: 'A' }
    ],
    'cat-math': [
      { text: 'What is 8 × 7?', options: ['56', '54', '58', '60'], correctOption: 'A' },
      { text: 'What is the value of π rounded to two decimal places?', options: ['3.14', '3.15', '3.13', '3.12'], correctOption: 'A' },
      { text: 'What is the square root of 81?', options: ['9', '7', '8', '10'], correctOption: 'A' },
      { text: 'What is 15% of 200?', options: ['30', '20', '25', '35'], correctOption: 'A' },
      { text: 'What is the next prime number after 7?', options: ['11', '9', '10', '13'], correctOption: 'A' },
      { text: 'What does 2 + 2 × 2 equal?', options: ['6', '8', '4', '10'], correctOption: 'A' },
      { text: 'How many degrees are in a right angle?', options: ['90', '45', '180', '360'], correctOption: 'A' },
      { text: 'What is the perimeter of a square with side length 5?', options: ['20', '25', '10', '15'], correctOption: 'A' },
      { text: 'What is the only even prime number?', options: ['2', '0', '4', '6'], correctOption: 'A' },
      { text: 'What is the sum of the interior angles of a triangle?', options: ['180°', '90°', '360°', '270°'], correctOption: 'A' }
    ],
    'cat-nature': [
      { text: 'What process do plants use to convert sunlight into energy?', options: ['Photosynthesis', 'Respiration', 'Digestion', 'Transpiration'], correctOption: 'A' },
      { text: 'What is the largest land animal?', options: ['African elephant', 'Giraffe', 'Hippopotamus', 'Rhinoceros'], correctOption: 'A' },
      { text: 'Which animal is often called the king of the jungle?', options: ['Lion', 'Tiger', 'Elephant', 'Gorilla'], correctOption: 'A' },
      { text: 'What do bees collect from flowers?', options: ['Nectar', 'Water', 'Sand', 'Leaf juice'], correctOption: 'A' },
      { text: 'What is the chemical name of table salt?', options: ['Sodium chloride', 'Calcium carbonate', 'Potassium nitrate', 'Magnesium sulfate'], correctOption: 'A' },
      { text: 'Which layer of Earth do we live on?', options: ['Crust', 'Mantle', 'Core', 'Outer core'], correctOption: 'A' },
      { text: 'What is the largest ocean on Earth?', options: ['Pacific Ocean', 'Atlantic Ocean', 'Indian Ocean', 'Arctic Ocean'], correctOption: 'A' },
      { text: 'Which type of tree keeps its leaves all year?', options: ['Evergreen', 'Deciduous', 'Birch', 'Maple'], correctOption: 'A' },
      { text: 'Which gas do humans breathe in and use for survival?', options: ['Oxygen', 'Carbon dioxide', 'Nitrogen', 'Helium'], correctOption: 'A' },
      { text: 'Which bird is known for being able to mimic human speech?', options: ['Parrot', 'Owl', 'Eagle', 'Pigeon'], correctOption: 'A' }
    ],
    'cat-technology': [
      { text: 'Which company created the iPhone?', options: ['Apple', 'Samsung', 'Nokia', 'Microsoft'], correctOption: 'A' },
      { text: 'What does CPU stand for?', options: ['Central Processing Unit', 'Computer Power Unit', 'Central Program Utility', 'Core Processing Utility'], correctOption: 'A' },
      { text: 'Which language is commonly used for web pages?', options: ['HTML', 'C++', 'Go', 'Rust'], correctOption: 'A' },
      { text: 'Which device stores data permanently?', options: ['Hard drive', 'RAM', 'CPU', 'Monitor'], correctOption: 'A' },
      { text: 'What does Wi-Fi stand for?', options: ['Wireless Fidelity', 'Wide Fiber', 'Web Internet Frequency', 'Wireless File Interface'], correctOption: 'A' },
      { text: 'Which programming language was created by Sun Microsystems?', options: ['Java', 'Python', 'Ruby', 'Swift'], correctOption: 'A' },
      { text: 'Which technology is used to send messages over the internet?', options: ['Packets', 'Bits of sound', 'Cables only', 'Bluetooth only'], correctOption: 'A' },
      { text: 'Which company created Windows?', options: ['Microsoft', 'Apple', 'Google', 'IBM'], correctOption: 'A' },
      { text: 'What is the main purpose of an operating system?', options: ['Manage hardware and software resources', 'Connect printers only', 'Store electricity', 'Display only text'], correctOption: 'A' },
      { text: 'Which network protocol is used for the web?', options: ['HTTP', 'FTP', 'SMTP', 'SSH'], correctOption: 'A' }
    ],
    'cat-art': [
      { text: 'Which artist painted the Mona Lisa?', options: ['Leonardo da Vinci', 'Vincent van Gogh', 'Pablo Picasso', 'Claude Monet'], correctOption: 'A' },
      { text: 'What type of art is created with pigments on canvas?', options: ['Painting', 'Sculpture', 'Photography', 'Architecture'], correctOption: 'A' },
      { text: 'Which color is created by mixing blue and yellow?', options: ['Green', 'Orange', 'Purple', 'Red'], correctOption: 'A' },
      { text: 'What art form uses clay and shaping tools?', options: ['Pottery', 'Painting', 'Printmaking', 'Drawing'], correctOption: 'A' },
      { text: 'Which building is famous for its glass pyramid entrance?', options: ['The Louvre', 'Eiffel Tower', 'Taj Mahal', 'Big Ben'], correctOption: 'A' },
      { text: 'Which artist is known for Starry Night?', options: ['Vincent van Gogh', 'Leonardo da Vinci', 'Michelangelo', 'Salvador Dalí'], correctOption: 'A' },
      { text: 'What is a sculpture?', options: ['A three-dimensional artwork', 'A flat drawing', 'A photograph', 'A poem'], correctOption: 'A' },
      { text: 'Which medium is used in charcoal drawing?', options: ['Charcoal', 'Oil', 'Watercolor', 'Marble'], correctOption: 'A' },
      { text: 'What kind of art is made by arranging colored pieces?', options: ['Mosaic', 'Woodcut', 'Calligraphy', 'Ceramics'], correctOption: 'A' },
      { text: 'Which art style uses bold, simplified forms and strong colors?', options: ['Modern art', 'Renaissance', 'Baroque', 'Classical'], correctOption: 'A' }
    ],
    'cat-food': [
      { text: 'Which fruit is commonly used to make guacamole?', options: ['Avocado', 'Mango', 'Berry', 'Orange'], correctOption: 'A' },
      { text: 'Which grain is used to make bread?', options: ['Wheat', 'Rice', 'Corn', 'Barley'], correctOption: 'A' },
      { text: 'Which cuisine is known for sushi?', options: ['Japanese', 'Italian', 'Mexican', 'French'], correctOption: 'A' },
      { text: 'Which vitamin is found in citrus fruits?', options: ['Vitamin C', 'Vitamin D', 'Vitamin A', 'Vitamin K'], correctOption: 'A' },
      { text: 'What is the main ingredient in hummus?', options: ['Chickpeas', 'Tomatoes', 'Beans', 'Rice'], correctOption: 'A' },
      { text: 'Which cooking method uses dry heat in an oven?', options: ['Baking', 'Boiling', 'Steaming', 'Fermenting'], correctOption: 'A' },
      { text: 'Which fruit is known for being red and sweet?', options: ['Strawberry', 'Lemon', 'Pear', 'Banana'], correctOption: 'A' },
      { text: 'What is the main ingredient in cheese?', options: ['Milk', 'Flour', 'Eggs', 'Rice'], correctOption: 'A' },
      { text: 'Which vegetable is commonly used in tomato soup?', options: ['Tomato', 'Potato', 'Carrot', 'Cabbage'], correctOption: 'A' },
      { text: 'Which spice is commonly used to flavor pizza?', options: ['Oregano', 'Cinnamon', 'Nutmeg', 'Cardamom'], correctOption: 'A' }
    ],
    'cat-animals': [
      { text: 'Which mammal is famous for being able to fly?', options: ['Bat', 'Mouse', 'Squirrel', 'Rabbit'], correctOption: 'A' },
      { text: 'Which animal is the largest on Earth?', options: ['Blue whale', 'Elephant', 'Giraffe', 'Hippopotamus'], correctOption: 'A' },
      { text: 'Which animal is known for carrying its house on its back?', options: ['Snail', 'Turtle', 'Crab', 'Lobster'], correctOption: 'A' },
      { text: 'Which animal is the fastest land mammal?', options: ['Cheetah', 'Horse', 'Lion', 'Gazelle'], correctOption: 'A' },
      { text: 'Which animal is famous for hibernating in winter?', options: ['Bear', 'Rabbit', 'Fox', 'Wolf'], correctOption: 'A' },
      { text: 'Which marine animal uses a shell for protection?', options: ['Turtle', 'Dolphin', 'Shark', 'Seal'], correctOption: 'A' },
      { text: 'Which animal is known for its black-and-white stripes?', options: ['Zebra', 'Tiger', 'Leopard', 'Panda'], correctOption: 'A' },
      { text: 'Which animal is a symbol of wisdom in many cultures?', options: ['Owl', 'Eagle', 'Falcon', 'Sparrow'], correctOption: 'A' },
      { text: 'Which animal is often called the king of the jungle?', options: ['Lion', 'Elephant', 'Gorilla', 'Tiger'], correctOption: 'A' },
      { text: 'Which large cat is famous for its orange fur and black stripes?', options: ['Tiger', 'Leopard', 'Cheetah', 'Lynx'], correctOption: 'A' }
    ],
    'cat-space': [
      { text: 'Which planet is closest to the Sun?', options: ['Mercury', 'Venus', 'Mars', 'Earth'], correctOption: 'A' },
      { text: 'What is the name of Earth\'s natural satellite?', options: ['Moon', 'Titan', 'Europa', 'Phobos'], correctOption: 'A' },
      { text: 'Which planet is famous for its rings?', options: ['Saturn', 'Jupiter', 'Neptune', 'Mars'], correctOption: 'A' },
      { text: 'What force keeps the planets in orbit around the Sun?', options: ['Gravity', 'Magnetism', 'Friction', 'Radiation'], correctOption: 'A' },
      { text: 'Which galaxy contains our Solar System?', options: ['Milky Way', 'Andromeda', 'Whirlpool', 'Sombrero'], correctOption: 'A' },
      { text: 'What is a comet mostly made of?', options: ['Ice and dust', 'Metal and lava', 'Glass and gas', 'Rock and water'], correctOption: 'A' },
      { text: 'Which planet is known as the Red Planet?', options: ['Mars', 'Venus', 'Mercury', 'Jupiter'], correctOption: 'A' },
      { text: 'Which spacecraft landed humans on the Moon?', options: ['Apollo', 'Voyager', 'Gemini', 'Sputnik'], correctOption: 'A' },
      { text: 'What is the name of the star at the center of our Solar System?', options: ['Sun', 'Proxima Centauri', 'Sirius', 'Betelgeuse'], correctOption: 'A' },
      { text: 'Which planet is the largest in the Solar System?', options: ['Jupiter', 'Saturn', 'Earth', 'Neptune'], correctOption: 'A' }
    ],
    'cat-health': [
      { text: 'Which activity improves heart health?', options: ['Regular exercise', 'Sleeping all day', 'Skipping meals', 'Smoking'], correctOption: 'A' },
      { text: 'What is the main purpose of sleep?', options: ['Rest and recovery', 'Digesting food', 'Producing blood', 'Cooling the body'], correctOption: 'A' },
      { text: 'Which food group is essential for strong bones?', options: ['Dairy and calcium-rich foods', 'Candy', 'Soda', 'Fast food'], correctOption: 'A' },
      { text: 'Which habit is most harmful to lungs?', options: ['Smoking', 'Drinking water', 'Walking', 'Stretching'], correctOption: 'A' },
      { text: 'What is a healthy way to manage stress?', options: ['Exercise and rest', 'Skipping sleep', 'Smoking', 'Eating junk food'], correctOption: 'A' },
      { text: 'Which organ is most directly responsible for filtering blood?', options: ['Kidneys', 'Stomach', 'Lungs', 'Liver'], correctOption: 'A' },
      { text: 'What should you do before exercising?', options: ['Warm up', 'Drink soda', 'Skip hydration', 'Lie down'], correctOption: 'A' },
      { text: 'Which nutrient helps build muscles?', options: ['Protein', 'Sugar', 'Salt', 'Caffeine'], correctOption: 'A' },
      { text: 'How can you keep your immune system strong?', options: ['Eat well and sleep enough', 'Eat only sweets', 'Avoid all exercise', 'Stay up late'], correctOption: 'A' },
      { text: 'Which is the best way to stay hydrated?', options: ['Drink water regularly', 'Only drink soda', 'Skip fluids', 'Drink coffee only'], correctOption: 'A' }
    ],
    'cat-language': [
      { text: 'What is the main purpose of grammar?', options: ['To structure sentences clearly', 'To add colors', 'To count numbers', 'To change spelling randomly'], correctOption: 'A' },
      { text: 'Which word is a noun?', options: ['Table', 'Quickly', 'Beautifully', 'Run'], correctOption: 'A' },
      { text: 'What is the opposite of “begin”?', options: ['End', 'Open', 'Start', 'Finish'], correctOption: 'A' },
      { text: 'Which sentence is written correctly?', options: ['She walks to school every day.', 'She walk to school every day.', 'She walking to school every day.', 'She walked to school every days.'], correctOption: 'A' },
      { text: 'What is a synonym for “happy”?', options: ['Joyful', 'Sad', 'Angry', 'Silent'], correctOption: 'A' },
      { text: 'What is an antonym of “hot”?', options: ['Cold', 'Warm', 'Bright', 'Fast'], correctOption: 'A' },
      { text: 'Which part of speech describes a verb, adjective, or other adverb?', options: ['Adverb', 'Noun', 'Pronoun', 'Preposition'], correctOption: 'A' },
      { text: 'What is a paragraph?', options: ['A group of related sentences', 'A single word', 'A punctuation mark', 'A number'], correctOption: 'A' },
      { text: 'Which punctuation mark ends a statement?', options: ['Period', 'Comma', 'Question mark', 'Colon'], correctOption: 'A' },
      { text: 'What does a dictionary help you do?', options: ['Find word meanings', 'Measure height', 'Cook food', 'Build a chair'], correctOption: 'A' }
    ],
    'cat-psychology': [
      { text: 'What is memory?', options: ['The ability to store and recall information', 'A feeling of hunger', 'A type of plant', 'A kind of light'], correctOption: 'A' },
      { text: 'Which emotion is often linked with fear?', options: ['Anxiety', 'Joy', 'Calm', 'Excitement'], correctOption: 'A' },
      { text: 'What is empathy?', options: ['Understanding and sharing another\'s feelings', 'Ignoring others', 'Feeling angry', 'Being silent'], correctOption: 'A' },
      { text: 'What is a common goal of mindfulness?', options: ['To stay present and focused', 'To ignore thoughts', 'To sleep constantly', 'To avoid all emotions'], correctOption: 'A' },
      { text: 'Which behavior is usually associated with good mental health?', options: ['Healthy coping skills', 'Avoiding all responsibilities', 'Constant panic', 'Feeling no emotion'], correctOption: 'A' },
      { text: 'What does perception mean?', options: ['How we interpret sensory information', 'A type of memory', 'A social norm', 'A mood'], correctOption: 'A' },
      { text: 'Which factor often influences decision-making?', options: ['Emotions and experience', 'Only random chance', 'Only weather', 'Only calendar dates'], correctOption: 'A' },
      { text: 'What is motivation?', options: ['The drive that causes action', 'A type of food', 'A shape of box', 'A speed limit'], correctOption: 'A' },
      { text: 'What does stress usually do to focus?', options: ['It can reduce it', 'It always improves it', 'It makes you invisible', 'It removes all thoughts'], correctOption: 'A' },
      { text: 'What is a habit?', options: ['A repeated behavior', 'A random idea', 'A type of fruit', 'A physical object'], correctOption: 'A' }
    ],
    'cat-religion': [
      { text: 'Which book is sacred in Christianity?', options: ['The Bible', 'The Torah', 'The Quran', 'The Vedas'], correctOption: 'A' },
      { text: 'What is the holy city of Islam?', options: ['Mecca', 'Jerusalem', 'Rome', 'Cairo'], correctOption: 'A' },
      { text: 'Which religion follows the Five Pillars?', options: ['Islam', 'Christianity', 'Hinduism', 'Judaism'], correctOption: 'A' },
      { text: 'What is the sacred text of Judaism?', options: ['Torah', 'Bible', 'Quran', 'Tripitaka'], correctOption: 'A' },
      { text: 'Which religion is associated with Karma and Dharma?', options: ['Hinduism', 'Islam', 'Buddhism', 'Christianity'], correctOption: 'A' },
      { text: 'What is a temple?', options: ['A place of worship', 'A school', 'A farm', 'A market'], correctOption: 'A' },
      { text: 'Which faith is centered on the teachings of Buddha?', options: ['Buddhism', 'Christianity', 'Judaism', 'Shinto'], correctOption: 'A' },
      { text: 'Which city is holy to Christians, Jews, and Muslims?', options: ['Jerusalem', 'Cairo', 'Mecca', 'Rome'], correctOption: 'A' },
      { text: 'Which religion is based on the life and teachings of Jesus?', options: ['Christianity', 'Islam', 'Hinduism', 'Sikhism'], correctOption: 'A' },
      { text: 'What is a shrine?', options: ['A holy or sacred place', 'A library', 'A sports stadium', 'A hospital'], correctOption: 'A' }
    ],
    'cat-business': [
      { text: 'What is the main goal of a business?', options: ['Create value and earn profit', 'Only spend money', 'Avoid customers', 'Ignore innovation'], correctOption: 'A' },
      { text: 'What is a budget?', options: ['A plan for income and spending', 'A type of machine', 'A legal rule', 'A product design'], correctOption: 'A' },
      { text: 'Which term means the amount of money a company makes?', options: ['Revenue', 'Debt', 'Inventory', 'Expense'], correctOption: 'A' },
      { text: 'What is marketing?', options: ['Promoting and selling products', 'Only counting stock', 'Building roads', 'Changing laws'], correctOption: 'A' },
      { text: 'What is branding?', options: ['Creating a recognizable identity', 'Making a legal contract', 'Writing code', 'Shipping products'], correctOption: 'A' },
      { text: 'Which document describes a business idea and plan?', options: ['Business plan', 'Receipt', 'Invoice', 'Bank statement'], correctOption: 'A' },
      { text: 'What is a stakeholder?', options: ['A person affected by a business decision', 'A machine part', 'A marketing campaign', 'A product feature'], correctOption: 'A' },
      { text: 'What does ROI measure?', options: ['Return on investment', 'Running operating inventory', 'Revenue of interest', 'Rate of income'], correctOption: 'A' },
      { text: 'What is customer service about?', options: ['Helping customers before and after a purchase', 'Ignoring customer issues', 'Lowering product prices only', 'Producing raw materials'], correctOption: 'A' },
      { text: 'Which financial statement shows profit and loss?', options: ['Income statement', 'Balance sheet', 'Cash register', 'Payroll form'], correctOption: 'A' }
    ],
    'cat-gaming': [
      { text: 'What is the objective of most racing games?', options: ['Finish the track first', 'Solve puzzles', 'Collect plants', 'Write code'], correctOption: 'A' },
      { text: 'What does FPS stand for in gaming?', options: ['First-person shooter', 'Fast play score', 'Full power system', 'Frame priority setting'], correctOption: 'A' },
      { text: 'Which action often wins a strategy game?', options: ['Planning ahead', 'Spamming random clicks', 'Ignoring resources', 'Skipping turns'], correctOption: 'A' },
      { text: 'What is loot in an RPG?', options: ['Rewards such as gear or items', 'A type of map', 'A weapon sound', 'A difficulty level'], correctOption: 'A' },
      { text: 'What is a game controller used for?', options: ['Input control in games', 'Listening to music', 'Reading books', 'Making coffee'], correctOption: 'A' },
      { text: 'What does multiplayer mean?', options: ['More than one player participates', 'A single-player mode', 'A high-resolution screen', 'A new game engine'], correctOption: 'A' },
      { text: 'Which genre focuses on solving logic puzzles?', options: ['Puzzle games', 'Racing games', 'Sports games', 'Platformers'], correctOption: 'A' },
      { text: 'What is a save point?', options: ['A point where progress is stored', 'A cheat code', 'A game title', 'A map icon'], correctOption: 'A' },
      { text: 'Which feature usually improves performance in a game?', options: ['Better graphics settings', 'Slower loading times', 'More random crashes', 'No sound'], correctOption: 'A' },
      { text: 'What is a game mechanic?', options: ['A rule or system that defines gameplay', 'A player name', 'A hard drive', 'A photo filter'], correctOption: 'A' }
    ],
    'cat-cars': [
      { text: 'What is the main purpose of a car\'s engine?', options: ['Generate power to move the vehicle', 'Store fuel only', 'Play music', 'Control traffic lights'], correctOption: 'A' },
      { text: 'What does ABS help a car do?', options: ['Prevent wheels from locking during braking', 'Increase tire pressure', 'Turn the steering wheel', 'Cool the cabin'], correctOption: 'A' },
      { text: 'Which fuel is commonly used in most gasoline cars?', options: ['Petrol', 'Diesel only', 'Coal', 'Natural gas'], correctOption: 'A' },
      { text: 'What is a transmission?', options: ['A system that transfers power to the wheels', 'A car radio', 'A seat cushion', 'A steering wheel'], correctOption: 'A' },
      { text: 'What does a tachometer measure?', options: ['Engine revolutions per minute', 'Fuel level', 'Temperature outside', 'Seat position'], correctOption: 'A' },
      { text: 'Which component is responsible for steering?', options: ['Steering wheel and linkage', 'Dashboard', 'Headlights', 'Roof rack'], correctOption: 'A' },
      { text: 'What is the purpose of a car battery?', options: ['Start the engine and power electronics', 'Move the wheels directly', 'Measure tire pressure', 'Cool the brakes'], correctOption: 'A' },
      { text: 'Which type of vehicle is designed mainly for off-road driving?', options: ['SUV', 'Sedan', 'Hatchback', 'Coupe'], correctOption: 'A' },
      { text: 'What does MPG measure?', options: ['Miles per gallon of fuel', 'Miles per hour', 'Milligrams per gram', 'Meters per gallon'], correctOption: 'A' },
      { text: 'Which part of the car is responsible for braking?', options: ['Brake pads', 'Headlights', 'Roof', 'Spoiler'], correctOption: 'A' }
    ],
    'cat-astronomy': [
      { text: 'What is a galaxy?', options: ['A vast collection of stars, gas, and dust', 'A single star', 'A planet', 'A moon'], correctOption: 'A' },
      { text: 'Which planet is known for its prominent rings?', options: ['Saturn', 'Mars', 'Mercury', 'Venus'], correctOption: 'A' },
      { text: 'What is a comet mainly composed of?', options: ['Ice, dust, and rock', 'Hot metal', 'Water vapor only', 'Copper and steel'], correctOption: 'A' },
      { text: 'What do astronomers use to study distant stars?', options: ['Telescopes', 'Microscopes', 'Magnets', 'Compasses'], correctOption: 'A' },
      { text: 'Which planet is closest to the Sun?', options: ['Mercury', 'Venus', 'Earth', 'Mars'], correctOption: 'A' },
      { text: 'What is a lunar eclipse?', options: ['Earth passes between the Sun and the Moon', 'The Moon passes between Earth and Sun', 'A planet collides with the Moon', 'The Sun disappears'], correctOption: 'A' },
      { text: 'What is a black hole?', options: ['A region of space with extremely strong gravity', 'A giant star', 'A planet with no atmosphere', 'An empty nebula'], correctOption: 'A' },
      { text: 'Which planet has the Great Red Spot?', options: ['Jupiter', 'Mars', 'Neptune', 'Venus'], correctOption: 'A' },
      { text: 'What is the name of Earth\'s star?', options: ['Sun', 'Sirius', 'Polaris', 'Vega'], correctOption: 'A' },
      { text: 'What does an orbit describe?', options: ['The path an object follows around another object', 'A burst of light', 'A weather pattern', 'A mountain shape'], correctOption: 'A' }
    ],
    'cat-philosophy': [
      { text: 'What is ethics mainly concerned with?', options: ['Right and wrong behavior', 'Chemical reactions', 'Weather patterns', 'Historic dates'], correctOption: 'A' },
      { text: 'Who is known for the statement “I think, therefore I am”?', options: ['Descartes', 'Plato', 'Aristotle', 'Socrates'], correctOption: 'A' },
      { text: 'What is logic?', options: ['Reasoning according to valid rules', 'A physical experiment', 'A shape of building', 'A form of medicine'], correctOption: 'A' },
      { text: 'What does metaphysics study?', options: ['The nature of reality', 'Numbers only', 'Weather systems', 'Human anatomy'], correctOption: 'A' },
      { text: 'What is truth?', options: ['A concept about what is correct and real', 'A kind of stone', 'A type of music', 'A form of light'], correctOption: 'A' },
      { text: 'What is an argument in philosophy?', options: ['A set of reasons supporting a conclusion', 'A loud disagreement', 'A type of building', 'A political event'], correctOption: 'A' },
      { text: 'Which philosopher taught in Athens and questioned many assumptions?', options: ['Socrates', 'Descartes', 'Hume', 'Kant'], correctOption: 'A' },
      { text: 'What is free will?', options: ['The ability to choose one\'s own actions', 'A law of physics', 'An animal instinct', 'A mood state'], correctOption: 'A' },
      { text: 'What is epistemology?', options: ['The study of knowledge', 'The study of numbers', 'The study of cells', 'The study of language'], correctOption: 'A' },
      { text: 'What does existentialism focus on?', options: ['Human existence and personal choice', 'Mathematical proofs', 'Plant biology', 'Mechanical design'], correctOption: 'A' }
    ],
    'cat-design': [
      { text: 'What is user experience design mainly concerned with?', options: ['How a person feels while using a product', 'Only making it colorful', 'Writing long documents', 'Printing on paper'], correctOption: 'A' },
      { text: 'What is a wireframe?', options: ['A simple layout blueprint', 'A code compiler', 'A database query', 'A marketing report'], correctOption: 'A' },
      { text: 'Why is contrast important in design?', options: ['It helps elements stand out and improve readability', 'It reduces creativity', 'It always slows users down', 'It removes color'], correctOption: 'A' },
      { text: 'What does UI stand for?', options: ['User Interface', 'Universal Input', 'User Information', 'Unified Internet'], correctOption: 'A' },
      { text: 'What is hierarchy in design?', options: ['The arrangement of elements by importance', 'A set of colors', 'A storage method', 'A layout grid only'], correctOption: 'A' },
      { text: 'Which principle helps users understand where to look first?', options: ['Visual hierarchy', 'Random placement', 'Overcrowding', 'No spacing'], correctOption: 'A' },
      { text: 'What is a prototype?', options: ['An early model of a design', 'A final product only', 'A marketing slogan', 'A code library'], correctOption: 'A' },
      { text: 'Why is whitespace useful in design?', options: ['It improves readability and focus', 'It makes text invisible', 'It removes meaning', 'It adds noise'], correctOption: 'A' },
      { text: 'What is accessibility in design?', options: ['Designing for people with different needs and abilities', 'Making content shorter only', 'Using only one color', 'Ignoring users'], correctOption: 'A' },
      { text: 'Which design tool is commonly used for interfaces?', options: ['Figma', 'Excel', 'Photoshop without design tools', 'A command line'], correctOption: 'A' }
    ]
  };

  function buildQuestionPoolFromBank(questionBank) {
    const questions = {};
    let questionNumber = 1;

    categories.forEach(({ id }) => {
      const pool = questionBank[id] || [];
      for (let index = 0; index < 42; index++) {
        const question = pool[index % pool.length] || {
          text: `Which answer best matches ${id}?`,
          options: ['Study', 'Guessing', 'Confusion', 'Delay'],
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

  window.quizLocaleData.en = {
    categories: translatedCategories,
    questions: buildQuestionPoolFromBank(categoryQuestionBanks)
  };
})();
