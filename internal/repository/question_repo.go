package repository

import (
	"database/sql"
	"fmt"

	"go_simple_quiz_adventure/internal/models"

	_ "github.com/glebarez/sqlite"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) InitSchema() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS categories (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS questions (
			id TEXT PRIMARY KEY,
			category_id TEXT NOT NULL,
			text TEXT NOT NULL,
			option_a TEXT NOT NULL,
			option_b TEXT NOT NULL,
			option_c TEXT NOT NULL,
			option_d TEXT NOT NULL,
			correct_option TEXT NOT NULL,
			FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
		)`,
	}

	for _, query := range queries {
		if _, err := r.db.Exec(query); err != nil {
			return fmt.Errorf("execute schema query: %w", err)
		}
	}

	return nil
}

func (r *QuestionRepository) Seed() error {
	categories := buildSeedCategories()
	for _, category := range categories {
		if _, err := r.db.Exec(`INSERT OR IGNORE INTO categories (id, name) VALUES (?, ?)`, category.ID, category.Name); err != nil {
			return fmt.Errorf("insert category %s: %w", category.ID, err)
		}
	}

	questions := buildSeedQuestions()
	for _, question := range questions {
		if _, err := r.db.Exec(`INSERT OR IGNORE INTO questions (id, category_id, text, option_a, option_b, option_c, option_d, correct_option) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, question.ID, question.CategoryID, question.Text, question.OptionA, question.OptionB, question.OptionC, question.OptionD, question.CorrectOption); err != nil {
			return fmt.Errorf("insert question %s: %w", question.ID, err)
		}
	}

	return nil
}

func buildSeedCategories() []models.Category {
	return []models.Category{
		{ID: "cat-programming", Name: "Programming"},
		{ID: "cat-history", Name: "History"},
		{ID: "cat-science", Name: "Science"},
		{ID: "cat-sports", Name: "Sports"},
		{ID: "cat-geography", Name: "Geography"},
		{ID: "cat-literature", Name: "Literature"},
		{ID: "cat-movies", Name: "Movies"},
		{ID: "cat-music", Name: "Music"},
		{ID: "cat-math", Name: "Math"},
		{ID: "cat-nature", Name: "Nature"},
		{ID: "cat-technology", Name: "Technology"},
		{ID: "cat-art", Name: "Art"},
		{ID: "cat-food", Name: "Food"},
		{ID: "cat-animals", Name: "Animals"},
		{ID: "cat-space", Name: "Space"},
		{ID: "cat-health", Name: "Health"},
		{ID: "cat-language", Name: "Language"},
		{ID: "cat-psychology", Name: "Psychology"},
		{ID: "cat-religion", Name: "Religion"},
		{ID: "cat-business", Name: "Business"},
		{ID: "cat-gaming", Name: "Gaming"},
		{ID: "cat-cars", Name: "Cars"},
		{ID: "cat-astronomy", Name: "Astronomy"},
		{ID: "cat-philosophy", Name: "Philosophy"},
		{ID: "cat-design", Name: "Design"},
	}
}

func buildSeedQuestions() []models.Question {
	categories := buildSeedCategories()
	questions := make([]models.Question, 0, 1008)
	questionNumber := 1
	for _, category := range categories {
		for i := 0; i < 42; i++ {
			question := generateSeedQuestion(category, i, questionNumber)
			questions = append(questions, question)
			questionNumber++
		}
	}
	return questions
}

type seedQuestion struct {
	Text          string
	Options       []string
	CorrectOption string
}

