package checktcpstatus

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type options struct {
	Status string `short:"s" long:"status" description:"tcp status\nTCP_ESTABLISHED|TCP_SYN_SENT|TCP_SYN_RECV|TCP_FIN_WAIT1|TCP_FIN_WAIT2\nTCP_TIME_WAIT|TCP_CLOSE|TCP_CLOSE_WAIT|TCP_LAST_ACK|TCP_LISTEN|TCP_CLOSING|TCP_NEW_SYN_RECV" required:"true"`
	// SrcPort  string  `short:"p" long:"port" description:"source pord number"`
	Warning  float64 `short:"w" long:"warning" description:"Warning threshold (num)"`
	Critical float64 `short:"c" long:"critical" description:"Critical threshold (num)"`
	Debug    bool    `short:"d" long:"debug" description:""`
}

type TCPStatus struct {
	Status map[string]float64
}

func (t *TCPStatus) parse(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var st string
	for n := 0; s.Scan(); n++ {
		// Skip the 1 header lines.
		if n < 1 {
			continue
		}

		// refs: https://elixir.bootlin.com/linux/latest/source/include/net/tcp_states.h#
		switch strings.Split(strings.Split(s.Text(), ":")[3], " ")[1] {
		case "01":
			st = "TCP_ESTABLISHED"
		case "02":
			st = "TCP_SYN_SENT"
		case "03":
			st = "TCP_SYN_RECV"
		case "04":
			st = "TCP_FIN_WAIT1"
		case "05":
			st = "TCP_FIN_WAIT2"
		case "06":
			st = "TCP_TIME_WAIT"
		case "07":
			st = "TCP_CLOSE"
		case "08":
			st = "TCP_CLOSE_WAIT"
		case "09":
			st = "TCP_LAST_ACK"
		case "0A":
			st = "TCP_LISTEN"
		case "0B":
			st = "TCP_CLOSING"
		case "0C":
			st = "TCP_NEW_SYN_RECV"
		}

		t.Status[st]++

	}

	return nil
}

func (t *TCPStatus) checkThreshold(opts options) {
	if opts.Debug {
		fmt.Println(t.Status)
		fmt.Println(t.Status[opts.Status])
	}

	chkSt := checkers.OK
	msg := "OK"
	if opts.Warning > 0 && t.Status[opts.Status] > opts.Warning {
		chkSt = checkers.WARNING
		msg = fmt.Sprintf("%s is %.0f", opts.Status, t.Status[opts.Status])
	}
	if opts.Critical > 0 && t.Status[opts.Status] > opts.Critical {
		chkSt = checkers.CRITICAL
		msg = fmt.Sprintf("%s is %.0f", opts.Status, t.Status[opts.Status])
	}

	checkers.NewChecker(chkSt, msg).Exit()
}

func Do() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}
	t := new(TCPStatus)
	t.Status = map[string]float64{}

	procs := []string{"/proc/net/tcp", "/proc/net/tcp6"}
	for _, f := range procs {
		t.parse(f)
	}

	t.checkThreshold(opts)
}
