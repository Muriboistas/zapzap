package img

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
	"github.com/gocolly/colly/v2"
)

var errImgNotFound = errors.New("image not founded")

func init() {
	commands.New(
		"img", img,
	).SetArgs(
		"text",
	).SetHelp(
		"Find some image",
	).SetCooldown(5).Add()
}

func img(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	text := args["text"]
	if text == "" {
		return errors.New("You can not search for a blank message")
	}

	c := colly.NewCollector()
	var find bool
	// Find and visit all links
	var body io.Reader
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("src"), "images?") && !find {
			find = true
			//Get the response bytes from the url
			res, err := http.Get(e.Attr("src"))
			if err != nil {
				return
			}
			defer res.Body.Close()
			//check if download works correctly
			if res.StatusCode != 200 {
				return
			}

			body = res.Body
			message.ReplyImg(body, wac, msg)

		}
	})

	// Set url queries
	v := url.Values{}
	v.Set("q", text)

	c.Visit(fmt.Sprintf("https://www.google.com/search?tbm=isch&%s", v.Encode()))

	if body == nil {
		return errImgNotFound
	}
	return nil
}
