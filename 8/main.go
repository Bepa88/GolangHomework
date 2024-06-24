package main

import (
	"context"
	"fmt"
	"time"
)

type Answer struct {
	Id        int
	Text      string
	IsCorrect bool
}

type Question struct {
	Id      int
	Text    string
	Answers []Answer
}

type Player struct {
	Name       string
	QuestionId int
	AnswerId   int
	IsCorrect  bool
}

type Quiz struct {
	Questions []Question
	Players   []Player
}

func main() {
	ctx := context.Background()
	ctxWithDeadline, cancelCtxWithDeadline := context.WithDeadline(ctx, time.Now().Add(time.Second*100))
	defer cancelCtxWithDeadline()

	answers_1 := []Answer{
		{Id: 1, Text: "A. Watching", IsCorrect: false},
		{Id: 2, Text: "B. Watch", IsCorrect: true},
		{Id: 3, Text: "C. Watches", IsCorrect: false},
		{Id: 4, Text: "D. Doesn't watch", IsCorrect: false},
	}

	answers_2 := []Answer{
		{Id: 1, Text: "A. Studying", IsCorrect: false},
		{Id: 2, Text: "B. Is studying", IsCorrect: true},
		{Id: 3, Text: "C. Is studing", IsCorrect: false},
		{Id: 4, Text: "D. Studies", IsCorrect: false},
	}

	answers_3 := []Answer{
		{Id: 1, Text: "A. Visits", IsCorrect: false},
		{Id: 2, Text: "B. Is visits", IsCorrect: false},
		{Id: 3, Text: "C. Is visiting", IsCorrect: true},
		{Id: 4, Text: "D. Visit", IsCorrect: false},
	}

	answers_4 := []Answer{
		{Id: 1, Text: "A. Download", IsCorrect: false},
		{Id: 2, Text: "B. Downloading", IsCorrect: false},
		{Id: 3, Text: "C. Are Downloading", IsCorrect: true},
		{Id: 4, Text: "D. Am Downloading", IsCorrect: false},
	}

	answers_5 := []Answer{
		{Id: 1, Text: "A. Smells", IsCorrect: true},
		{Id: 2, Text: "B. Is smelling", IsCorrect: false},
		{Id: 3, Text: "C. Smell", IsCorrect: false},
		{Id: 4, Text: "D. Are Smelling", IsCorrect: false},
	}

	questions := []Question{
		{Id: 1, Text: "  I always .... TV at home", Answers: answers_1},
		{Id: 2, Text: "  At the moment, Jerry .... at university", Answers: answers_2},
		{Id: 3, Text: "  My sisters .... next week", Answers: answers_3},
		{Id: 4, Text: "  Peter and I .... songs today", Answers: answers_4},
		{Id: 5, Text: "  Look! The food .... delicious", Answers: answers_5},
	}

	players := []Player{
		{Name: "Viktor"},
		{Name: "Ira"},
		{Name: "Oleksandr"},
	}

	quiz := Quiz{
		questions,
		players,
	}

	pChan := make(chan []Player)
	resultChan := make(chan []Player)
	qChan := make(chan Question, len(quiz.Questions))

	go generator(ctxWithDeadline, quiz, qChan)
	go sendQuestionsToPlayer(ctxWithDeadline, qChan, pChan, players)
	go checkAnswers(ctxWithDeadline, pChan, resultChan, questions)

	answeredPlayers := <-resultChan
	playerCorrectCount := make(map[string]int)
	questionCorrectCount := make(map[int]int)

	for _, pl := range answeredPlayers {
		if pl.IsCorrect {
			playerCorrectCount[pl.Name]++
			questionCorrectCount[pl.QuestionId]++
		}
	}

	for name, correctCount := range playerCorrectCount {
		fmt.Printf("%s answered correctly %d times.\n", name, correctCount)
	}

	for qId, correctCount := range questionCorrectCount {
		fmt.Printf("Question ID %d had %d correct answers.\n", qId, correctCount)
	}
}

func generator(ctx context.Context, q Quiz, questChan chan Question) {
	defer close(questChan)
	t := time.NewTicker(10 * time.Second)
	for _, val := range q.Questions {
		select {
		case <-ctx.Done():
			return
		case questChan <- val:
			<-t.C
		}
	}
}

func sendQuestionsToPlayer(ctx context.Context, questChan chan Question, playersChan chan []Player, players []Player) {
	defer close(playersChan)
	select {
	case <-ctx.Done():
		return
	default:
		var playersAnswer []Player
		for i := 1; i < 6; i++ {
			sendingQuestion := <-questChan

			fmt.Println(sendingQuestion.Text)
			for _, ans := range sendingQuestion.Answers {
				fmt.Println(ans.Text)
			}

			for i := range players {
				fmt.Printf("%s, your answer: ", players[i].Name)
				fmt.Scan(&players[i].AnswerId)
				players[i].QuestionId = sendingQuestion.Id
				playersAnswer = append(playersAnswer, players[i])
			}
		}
		playersChan <- playersAnswer
	}
}

func checkAnswers(ctx context.Context, playersChan chan []Player, resultChan chan []Player, questions []Question) {
	defer close(resultChan)
	var playersWithWriteAnswers []Player
	select {
	case <-ctx.Done():
		return
	case allPlayers := <-playersChan:
		for _, player := range allPlayers {
			for _, q := range questions {
				if player.QuestionId == q.Id {
					for _, ans := range q.Answers {
						if player.AnswerId == ans.Id {
							player.IsCorrect = ans.IsCorrect
							break
						}
					}
					break
				}
			}
			if player.IsCorrect {
				playersWithWriteAnswers = append(playersWithWriteAnswers, player)
			}
		}
		resultChan <- playersWithWriteAnswers
	}
}
