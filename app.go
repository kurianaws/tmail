package main

import "github.com/go-mail/mail"
import "os"
import "fmt"
import "strings"

// TODO this is a function for sending group mail
// TODO specify the list of recipients in csv file
// TODO set email templates to be sent as mail
// TODO set parameters
// TODO get parameters from google
//const template = []string{`<h1>hello</h1>`, `<h1>Bye</h1>`}

func main() {
	//path, _ := getPath()
	//list, _ := readEmailList(path)
	//template := getTemplate()
	list := []string{"kurianck.mail@gmail.com"}
	err := sendEmail(list)
	if err != nil {
		fmt.Println(err)
	}
}

func getTemplate() (string, error) {
	_, err := getPathFromEnv()
	if err != nil {
		return "", err
	}
	return "", nil
	//return template[2], nil
}

func getPath() (string, error) {
	vars, err := getPathFromEnv()
	if err != nil {
		return "", err
	}
	return string(vars[1]), nil
}

func getPathFromEnv() ([]string, error) {
	// this returns the path of file from args
	list := make([]string, 2)
	for _, entry := range os.Args {
		list = append(list, entry)
	}
	return list, nil
}
func readEmailList(path string) (*[]string, error) {
	// this function reads from csv file, from the file path
	// it returns email list as a slice of strings
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data := make([]byte, 100)
	count, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	var list = strings.Split(fmt.Sprintf("%s", data), ",")
	fmt.Printf("%s %d %v", data, count, list)
	return &list, nil
}

func sendEmail(list []string) error {
	// this function sends email to all recipients

	for _, recept := range list {
		m := mail.NewMessage()
		m.SetHeader("From", "kurianck.mail@gmail.com")
		m.SetHeader("To", recept)
		m.SetHeader("Subject", "Test mail")
		m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
		d := mail.NewDialer("smtp.gmail.com", 587, "kurianck.mail@gmail.com", "cqamwilingcswxok")
		d.StartTLSPolicy = mail.MandatoryStartTLS
		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
	return nil
}
