package auth

import (
    "log"
    "regexp"
    "github.com/hpcloud/tail"
)

func Loop(events chan<- string) {
    t, err := tail.TailFile("/var/log/auth.log", tail.Config{
        Follow: true})
    for line := range t.Lines {
        m := extractCommand(line.Text)
        if m != "" {
            events <- m
        }
    }
    if err != nil {
        log.Fatal(err)
    }
}

func extractCommand(s string) string {
    re := regexp.MustCompile("COMMAND=.*$")
    m := re.FindString(s)
    return m
}
