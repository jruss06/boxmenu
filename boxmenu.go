package main

import (
    "os"
    "path/filepath"
    "fmt"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type desktopEntry struct {
    name string
    cat string
    command string
    used bool
}

type ConfigJson struct {
    OtherMenu bool
    Favorites [][]string
    SystemName string
    System [][]string
    Categories [][]string
}

var desktopMenu []desktopEntry

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func loadConf() ConfigJson {
    path := os.ExpandEnv("$HOME/.config/boxmenu/config.json")
    jsonRaw, err := ioutil.ReadFile(path)
    check(err)

    conf := ConfigJson{}

    err = json.Unmarshal(jsonRaw, &conf)
    check(err)

    return conf
}

func cleanCommand(command string) string {
    words := strings.Split(command, " ")
    for i, word := range words {
        if strings.Contains(word, "%") || word == "--" {
            words[i] = ""
        }
    }

    return strings.Join(words[:], " ")
}

func parseEntry(entry string) desktopEntry {
    data, err := ioutil.ReadFile(entry)
    check(err)

    strData := string(data)
    lines := strings.Split(strData, "\n")

    newItem := desktopEntry{}

    for _, line := range lines {
        keyValue := strings.SplitN(line, "=", 2)

        if keyValue[0] == "Name" && newItem.name == "" {
            newItem.name = strings.TrimSpace(keyValue[1])
        }
        if keyValue[0] == "Exec" && newItem.command == "" {
            newItem.command = keyValue[1]
        }
        if keyValue[0] == "Categories" && newItem.cat == "" {
            newItem.cat = strings.ReplaceAll(keyValue[1], ";", "")
        }
        if keyValue[0] == "NoDisplay" {
            newItem.used = true
        }
    }

    newCommand := cleanCommand(newItem.command)
    newItem.command = strings.TrimSpace(newCommand)

    return newItem
}

func countCatItems(catName string) int {
    count := 0
    for _, item := range desktopMenu {
       if strings.Contains(item.cat, catName) {
        count++
       }
    }
    return count
}

func getDesktopFiles() {
    entries,_ := filepath.Glob("/usr/share/applications/*.desktop")

    for _, entry := range entries {
        newItem := parseEntry(entry)
        if newItem.used == false {
            desktopMenu = append(desktopMenu, newItem)
        }
    }
}

func generate(confJson ConfigJson) {

    fmt.Println("<openbox_pipe_menu>")

    for _,fav := range confJson.Favorites {
        fmt.Printf("<item label=\"%s\"><action name=\"Execute\"><command><![CDATA[%s]]></command></action></item>\n", fav[0], fav[1])
    }

    fmt.Println("<separator label=\"Categories\"/>")
    for _, cat := range confJson.Categories {
        if countCatItems(cat[1]) > 0 {
            fmt.Printf("<menu id=\"%s\" label=\"%s\">\n", cat[1], cat[0])
            for i,item := range desktopMenu {
                if strings.Contains(item.cat, cat[1]) {
                    fmt.Printf("<item label=\"%s\"><action name=\"Execute\"><command><![CDATA[%s]]></command></action></item>\n", item.name, item.command)
                    desktopMenu[i].used = true
                }
            }
            fmt.Println("</menu>")
        }
    }

    fmt.Printf("<menu id=\"%s\" label=\"%s\">\n", "Other", "Other")
    for _, item := range desktopMenu {
        if item.used == false {
            fmt.Printf("<item label=\"%s\"><action name=\"Execute\"><command><![CDATA[%s]]></command></action></item>\n", item.name, item.command)
        }
    }
    fmt.Println("</menu>")

    fmt.Println("<separator/>")
    fmt.Printf("<menu id=\"%s\" label=\"%s\">\n", "System", "Openbox")
    for _, item := range confJson.System {
        fmt.Printf("<item label=\"%s\"><action name=\"Execute\"><command><![CDATA[%s]]></command></action></item>\n", item[0], item[1])
    }
    fmt.Println("</menu>")

    fmt.Println("</openbox_pipe_menu>")
}

func main() {
    confJson := loadConf()
    getDesktopFiles()
    generate(confJson)
}
