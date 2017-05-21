package repo

import (
	"time"

	"github.com/goodplayer/yaya/global"
)

type Post struct {
	_Id         string
	User        *User
	Status      int64
	_CreateTime time.Time
	_UpdateTime time.Time
	Rev         int64
	Parent      *Post
	Type        int64
	Title       string
	Summary     string
	Content     string
}

func defaultParent() *Post {
	p := new(Post)
	p._Id = "00000000000000000000000000000000"
	return p
}

func (this *Post) GetId() string {
	return this._Id
}

func (this *Post) GetCreateTime() time.Time {
	return this._CreateTime
}

func (this *Post) GetUpdateTime() time.Time {
	return this._UpdateTime
}

func (this *Post) SaveNewPost() error {
	this._Id = global.NextId()
	now := time.Now()
	this._CreateTime = now
	this._UpdateTime = now
	if this.Parent == nil {
		this.Parent = defaultParent()
	}

	_, err := _pool.Prepare(
		"saveNewPost",
		`insert into post(id, "user", status, createtime, updatetime, rev, parentid, type, title, summary, content) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
	)

	if err != nil {
		return err
	}

	_, err = _pool.Exec(
		"saveNewPost",
		this._Id,
		this.User.Id,
		this.Status,
		this._CreateTime,
		this._UpdateTime,
		this.Rev,
		this.Parent._Id,
		this.Type,
		this.Title,
		this.Summary,
		this.Content,
	)

	return err
}
