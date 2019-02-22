package json

type Article struct {
	ID         int
	Title      int
	Content    string
	CreateTime string
	UpdateTime string
	Category   string
	Tags       []string
}

func (article *Article) ArticleJson() {

}
