package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nicholaskh/dist-db/config"
	"github.com/nicholaskh/dist-db/storage"
	log "github.com/nicholaskh/log4go"
)

type cmd struct {
	cmdLine string
	op      string
	params  []string
	storage storage.Storage
}

func NewCmd(cmdLine string) *cmd {
	this := new(cmd)
	this.cmdLine = this.trimCmdline(cmdLine)
	this.storage = storage.Factory(config.DistDb.Storage.StorageType)
	return this
}

func (this *cmd) Process() (string, error) {
	this.parse()
	return this.execute()
}

func (this *cmd) parse() {
	var parts []string
	//parts := make([]string, 20)
	for _, word := range strings.Split(this.cmdLine, " ") {
		if word != "" {
			parts = append(parts, word)
		}
	}
	log.Debug("Command: %s", parts)
	this.op = parts[0]
	this.params = parts[1:]
}

func (this *cmd) execute() (string, error) {
	var res string
	switch this.op {
	case OP_GET:
		// params: key
		var err error
		if len(this.params) != 1 {
			return "", ErrSyntaxError
		}
		res, err = this.get(this.params[0])
		if err != nil {
			return "", err
		}
	case OP_SET:
		// params: key value
		err := this.set(this.params[0], this.params[1])
		if err != nil {
			return "", err
		}
		res = OUTPUT_OK
	case OP_EMPTY:
		return "", nil
	default:
		return "", errors.New(fmt.Sprintf("wrong op \"%s\"", this.op))
	}

	return res, nil
}

func (this *cmd) trimCmdline(str string) string {
	return strings.TrimRight(str, string([]rune{0, 13, 10}))
}