var categoryQuestionBanks = map[string][]seedQuestion{
	"cat-programming": {
		{Text: "Which keyword declares a function in Go?", Options: []string{"func", "let", "class", "def"}, CorrectOption: "A"},
		{Text: "What does Go use to run code concurrently?", Options: []string{"goroutines", "threads", "classes", "modules"}, CorrectOption: "A"},
		{Text: "Which data structure keeps key-value pairs in Go?", Options: []string{"map", "stack", "queue", "array"}, CorrectOption: "A"},
		{Text: "Which command runs a Go program from the module root?", Options: []string{"go run .", "npm start", "cargo run", "python app.py"}, CorrectOption: "A"},
		{Text: "Which keyword is used to import a package in Go?", Options: []string{"import", "include", "require", "using"}, CorrectOption: "B"},
		{Text: "Which type is commonly used for text in Go?", Options: []string{"string", "int", "bool", "char"}, CorrectOption: "A"},
		{Text: "Which Go tool compiles packages and dependencies?", Options: []string{"go build", "gcc", "javac", "tsc"}, CorrectOption: "A"},
		{Text: "What is a slice in Go?", Options: []string{"a dynamic sequence of values", "a single integer", "a compiled file", "a database record"}, CorrectOption: "A"},
		{Text: "Which type in Go holds a single true/false value?", Options: []string{"bool", "string", "int", "byte"}, CorrectOption: "A"},
		{Text: "Which Go keyword creates a new variable with an inferred type?", Options: []string{"var", "class", "def", "new"}, CorrectOption: "A"},
	},
	"cat-history": {
		{Text: "Who was the first President of the United States?", Options: []string{"George Washington", "Thomas Jefferson", "John Adams", "Abraham Lincoln"}, CorrectOption: "A"},
		{Text: "In which year did the Berlin Wall fall?", Options: []string{"1989", "1987", "1991", "1979"}, CorrectOption: "A"},
		{Text: "Which empire was ruled by Genghis Khan?", Options: []string{"Mongol Empire", "Roman Empire", "Ottoman Empire", "Byzantine Empire"}, CorrectOption: "A"},
		{Text: "The Magna Carta was signed in which year?", Options: []string{"1215", "1315", "1415", "1115"}, CorrectOption: "A"},
		{Text: "Which ancient civilization built the pyramids of Giza?", Options: []string{"Egyptians", "Romans", "Greeks", "Vikings"}, CorrectOption: "A"},
		{Text: "Which conflict was fought between the North and South in the United States?", Options: []string{"American Civil War", "World War I", "Revolutionary War", "War of 1812"}, CorrectOption: "A"},
		{Text: "Who discovered the Americas in 1492?", Options: []string{"Christopher Columbus", "Ferdinand Magellan", "Marco Polo", "Amerigo Vespucci"}, CorrectOption: "A"},
		{Text: "Which city was destroyed by the eruption of Mount Vesuvius?", Options: []string{"Pompeii", "Rome", "Athens", "Naples"}, CorrectOption: "A"},
		{Text: "Which document established the U.S. Constitution's framework of government?", Options: []string{"The Constitution", "The Declaration", "The Articles of Confederation", "The Bill of Rights"}, CorrectOption: "A"},
		{Text: "Which king was forced to sign the Magna Carta?", Options: []string{"King John", "Henry VIII", "Edward I", "Richard III"}, CorrectOption: "A"},
	},
	"cat-science": {
		{Text: "What planet is known as the Red Planet?", Options: []string{"Mars", "Venus", "Mercury", "Jupiter"}, CorrectOption: "A"},
		{Text: "What is the chemical formula of water?", Options: []string{"H2O", "CO2", "NaCl", "O2"}, CorrectOption: "A"},
		{Text: "What force keeps planets in orbit around the Sun?", Options: []string{"Gravity", "Magnetism", "Electricity", "Friction"}, CorrectOption: "A"},
		{Text: "Which part of the cell contains genetic material?", Options: []string{"Nucleus", "Cell membrane", "Ribosome", "Cytoplasm"}, CorrectOption: "A"},
		{Text: "Which gas do plants absorb from the atmosphere?", Options: []string{"Carbon dioxide", "Oxygen", "Nitrogen", "Hydrogen"}, CorrectOption: "A"},
		{Text: "What is the center of an atom called?", Options: []string{"Nucleus", "Electron", "Proton", "Neutron"}, CorrectOption: "A"},
		{Text: "What is the boiling point of water at sea level?", Options: []string{"100°C", "90°C", "80°C", "110°C"}, CorrectOption: "A"},
		{Text: "Which organ pumps blood through the body?", Options: []string{"Heart", "Liver", "Lung", "Kidney"}, CorrectOption: "A"},
		{Text: "What gas do humans exhale in large amounts?", Options: []string{"Carbon dioxide", "Oxygen", "Helium", "Argon"}, CorrectOption: "A"},
		{Text: "Which bird is famous for mimicking sounds?", Options: []string{"Parrot", "Eagle", "Penguin", "Owl"}, CorrectOption: "A"},
	},
	"cat-sports": {
		{Text: "How many players are on the field for one soccer team?", Options: []string{"11", "9", "10", "12"}, CorrectOption: "A"},
		{Text: "In baseball, what is a home run?", Options: []string{"A run scored by hitting the ball out of play", "A strikeout", "A foul ball", "A double play"}, CorrectOption: "A"},
		{Text: "How many minutes are in a standard soccer match?", Options: []string{"90", "60", "75", "120"}, CorrectOption: "A"},
		{Text: "Which sport uses a shuttlecock?", Options: []string{"Badminton", "Tennis", "Squash", "Cricket"}, CorrectOption: "A"},
		{Text: "How many rings appear on the Olympic flag?", Options: []string{"5", "4", "6", "7"}, CorrectOption: "A"},
		{Text: "Which sport is played with five players per team on a court?", Options: []string{"Basketball", "Volleyball", "Baseball", "Rugby"}, CorrectOption: "A"},
		{Text: "Which country won the FIFA World Cup in 2018?", Options: []string{"France", "Germany", "Brazil", "Argentina"}, CorrectOption: "A"},
		{Text: "In tennis, a score of zero is called what?", Options: []string{"Love", "Nil", "Blank", "Zero"}, CorrectOption: "A"},
		{Text: "What is the maximum score in a frame of ten-pin bowling?", Options: []string{"300", "250", "200", "275"}, CorrectOption: "A"},
		{Text: "Which sport uses a puck?", Options: []string{"Ice hockey", "Golf", "Cricket", "Tennis"}, CorrectOption: "A"},
	},
	"cat-geography": {
		{Text: "What is the largest continent by land area?", Options: []string{"Asia", "Africa", "Europe", "North America"}, CorrectOption: "A"},
		{Text: "What is the capital of France?", Options: []string{"Paris", "Berlin", "Madrid", "Rome"}, CorrectOption: "A"},
		{Text: "Which river is the longest in the world?", Options: []string{"Nile", "Amazon", "Yangtze", "Mississippi"}, CorrectOption: "A"},
		{Text: "Which ocean lies east of the United States?", Options: []string{"Atlantic Ocean", "Pacific Ocean", "Indian Ocean", "Arctic Ocean"}, CorrectOption: "A"},
		{Text: "Which country has the highest population?", Options: []string{"India", "China", "United States", "Brazil"}, CorrectOption: "A"},
		{Text: "Which mountain range contains Mount Everest?", Options: []string{"Himalayas", "Andes", "Alps", "Rockies"}, CorrectOption: "A"},
		{Text: "What is the capital of Japan?", Options: []string{"Tokyo", "Seoul", "Beijing", "Bangkok"}, CorrectOption: "A"},
		{Text: "Which desert is the largest hot desert on Earth?", Options: []string{"Sahara", "Gobi", "Kalahari", "Mojave"}, CorrectOption: "A"},
		{Text: "Which U.S. state is nicknamed the Sunshine State?", Options: []string{"Florida", "California", "Texas", "Arizona"}, CorrectOption: "A"},
		{Text: "What is the smallest country in the world by area?", Options: []string{"Vatican City", "Monaco", "San Marino", "Liechtenstein"}, CorrectOption: "A"},
	},
	"cat-literature": {
		{Text: "Who wrote Romeo and Juliet?", Options: []string{"William Shakespeare", "Charles Dickens", "Jane Austen", "Mark Twain"}, CorrectOption: "A"},
		{Text: "What is the title of Orwell's dystopian novel?", Options: []string{"1984", "Brave New World", "Fahrenheit 451", "The Giver"}, CorrectOption: "A"},
		{Text: "Who wrote Pride and Prejudice?", Options: []string{"Jane Austen", "Emily Brontë", "Charlotte Brontë", "Louisa May Alcott"}, CorrectOption: "A"},
		{Text: "Which novel begins with Call me Ishmael?", Options: []string{"Moby-Dick", "The Odyssey", "Treasure Island", "The Iliad"}, CorrectOption: "A"},
		{Text: "Which school is central to Harry Potter?", Options: []string{"Hogwarts", "Beauxbatons", "Durmstrang", "Ilvermorny"}, CorrectOption: "A"},
		{Text: "Who wrote The Hobbit?", Options: []string{"J.R.R. Tolkien", "C.S. Lewis", "J.K. Rowling", "George R.R. Martin"}, CorrectOption: "A"},
		{Text: "Who is Sherlock Holmes's assistant?", Options: []string{"Dr. Watson", "Inspector Lestrade", "Mrs. Hudson", "Professor Moriarty"}, CorrectOption: "A"},
		{Text: "Which novel features Atticus Finch?", Options: []string{"To Kill a Mockingbird", "The Great Gatsby", "Of Mice and Men", "1984"}, CorrectOption: "A"},
		{Text: "What type of poem has 14 lines?", Options: []string{"Sonnet", "Haiku", "Limerick", "Epic"}, CorrectOption: "A"},
		{Text: "Which book opens with It was the best of times, it was the worst of times?", Options: []string{"A Tale of Two Cities", "Great Expectations", "Oliver Twist", "David Copperfield"}, CorrectOption: "A"},
	},
	"cat-movies": {
		{Text: "Which movie features the character Indiana Jones?", Options: []string{"Raiders of the Lost Ark", "The Matrix", "Jurassic Park", "Back to the Future"}, CorrectOption: "A"},
		{Text: "Who directed Jurassic Park?", Options: []string{"Steven Spielberg", "James Cameron", "Christopher Nolan", "Ridley Scott"}, CorrectOption: "A"},
		{Text: "Which film won Best Picture at the 2020 Academy Awards?", Options: []string{"Parasite", "1917", "Joker", "Once Upon a Time in Hollywood"}, CorrectOption: "A"},
		{Text: "Which superhero appears in The Dark Knight?", Options: []string{"Batman", "Spider-Man", "Iron Man", "Superman"}, CorrectOption: "A"},
		{Text: "Which movie features the ship Titanic?", Options: []string{"Titanic", "The Poseidon Adventure", "Pirates of the Caribbean", "Life of Pi"}, CorrectOption: "A"},
		{Text: "Which animated movie features Olaf the snowman?", Options: []string{"Frozen", "Toy Story", "Shrek", "Moana"}, CorrectOption: "A"},
		{Text: "What color pill does Neo take in The Matrix?", Options: []string{"Red", "Blue", "Green", "Yellow"}, CorrectOption: "B"},
		{Text: "Which Tom Hanks movie is set on a deserted island?", Options: []string{"Cast Away", "The Beach", "Life of Pi", "Robin Crusoe"}, CorrectOption: "A"},
		{Text: "Which film series includes Darth Vader?", Options: []string{"Star Wars", "Star Trek", "Avatar", "The Matrix"}, CorrectOption: "A"},
		{Text: "Which film includes the song My Heart Will Go On?", Options: []string{"Titanic", "The Bodyguard", "Ghost", "Dirty Dancing"}, CorrectOption: "A"},
	},
	"cat-music": {
		{Text: "Which band released Hey Jude?", Options: []string{"The Beatles", "The Rolling Stones", "Queen", "Pink Floyd"}, CorrectOption: "A"},
		{Text: "Which instrument has 88 keys?", Options: []string{"Piano", "Guitar", "Violin", "Flute"}, CorrectOption: "A"},
		{Text: "Who is known as the King of Pop?", Options: []string{"Michael Jackson", "Elvis Presley", "Prince", "Madonna"}, CorrectOption: "A"},
		{Text: "Which singer released Rolling in the Deep?", Options: []string{"Adele", "Beyoncé", "Taylor Swift", "Rihanna"}, CorrectOption: "A"},
		{Text: "What instrument does a drummer typically play?", Options: []string{"Drums", "Violin", "Flute", "Cello"}, CorrectOption: "A"},
		{Text: "Which band is known for Bohemian Rhapsody?", Options: []string{"Queen", "The Beatles", "Nirvana", "U2"}, CorrectOption: "A"},
		{Text: "Which genre is Taylor Swift best known for?", Options: []string{"Pop and country", "Classical", "Jazz", "Hip hop"}, CorrectOption: "A"},
		{Text: "How many strings does a standard guitar usually have?", Options: []string{"6", "4", "7", "8"}, CorrectOption: "A"},
		{Text: "Which city is closely associated with Country music in the United States?", Options: []string{"Nashville", "Austin", "New Orleans", "Los Angeles"}, CorrectOption: "A"},
		{Text: "Which artist released the album Thriller?", Options: []string{"Michael Jackson", "Prince", "Madonna", "Whitney Houston"}, CorrectOption: "A"},
	},
	"cat-math": {
		{Text: "What is 8 × 7?", Options: []string{"56", "54", "58", "60"}, CorrectOption: "A"},
		{Text: "What is the value of π rounded to two decimal places?", Options: []string{"3.14", "3.15", "3.13", "3.12"}, CorrectOption: "A"},
		{Text: "What is the square root of 81?", Options: []string{"9", "7", "8", "10"}, CorrectOption: "A"},
		{Text: "What is 15% of 200?", Options: []string{"30", "20", "25", "35"}, CorrectOption: "A"},
		{Text: "What is the next prime number after 7?", Options: []string{"11", "9", "10", "13"}, CorrectOption: "A"},
		{Text: "What does 2 + 2 × 2 equal?", Options: []string{"6", "8", "4", "10"}, CorrectOption: "A"},
		{Text: "How many degrees are in a right angle?", Options: []string{"90", "45", "180", "360"}, CorrectOption: "A"},
		{Text: "What is the perimeter of a square with side length 5?", Options: []string{"20", "25", "10", "15"}, CorrectOption: "A"},
		{Text: "What is the only even prime number?", Options: []string{"2", "0", "4", "6"}, CorrectOption: "A"},
		{Text: "What is the sum of the interior angles of a triangle?", Options: []string{"180°", "90°", "360°", "270°"}, CorrectOption: "A"},
	},
	"cat-nature": {
		{Text: "What process do plants use to convert sunlight into energy?", Options: []string{"Photosynthesis", "Respiration", "Digestion", "Transpiration"}, CorrectOption: "A"},
		{Text: "What is the largest land animal?", Options: []string{"African elephant", "Giraffe", "Hippopotamus", "Rhinoceros"}, CorrectOption: "A"},
		{Text: "Which animal is often called the king of the jungle?", Options: []string{"Lion", "Tiger", "Elephant", "Gorilla"}, CorrectOption: "A"},
		{Text: "What do bees collect from flowers?", Options: []string{"Nectar", "Water", "Sand", "Leaf juice"}, CorrectOption: "A"},
		{Text: "What is the chemical name of table salt?", Options: []string{"Sodium chloride", "Calcium carbonate", "Potassium nitrate", "Magnesium sulfate"}, CorrectOption: "A"},
		{Text: "Which layer of Earth do we live on?", Options: []string{"Crust", "Mantle", "Core", "Outer core"}, CorrectOption: "A"},
		{Text: "What is the largest ocean on Earth?", Options: []string{"Pacific Ocean", "Atlantic Ocean", "Indian Ocean", "Arctic Ocean"}, CorrectOption: "A"},
		{Text: "Which type of tree keeps its leaves all year?", Options: []string{"Evergreen", "Deciduous", "Birch", "Maple"}, CorrectOption: "A"},
		{Text: "Which gas do humans breathe in and use for survival?", Options: []string{"Oxygen", "Carbon dioxide", "Nitrogen", "Helium"}, CorrectOption: "A"},
		{Text: "Which bird is known for being able to mimic human speech?", Options: []string{"Parrot", "Owl", "Eagle", "Pigeon"}, CorrectOption: "A"},
	},
	"cat-technology": {
		{Text: "Which company created the iPhone?", Options: []string{"Apple", "Samsung", "Nokia", "Microsoft"}, CorrectOption: "A"},
		{Text: "What does CPU stand for?", Options: []string{"Central Processing Unit", "Computer Power Unit", "Central Program Utility", "Core Processing Utility"}, CorrectOption: "A"},
		{Text: "Which language is commonly used for web pages?", Options: []string{"HTML", "C++", "Go", "Rust"}, CorrectOption: "A"},
		{Text: "Which device stores data permanently?", Options: []string{"Hard drive", "RAM", "CPU", "Monitor"}, CorrectOption: "A"},
		{Text: "What does Wi-Fi stand for?", Options: []string{"Wireless Fidelity", "Wide Fiber", "Web Internet Frequency", "Wireless File Interface"}, CorrectOption: "A"},
		{Text: "Which programming language was created by Sun Microsystems?", Options: []string{"Java", "Python", "Ruby", "Swift"}, CorrectOption: "A"},
		{Text: "Which technology is used to send messages over the internet?", Options: []string{"Packets", "Bits of sound", "Cables only", "Bluetooth only"}, CorrectOption: "A"},
		{Text: "Which company created Windows?", Options: []string{"Microsoft", "Apple", "Google", "IBM"}, CorrectOption: "A"},
		{Text: "What is the main purpose of an operating system?", Options: []string{"Manage hardware and software resources", "Connect printers only", "Store electricity", "Display only text"}, CorrectOption: "A"},
		{Text: "Which network protocol is used for the web?", Options: []string{"HTTP", "FTP", "SMTP", "SSH"}, CorrectOption: "A"},
	},
	"cat-art": {
		{Text: "Which artist painted the Mona Lisa?", Options: []string{"Leonardo da Vinci", "Vincent van Gogh", "Pablo Picasso", "Claude Monet"}, CorrectOption: "A"},
		{Text: "What type of art is created with pigments on canvas?", Options: []string{"Painting", "Sculpture", "Photography", "Architecture"}, CorrectOption: "A"},
		{Text: "Which color is created by mixing blue and yellow?", Options: []string{"Green", "Orange", "Purple", "Red"}, CorrectOption: "A"},
		{Text: "What art form uses clay and shaping tools?", Options: []string{"Pottery", "Painting", "Printmaking", "Drawing"}, CorrectOption: "A"},
		{Text: "Which building is famous for its glass pyramid entrance?", Options: []string{"The Louvre", "Eiffel Tower", "Taj Mahal", "Big Ben"}, CorrectOption: "A"},
		{Text: "Which artist is known for Starry Night?", Options: []string{"Vincent van Gogh", "Leonardo da Vinci", "Michelangelo", "Salvador Dalí"}, CorrectOption: "A"},
		{Text: "What is a sculpture?", Options: []string{"A three-dimensional artwork", "A flat drawing", "A photograph", "A poem"}, CorrectOption: "A"},
		{Text: "Which medium is used in charcoal drawing?", Options: []string{"Charcoal", "Oil", "Watercolor", "Marble"}, CorrectOption: "A"},
		{Text: "What kind of art is made by arranging colored pieces?", Options: []string{"Mosaic", "Woodcut", "Calligraphy", "Ceramics"}, CorrectOption: "A"},
		{Text: "Which art style uses bold, simplified forms and strong colors?", Options: []string{"Modern art", "Renaissance", "Baroque", "Classical"}, CorrectOption: "A"},
	},
	"cat-food": {
		{Text: "Which fruit is commonly used to make guacamole?", Options: []string{"Avocado", "Mango", "Berry", "Orange"}, CorrectOption: "A"},
		{Text: "Which grain is used to make bread?", Options: []string{"Wheat", "Rice", "Corn", "Barley"}, CorrectOption: "A"},
		{Text: "Which cuisine is known for sushi?", Options: []string{"Japanese", "Italian", "Mexican", "French"}, CorrectOption: "A"},
		{Text: "Which vitamin is found in citrus fruits?", Options: []string{"Vitamin C", "Vitamin D", "Vitamin A", "Vitamin K"}, CorrectOption: "A"},
		{Text: "What is the main ingredient in hummus?", Options: []string{"Chickpeas", "Tomatoes", "Beans", "Rice"}, CorrectOption: "A"},
		{Text: "Which cooking method uses dry heat in an oven?", Options: []string{"Baking", "Boiling", "Steaming", "Fermenting"}, CorrectOption: "A"},
		{Text: "Which fruit is known for being red and sweet?", Options: []string{"Strawberry", "Lemon", "Pear", "Banana"}, CorrectOption: "A"},
		{Text: "What is the main ingredient in cheese?", Options: []string{"Milk", "Flour", "Eggs", "Rice"}, CorrectOption: "A"},
		{Text: "Which vegetable is commonly used in tomato soup?", Options: []string{"Tomato", "Potato", "Carrot", "Cabbage"}, CorrectOption: "A"},
		{Text: "Which spice is commonly used to flavor pizza?", Options: []string{"Oregano", "Cinnamon", "Nutmeg", "Cardamom"}, CorrectOption: "A"},
	},
	"cat-animals": {
		{Text: "Which mammal is famous for being able to fly?", Options: []string{"Bat", "Mouse", "Squirrel", "Rabbit"}, CorrectOption: "A"},
		{Text: "Which animal is the largest on Earth?", Options: []string{"Blue whale", "Elephant", "Giraffe", "Hippopotamus"}, CorrectOption: "A"},
		{Text: "Which animal is known for carrying its house on its back?", Options: []string{"Snail", "Turtle", "Crab", "Lobster"}, CorrectOption: "A"},
		{Text: "Which animal is the fastest land mammal?", Options: []string{"Cheetah", "Horse", "Lion", "Gazelle"}, CorrectOption: "A"},
		{Text: "Which animal is famous for hibernating in winter?", Options: []string{"Bear", "Rabbit", "Fox", "Wolf"}, CorrectOption: "A"},
		{Text: "Which marine animal uses a shell for protection?", Options: []string{"Turtle", "Dolphin", "Shark", "Seal"}, CorrectOption: "A"},
		{Text: "Which animal is known for its black-and-white stripes?", Options: []string{"Zebra", "Tiger", "Leopard", "Panda"}, CorrectOption: "A"},
		{Text: "Which animal is a symbol of wisdom in many cultures?", Options: []string{"Owl", "Eagle", "Falcon", "Sparrow"}, CorrectOption: "A"},
		{Text: "Which animal is often called the king of the jungle?", Options: []string{"Lion", "Elephant", "Gorilla", "Tiger"}, CorrectOption: "A"},
		{Text: "Which large cat is famous for its orange fur and black stripes?", Options: []string{"Tiger", "Leopard", "Cheetah", "Lynx"}, CorrectOption: "A"},
	},
	"cat-space": {
		{Text: "Which planet is closest to the Sun?", Options: []string{"Mercury", "Venus", "Mars", "Earth"}, CorrectOption: "A"},
		{Text: "What is the name of Earth's natural satellite?", Options: []string{"Moon", "Titan", "Europa", "Phobos"}, CorrectOption: "A"},
		{Text: "Which planet is famous for its rings?", Options: []string{"Saturn", "Jupiter", "Neptune", "Mars"}, CorrectOption: "A"},
		{Text: "What force keeps the planets in orbit around the Sun?", Options: []string{"Gravity", "Magnetism", "Friction", "Radiation"}, CorrectOption: "A"},
		{Text: "Which galaxy contains our Solar System?", Options: []string{"Milky Way", "Andromeda", "Whirlpool", "Sombrero"}, CorrectOption: "A"},
		{Text: "What is a comet mostly made of?", Options: []string{"Ice and dust", "Metal and lava", "Glass and gas", "Rock and water"}, CorrectOption: "A"},
		{Text: "Which planet is known as the Red Planet?", Options: []string{"Mars", "Venus", "Mercury", "Jupiter"}, CorrectOption: "A"},
		{Text: "Which spacecraft landed humans on the Moon?", Options: []string{"Apollo", "Voyager", "Gemini", "Sputnik"}, CorrectOption: "A"},
		{Text: "What is the name of the star at the center of our Solar System?", Options: []string{"Sun", "Proxima Centauri", "Sirius", "Betelgeuse"}, CorrectOption: "A"},
		{Text: "Which planet is the largest in the Solar System?", Options: []string{"Jupiter", "Saturn", "Earth", "Neptune"}, CorrectOption: "A"},
	},
	"cat-health": {
		{Text: "Which activity improves heart health?", Options: []string{"Regular exercise", "Sleeping all day", "Skipping meals", "Smoking"}, CorrectOption: "A"},
		{Text: "What is the main purpose of sleep?", Options: []string{"Rest and recovery", "Digesting food", "Producing blood", "Cooling the body"}, CorrectOption: "A"},
		{Text: "Which food group is essential for strong bones?", Options: []string{"Dairy and calcium-rich foods", "Candy", "Soda", "Fast food"}, CorrectOption: "A"},
		{Text: "Which habit is most harmful to lungs?", Options: []string{"Smoking", "Drinking water", "Walking", "Stretching"}, CorrectOption: "A"},
		{Text: "What is a healthy way to manage stress?", Options: []string{"Exercise and rest", "Skipping sleep", "Smoking", "Eating junk food"}, CorrectOption: "A"},
		{Text: "Which organ is most directly responsible for filtering blood?", Options: []string{"Kidneys", "Stomach", "Lungs", "Liver"}, CorrectOption: "A"},
		{Text: "What should you do before exercising?", Options: []string{"Warm up", "Drink soda", "Skip hydration", "Lie down"}, CorrectOption: "A"},
		{Text: "Which nutrient helps build muscles?", Options: []string{"Protein", "Sugar", "Salt", "Caffeine"}, CorrectOption: "A"},
		{Text: "How can you keep your immune system strong?", Options: []string{"Eat well and sleep enough", "Eat only sweets", "Avoid all exercise", "Stay up late"}, CorrectOption: "A"},
		{Text: "Which is the best way to stay hydrated?", Options: []string{"Drink water regularly", "Only drink soda", "Skip fluids", "Drink coffee only"}, CorrectOption: "A"},
	},
	"cat-language": {
		{Text: "What is the main purpose of grammar?", Options: []string{"To structure sentences clearly", "To add colors", "To count numbers", "To change spelling randomly"}, CorrectOption: "A"},
		{Text: "Which word is a noun?", Options: []string{"Table", "Quickly", "Beautifully", "Run"}, CorrectOption: "A"},
		{Text: "What is the opposite of 'begin'?", Options: []string{"End", "Open", "Start", "Finish"}, CorrectOption: "A"},
		{Text: "Which sentence is written correctly?", Options: []string{"She walks to school every day.", "She walk to school every day.", "She walking to school every day.", "She walked to school every days."}, CorrectOption: "A"},
		{Text: "What is a synonym for 'happy'?", Options: []string{"Joyful", "Sad", "Angry", "Silent"}, CorrectOption: "A"},
		{Text: "What is an antonym of 'hot'?", Options: []string{"Cold", "Warm", "Bright", "Fast"}, CorrectOption: "A"},
		{Text: "Which part of speech describes a verb, adjective, or other adverb?", Options: []string{"Adverb", "Noun", "Pronoun", "Preposition"}, CorrectOption: "A"},
		{Text: "What is a paragraph?", Options: []string{"A group of related sentences", "A single word", "A punctuation mark", "A number"}, CorrectOption: "A"},
		{Text: "Which punctuation mark ends a statement?", Options: []string{"Period", "Comma", "Question mark", "Colon"}, CorrectOption: "A"},
		{Text: "What does a dictionary help you do?", Options: []string{"Find word meanings", "Measure height", "Cook food", "Build a chair"}, CorrectOption: "A"},
	},
	"cat-psychology": {
		{Text: "What is memory?", Options: []string{"The ability to store and recall information", "A feeling of hunger", "A type of plant", "A kind of light"}, CorrectOption: "A"},
		{Text: "Which emotion is often linked with fear?", Options: []string{"Anxiety", "Joy", "Calm", "Excitement"}, CorrectOption: "A"},
		{Text: "What is empathy?", Options: []string{"Understanding and sharing another's feelings", "Ignoring others", "Feeling angry", "Being silent"}, CorrectOption: "A"},
		{Text: "What is a common goal of mindfulness?", Options: []string{"To stay present and focused", "To ignore thoughts", "To sleep constantly", "To avoid all emotions"}, CorrectOption: "A"},
		{Text: "Which behavior is usually associated with good mental health?", Options: []string{"Healthy coping skills", "Avoiding all responsibilities", "Constant panic", "Feeling no emotion"}, CorrectOption: "A"},
		{Text: "What does perception mean?", Options: []string{"How we interpret sensory information", "A type of memory", "A social norm", "A mood"}, CorrectOption: "A"},
		{Text: "Which factor often influences decision-making?", Options: []string{"Emotions and experience", "Only random chance", "Only weather", "Only calendar dates"}, CorrectOption: "A"},
		{Text: "What is motivation?", Options: []string{"The drive that causes action", "A type of food", "A shape of box", "A speed limit"}, CorrectOption: "A"},
		{Text: "What does stress usually do to focus?", Options: []string{"It can reduce it", "It always improves it", "It makes you invisible", "It removes all thoughts"}, CorrectOption: "A"},
		{Text: "What is a habit?", Options: []string{"A repeated behavior", "A random idea", "A type of fruit", "A physical object"}, CorrectOption: "A"},
	},
	"cat-religion": {
		{Text: "Which book is sacred in Christianity?", Options: []string{"The Bible", "The Torah", "The Quran", "The Vedas"}, CorrectOption: "A"},
		{Text: "What is the holy city of Islam?", Options: []string{"Mecca", "Jerusalem", "Rome", "Cairo"}, CorrectOption: "A"},
		{Text: "Which religion follows the Five Pillars?", Options: []string{"Islam", "Christianity", "Hinduism", "Judaism"}, CorrectOption: "A"},
		{Text: "What is the sacred text of Judaism?", Options: []string{"Torah", "Bible", "Quran", "Tripitaka"}, CorrectOption: "A"},
		{Text: "Which religion is associated with Karma and Dharma?", Options: []string{"Hinduism", "Islam", "Buddhism", "Christianity"}, CorrectOption: "A"},
		{Text: "What is a temple?", Options: []string{"A place of worship", "A school", "A farm", "A market"}, CorrectOption: "A"},
		{Text: "Which faith is centered on the teachings of Buddha?", Options: []string{"Buddhism", "Christianity", "Judaism", "Shinto"}, CorrectOption: "A"},
		{Text: "Which city is holy to Christians, Jews, and Muslims?", Options: []string{"Jerusalem", "Cairo", "Mecca", "Rome"}, CorrectOption: "A"},
		{Text: "Which religion is based on the life and teachings of Jesus?", Options: []string{"Christianity", "Islam", "Hinduism", "Sikhism"}, CorrectOption: "A"},
		{Text: "What is a shrine?", Options: []string{"A holy or sacred place", "A library", "A sports stadium", "A hospital"}, CorrectOption: "A"},
	},
	"cat-business": {
		{Text: "What is the main goal of a business?", Options: []string{"Create value and earn profit", "Only spend money", "Avoid customers", "Ignore innovation"}, CorrectOption: "A"},
		{Text: "What is a budget?", Options: []string{"A plan for income and spending", "A type of machine", "A legal rule", "A product design"}, CorrectOption: "A"},
		{Text: "Which term means the amount of money a company makes?", Options: []string{"Revenue", "Debt", "Inventory", "Expense"}, CorrectOption: "A"},
		{Text: "What is marketing?", Options: []string{"Promoting and selling products", "Only counting stock", "Building roads", "Changing laws"}, CorrectOption: "A"},
		{Text: "What is branding?", Options: []string{"Creating a recognizable identity", "Making a legal contract", "Writing code", "Shipping products"}, CorrectOption: "A"},
		{Text: "Which document describes a business idea and plan?", Options: []string{"Business plan", "Receipt", "Invoice", "Bank statement"}, CorrectOption: "A"},
		{Text: "What is a stakeholder?", Options: []string{"A person affected by a business decision", "A machine part", "A marketing campaign", "A product feature"}, CorrectOption: "A"},
		{Text: "What does ROI measure?", Options: []string{"Return on investment", "Running operating inventory", "Revenue of interest", "Rate of income"}, CorrectOption: "A"},
		{Text: "What is customer service about?", Options: []string{"Helping customers before and after a purchase", "Ignoring customer issues", "Lowering product prices only", "Producing raw materials"}, CorrectOption: "A"},
		{Text: "Which financial statement shows profit and loss?", Options: []string{"Income statement", "Balance sheet", "Cash register", "Payroll form"}, CorrectOption: "A"},
	},
	"cat-gaming": {
		{Text: "What is the objective of most racing games?", Options: []string{"Finish the track first", "Solve puzzles", "Collect plants", "Write code"}, CorrectOption: "A"},
		{Text: "What does FPS stand for in gaming?", Options: []string{"First-person shooter", "Fast play score", "Full power system", "Frame priority setting"}, CorrectOption: "A"},
		{Text: "Which action often wins a strategy game?", Options: []string{"Planning ahead", "Spamming random clicks", "Ignoring resources", "Skipping turns"}, CorrectOption: "A"},
		{Text: "What is loot in an RPG?", Options: []string{"Rewards such as gear or items", "A type of map", "A weapon sound", "A difficulty level"}, CorrectOption: "A"},
		{Text: "What is a game controller used for?", Options: []string{"Input control in games", "Listening to music", "Reading books", "Making coffee"}, CorrectOption: "A"},
		{Text: "What does multiplayer mean?", Options: []string{"More than one player participates", "A single-player mode", "A high-resolution screen", "A new game engine"}, CorrectOption: "A"},
		{Text: "Which genre focuses on solving logic puzzles?", Options: []string{"Puzzle games", "Racing games", "Sports games", "Platformers"}, CorrectOption: "A"},
		{Text: "What is a save point?", Options: []string{"A point where progress is stored", "A cheat code", "A game title", "A map icon"}, CorrectOption: "A"},
		{Text: "Which feature usually improves performance in a game?", Options: []string{"Better graphics settings", "Slower loading times", "More random crashes", "No sound"}, CorrectOption: "A"},
		{Text: "What is a game mechanic?", Options: []string{"A rule or system that defines gameplay", "A player name", "A hard drive", "A photo filter"}, CorrectOption: "A"},
	},
	"cat-cars": {
		{Text: "What is the main purpose of a car's engine?", Options: []string{"Generate power to move the vehicle", "Store fuel only", "Play music", "Control traffic lights"}, CorrectOption: "A"},
		{Text: "What does ABS help a car do?", Options: []string{"Prevent wheels from locking during braking", "Increase tire pressure", "Turn the steering wheel", "Cool the cabin"}, CorrectOption: "A"},
		{Text: "Which fuel is commonly used in most gasoline cars?", Options: []string{"Petrol", "Diesel only", "Coal", "Natural gas"}, CorrectOption: "A"},
		{Text: "What is a transmission?", Options: []string{"A system that transfers power to the wheels", "A car radio", "A seat cushion", "A steering wheel"}, CorrectOption: "A"},
		{Text: "What does a tachometer measure?", Options: []string{"Engine revolutions per minute", "Fuel level", "Temperature outside", "Seat position"}, CorrectOption: "A"},
		{Text: "Which component is responsible for steering?", Options: []string{"Steering wheel and linkage", "Dashboard", "Headlights", "Roof rack"}, CorrectOption: "A"},
		{Text: "What is the purpose of a car battery?", Options: []string{"Start the engine and power electronics", "Move the wheels directly", "Measure tire pressure", "Cool the brakes"}, CorrectOption: "A"},
		{Text: "Which type of vehicle is designed mainly for off-road driving?", Options: []string{"SUV", "Sedan", "Hatchback", "Coupe"}, CorrectOption: "A"},
		{Text: "What does MPG measure?", Options: []string{"Miles per gallon of fuel", "Miles per hour", "Milligrams per gram", "Meters per gallon"}, CorrectOption: "A"},
		{Text: "Which part of the car is responsible for braking?", Options: []string{"Brake pads", "Headlights", "Roof", "Spoiler"}, CorrectOption: "A"},
	},
	"cat-astronomy": {
		{Text: "What is a galaxy?", Options: []string{"A vast collection of stars, gas, and dust", "A single star", "A planet", "A moon"}, CorrectOption: "A"},
		{Text: "Which planet is known for its prominent rings?", Options: []string{"Saturn", "Mars", "Mercury", "Venus"}, CorrectOption: "A"},
		{Text: "What is a comet mainly composed of?", Options: []string{"Ice, dust, and rock", "Hot metal", "Water vapor only", "Copper and steel"}, CorrectOption: "A"},
		{Text: "What do astronomers use to study distant stars?", Options: []string{"Telescopes", "Microscopes", "Magnets", "Compasses"}, CorrectOption: "A"},
		{Text: "Which planet is closest to the Sun?", Options: []string{"Mercury", "Venus", "Earth", "Mars"}, CorrectOption: "A"},
		{Text: "What is a lunar eclipse?", Options: []string{"Earth passes between the Sun and the Moon", "The Moon passes between Earth and Sun", "A planet collides with the Moon", "The Sun disappears"}, CorrectOption: "A"},
		{Text: "What is a black hole?", Options: []string{"A region of space with extremely strong gravity", "A giant star", "A planet with no atmosphere", "An empty nebula"}, CorrectOption: "A"},
		{Text: "Which planet has the Great Red Spot?", Options: []string{"Jupiter", "Mars", "Neptune", "Venus"}, CorrectOption: "A"},
		{Text: "What is the name of Earth's star?", Options: []string{"Sun", "Sirius", "Polaris", "Vega"}, CorrectOption: "A"},
		{Text: "What does an orbit describe?", Options: []string{"The path an object follows around another object", "A burst of light", "A weather pattern", "A mountain shape"}, CorrectOption: "A"},
	},
	"cat-philosophy": {
		{Text: "What is ethics mainly concerned with?", Options: []string{"Right and wrong behavior", "Chemical reactions", "Weather patterns", "Historic dates"}, CorrectOption: "A"},
		{Text: "Who is known for the statement I think, therefore I am?", Options: []string{"Descartes", "Plato", "Aristotle", "Socrates"}, CorrectOption: "A"},
		{Text: "What is logic?", Options: []string{"Reasoning according to valid rules", "A physical experiment", "A shape of building", "A form of medicine"}, CorrectOption: "A"},
		{Text: "What does metaphysics study?", Options: []string{"The nature of reality", "Numbers only", "Weather systems", "Human anatomy"}, CorrectOption: "A"},
		{Text: "What is truth?", Options: []string{"A concept about what is correct and real", "A kind of stone", "A type of music", "A form of light"}, CorrectOption: "A"},
		{Text: "What is an argument in philosophy?", Options: []string{"A set of reasons supporting a conclusion", "A loud disagreement", "A type of building", "A political event"}, CorrectOption: "A"},
		{Text: "Which philosopher taught in Athens and questioned many assumptions?", Options: []string{"Socrates", "Descartes", "Hume", "Kant"}, CorrectOption: "A"},
		{Text: "What is free will?", Options: []string{"The ability to choose one's own actions", "A law of physics", "An animal instinct", "A mood state"}, CorrectOption: "A"},
		{Text: "What is epistemology?", Options: []string{"The study of knowledge", "The study of numbers", "The study of cells", "The study of language"}, CorrectOption: "A"},
		{Text: "What does existentialism focus on?", Options: []string{"Human existence and personal choice", "Mathematical proofs", "Plant biology", "Mechanical design"}, CorrectOption: "A"},
	},
	"cat-design": {
		{Text: "What is user experience design mainly concerned with?", Options: []string{"How a person feels while using a product", "Only making it colorful", "Writing long documents", "Printing on paper"}, CorrectOption: "A"},
		{Text: "What is a wireframe?", Options: []string{"A simple layout blueprint", "A code compiler", "A database query", "A marketing report"}, CorrectOption: "A"},
		{Text: "Why is contrast important in design?", Options: []string{"It helps elements stand out and improve readability", "It reduces creativity", "It always slows users down", "It removes color"}, CorrectOption: "A"},
		{Text: "What does UI stand for?", Options: []string{"User Interface", "Universal Input", "User Information", "Unified Internet"}, CorrectOption: "A"},
		{Text: "What is hierarchy in design?", Options: []string{"The arrangement of elements by importance", "A set of colors", "A storage method", "A layout grid only"}, CorrectOption: "A"},
		{Text: "Which principle helps users understand where to look first?", Options: []string{"Visual hierarchy", "Random placement", "Overcrowding", "No spacing"}, CorrectOption: "A"},
		{Text: "What is a prototype?", Options: []string{"An early model of a design", "A final product only", "A marketing slogan", "A code library"}, CorrectOption: "A"},
		{Text: "Why is whitespace useful in design?", Options: []string{"It improves readability and focus", "It makes text invisible", "It removes meaning", "It adds noise"}, CorrectOption: "A"},
		{Text: "What is accessibility in design?", Options: []string{"Designing for people with different needs and abilities", "Making content shorter only", "Using only one color", "Ignoring users"}, CorrectOption: "A"},
		{Text: "Which design tool is commonly used for interfaces?", Options: []string{"Figma", "Excel", "Photoshop without design tools", "A command line"}, CorrectOption: "A"},
	},
}

