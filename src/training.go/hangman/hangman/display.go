package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`

/$$   /$$                                                                
| $$  | $$                                                                
| $$  | $$  /$$$$$$  /$$$$$$$   /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$$$$$$ 
| $$$$$$$$ |____  $$| $$__  $$ /$$__  $$| $$_  $$_  $$ |____  $$| $$__  $$
| $$__  $$  /$$$$$$$| $$  \ $$| $$  \ $$| $$ \ $$ \ $$  /$$$$$$$| $$  \ $$
| $$  | $$ /$$__  $$| $$  | $$| $$  | $$| $$ | $$ | $$ /$$__  $$| $$  | $$
| $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$| $$  | $$
|__/  |__/ \_______/|__/  |__/ \____  $$|__/ |__/ |__/ \_______/|__/  |__/
                               /$$  \ $$                                  
                              |  $$$$$$/                                  
                               \______/                                   
 `)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter %v was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%v' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost :( The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You won! The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		____
	   |    |      
	   |    o      
	   |   /|\     
	   |    |
	   |   / \
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 1:
		draw = `
		____
	   |    |      
	   |    o      
	   |   /|\     
	   |    |
	   |    
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 2:
		draw = `
		____
	   |    |      
	   |    o      
	   |    |
	   |    |
	   |     
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 3:
		draw = `
		____
	   |    |      
	   |    o      
	   |        
	   |   
	   |   
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 4:
		draw = `
		____
	   |    |      
	   |      
	   |      
	   |  
	   |  
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 5:
		draw = `
		____
	   |        
	   |        
	   |        
	   |   
	   |   
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 6:
		draw = `
		
	   |     
	   |     
	   |     
	   |
	   |
	  _|_
	 |   |______
	 |          |
	 |__________|
			`
	case 7:
		draw = `
	  _ _
	 |   |______
	 |          |
	 |__________|
			`
	case 8:
		draw = `
	
			`
	}
	fmt.Println(draw)
	fmt.Printf("Turns left: %v\n", l)
}
