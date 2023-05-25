package battle

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

var tempPosition pixel.Vec
var dropRandomItem = ""

func BattleEndScreen(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus, enemy *enemy.EnemyStatus) myGame.GameState {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(0, 0)
	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, "リザルト  再戦 : Press Tab | 町に戻る : Press BackSpace")
	tempPosition = myPos.TopCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)
	if dropRandomItem == "" {
		dropRandomItem = enemy.DropItems[rand.Intn(3)]
		myGame.SaveFileItemsLoad(myGame.SaveFilePathItems)
		myGame.SaveGameItems(myGame.SaveFilePathItems, player, dropRandomItem)
	}
	ClearTxt := []string{"獲得ゴールド:" + strconv.Itoa(gainGold) + "S", "入力単語数:" + strconv.Itoa(score), "正解タイプ数:" + strconv.Itoa(collectType),
		"正解タイプ数:" + strconv.Itoa(collectType), "ミスタイプ数:" + strconv.Itoa(missType), "ドロップ素材:" + dropRandomItem, "獲得AP:" + strconv.Itoa(enemy.DropAP),
	}

	if player.HP > 0 {
		for _, value := range ClearTxt {
			myUtil.ScreenTxt.Clear()
			myUtil.ScreenTxt.Color = colornames.White
			fmt.Fprintln(myUtil.ScreenTxt, value)
			yOffSet -= myUtil.ScreenTxt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			myUtil.ScreenTxt.Draw(win, tempPosition)
		}
		//TODO: 最初の討伐時だけ武器の追加を教える.
		// tempInt, _ := strconv.Atoi(player.PossessedWeapon[0])
		// if tempInt >= 1 {
		// 	endLines = append(endLines, "武器屋に新しい武器が追加された<-New!!")
		// }
		// log.Println(endLines)
	} else {
		yourTimeString := fmt.Sprintf("%0.3f", yourTime)
		//平均キータイプ数 回/秒 Escでもう一度,Tabでタイトル
		endLines := []string{
			"GameOver...",
			"You lost " + strconv.Itoa(lostGold) + " gold",
			"YourScore : " + strconv.Itoa(score),
			"\n",
			"yourTime =" + yourTimeString,
			"collectType = " + strconv.Itoa(collectType) + " missType = " + strconv.Itoa(missType),
			"\n\n",
			"ReSTART : Press Tab | Back : Press BackSpace",
		}
		myPos.LineCenterAlign(win, endLines, Txt, "center")
	}

	//画面遷移,いろいろリセット
	if win.JustPressed(pixelgl.KeyTab) {
		myGame.CurrentGS = myGame.PlayingScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		enemy.HP = enemy.MaxHP
		dropRandomItem = ""
		Shuffle(words)
		myUtil.SetSaveReset(false)
		log.Println("Press:Enter -> GameState:Playing")
	} else if win.JustPressed(pixelgl.KeyBackspace) {
		myGame.CurrentGS = myGame.GoToScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		enemy.HP = enemy.MaxHP
		dropRandomItem = ""
		Shuffle(words)
		myUtil.SetSaveReset(false)
		log.Println("Press:Enter -> GameState:GoToScreen")
	}

	return myGame.CurrentGS
}
