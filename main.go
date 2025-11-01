package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "8375922369:AAExuMsDGpvdBR8g40Lx6FRH6R0AxyzSYY8",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello user! This bot can manage the folder \"folder\". To wiev all comand please enter /help")
	})
	b.Handle("/help", func(c tele.Context) error {
		return c.Send("/createfile - create file with your name \n /rmfile - remove file \n /lsdir - shows all files in a directory")
	})
	var userinput string
	b.Handle("/createfile", func(c tele.Context) error {
		userinput = "createfile"
		return c.Send("Enter file name:")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		filename := c.Text()
		if userinput == "createfile" {
			filepath := "folder/" + filename
			os.Create(filepath)
			return c.Send("File is create!")
		} else if userinput == "rmfile" {
			filepath := "folder/" + filename
			err := os.Remove(filepath)
			if err != nil {
				return c.Send("No such file")
			}
			return c.Send("File is remove!")
		}
		return c.Send("")
	})
	b.Handle("/rmfile", func(c tele.Context) error {
		userinput = "rmfile"
		return c.Send("Enter file name:")
	})
	b.Handle("/lsdir", func(c tele.Context) error {
		list, err := os.ReadDir("folder/")
		if err != nil {
			return c.Send("No such directory!")
		}
		var fileList string
		for _, file := range list {
			fileList += "File: " + file.Name() + "\n"
			c.Send(fileList)
		}

		return c.Send("These file in folder")
	})

	fmt.Println("Bot is running")
	b.Start()
}
