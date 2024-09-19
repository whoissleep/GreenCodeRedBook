package storage

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/GreenCodeBook/src/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dialer struct {
	client *ssh.Client
}

type SSH struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type Postgres struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (m *Postgres) New() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s database=%s postgres+ssh(%s:%d)?charset=utf8&parseTime=True&loc=Local",
		m.User, m.Password, m.Database, m.Host, m.Port)
	fmt.Println(dsn)
	sqlDB, err := sql.Open("postgres+ssh", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", m.User, m.Password, m.Host, m.Database))
	if err != nil {
		fmt.Println("here1")
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("here2")
		panic(err)
	}
	fmt.Println("Conn to db done")
	return

}
func (sel *Dialer) Dial(network, address string) (net.Conn, error) {
	return sel.client.Dial(network, address)
}
func (sel *Dialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(sel, s)
}
func (sel *Dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return sel.client.Dial(network, address)
}

func (s *SSH) DialWithPassword() (*ssh.Client, error) {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", address, config)
}

var dial *ssh.Client

func Init() *gorm.DB {
	sshPort, _ := strconv.Atoi(os.Getenv("SSH_PORT"))
	client := SSH{
		Host:     os.Getenv("SSH_HOST"),
		User:     os.Getenv("SSH_USER"),
		Port:     sshPort,
		Password: os.Getenv("SSH_PASSWORD"),
	}
	dbPort, _ := strconv.Atoi(os.Getenv("BD_PORT"))
	mydb := Postgres{
		Host:     os.Getenv("BD_HOST"),
		User:     os.Getenv("BD_USER"),
		Password: os.Getenv("BD_PASSWORD"),
		Port:     dbPort,
		Database: os.Getenv("BD_DATABASE"),
	}
	dial, er := client.DialWithPassword()

	if er != nil {
		panic("NO SSH connection")
	}
	fmt.Println("SSH conn done")
	//defer dial.Close()

	sql.Register("postgres+ssh", &Dialer{dial})
	db, err := mydb.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("ok")
	err = db.AutoMigrate(
		&models.User{},
		&models.Note{},
	)
	if err != nil {
		fmt.Println("here3")
		panic(err)
	}
	fmt.Println("ok")
	return db
}

func Close() {
	dial.Close()
}
