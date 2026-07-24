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
	categories := []models.Category{
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
	}

	for _, category := range categories {
		if _, err := r.db.Exec(`INSERT OR IGNORE INTO categories (id, name) VALUES (?, ?)`, category.ID, category.Name); err != nil {
			return fmt.Errorf("insert category %s: %w", category.ID, err)
		}
	}

	questions := []models.Question{
		{ID: "q-1", CategoryID: "cat-programming", Text: "Which keyword starts a Go function declaration?", OptionA: "func", OptionB: "def", OptionC: "class", OptionD: "method", CorrectOption: "A"},
		{ID: "q-2", CategoryID: "cat-programming", Text: "What does Go use for concurrent execution?", OptionA: "threads", OptionB: "goroutines", OptionC: "processes", OptionD: "coroutines", CorrectOption: "B"},
		{ID: "q-3", CategoryID: "cat-programming", Text: "What type does Go use for a sequence of bytes?", OptionA: "string", OptionB: "[]byte", OptionC: "byte[]", OptionD: "bytes", CorrectOption: "B"},
		{ID: "q-4", CategoryID: "cat-programming", Text: "Which Go keyword is used to import packages?", OptionA: "package", OptionB: "import", OptionC: "include", OptionD: "using", CorrectOption: "B"},
		{ID: "q-5", CategoryID: "cat-programming", Text: "Which operator is used to create a pointer in Go?", OptionA: "&", OptionB: "*", OptionC: "^", OptionD: "%", CorrectOption: "A"},
		{ID: "q-6", CategoryID: "cat-programming", Text: "What is the zero value of a bool in Go?", OptionA: "true", OptionB: "false", OptionC: "nil", OptionD: "0", CorrectOption: "B"},
		{ID: "q-7", CategoryID: "cat-programming", Text: "Which keyword declares a new package in Go?", OptionA: "package", OptionB: "module", OptionC: "import", OptionD: "namespace", CorrectOption: "A"},
		{ID: "q-8", CategoryID: "cat-programming", Text: "Which built-in function allocates memory for a slice?", OptionA: "new", OptionB: "make", OptionC: "alloc", OptionD: "slice", CorrectOption: "B"},
		{ID: "q-9", CategoryID: "cat-programming", Text: "What is the Go keyword used to handle panics?", OptionA: "try", OptionB: "catch", OptionC: "recover", OptionD: "finally", CorrectOption: "C"},
		{ID: "q-10", CategoryID: "cat-programming", Text: "Which type represents an integer of unspecified size?", OptionA: "int", OptionB: "uint8", OptionC: "intptr", OptionD: "bigint", CorrectOption: "A"},
		{ID: "q-11", CategoryID: "cat-history", Text: "Who was the first President of the United States?", OptionA: "Thomas Jefferson", OptionB: "George Washington", OptionC: "Abraham Lincoln", OptionD: "John Adams", CorrectOption: "B"},
		{ID: "q-12", CategoryID: "cat-history", Text: "In what year did the Berlin Wall fall?", OptionA: "1987", OptionB: "1988", OptionC: "1989", OptionD: "1990", CorrectOption: "C"},
		{ID: "q-13", CategoryID: "cat-history", Text: "Which empire was led by Genghis Khan?", OptionA: "Roman Empire", OptionB: "Mongol Empire", OptionC: "Ottoman Empire", OptionD: "Persian Empire", CorrectOption: "B"},
		{ID: "q-14", CategoryID: "cat-history", Text: "The Magna Carta was signed in which year?", OptionA: "1215", OptionB: "1315", OptionC: "1415", OptionD: "1515", CorrectOption: "A"},
		{ID: "q-15", CategoryID: "cat-history", Text: "Which civilization built the pyramids at Giza?", OptionA: "Romans", OptionB: "Greeks", OptionC: "Egyptians", OptionD: "Persians", CorrectOption: "C"},
		{ID: "q-16", CategoryID: "cat-history", Text: "Which war was fought between the North and the South in the United States?", OptionA: "World War I", OptionB: "American Revolution", OptionC: "Civil War", OptionD: "War of 1812", CorrectOption: "C"},
		{ID: "q-17", CategoryID: "cat-history", Text: "Who discovered the Americas in 1492?", OptionA: "Christopher Columbus", OptionB: "Marco Polo", OptionC: "Ferdinand Magellan", OptionD: "Amerigo Vespucci", CorrectOption: "A"},
		{ID: "q-18", CategoryID: "cat-history", Text: "Which ancient city was buried by the eruption of Mount Vesuvius?", OptionA: "Rome", OptionB: "Pompeii", OptionC: "Athens", OptionD: "Carthage", CorrectOption: "B"},
		{ID: "q-19", CategoryID: "cat-history", Text: "Which document began the United States government in 1789?", OptionA: "The Declaration of Independence", OptionB: "The Articles of Confederation", OptionC: "The Constitution", OptionD: "The Bill of Rights", CorrectOption: "C"},
		{ID: "q-20", CategoryID: "cat-history", Text: "Which king signed the Magna Carta?", OptionA: "Henry VIII", OptionB: "John", OptionC: "Richard III", OptionD: "Edward I", CorrectOption: "B"},
		{ID: "q-21", CategoryID: "cat-science", Text: "What planet is known as the Red Planet?", OptionA: "Venus", OptionB: "Mars", OptionC: "Mercury", OptionD: "Jupiter", CorrectOption: "B"},
		{ID: "q-22", CategoryID: "cat-science", Text: "What is the chemical symbol for water?", OptionA: "CO2", OptionB: "O2", OptionC: "H2O", OptionD: "HO2", CorrectOption: "C"},
		{ID: "q-23", CategoryID: "cat-science", Text: "What force keeps planets in orbit around the sun?", OptionA: "Magnetism", OptionB: "Gravity", OptionC: "Friction", OptionD: "Electricity", CorrectOption: "B"},
		{ID: "q-24", CategoryID: "cat-science", Text: "What part of the cell contains genetic material?", OptionA: "Cell membrane", OptionB: "Mitochondria", OptionC: "Nucleus", OptionD: "Ribosome", CorrectOption: "C"},
		{ID: "q-25", CategoryID: "cat-science", Text: "What gas do plants absorb from the atmosphere?", OptionA: "Oxygen", OptionB: "Nitrogen", OptionC: "Carbon dioxide", OptionD: "Helium", CorrectOption: "C"},
		{ID: "q-26", CategoryID: "cat-science", Text: "What is the center of an atom called?", OptionA: "Electron", OptionB: "Proton", OptionC: "Nucleus", OptionD: "Neutron", CorrectOption: "C"},
		{ID: "q-27", CategoryID: "cat-science", Text: "What is the boiling point of water at sea level?", OptionA: "90°C", OptionB: "100°C", OptionC: "110°C", OptionD: "120°C", CorrectOption: "B"},
		{ID: "q-28", CategoryID: "cat-science", Text: "Which organ pumps blood through the body?", OptionA: "Lungs", OptionB: "Heart", OptionC: "Liver", OptionD: "Kidney", CorrectOption: "B"},
		{ID: "q-29", CategoryID: "cat-science", Text: "What is the chemical symbol for gold?", OptionA: "Au", OptionB: "Ag", OptionC: "Gd", OptionD: "Go", CorrectOption: "A"},
		{ID: "q-30", CategoryID: "cat-science", Text: "What is the most abundant gas in Earth's atmosphere?", OptionA: "Oxygen", OptionB: "Nitrogen", OptionC: "Carbon dioxide", OptionD: "Hydrogen", CorrectOption: "B"},
		{ID: "q-31", CategoryID: "cat-sports", Text: "How many players are on a soccer team on the field?", OptionA: "9", OptionB: "10", OptionC: "11", OptionD: "12", CorrectOption: "C"},
		{ID: "q-32", CategoryID: "cat-sports", Text: "In which sport is the term 'home run' used?", OptionA: "Baseball", OptionB: "Basketball", OptionC: "Soccer", OptionD: "Tennis", CorrectOption: "A"},
		{ID: "q-33", CategoryID: "cat-sports", Text: "How long is a standard NBA game in minutes?", OptionA: "32", OptionB: "36", OptionC: "40", OptionD: "48", CorrectOption: "D"},
		{ID: "q-34", CategoryID: "cat-sports", Text: "Which sport uses a shuttlecock?", OptionA: "Tennis", OptionB: "Badminton", OptionC: "Squash", OptionD: "Ping pong", CorrectOption: "B"},
		{ID: "q-35", CategoryID: "cat-sports", Text: "How many rings are on the Olympic flag?", OptionA: "4", OptionB: "5", OptionC: "6", OptionD: "7", CorrectOption: "B"},
		{ID: "q-36", CategoryID: "cat-sports", Text: "What sport uses a net, ball, and court and has five players per team?", OptionA: "Volleyball", OptionB: "Basketball", OptionC: "Handball", OptionD: "Lacrosse", CorrectOption: "B"},
		{ID: "q-37", CategoryID: "cat-sports", Text: "Which country won the FIFA World Cup in 2018?", OptionA: "Germany", OptionB: "Brazil", OptionC: "France", OptionD: "Argentina", CorrectOption: "C"},
		{ID: "q-38", CategoryID: "cat-sports", Text: "In tennis, what is a score of zero called?", OptionA: "Love", OptionB: "None", OptionC: "Zero", OptionD: "Nil", CorrectOption: "A"},
		{ID: "q-39", CategoryID: "cat-sports", Text: "How many minutes are in a standard soccer match?", OptionA: "60", OptionB: "80", OptionC: "90", OptionD: "120", CorrectOption: "C"},
		{ID: "q-40", CategoryID: "cat-sports", Text: "What is the term for a score of one in golf?", OptionA: "Birdie", OptionB: "Eagle", OptionC: "Albatross", OptionD: "Hole-in-one", CorrectOption: "D"},
		{ID: "q-41", CategoryID: "cat-geography", Text: "What is the largest continent by land area?", OptionA: "Africa", OptionB: "Asia", OptionC: "Europe", OptionD: "Antarctica", CorrectOption: "B"},
		{ID: "q-42", CategoryID: "cat-geography", Text: "What is the capital city of France?", OptionA: "Berlin", OptionB: "Madrid", OptionC: "Paris", OptionD: "Rome", CorrectOption: "C"},
		{ID: "q-43", CategoryID: "cat-geography", Text: "Which river is the longest in the world?", OptionA: "Amazon", OptionB: "Nile", OptionC: "Yangtze", OptionD: "Mississippi", CorrectOption: "B"},
		{ID: "q-44", CategoryID: "cat-geography", Text: "Which country has the largest population?", OptionA: "United States", OptionB: "India", OptionC: "China", OptionD: "Russia", CorrectOption: "C"},
		{ID: "q-45", CategoryID: "cat-geography", Text: "What ocean is east of the United States?", OptionA: "Pacific", OptionB: "Atlantic", OptionC: "Indian", OptionD: "Arctic", CorrectOption: "B"},
		{ID: "q-46", CategoryID: "cat-geography", Text: "Which mountain range includes Mount Everest?", OptionA: "Andes", OptionB: "Rockies", OptionC: "Himalayas", OptionD: "Alps", CorrectOption: "C"},
		{ID: "q-47", CategoryID: "cat-geography", Text: "What is the capital of Japan?", OptionA: "Seoul", OptionB: "Tokyo", OptionC: "Beijing", OptionD: "Bangkok", CorrectOption: "B"},
		{ID: "q-48", CategoryID: "cat-geography", Text: "Which desert is the largest hot desert on Earth?", OptionA: "Sahara", OptionB: "Gobi", OptionC: "Kalahari", OptionD: "Mojave", CorrectOption: "A"},
		{ID: "q-49", CategoryID: "cat-geography", Text: "Which US state is known as the 'Sunshine State'?", OptionA: "California", OptionB: "Florida", OptionC: "Texas", OptionD: "Arizona", CorrectOption: "B"},
		{ID: "q-50", CategoryID: "cat-geography", Text: "What is the smallest country in the world by land area?", OptionA: "Monaco", OptionB: "Nauru", OptionC: "Vatican City", OptionD: "San Marino", CorrectOption: "C"},
		{ID: "q-51", CategoryID: "cat-literature", Text: "Who wrote 'Romeo and Juliet'?", OptionA: "William Shakespeare", OptionB: "Charles Dickens", OptionC: "Jane Austen", OptionD: "Mark Twain", CorrectOption: "A"},
		{ID: "q-52", CategoryID: "cat-literature", Text: "What is the title of George Orwell's novel about a dystopian future?", OptionA: "Brave New World", OptionB: "1984", OptionC: "Fahrenheit 451", OptionD: "The Giver", CorrectOption: "B"},
		{ID: "q-53", CategoryID: "cat-literature", Text: "Who is the author of 'Pride and Prejudice'?", OptionA: "Emily Bronte", OptionB: "Jane Austen", OptionC: "Charlotte Bronte", OptionD: "Louisa May Alcott", CorrectOption: "B"},
		{ID: "q-54", CategoryID: "cat-literature", Text: "Which novel begins with the line 'Call me Ishmael'?", OptionA: "Moby-Dick", OptionB: "The Odyssey", OptionC: "Treasure Island", OptionD: "The Iliad", CorrectOption: "A"},
		{ID: "q-55", CategoryID: "cat-literature", Text: "What is the name of the wizard school in the Harry Potter series?", OptionA: "Hogwarts", OptionB: "Beauxbatons", OptionC: "Durmstrang", OptionD: "Ilvermorny", CorrectOption: "A"},
		{ID: "q-56", CategoryID: "cat-literature", Text: "Who wrote 'The Hobbit'?", OptionA: "C.S. Lewis", OptionB: "J.R.R. Tolkien", OptionC: "J.K. Rowling", OptionD: "George R.R. Martin", CorrectOption: "B"},
		{ID: "q-57", CategoryID: "cat-literature", Text: "What is the name of Sherlock Holmes' assistant?", OptionA: "Dr. Watson", OptionB: "Inspector Lestrade", OptionC: "Mrs. Hudson", OptionD: "Professor Moriarty", CorrectOption: "A"},
		{ID: "q-58", CategoryID: "cat-literature", Text: "Which novel features the character Atticus Finch?", OptionA: "To Kill a Mockingbird", OptionB: "The Great Gatsby", OptionC: "Of Mice and Men", OptionD: "1984", CorrectOption: "A"},
		{ID: "q-59", CategoryID: "cat-literature", Text: "What type of poem has 14 lines?", OptionA: "Haiku", OptionB: "Sonnet", OptionC: "Limerick", OptionD: "Epic", CorrectOption: "B"},
		{ID: "q-60", CategoryID: "cat-literature", Text: "Which book begins with 'It was the best of times, it was the worst of times'?", OptionA: "A Tale of Two Cities", OptionB: "Great Expectations", OptionC: "Oliver Twist", OptionD: "David Copperfield", CorrectOption: "A"},
		{ID: "q-61", CategoryID: "cat-movies", Text: "Which movie features the character Indiana Jones?", OptionA: "Star Wars", OptionB: "Indiana Jones and the Raiders of the Lost Ark", OptionC: "Jurassic Park", OptionD: "Back to the Future", CorrectOption: "B"},
		{ID: "q-62", CategoryID: "cat-movies", Text: "Who directed 'Jurassic Park'?", OptionA: "Steven Spielberg", OptionB: "James Cameron", OptionC: "Christopher Nolan", OptionD: "Ridley Scott", CorrectOption: "A"},
		{ID: "q-63", CategoryID: "cat-movies", Text: "Which film won Best Picture at the 2020 Academy Awards?", OptionA: "1917", OptionB: "Parasite", OptionC: "Joker", OptionD: "Once Upon a Time in Hollywood", CorrectOption: "B"},
		{ID: "q-64", CategoryID: "cat-movies", Text: "What is the name of the superhero in 'The Dark Knight'?", OptionA: "Spider-Man", OptionB: "Batman", OptionC: "Iron Man", OptionD: "Superman", CorrectOption: "B"},
		{ID: "q-65", CategoryID: "cat-movies", Text: "Which movie features a ship called the Titanic?", OptionA: "Titanic", OptionB: "The Poseidon Adventure", OptionC: "Pirates of the Caribbean", OptionD: "Life of Pi", CorrectOption: "A"},
		{ID: "q-66", CategoryID: "cat-movies", Text: "Which animated movie features a talking snowman named Olaf?", OptionA: "Frozen", OptionB: "Toy Story", OptionC: "Shrek", OptionD: "Moana", CorrectOption: "A"},
		{ID: "q-67", CategoryID: "cat-movies", Text: "What color pill does Neo take in 'The Matrix'?", OptionA: "Blue", OptionB: "Red", OptionC: "Green", OptionD: "Yellow", CorrectOption: "B"},
		{ID: "q-68", CategoryID: "cat-movies", Text: "Which film is set on a deserted island and stars Tom Hanks?", OptionA: "Cast Away", OptionB: "The Beach", OptionC: "Life of Pi", OptionD: "Robinson Crusoe", CorrectOption: "A"},
		{ID: "q-69", CategoryID: "cat-movies", Text: "Which film series features a character named Darth Vader?", OptionA: "Star Wars", OptionB: "Star Trek", OptionC: "Avatar", OptionD: "The Matrix", CorrectOption: "A"},
		{ID: "q-70", CategoryID: "cat-movies", Text: "Which movie features the song 'My Heart Will Go On'?", OptionA: "The Bodyguard", OptionB: "Titanic", OptionC: "Dirty Dancing", OptionD: "Ghost", CorrectOption: "B"},
		{ID: "q-71", CategoryID: "cat-music", Text: "Which band released the song 'Hey Jude'?", OptionA: "The Beatles", OptionB: "The Rolling Stones", OptionC: "Queen", OptionD: "Pink Floyd", CorrectOption: "A"},
		{ID: "q-72", CategoryID: "cat-music", Text: "Which musical instrument has 88 keys?", OptionA: "Guitar", OptionB: "Piano", OptionC: "Violin", OptionD: "Drums", CorrectOption: "B"},
		{ID: "q-73", CategoryID: "cat-music", Text: "Who is known as the 'King of Pop'?", OptionA: "Elvis Presley", OptionB: "Michael Jackson", OptionC: "Prince", OptionD: "Madonna", CorrectOption: "B"},
		{ID: "q-74", CategoryID: "cat-music", Text: "Which singer is famous for the song 'Rolling in the Deep'?", OptionA: "Adele", OptionB: "Beyoncé", OptionC: "Taylor Swift", OptionD: "Rihanna", CorrectOption: "A"},
		{ID: "q-75", CategoryID: "cat-music", Text: "What instrument does a drummer play?", OptionA: "Violin", OptionB: "Drums", OptionC: "Flute", OptionD: "Cello", CorrectOption: "B"},
		{ID: "q-76", CategoryID: "cat-music", Text: "Which band is known for the song 'Bohemian Rhapsody'?", OptionA: "Queen", OptionB: "The Beatles", OptionC: "Nirvana", OptionD: "U2", CorrectOption: "A"},
		{ID: "q-77", CategoryID: "cat-music", Text: "Which genre is Taylor Swift best known for?", OptionA: "Classical", OptionB: "Country/Pop", OptionC: "Hip hop", OptionD: "Jazz", CorrectOption: "B"},
		{ID: "q-78", CategoryID: "cat-music", Text: "How many strings does a standard guitar have?", OptionA: "4", OptionB: "6", OptionC: "7", OptionD: "8", CorrectOption: "B"},
		{ID: "q-79", CategoryID: "cat-music", Text: "Which city is home to the Nashville music scene?", OptionA: "Nashville", OptionB: "Austin", OptionC: "New Orleans", OptionD: "Los Angeles", CorrectOption: "A"},
		{ID: "q-80", CategoryID: "cat-music", Text: "Which musician is famous for the song 'Thriller'?", OptionA: "Prince", OptionB: "Michael Jackson", OptionC: "Madonna", OptionD: "Whitney Houston", CorrectOption: "B"},
		{ID: "q-81", CategoryID: "cat-math", Text: "What is 8 times 7?", OptionA: "54", OptionB: "56", OptionC: "58", OptionD: "60", CorrectOption: "B"},
		{ID: "q-82", CategoryID: "cat-math", Text: "What is the value of pi rounded to two decimal places?", OptionA: "3.14", OptionB: "3.15", OptionC: "3.13", OptionD: "3.12", CorrectOption: "A"},
		{ID: "q-83", CategoryID: "cat-math", Text: "What is the square root of 81?", OptionA: "7", OptionB: "8", OptionC: "9", OptionD: "10", CorrectOption: "C"},
		{ID: "q-84", CategoryID: "cat-math", Text: "What is 15% of 200?", OptionA: "20", OptionB: "25", OptionC: "30", OptionD: "35", CorrectOption: "C"},
		{ID: "q-85", CategoryID: "cat-math", Text: "What is the next prime number after 7?", OptionA: "9", OptionB: "10", OptionC: "11", OptionD: "13", CorrectOption: "C"},
		{ID: "q-86", CategoryID: "cat-math", Text: "What does 2 + 2 × 2 equal?", OptionA: "6", OptionB: "8", OptionC: "4", OptionD: "10", CorrectOption: "A"},
		{ID: "q-87", CategoryID: "cat-math", Text: "How many degrees are in a right angle?", OptionA: "45", OptionB: "90", OptionC: "180", OptionD: "360", CorrectOption: "B"},
		{ID: "q-88", CategoryID: "cat-math", Text: "What is the perimeter of a square with side length 5?", OptionA: "20", OptionB: "25", OptionC: "10", OptionD: "15", CorrectOption: "A"},
		{ID: "q-89", CategoryID: "cat-math", Text: "What is the first even prime number?", OptionA: "0", OptionB: "2", OptionC: "4", OptionD: "6", CorrectOption: "B"},
		{ID: "q-90", CategoryID: "cat-math", Text: "What is the sum of the angles in a triangle?", OptionA: "180", OptionB: "90", OptionC: "360", OptionD: "270", CorrectOption: "A"},
		{ID: "q-91", CategoryID: "cat-nature", Text: "What process do plants use to convert sunlight into energy?", OptionA: "Respiration", OptionB: "Photosynthesis", OptionC: "Digestion", OptionD: "Transpiration", CorrectOption: "B"},
		{ID: "q-92", CategoryID: "cat-nature", Text: "What is the largest land animal?", OptionA: "Elephant", OptionB: "Giraffe", OptionC: "Hippopotamus", OptionD: "Rhino", CorrectOption: "A"},
		{ID: "q-93", CategoryID: "cat-nature", Text: "Which animal is known as the king of the jungle?", OptionA: "Tiger", OptionB: "Lion", OptionC: "Elephant", OptionD: "Gorilla", CorrectOption: "B"},
		{ID: "q-94", CategoryID: "cat-nature", Text: "What do bees collect from flowers?", OptionA: "Nectar", OptionB: "Water", OptionC: "Pollen", OptionD: "Sap", CorrectOption: "A"},
		{ID: "q-95", CategoryID: "cat-nature", Text: "What is the chemical name for table salt?", OptionA: "Sodium chloride", OptionB: "Calcium carbonate", OptionC: "Potassium nitrate", OptionD: "Magnesium sulfate", CorrectOption: "A"},
		{ID: "q-96", CategoryID: "cat-nature", Text: "What layer of Earth do we live on?", OptionA: "Mantle", OptionB: "Core", OptionC: "Crust", OptionD: "Outer core", CorrectOption: "C"},
		{ID: "q-97", CategoryID: "cat-nature", Text: "What is the largest ocean on Earth?", OptionA: "Atlantic", OptionB: "Indian", OptionC: "Pacific", OptionD: "Arctic", CorrectOption: "C"},
		{ID: "q-98", CategoryID: "cat-nature", Text: "Which type of tree keeps its leaves year-round?", OptionA: "Deciduous", OptionB: "Coniferous", OptionC: "Maple", OptionD: "Birch", CorrectOption: "B"},
		{ID: "q-99", CategoryID: "cat-nature", Text: "What gas do humans exhale?", OptionA: "Oxygen", OptionB: "Nitrogen", OptionC: "Carbon dioxide", OptionD: "Hydrogen", CorrectOption: "C"},
		{ID: "q-100", CategoryID: "cat-nature", Text: "Which bird is known for mimicking sounds?", OptionA: "Owl", OptionB: "Parrot", OptionC: "Eagle", OptionD: "Pigeon", CorrectOption: "B"},
	}

	for _, question := range questions {
		if _, err := r.db.Exec(`INSERT OR IGNORE INTO questions (id, category_id, text, option_a, option_b, option_c, option_d, correct_option) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, question.ID, question.CategoryID, question.Text, question.OptionA, question.OptionB, question.OptionC, question.OptionD, question.CorrectOption); err != nil {
			return fmt.Errorf("insert question %s: %w", question.ID, err)
		}
	}

	return nil
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
