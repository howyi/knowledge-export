package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

type Post struct {
	Id       int    `json:"knowledgeId"`
	Title     string `json:"title"`
}

func main() {
	app := cli.NewApp()
	app.Name = "export"
	app.Usage = "export json from knowledge"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "group, g",
			Usage: "group name",
		},
		cli.StringFlag{
			Name:  "text, t",
			Usage: "search text",
		},
		cli.StringFlag{
			Name:  "path, p",
			Value: "./",
			Usage: "export json path",
		},
	}

	app.Action = func(c *cli.Context) error {
		endpoint := c.Args().Get(0) + "/api/knowledges"
		token := c.Args().Get(1)
		group := c.String("group")
		//text := c.String("text")
		// TODO

		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}

		q := u.Query()
		q.Set("private_token", token)
		q.Set("groups", group)
		u.RawQuery = q.Encode()

		res, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		byteArray, _ := ioutil.ReadAll(res.Body)

		var posts []Post
		if err := json.Unmarshal(byteArray, &posts); err != nil {
			log.Fatal(err)
		}

		for _, p := range posts {
			fmt.Printf("%d : %s\n", p.Id, p.Title)
		}

		// TODO pathにJSONを保存する

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}