func generateSeedQuestion(category models.Category, index, questionNumber int) models.Question {
	bank := categoryQuestionBanks[category.ID]
	if len(bank) == 0 {
		bank = []seedQuestion{{
			Text:          fmt.Sprintf("Which answer best matches %s?", category.Name),
			Options:       []string{"Study", "Guessing", "Confusion", "Delay"},
			CorrectOption: "A",
		}}
	}

	base := bank[index%len(bank)]
	shift := (index + 1) % len(base.Options)
	letters := []string{"A", "B", "C", "D"}
	rotated := make([]string, len(base.Options))
	for i := range base.Options {
		rotated[(i+shift)%len(base.Options)] = base.Options[i]
	}
	correctPos := (indexOfLetter(base.CorrectOption) + shift) % len(letters)
	correctOption := letters[correctPos]

	return models.Question{
		ID:            fmt.Sprintf("q-%d", questionNumber),
		CategoryID:    category.ID,
		Text:          base.Text,
		OptionA:       rotated[0],
		OptionB:       rotated[1],
		OptionC:       rotated[2],
		OptionD:       rotated[3],
		CorrectOption: correctOption,
	}
}

func indexOfLetter(letter string) int {
	letters := []string{"A", "B", "C", "D"}
	for i, value := range letters {
		if value == letter {
			return i
		}
	}
	return 0
}

func (r *QuestionRepository) ListCategories() ([]models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name FROM categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (r *QuestionRepository) ListQuestionsByCategory(categoryID string) ([]models.Question, error) {
	rows, err := r.db.Query(`SELECT id, category_id, text, option_a, option_b, option_c, option_d, correct_option FROM questions WHERE category_id = ? ORDER BY id`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var question models.Question
		if err := rows.Scan(&question.ID, &question.CategoryID, &question.Text, &question.OptionA, &question.OptionB, &question.OptionC, &question.OptionD, &question.CorrectOption); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, rows.Err()
}

func (r *QuestionRepository) GetQuestionByID(id string) (*models.Question, error) {
	var question models.Question
	err := r.db.QueryRow(`SELECT id, category_id, text, option_a, option_b, option_c, option_d, correct_option FROM questions WHERE id = ?`, id).Scan(&question.ID, &question.CategoryID, &question.Text, &question.OptionA, &question.OptionB, &question.OptionC, &question.OptionD, &question.CorrectOption)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &question, nil
}
