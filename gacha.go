package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	rank := rand.Intn(100)
	name := rand.Intn(100)

	switch {
	case rank < 80:
		switch {
		case name < 10:
			return &card{rarity: rarityN, name: "スライム"}
		case name < 20:
			return &card{rarity: rarityN, name: "スライムベス"}
		case name < 30:
			return &card{rarity: rarityN, name: "いたずらもぐら"}
		case name < 40:
			return &card{rarity: rarityN, name: "モーモン"}
		case name < 50:
			return &card{rarity: rarityN, name: "いっかくウサギ"}
		case name < 60:
			return &card{rarity: rarityN, name: "ドラキー"}
		case name < 70:
			return &card{rarity: rarityN, name: "おおきづち"}
		case name < 80:
			return &card{rarity: rarityN, name: "ゴースト"}
		case name < 90:
			return &card{rarity: rarityN, name: "リリパット"}
		default:
			return &card{rarity: rarityN, name: "ぶちスライム"}
		}
	case rank < 95:
		switch {
		case name < 25:
			return &card{rarity: rarityR, name: "ばくだん岩"}
		case name < 50:
			return &card{rarity: rarityR, name: "わらいぶくろ"}
		case name < 75:
			return &card{rarity: rarityR, name: "スライムナイト"}
		case name < 95:
			return &card{rarity: rarityR, name: "メタルスライム"}
		default:
			return &card{rarity: rarityR, name: "はぐれメタル"}
		}
	case rank < 99:
		switch {
		case name < 34:
			return &card{rarity: raritySR, name: "ゴールデンスライム"}
		case name < 67:
			return &card{rarity: raritySR, name: "キャプテン・クロウ"}
		default:
			return &card{rarity: raritySR, name: "バベルボブル"}
		}
	default:
		switch {
		case name < 50:
			return &card{rarity: rarityXR, name: "マスタードラゴン"}
		default:
			return &card{rarity: rarityXR, name: "オムド・ロレス"}
		}
	}
}
