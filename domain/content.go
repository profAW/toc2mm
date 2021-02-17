package domain


type Content []string

type IContentManagement interface{
	GetContent() Content
	AddContent(string)
	SetContent(Content)
	ClearContent()
}

var localContent Content

func (c Content) GetContent() Content {
	return localContent
}

func (c Content) SetContent(newContent Content) {
	localContent = newContent
}

func (c Content) AddContent(line string) {
	localContent = append(localContent,line)
}

func (c Content) ClearContent(){
	localContent = []string{}
}







