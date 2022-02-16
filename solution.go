package solution
import "github.com/kyokomi/emoji/v2" 

func GetMessage() string {
	resultMessage := emoji.Sprint("Hello :world_map:!")
	return resultMessage
}
