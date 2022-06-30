package storage

import (
	"fmt"
	"github.com/hirochachacha/go-smb2"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"net"
	"os"
)

var mediaStorage *sambaDriver = nil

func NewMediaSambaStorage() Driver {
	if mediaStorage != nil {
		return mediaStorage
	}

	fmt.Println(
		config.Env.MediaSambaHostName,
		config.Env.MediaSambaUserName,
		config.Env.MediaSambaPassword,
		config.Env.MediaSambaShareName,
	)
	d := &sambaDriver{}
	if err := d.connect(
		"tcp",
		config.Env.MediaSambaHostName,
		config.Env.MediaSambaUserName,
		config.Env.MediaSambaPassword,
		config.Env.MediaSambaShareName,
	); err != nil {
		panic(err)
	}
	fmt.Println("connected smb server", config.Env.MediaSambaHostName)

	mediaStorage = d
	return mediaStorage
}

type sambaDriver struct {
	conn net.Conn
	dial *smb2.Dialer
	sess *smb2.Session
	fs   *smb2.Share
}

func (d *sambaDriver) connect(protocol, hostname, user, password, share string) error {
	conn, err := net.Dial(protocol, hostname)
	if err != nil {
		return errors.New(errors.SambaConnectFatal, err)
	}

	dial := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:        user,
			Password:    password,
			Hash:        nil,
			Domain:      "",
			Workstation: "",
			TargetSPN:   "",
		},
	}
	sess, err := dial.Dial(conn)
	if err != nil {
		return errors.New(errors.SambaConnectFatal, err)
	}

	fs, err := sess.Mount(share)
	if err != nil {
		return errors.New(errors.SambaConnectFatal, err)
	}

	d.conn = conn
	d.dial = dial
	d.sess = sess
	d.fs = fs

	return nil
}

func (d *sambaDriver) CreateFile(filePath string, data []byte) error {
	f, err := d.fs.Create(filePath)
	defer f.Close()
	if err != nil {
		return errors.New(errors.SambaCreateFatal, err)
	}
	if _, err := f.Write(data); err != nil {
		return errors.New(errors.SambaCreateFatal, err)
	}

	return nil
}

func (d *sambaDriver) CreateDir(dirPath string, perm os.FileMode) error {
	if err := d.fs.MkdirAll(dirPath, perm); err != nil {
		return errors.New(errors.SambaCreateDirFatal, err)
	}
	return nil
}

func (d *sambaDriver) Rename(old, file string) error {
	if err := d.fs.Rename(old, file); err != nil {
		return errors.New(errors.SambaRenameFatal, err)
	}
	return nil
}

func (d *sambaDriver) ReadDir(dirPath string) ([]os.FileInfo, error) {
	list, err := d.fs.ReadDir(dirPath)
	if err != nil {
		return nil, errors.New(errors.SambaReadDirFatal, err)
	}
	return list, nil
}

func (d *sambaDriver) ReadFile(filePath string) ([]byte, error) {
	data, err := d.fs.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(errors.SambaReadFatal, err)
	}
	return data, nil
}

func (d *sambaDriver) Glob(pattern string) ([]string, error) {
	matches, err := d.fs.Glob(pattern)
	if err != nil {
		return nil, errors.New(errors.SambaGlobFatal, err)
	}
	return matches, nil
}

func (d *sambaDriver) Exist(filePath string) bool {
	_, err := d.fs.Stat(filePath)
	return err == nil
}

func (d *sambaDriver) Delete(filePath string) error {
	if err := d.fs.Remove(filePath); err != nil {
		return errors.New(errors.SambaDeleteFatal, err)
	}
	return nil
}

func (d *sambaDriver) DeleteAll(path string) error {
	if err := d.fs.RemoveAll(path); err != nil {
		return errors.New(errors.SambaDeleteAllFatal, err)
	}
	return nil
}

func (d *sambaDriver) Close() {
	if d.fs != nil {
		d.fs.Umount()
	}
	if d.sess != nil {
		d.sess.Logoff()
	}
	if d.conn != nil {
		d.conn.Close()
	}
}
