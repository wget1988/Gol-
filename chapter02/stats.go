package main

import "fmt"

func main() {
	// users := []string{"sql", "hyy", "waq", "sss", "sql", "sss"}
	// user_stat := make(map[string]int)

	// for _, user := range users {
	// 	if _, ok := user_stat[user]; !ok {
	// 		user_stat[user] = 1
	// 	} else {
	// 		user_stat[user]++
	// 	}
	// }

	// //user_stat[user]++
	// fmt.Println(user_stat)

	article :=
		`
	Five score years ago, a great American, in whose symbolic shadow we stand today, signed the Emancipation Proclamation. 
	This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. 
	It came as a joyous daybreak to end the long night of bad captivity.
	But one hundred years later, the Negro still is not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languished in the corners of American society and finds himself an exile in his own land. So we’ve come here today to dramatize a shameful condition.
	I am not unmindful that some of you have come here out of great trials and tribulations. Some of you have come fresh from narrow jail cells. Some of you have come from areas where your quest for freedom left you battered by the storms of persecution and staggered by the winds of police brutality. You have been the veterans of creative suffering. Continue to work with the faith that unearned suffering is redemptive.
	Go back to Mississippi, go back to Alabama, go back to South Carolina, go back to Georgia, go back to Louisiana, go back to the slums and ghettos of our northern cities, knowing that somehow this situation can and will be changed. Let us not wallow in the valley of despair.
	I say to you today, my friends, so even though we face the difficulties of today and tomorrow, I still have a dream. It is a dream deeply rooted in the American dream.
	I have a dream that one day this nation will rise up, live up to the true meaning of its creed: “We hold these truths to be self-evident; that all men are created equal.”
	I have a dream that one day on the red hills of Georgia the sons of former slaves and the sons of former slave-owners will be able to sit down together at the table of brotherhood.
	I have a dream that one day even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice.
	I have a dream that my four children will one day live in a nation where they will not be judged by the color if their skin but by the content of their character.
	I have a dream today.
	I have a dream that one day down in Alabama with its governor having his lips dripping with the words of interposition and nullification, one day right down in Alabama little black boys and black girls will be able to join hands with little white boys and white girls as sisters and brothers.
	I have a dream today.
	I have a dream that one day every valley shall be exalted, every hill and mountain shall be made low, the rough places will be made plain, and the crooked places will be made straight, and the glory of the Lord shall be revealed, and all flesh shall see it together.
	This is our hope. This is the faith that I go back to the South with. With this faith we will be able to hew out of the mountain of despair a stone of hope. With this faith we will be able to transform the jangling discords of our nation into a beautiful symphony of brotherhood. With this faith we will be able to work together, to pray together, to struggle together, to go to jail together, to stand up for freedom together, knowing that we will be free one day.
	This will be the day when all of God’s children will be able to sing with new meaning.
	`

	article_stats := map[rune]int{}

	for _, ch := range article {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			article_stats[ch]++
		}
	}
	fmt.Println(article_stats)

	for ch, cnt := range article_stats {

		fmt.Printf("%c:%d\n", ch, cnt)
	}
}
