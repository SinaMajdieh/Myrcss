package rcss

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/chewxy/chexySexp"
	"github.com/nsf/sexp"
)

type Server interface {
	Match

	Stop() error
}

type server struct {
	raddr *net.UDPAddr
	conn  net.PacketConn
}

// type Init struct {
// 	init []string `sexp:"init,siblings"`
// }

func NewServer(addr string) (Server, error) {
	if raddr, err := net.ResolveUDPAddr("udp", addr); err != nil {
		return nil, err
	} else if conn, err := createUdpConn(); err != nil {
		return nil, err
	} else {
		srv := &server{
			raddr: raddr,
			conn:  conn,
		}

		return srv, nil
	}
}

func (s server) Stop() error {
	return s.conn.Close()
}

func (s server) bind(team Team) {
	l := make([]byte, 4096)
	var str string
	var ast *sexp.Node
	for {
		if _, _, err := s.conn.ReadFrom(l); err != nil {
			fmt.Printf("error: %s\n", err)
			return
		} else {
			str = string(l[:])
			ast, err := sexp.Parse(strings.NewReader(str), nil)
			if nil != err {
				fmt.Printf("error on : %s\n", err)
			}
			// var msg Message
			// if err := msg.UnmarshalBinary(l); err != nil {
			// 	fmt.Printf("message parse error: %s\n", err)
			// 	continue
			// }

			//fmt.Printf("%#v\n", msg)
			Sexp, err := chewxySexp.ParseString(str)
			if nil != err {
				fmt.Printf("Error : %s\n", err)
			}

			switch fmt.Sprint(Sexp[0].Head()) {
			case "init":
				fmt.Println(fmt.Sprint(Sexp[0].Head()))
				var m Init
				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal Init message: %s\n", err)

				// 	continue
				// }

				//go team.Init(s, m.Side, m.UniformNumber, m.PlayMode)
				err := ast.Unmarshal(&m.Init)
				if nil != err {
					panic(err)
				} else {
					m.SetValues()
				}

			case "server_param":
				fmt.Println(fmt.Sprint(Sexp[0].Head()))
				var m ServerParameters
				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal ServerParameters message: %s\n", err)

				// 	continue
				// }

				// go team.ServerParam(m)
				err := ast.Unmarshal(&m.ServerParameters)
				if nil != err {
					panic(err)
				} else {
					fmt.Println("vals")
					m.SetValues()
				}
			case "player_param":
				fmt.Println(fmt.Sprint(Sexp[0].Head()))
				var m PlayerParameters
				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal PlayerParameters message: %s\n", err)

				// 	continue
				// }

				// go team.PlayerParam(m)
				err := ast.Unmarshal(&m.PlayerParameters)
				if nil != err {
					panic(err)
				} else {
					m.SetValues()
				}

			case "player_type":
				fmt.Println(fmt.Sprint(Sexp[0].Head()))
				var m PlayerType
				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal PlayerType message: %s\n", err)

				// 	continue
				// }

				// go team.PlayerType(m)
				err := ast.Unmarshal(&m.PlayerType)
				if nil != err {
					panic(err)
				} else {
					m.SetValues()
				}

			case "see":

			case "hear":

			case "sense_body":

			case "score":

			case "error":

			default:
				//fmt.Printf("unhandled server input: `%s`\n", fmt.Sprint(Sexp[0]))
			}
		}
	}
	//Should be ommited , I just put it to get rid og "ast declered and not used" error!
	ast.Unmarshal()
}

func newInitCommand(teamName string, goalie bool, version int) Message {
	msg := Message{name: "init"}

	msg.AddValues(teamName)
	if version > 0 {
		ver := Message{name: "version"}
		ver.AddValues(strconv.Itoa(version))
		msg.AddSubmessages(ver)
	}
	if goalie {
		g := Message{name: "goalie"}
		msg.AddSubmessages(g)
	}

	return msg
}

func (s server) Join(team Team) error {
	go s.bind(team)

	cmd := newInitCommand(team.Name(), false, 15)

	if b, err := cmd.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Reconnect(team Team, unum UniformNumber) error {
	return nil
}

func (s server) Bye() error {
	return nil
}

func (s server) Catch(dir Direction) error {
	return nil
}

func (s server) ChangeView(w SightWidth, q SightQuality) error {
	return nil
}

func (s server) Dash() error {
	return nil
}

func (s server) Kick() error {
	return nil
}

func (s server) Move(x, y int) error {
	_, err := s.conn.WriteTo([]byte(fmt.Sprintf("(move %d %d)", x, y)), s.raddr)

	return err
}

func (s server) Say() error {
	return nil
}

func (s server) Turn() error {
	return nil
}

func (s server) TurnNeck() error {
	return nil
}

func (s server) Score() error {
	return nil
}

func (s server) See() error {
	return nil
}

func (s server) SenseBody() error {
	return nil
}
