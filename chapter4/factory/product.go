package factory

/*
product: 抽象メソッドuseのみ定義されている抽クラス
factory: メソッドcreatを実装している抽象クラス
IDCard: メソッドuseを実装している
IdCardFactory: メソッドcreateProduct, registerProductを実装しているクラス
 */


// 製品を表現したクラス。
// 抽象メソッドのuse()のみ宣言されている。
// 製品は何はともあれ使える
type Product interface {
	Use()
}
