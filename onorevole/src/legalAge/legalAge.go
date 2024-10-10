package main
import (
	"fmt"
	"time"
)

func main() {
	//Initializing
	var birthDay, birthMonth, birthYear, currentDay, currentMonth, currentYear int
	var name string
	tm := time.Now()
	currentYear = tm.Year()
	currentMonth = int(tm.Month())
	currentDay = tm.Day()

	// The clownery
	fmt.Print("Have you ever been to a liquor shop, but you didn't buy anything because didn't know if you were old enough to be chuggin' that sweet sweet freedom nectar? (Oh no, how tragic!)\n")
	fmt.Print("Have you, yes you, ever met an attractive person of ripe and mature age (also known as P.I.L.F.), but you passed because you were afraid you would send them to jail? (Another chance wasted!)\n")
	fmt.Print("Well, look no longer! Here at VEN-Z, we have developed a cure to your chronic illness of doubt! (Really? It's for real?)\n")
	fmt.Print("Wwwwwwelcome ladies'n'gentlemens to the one! And only! LEGAL AGE TEST!!\n")
	fmt.Print("Now, who shall begin? (Me! Me! Pick Me!) Welllll, would you look at that cutie we got over there, on the other side of the screen! What's your name, sweetheart? (They went first! Oh the envy!)\n")
	fmt.Scan(&name)
	fmt.Print(name,"! What a WONDERFUL NAME (crowd cheering), no, what do I say? MONUMENTAL! So let's get started, ",name,", darling... Time for a quiz! (GASP)\n")
	fmt.Print("Oh, I know you're scared. But don't worry, no no no no no no no! Your brain is stuff for legends,",name," I'm ~sure~ you will answer in no time! (Give us the question! GIVE US THE QUESTION!)\nOhohohoh, since the crowd's getting impatient, let's get started! Give them a show, a-ok?\n")
	fmt.Print("Very well, ",name,". The first question is: ON WHAT DAY WERE YOU BORN?\n")
	fmt.Scan(&birthDay)
	fmt.Print("Amazing! Simply spectacular! No hesitation whatsoever! We don't want to lose the momentum, so let's get onto the next question! What MONTH were you born? (Please type it in numeric format, hehe, our translator got shot yesterday and we haven't found a replacement yet)\n")
	fmt.Scan(&birthMonth)
	fmt.Print("You're on a roll! Can you hear the crowd cheering, sugar? They're going at it like crazy! (",name,"!",name,"!",name,"!) Onto the final question, then. Are you ready?\n")
	fmt.Print("What... YEAR were you born in? (Le-Gasp! So tough! It must have been so long ago!) Shush folks, I'm sure ", name," here finds it easy-as-pie!\n")
	fmt.Scan(&birthYear)
	fmt.Print("Yes! You have done it! Splendid!\n")
	fmt.Print("It is with great pleasure that I inform you... YOU HAVE WON THE QUIZ!! (crowd screaming in pain)\n")
	fmt.Print("Are you ready then, for the answer? The final revelation that will reshape your reality, remake you a new person, restructure your bounds, your anus, your liver and potentiallyyourwholebodyifyougetintoacarcrash?\n")
	fmt.Print("You are... (everyone shuts the fuck up)\n")
	fmt.Print("...In fact...\n")
	if currentYear < birthYear || currentYear == birthYear && currentMonth < birthMonth || currentYear == birthYear && currentMonth == birthMonth && currentDay < birthDay {
		fmt.Print("From... the... future? Hold on a second, folks, we might need to double check. We'll... we'll let you know. Maybe stab this little fucker in the meantime, cuz that's what you get for ruining our show. WELL, WHO'S NEXT? LET'S GO, BABY!\n")
	} else if currentYear - birthYear > 18 || currentYear - birthYear == 18 && currentMonth >= birthMonth || currentYear - birthYear == 18 && currentMonth == birthMonth && currentDay >= birthDay  {
		fmt.Print("NOT!! UNDERAGE!!! Yes, you heard me right! You can drink, have sex, vote, drive, go do drugs for all I care, time flies and we have a long line of people who want their lives enlightened! Still it was a truly fantastic start to the night! Stay tuned and... see you next week!\n")
	} else if birthYear == 0 || birthMonth == 0 || birthDay == 0 {
		fmt.Print("*gets shot* (audience gasps)\n")
	} else {
		fmt.Print("UNDERAGE!! Sorry pal, your life still sucks. I'd also like to inform you that minors are BANNED from this show, so please be a doll and go give this $10000 fine to your closest parent or legal guardian, hmm? A-HEM SO, WHO WANTS TO TRY THEIR LUCK NEXT?\n")
	}
	if currentMonth == birthMonth && currentDay == birthDay {
		fmt.Print("Oh, by the way, happy birthday sweetie!\n")
	}
	fmt.Print("\nwow that was weird\n")
}